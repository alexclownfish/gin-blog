package model

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/wonderivan/logger"
	"sync"

	//"math/rand"
	//"strings"
	//"fmt"
	"gin-blog/utils"
	"gin-blog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

////随机生成字符串
//var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
//	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
//	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
//
//func RandString(lenNum int) string {
//	str := strings.Builder{}
//	length := 52
//	for i := 0; i < lenNum; i++ {
//		str.WriteString(CHARS[rand.Intn(length)])
//	}
//	return str.String()
//}

//var FileUrl = utils.QiniuSoftServer
//var VideoUrl = utils.QiniuVideoServer

//文件上传至七牛云

func Upload(file *multipart.FileHeader, wg *sync.WaitGroup) (string, int) {
	src, err := file.Open()
	if err != nil {
		return err.Error(), 500
	}
	defer src.Close()
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	fromUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	key := file.Filename
	err = fromUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}

	url := ImgUrl + key
	wg.Done()
	return url, errmsg.SUCCESS
}

//func Uploadfiles(files [])  {
//
//}

func GetImages(prefix, delimiter, marker string, limit int) (imgUrls []map[string]string, code int, err error) {
	mac := qbox.NewMac(utils.AccessKey, utils.SecretKey)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	//img_urls = make(map[string]string)
	bucket := utils.Bucket

	var ts map[string]string

	//初始列举marker为空
	for {
		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			code = errmsg.ERROR
			break
		}
		for _, data := range entries {
			url := "http://blog-img.alexcld.com/" + data.Key
			//组装map
			ts = map[string]string{
				"src": url,
			}
			imgUrls = append(imgUrls, ts)
		}
		if hasNext {
			marker = nextMarker
		} else {
			//list end
			break
		}
	}
	return imgUrls, errmsg.SUCCESS, err
}

func DeleteQNFiles(keys []string) (code int, err error) {
	mac := auth.New(utils.AccessKey, utils.SecretKey)
	cfg := storage.Config{
		Zone: &storage.ZoneHuadong,
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(utils.Bucket, key))
	}

	rets, err := bucketManager.Batch(deleteOps)
	if err != nil {
		if _, ok := err.(*storage.ErrorInfo); ok {
			for _, ret := range rets {
				code = ret.Code
				logger.Info("%s", ret)
				if ret.Code != 200 {
					code = ret.Code
					logger.Error("%s", ret.Data.Error)
				}
			}
		} else {
			logger.Error("batch error,%s", err)
		}
	} else {
		for _, ret := range rets {
			code = ret.Code
			logger.Info("Code: %d  Status: 删除成功", ret.Code)
		}
	}
	return code, err
}
