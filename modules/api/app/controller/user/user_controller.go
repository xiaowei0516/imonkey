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
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	h "github.com/open-falcon/falcon-plus/modules/api/app/helper"
	f "github.com/open-falcon/falcon-plus/modules/api/app/model/falcon_portal"
	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/app"
	"github.com/open-falcon/falcon-plus/modules/api/app/pkg/e"
)

func GetUserList(c *gin.Context) {
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

func GetUser(c *gin.Context) {
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
	user := f.User{Id: int64(nid)}
	if dt := db.Falcon.Find(&user); dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	h.JSONR(c, user)
	return
}

type APICreateUserInputs struct {
	BorrowUser         string  `json:"borrow_user" binding:"required"`
	BorrowPhone        string  `json:"borrow_phone" binding:"required"`
	MateUser           string  `json:"mate_user"`
	MatePhone          string  `json:"mate_phone"`
	JointBorrowUser1   string  `json:"joint_borrow_user_1"`
	JointBorrowPhone1  string  `json:"joint_borrow_phone_1"`
	JointBorrowUser2   string  `json:"joint_borrow_user_2"`
	JointBorrowPhone2  string  `json:"joint_borrow_phone_2"`
	Guarantor          string  `json:"guarantor"`
	GuarantorPhone     string  `json:"guarantor_phone"`
	Address            string  `json:"address"`
	PawnOwner          string  `json:"pawn_owner"`
	PawnNumber         string  `json:"pawn_number"`
	PawnLocation       string  `json:"pawn_location"`
	PawnTime           int64   `json:"pawn_time"`
	PawnArea           float32 `json:"pawn_area"`
	PawnPawner         string  `json:"pawn_pawner"`
	PawnLoanTotal      float32 `json:"pawn_loan_total"`
	PawnLoanRemain     float32 `json:"pawn_loan_remain"`
	PawnUnitPerice     float32 `json:"pawn_unit_perice"`
	PawnRemain         float32 `json:"pawn_remain"`
	PawnProperty       string  `json:"pawn_property"`
	PawnGuarantorTotal float32 `json:"pawn_guarantor_total"`
	LoanPrincipal      float32 `json:"loan_principal"`
	LoanPeriods        int     `json:"loan_periods"`
	LoanRate           float32 `json:"loan_rate"`
	LoanStillPrincipal float32 `json:"loan_still_principal"`
	LoanStillRate      float32 `json:"loan_still_rate"`
	LoanReason         string  `json:"loan_reason"`
	LoanReturnEvaluate string  `json:"loan_return_evaluate"`
	ImageContract      string  `json:"image_contract"`
	ImageHouse         string  `json:"image_house"`
	ImageCredit        string  `json:"image_credit"`
	ImageOther         string  `json:"image_other"`
	Other1             string  `json:"other_1"`
	Other2             string  `json:"other_2"`
	State              int     `json:"state"`
}

func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}

	var inputs APICreateUserInputs
	if err := c.Bind(&inputs); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error())
		return
	}

	user := f.User{
		BorrowUser:         inputs.BorrowUser,
		BorrowPhone:        inputs.BorrowPhone,
		MateUser:           inputs.MateUser,
		MatePhone:          inputs.MatePhone,
		JointBorrowUser1:   inputs.JointBorrowUser1,
		JointBorrowPhone1:  inputs.JointBorrowPhone1,
		JointBorrowUser2:   inputs.JointBorrowUser2,
		JointBorrowPhone2:  inputs.JointBorrowPhone2,
		Guarantor:          inputs.Guarantor,
		GuarantorPhone:     inputs.GuarantorPhone,
		Address:            inputs.Address,
		PawnOwner:          inputs.PawnOwner,
		PawnNumber:         inputs.PawnNumber,
		PawnLocation:       inputs.PawnLocation,
		PawnTime:           time.Unix(inputs.PawnTime, 0),
		PawnArea:           inputs.PawnArea,
		PawnPawner:         inputs.PawnPawner,
		PawnLoanTotal:      inputs.PawnLoanTotal,
		PawnLoanRemain:     inputs.PawnLoanRemain,
		PawnUnitPerice:     inputs.PawnUnitPerice,
		PawnRemain:         inputs.PawnRemain,
		PawnProperty:       inputs.PawnProperty,
		PawnGuarantorTotal: inputs.PawnGuarantorTotal,
		LoanPrincipal:      inputs.LoanPrincipal,
		LoanPeriods:        inputs.LoanPeriods,
		LoanRate:           inputs.LoanRate,
		LoanStillPrincipal: inputs.LoanStillPrincipal,
		LoanStillRate:      inputs.LoanStillRate,
		LoanReason:         inputs.LoanReason,
		LoanReturnEvaluate: inputs.LoanReturnEvaluate,
		ImageContract:      inputs.ImageContract,
		ImageHouse:         inputs.ImageHouse,
		ImageCredit:        inputs.ImageCredit,
		ImageOther:         inputs.ImageOther,
		Other1:             inputs.Other1,
		Other2:             inputs.Other2,
		State:              2,
	}

	if dt := db.Falcon.Save(&user); dt.Error != nil {
		appG.Response(http.StatusOK, e.ERROR_ADD_USER_FAIL, dt.Error.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
	return
}

type APIUpdateUserInputs struct {
	ID                 string  `json:"id" binding:"required"`
	BorrowUser         string  `json:"borrow_user"`
	BorrowPhone        string  `json:"borrow_phone"`
	MateUser           string  `json:"mate_user"`
	MatePhone          string  `json:"mate_phone"`
	JointBorrowUser1   string  `json:"joint_borrow_user_1"`
	JointBorrowPhone1  string  `json:"joint_borrow_phone_1"`
	JointBorrowUser2   string  `json:"joint_borrow_user_2"`
	JointBorrowPhone2  string  `json:"joint_borrow_phone_2"`
	Guarantor          string  `json:"guarantor"`
	GuarantorPhone     string  `json:"guarantor_phone"`
	Address            string  `json:"address"`
	PawnOwner          string  `json:"pawn_owner"`
	PawnNumber         string  `json:"pawn_number"`
	PawnLocation       string  `json:"pawn_location"`
	PawnTime           int64   `json:"pawn_time"`
	PawnArea           float32 `json:"pawn_area"`
	PawnPawner         string  `json:"pawn_pawner"`
	PawnLoanTotal      float32 `json:"pawn_loan_total"`
	PawnLoanRemain     float32 `json:"pawn_loan_remain"`
	PawnUnitPerice     float32 `json:"pawn_unit_perice"`
	PawnRemain         float32 `json:"pawn_remain"`
	PawnProperty       string  `json:"pawn_property"`
	PawnGuarantorTotal float32 `json:"pawn_guarantor_total"`
	LoanPrincipal      float32 `json:"loan_principal"`
	LoanPeriods        int     `json:"loan_periods"`
	LoanRate           float32 `json:"loan_rate"`
	LoanStillPrincipal float32 `json:"loan_still_principal"`
	LoanStillRate      float32 `json:"loan_still_rate"`
	LoanReason         string  `json:"loan_reason"`
	LoanReturnEvaluate string  `json:"loan_return_evaluate"`
	ImageContract      string  `json:"image_contract"`
	ImageHouse         string  `json:"image_house"`
	ImageCredit        string  `json:"image_credit"`
	ImageOther         string  `json:"image_other"`
	Other1             string  `json:"other_1"`
	Other2             string  `json:"other_2"`
	State              int     `json:"state"`
}

func UpdateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var inputs APIUpdateUserInputs
	if err := c.Bind(&inputs); err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error())
		return
	}

	var user f.User

	uuser := f.User{
		BorrowUser:         inputs.BorrowUser,
		BorrowPhone:        inputs.BorrowPhone,
		MateUser:           inputs.MateUser,
		MatePhone:          inputs.MatePhone,
		JointBorrowUser1:   inputs.JointBorrowUser1,
		JointBorrowPhone1:  inputs.JointBorrowPhone1,
		JointBorrowUser2:   inputs.JointBorrowUser2,
		JointBorrowPhone2:  inputs.JointBorrowPhone2,
		Guarantor:          inputs.Guarantor,
		GuarantorPhone:     inputs.GuarantorPhone,
		Address:            inputs.Address,
		PawnOwner:          inputs.PawnOwner,
		PawnNumber:         inputs.PawnNumber,
		PawnLocation:       inputs.PawnLocation,
		PawnTime:           time.Unix(inputs.PawnTime, 0),
		PawnArea:           inputs.PawnArea,
		PawnPawner:         inputs.PawnPawner,
		PawnLoanTotal:      inputs.PawnLoanTotal,
		PawnLoanRemain:     inputs.PawnLoanRemain,
		PawnUnitPerice:     inputs.PawnUnitPerice,
		PawnRemain:         inputs.PawnRemain,
		PawnProperty:       inputs.PawnProperty,
		PawnGuarantorTotal: inputs.PawnGuarantorTotal,
		LoanPrincipal:      inputs.LoanPrincipal,
		LoanPeriods:        inputs.LoanPeriods,
		LoanRate:           inputs.LoanRate,
		LoanStillPrincipal: inputs.LoanStillPrincipal,
		LoanStillRate:      inputs.LoanStillRate,
		LoanReason:         inputs.LoanReason,
		LoanReturnEvaluate: inputs.LoanReturnEvaluate,
		ImageContract:      inputs.ImageContract,
		ImageHouse:         inputs.ImageHouse,
		ImageCredit:        inputs.ImageCredit,
		ImageOther:         inputs.ImageOther,
		Other1:             inputs.Other1,
		Other2:             inputs.Other2,
		State:              2,
	}
	if dt := db.Falcon.Where("id = ?", inputs.ID).Update(uuser).Find(&user); dt.Error != nil {
		appG.Response(http.StatusOK, e.ERROR_EDIT_USER_FAIL, dt.Error.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
	return
}

func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	nidtmp := c.Params.ByName("nid")
	if nidtmp == "" {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, "nid is missing")
		return
	}
	nid, err := strconv.Atoi(nidtmp)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error())
		return
	}
	user := f.User{Id: int64(nid)}
	if dt := db.Falcon.Delete(&user); dt.Error != nil {
		appG.Response(http.StatusOK, e.ERROR_DELETE_USER_FAIL, dt.Error.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
	return
}

func GetUserLists(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		limit int
		page  int
		err   error
	)
	pageTmp := c.DefaultQuery("page", "")
	limitTmp := c.DefaultQuery("limit", "")
	page, limit, err = h.PageParser(pageTmp, limitTmp)
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error())
		return
	}
	var dt *gorm.DB
	users := []f.User{}

	var count int
	if err := db.Falcon.Model(&f.User{}).Count(&count).Error; err != nil {
		count = 0
	}

	if limit != -1 && page != -1 {
		dt = db.Falcon.Raw(fmt.Sprintf("SELECT * from cost limit %d,%d", page, limit)).Scan(&users)
	} else {
		dt = db.Falcon.Find(&users)
	}
	if dt.Error != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_USER_FAIL, dt.Error.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": users,
		"total": count,
	})
	return
}
