package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mars_agent/configs/base"
	"mars_agent/utils"
)

var bucketName = "mars"

func OssCreateBucket(bucketName string) error {
	// 创建OSSClient实例
	client, err := getOssClient()
	if err != nil {
		utils.HandleError(err, "创建client实例失败")
		return err
	}
	// 创建存储空间。
	err = client.CreateBucket(bucketName)
	if err != nil {
		utils.HandleError(err, "创建buckt失败")
	}
	return err
}

func OssUpload(localFile string, remoteOssName string) (err error) {
	// 创建OSSClient实例。
	client, err := getOssClient()
	if err != nil {
		utils.HandleError(err, "创建client实例失败")
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		utils.HandleError(err, "获取存储空间")
		return
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(remoteOssName, localFile)
	if err != nil {
		utils.HandleError(err, "上传失败")
		return
	}
	return
}

func OssDownloadFile(remoteFileName, localStoreFile string) (err error) {
	// 创建OSSClient实例。
	client, err := getOssClient()
	if err != nil {
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		utils.HandleError(err, "获取存储buck失败")
		return
	}
	// 下载文件。
	err = bucket.GetObjectToFile(remoteFileName, localStoreFile)
	if err != nil {
		utils.HandleError(err, "下载失败")
	}
	return
}

func OssObjectExist(remoteFileName string) (bool, error) {
	client, err := getOssClient()
	if err != nil {
		return false, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		utils.HandleError(err, "获取存储buck失败")
		return false, err
	}
	// 判断文件是否存在。
	exits, err := bucket.IsObjectExist(remoteFileName)
	if err != nil {
		utils.HandleError(err, "判断文件是否存在失败")
	}
	return exits, err
}

func getOssClient() (client *oss.Client, err error) {
	endpoint := base.AppConfig.Oss.EndPoint
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId := base.AppConfig.Oss.AccessKey
	accessKeySecret := base.AppConfig.Oss.AccessSecret

	// 创建OSSClient实例。
	client, err = oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		utils.HandleError(err, "创建client实例失败")
	}
	return
}
