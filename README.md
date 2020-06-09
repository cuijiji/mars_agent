## mars agent


#### 构建
```
make VERSION=1.0.0 all
cd ./bin/xxx/
./mars_agent -version

#ouput:
mars agent 
git_tag:, version:1.0.0, build_time: (2019-06-21T14:19:48+0800)

./mars_agent -h

#output:
mars agent version: 0.0.2
Usage: mars_agent  [-s name] [-version] [-h]

Options:
  -h    this help
  -s name
        start consul mode by name : rec, play
  -version
        print agent version information
```

#### 运行
```
./mars_agent -s gor

or

./mars_agent -s jmeter
```
#### 需要软件 jmeter 与 gor 
```
下载地址 
```

#### 使用consul 需要加配置 
```
app:
  env: release
  port: 22011
oss:
  end-point: "http://oss-cn-beijing-internal.aliyuncs.com"
  access-key: ""
  access-secret: ""
commands:
  - name: ls
    exec: [echo]
  - name: gor
    exec: [/mars/gor]
```

#### 使用consul 需要加配置 
```
app:
  env: release
  port: 22001
oss:
  end-point: "http://oss-cn-beijing-internal.aliyuncs.com"
  access-key: ""
  access-secret: ""
commands:
  - name: ls
    exec: [echo]
  - name: jmeter
    exec: [/mars/apache-jmeter-5.0-pure/bin/jmeter]
```
