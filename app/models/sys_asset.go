package models

import (
	"time"
)

type SysAsset struct {
	M
	AssetId      string    `gorm:"index:idx_aid,unique;comment:asset id" json:"assetId"`
	Name         string    `gorm:"comment:asset name" json:"name"`
	Type         string    `gorm:"comment:asset type" json:"type"`
	ChildType    string    `gorm:"comment:asset child type" json:"childType"`
	ProjectName  string    `gorm:"comment:asset assign project name" json:"projectName"`
	Origin       string    `gorm:"comment:origin" json:"origin"`
	Model        string    `gorm:"comment:model" json:"model"`
	Value        float64   `gorm:"comment:asset value" json:"value"`
	Unit         string    `gorm:"comment:unit" json:"unit"`
	InDate       time.Time `gorm:"comment:put in date" json:"inDate"`
	OpDate       time.Time `gorm:"comment:op date" json:"opDate"`
	DepYear      int       `gorm:"comment:dep year" json:"depYear"`
	Location     string    `gorm:"comment:asset location info" json:"location"`
	Status       *uint     `gorm:"type:tinyint(1);default:0;comment:asset status" json:"status"`
	Image        string    `gorm:"comment:image url" json:"image"`
	AssignUserId uint      `gorm:"comment:assign user id" json:"assignUserId"`
}
