// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-falcon/falcon-plus/modules/api/app/controller/cost"
	"github.com/open-falcon/falcon-plus/modules/api/app/controller/image"
	"github.com/open-falcon/falcon-plus/modules/api/app/controller/user"
	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/export"
	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/upload"
	"github.com/open-falcon/falcon-plus/modules/api/app/utils"
)

func StartGin(port string, r *gin.Engine) {
	r.Use(utils.CORS())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm Falcon+ (｡A｡)")
	})

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.POST("/upload", image.UploadImage)
	r.POST("/uploads", image.UploadImages)

	cost.Routes(r)
	user.Routes(r)
	r.Run(port)
}
