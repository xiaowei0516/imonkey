package image

import (
	"github.com/gin-gonic/gin"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/app"
	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/e"
	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/upload"
)

// @Summary 上传单个图片
// @Produce  json
// @Param image post file true "图片文件"
// @Success 200 {string} json "{"code":200,"data":{"image_save_url":"upload/images/96a.jpg", "image_url": "http://..."}"
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		log.Printf("%+v", err.Error())
		log.Warn(err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		log.Warn(err)
		appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		log.Warn(err)
		appG.Response(http.StatusOK, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}

// @Summary 上传多个图片
// @Produce  json
// @Param image post file true "图片文件"
// @Success 200 {string} json "{"code":200,"data":{"image_save_url":"upload/images/96a.jpg", "image_url": "http://..."}"
// @Router /api/v1/tags/import [post]
func UploadImages(c *gin.Context) {
	appG := app.Gin{C: c}
	form, err := c.MultipartForm()
	if err != nil {
		log.Warn(err)
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	imageUrls := make([]string, 0)
	images := make([]string, 0)
	files := form.File["images"]
	for _, file := range files {
		imageName := upload.GetImageName(file.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + imageName
		log.Debugf("src :%s\n", src)

		if !upload.CheckImageExt(imageName) {
			appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
			return
		}

		err = upload.CheckImage(fullPath)
		if err != nil {
			log.Warn(err)
			appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
			return
		}
		if err := c.SaveUploadedFile(file, src); err != nil {
			log.Warn(err)
		} else {
			imageUrls = append(imageUrls, savePath+imageName)
			images = append(images, upload.GetImageFullUrl(imageName))
		}
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"image_url":      images,
		"image_save_url": imageUrls,
	})

}
