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

package falcon_portal

import (
	"time"
)

type User struct {
	Id                 int64     `gorm:"column:id" json:"id"`
	BorrowUser         string    `gorm:"column:borrow_user" json:"borrow_user"`
	BorrowPhone        string    `gorm:"column:borrow_phone" json:"borrow_phone"`
	MateUser           string    `gorm:"column:mate_user" json:"mate_user"`
	MatePhone          string    `gorm:"column:mate_phone" json:"mate_phone"`
	JointBorrowUser1   string    `gorm:"column:joint_borrow_user_1" json:"joint_borrow_user_1"`
	JointBorrowPhone1  string    `gorm:"column:joint_borrow_phone_1" json:"joint_borrow_phone_1"`
	JointBorrowUser2   string    `gorm:"column:joint_borrow_user_2" json:"joint_borrow_user_2"`
	JointBorrowPhone2  string    `gorm:"column:joint_borrow_phone_2" json:"joint_borrow_phone_2"`
	Address            string    `gorm:"column:address" json:"address"`
	Guarantor          string    `gorm:"column:guarantor" json:"guarantor"`
	GuarantorPhone     string    `gorm:"column:guarantor_phone" json:"guarantor_phone"`
	PawnOwner          string    `gorm:"column:pawn_owner" json:"pawn_owner"`
	PawnNumber         string    `gorm:"column:pawn_number" json:"pawn_number"`
	PawnLocation       string    `gorm:"column:pawn_location" json:"pawn_location"`
	PawnTime           time.Time `gorm:"column:pawn_time" json:"pawn_time"`
	PawnArea           float32   `gorm:"column:pawn_area" json:"pawn_area"`
	PawnPawner         string    `gorm:"column:pawn_pawner" json:"pawn_pawner"`
	PawnLoanTotal      float32   `gorm:"column:pawn_loan_total" json:"pawn_loan_total"`
	PawnLoanRemain     float32   `gorm:"column:pawn_loan_remain" json:"pawn_loan_remain"`
	PawnUnitPerice     float32   `gorm:"column:pawn_unit_perice" json:"pawn_unit_perice"`
	PawnRemain         float32   `gorm:"column:pawn_remain" json:"pawn_remain"`
	PawnProperty       string    `gorm:"column:pawn_property" json:"pawn_property"`
	PawnGuarantorTotal float32   `gorm:"column:pawn_guarantor_total" json:"pawn_guarantor_total"`
	LoanPrincipal      float32   `gorm:"column:loan_principal" json:"loan_principal"`
	LoanPeriods        int       `gorm:"column:loan_periods" json:"loan_periods"`
	LoanRate           float32   `gorm:"column:loan_rate" json:"loan_rate"`
	LoanStillPrincipal float32   `gorm:"column:loan_still_principal" json:"loan_still_principal"`
	LoanStillRate      float32   `gorm:"column:loan_still_rate" json:"loan_still_rate"`
	LoanReason         string    `gorm:"column:loan_reason" json:"loan_reason"`
	LoanReturnEvaluate string    `gorm:"column:loan_return_evaluate" json:"loan_return_evaluate"`
	ImageContract      string    `gorm:"column:image_contract" json:"image_contract"`
	ImageHouse         string    `gorm:"column:image_house" json:"image_house"`
	ImageCredit        string    `gorm:"column:image_credit" json:"image_credit"`
	ImageOther         string    `gorm:"column:image_other" json:"image_other"`
	Other1             string    `gorm:"column:other_1" json:"other_1"`
	Other2             string    `gorm:"column:other_2" json:"other_2"`
	State              int       `gorm:"column:state" json:"state"`
}

func (this User) TableName() string {
	return "user"
}
