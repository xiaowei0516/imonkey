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

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/open-falcon/falcon-plus/modules/api/app/utils"
	"github.com/open-falcon/falcon-plus/modules/api/config"
)

var db config.DBPool

const badstatus = http.StatusBadRequest
const expecstatus = http.StatusExpectationFailed

func Routes(r *gin.Engine) {
	db = config.Con()
	u := r.Group("/api/v1/user")
	//获取指定用户信息
	u.GET("/:nid", GetUser)
	//创建用户
	u.POST("/", CreateUser)
	//更新用户
	u.PUT("/", UpdateUser)
	//删除用户
	u.DELETE("/:nid", DeleteUser)
	//获取所有用户列表
	u.GET("", GetUserLists)

	//更改客户状态，从潜在客户到当前客户
	u.POST("/updateState", UpdateStatus)

	rep := r.Group("/api/v1/repay")
	//修改偿还记录
	rep.PUT("/", UpdateRepay)
	//获取某个用户的偿还记录
	rep.GET("/user/:uid", GetUserRepay)
	//根据ID获取偿还记录
	rep.GET("/repay/:rid", GetRepay)
}
