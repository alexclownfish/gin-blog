package model

import (
	"context"
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

func Upload(file *multipart.FileHeader) (string, int) {
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
	//var key string
	//var resultUrl string
	//switch kind {
	//case "img":
	//	key = "img/" + key
	//	resultUrl = ImgUrl
	//case "video":
	//	key = "video/" + key
	//	resultUrl = VideoUrl
	//case "soft":
	//	key = "soft/" + key
	//	resultUrl = FileUrl
	//}
	//if kind == "img" {
	//	key = "img/" + key
	//}
	//if kind == "video" {
	//	key = "video/" + key
	//}
	//if kind == "soft" {
	//	key = "soft/" + key
	//}
	key := file.Filename
	err = fromUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}

	url := ImgUrl + key
	return url, errmsg.SUCCESS
}

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
