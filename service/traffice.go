package service

import (
	"bytes"
	"mars_agent/gor"
	"mars_agent/utils"
)

func init() {

}

func GorFileName(taskId string) string {
	return "task_" + taskId + ".log"
}

func OssFileName(taskId string) string {
	return "oss_task_" + taskId + ".csv"
}

func JmxFileName(jmxId string) string {
	return "jmeter_" + jmxId + ".jmx"
}

func JmxSampleFileName(jmxId string) string {
	return "jmeter_sample_" + jmxId + ".jmx"
}

func TrafficDyeingToCsv(path string, taskId string) (err error) {
	fileName := GorFileName(taskId)
	reader, err := gor.NewFileInputReader(path + "/" + fileName)
	if err != nil {
		utils.HandleError(err, "file no find in "+path)
		return err
	}
	ossName := OssFileName(taskId)
	file, err := gor.NewFileOutPutWrite(path + "/" + ossName)
	if err != nil {
		utils.HandleError(err, "error cause by create file at  "+path)
		return err
	}
	for {
		payload := reader.ReadPayload()
		if payload == nil {
			break
		}
		join := bytes.Join(payload, []byte("||"))
		_ = file.Write(join)
		_ = file.Write([]byte("\n"))
	}
	return file.Close()
}
