package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/file"
	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/util"

	"github.com/spf13/viper"
)

//获取图片的完整访问路径
func GetImageFullUrl(name string) string {
	return viper.GetString("prefix_url") + "/" + GetImagePath() + name
}

//获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//获取图片路径
func GetImagePath() string {
	str := viper.GetString("image.image_save_path")
	log.Println("getImagePath ", str)
	return str
}

func GetImageFullPath() string {
	return viper.GetString("image.runtime_root_path") + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range strings.Split(viper.GetString("image.image_allow_exts"), ",") {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}

	return size <= viper.GetInt("image.image_max_size")*1024*1024
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
