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

package cost

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	h "github.com/open-falcon/falcon-plus/modules/api/app/helper"
	f "github.com/open-falcon/falcon-plus/modules/api/app/model/falcon_portal"
)

func GetCostList(c *gin.Context) {
	var (
		limit int
		page  int
		err   error
	)
	pageTmp := c.DefaultQuery("page", "")
	limitTmp := c.DefaultQuery("limit", "")
	page, limit, err = h.PageParser(pageTmp, limitTmp)
	if err != nil {
		h.JSONR(c, badstatus, err.Error())
		return
	}
	var dt *gorm.DB
	costs := []f.Cost{}

	var count int
	if err := db.Falcon.Model(&f.Cost{}).Count(&count).Error; err != nil {
		count = 0
	}

	if limit != -1 && page != -1 {
		dt = db.Falcon.Raw(fmt.Sprintf("SELECT * from cost limit %d,%d", page, limit)).Scan(&costs)
	} else {
		dt = db.Falcon.Find(&costs)
	}
	if dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  costs,
		"total": count,
	})
	return
}

func GetCost(c *gin.Context) {
	nidtmp := c.Params.ByName("nid")
	if nidtmp == "" {
		h.JSONR(c, badstatus, "nid is missing")
		return
	}
	nid, err := strconv.Atoi(nidtmp)
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	cost := f.Cost{ID: int64(nid)}
	if dt := db.Falcon.Find(&cost); dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	h.JSONR(c, cost)
	return
}

type APICreateCostInputs struct {
	Reason    string  `json:"reason" binding:"required"`
	Timestamp int64   `json:"timestamp" binding:"required"`
	Money     float64 `json:"money" binding:"required"`
}

func CreateCost(c *gin.Context) {
	var inputs APICreateCostInputs
	if err := c.Bind(&inputs); err != nil {
		h.JSONR(c, badstatus, err)
		return
	}

	mockcfg := f.Cost{
		Reason:    inputs.Reason,
		Timestamp: time.Unix(inputs.Timestamp, 0),
		Money:     inputs.Money,
	}

	if dt := db.Falcon.Save(&mockcfg); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, mockcfg)
	return
}

type APIUpdateCostInputs struct {
	ID        int64   `json:"id" binding:"required"`
	Reason    string  `json:"reason" binding:"required"`
	Timestamp int64   `json:"timestamp" binding:"required"`
	Money     float64 `json:"money" binding:"required"`
}

func UpdateCost(c *gin.Context) {
	var inputs APIUpdateCostInputs
	if err := c.Bind(&inputs); err != nil {
		h.JSONR(c, badstatus, err)
		return
	}

	cost := &f.Cost{ID: inputs.ID}
	ucost := map[string]interface{}{
		"reason":    inputs.Reason,
		"timestamp": time.Unix(inputs.Timestamp, 0),
		"money":     inputs.Money,
	}
	if dt := db.Falcon.Model(&cost).Where("id = ?", inputs.ID).Update(ucost).Find(&cost); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, cost)
	return
}

func DeleteCost(c *gin.Context) {
	nidtmp := c.Params.ByName("nid")
	if nidtmp == "" {
		h.JSONR(c, badstatus, "nid is missing")
		return
	}
	nid, err := strconv.Atoi(nidtmp)
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	cost := f.Cost{ID: int64(nid)}
	if dt := db.Falcon.Delete(&cost); dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	h.JSONR(c, fmt.Sprintf("cost:%d is deleted", nid))
	return
}
