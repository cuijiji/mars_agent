package consul

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/toolkits/net"
	"gopkg.in/yaml.v2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	. "mars_agent/configs/base"
	"mars_agent/configs/code"
	. "mars_agent/configs/log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ServiceInfo struct {
	ServiceID string
	IP        string
	Port      int
	Load      int
	Timestamp int //load updated ts
}

type ServiceList []ServiceInfo

type KVData struct {
	Load      int `json:"load"`
	Timestamp int `json:"ts"`
}

var (
	serviceMaps   = make(map[string]ServiceList)
	serviceLocker = new(sync.Mutex)
	consulClient  *api.Client
	serviceId     string
	serviceName   string
	kvKey         string
	LocalIp, _    = net.IntranetIP()
)

func InitConsulConfig(serverMode string) {
	keyName := "config/" + serverMode + "/data"
	consulData := GetKeyValueByName(keyName)
	if consulData == "" {
		fmt.Println("获取配置返回为空，key:" + keyName)
		os.Exit(code.ConsulConfigIsNull)
	}
	config := Config{}

	if err := yaml.Unmarshal([]byte(consulData), &config); err != nil {
		fmt.Printf("获取consul配置失败,key:%s,%s", keyName, err)
		os.Exit(code.ConsulConfigIsNull)
	}
	if config.App.Port <= 0 {
		config.App.Port = 22641
	}
	config.App.Model = serverMode
	AppConfig = config
}

func CheckErr(err error) {
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check status.")
	fmt.Fprint(w, "status ok!")
}

func StartService(addr string) {
	http.HandleFunc("/status", StatusHandler)
	fmt.Println("start listen...")
	err := http.ListenAndServe(addr, nil)
	CheckErr(err)
}

//func main() {
//	var status_monitor_addr, service_name, service_ip, consul_addr, found_service string
//	var service_port int
//	flag.StringVar(&consul_addr, "consul_addr", "localhost:8500", "host:port of the service stuats monitor interface")
//	flag.StringVar(&status_monitor_addr, "monitor_addr", "127.0.0.1:54321", "host:port of the service stuats monitor interface")
//	flag.StringVar(&service_name, "service_name", "worker", "name of the service")
//	flag.StringVar(&service_ip, "ip", "127.0.0.1", "service serve ip")
//	flag.StringVar(&found_service, "found_service", "worker", "found the target service")
//	flag.IntVar(&service_port, "port", 4300, "service serve port")
//	flag.Parse()
//
//	serviceName = service_name
//
//	DoRegisterService(consul_addr, status_monitor_addr, service_name, service_ip, service_port)
//
//	go DoDiscover(consul_addr, found_service)
//
//	go StartService(status_monitor_addr)
//
//	go WaitToUnRegisterService()
//
//	go DoUpdateKeyValue(consul_addr, service_name, service_ip, service_port)
//
//	select {}
//}

func DoRegisterGrpcService(serviceName string, ip string, port int) error {
	serviceId = serviceName + "-" + ip
	var tags []string
	service := &api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    serviceName,
		Port:    port,
		Address: ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "1s",
			GRPC:                           ip + ":" + strconv.Itoa(port),
			CheckID:                        serviceId,
			GRPCUseTLS:                     false,
			DeregisterCriticalServiceAfter: "10s",
			Name:                           serviceName,
		},
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		Logger.Sugar().Error("get consul http client error", err)
		return err
	}

	consulClient = client
	if err := consulClient.Agent().ServiceRegister(service); err != nil {
		Logger.Sugar().Error("register consul service error", err)
		return err
	}
	Logger.Sugar().Info(fmt.Sprintf("Registered grpc service %q in consul with tags %q", serviceName, strings.Join(tags, ",")))
	return nil
}

func DoRegisterService(consulAddr string, monitorAddr string, serviceName string, ip string, port int) {
	serviceId = serviceName + "-" + ip
	var tags []string
	service := &api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    serviceName,
		Port:    port,
		Address: ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			Interval: "5s",
			Timeout:  "1s",
			GRPC:     ip + ":" + strconv.Itoa(port),
		},
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	consulClient = client
	if err := consulClient.Agent().ServiceRegister(service); err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered service %q in consul with tags %q", serviceName, strings.Join(tags, ","))
}

func WaitToUnRegisterService() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	if consulClient == nil {
		return
	}
	if err := consulClient.Agent().ServiceDeregister(serviceId); err != nil {
		log.Fatal(err)
	}
}

func DoDiscover(consulAddr string, foundService string) {
	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-t.C:
			DiscoverServices(consulAddr, true, foundService)
		}
	}
}

func DiscoverServices(addr string, healthyOnly bool, serviceName string) {
	consulConf := api.DefaultConfig()
	consulConf.Address = addr
	client, err := api.NewClient(consulConf)
	CheckErr(err)

	services, _, err := client.Catalog().Services(&api.QueryOptions{})
	CheckErr(err)

	fmt.Println("--do discover ---:", addr)

	var sers ServiceList
	for name := range services {
		servicesData, _, err := client.Health().Service(name, "", healthyOnly,
			&api.QueryOptions{})
		CheckErr(err)
		for _, entry := range servicesData {
			if serviceName != entry.Service.Service {
				continue
			}
			for _, health := range entry.Checks {
				if health.ServiceName != serviceName {
					continue
				}
				fmt.Println("  health nodeid:", health.Node, " serviceName:", health.ServiceName, " service_id:", health.ServiceID, " status:", health.Status, " ip:", entry.Service.Address, " port:", entry.Service.Port)

				var node ServiceInfo
				node.IP = entry.Service.Address
				node.Port = entry.Service.Port
				node.ServiceID = health.ServiceID

				//get data from kv store
				s := GetKeyValue(serviceName, node.IP, node.Port)
				if len(s) > 0 {
					var data KVData
					err = json.Unmarshal([]byte(s), &data)
					if err == nil {
						node.Load = data.Load
						node.Timestamp = data.Timestamp
					}
				}
				fmt.Println("service node updated ip:", node.IP, " port:", node.Port, " serviceid:", node.ServiceID, " load:", node.Load, " ts:", node.Timestamp)
				sers = append(sers, node)
			}
		}
	}

	serviceLocker.Lock()
	serviceMaps[serviceName] = sers
	serviceLocker.Unlock()
}

func DoUpdateKeyValue(consulAddr string, serviceName string, ip string, port int) {
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			StoreKeyValue(consulAddr, serviceName, ip, port)
		}
	}
}

func StoreKeyValue(consulAddr string, serviceName string, ip string, port int) {

	kvKey = serviceName + "/" + ip + ":" + strconv.Itoa(port)

	var data KVData
	data.Load = rand.Intn(100)
	data.Timestamp = int(time.Now().Unix())
	bys, _ := json.Marshal(&data)

	kv := &api.KVPair{
		Key:   kvKey,
		Flags: 0,
		Value: bys,
	}

	_, err := consulClient.KV().Put(kv, nil)
	CheckErr(err)
	fmt.Println(" store data key:", kv.Key, " value:", string(bys))
}

func GetKeyValue(serviceName string, ip string, port int) string {
	key := serviceName + "/" + ip + ":" + strconv.Itoa(port)

	kv, _, err := consulClient.KV().Get(key, nil)
	if kv == nil {
		return ""
	}
	CheckErr(err)

	return string(kv.Value)
}

func GetKeyValueByName(kvName string) string {
	if consulClient == nil {
		client, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			return ""
		}
		consulClient = client

	}
	kv, _, err := consulClient.KV().Get(kvName, nil)
	if kv == nil {
		return ""
	}

	CheckErr(err)

	return string(kv.Value)
}

//func Wactch(kvName string) string {
//	watch.Parse()
//}
