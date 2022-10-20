package models

import "time"

type M struct {
	Id             uint      `gorm:"primaryKey;comment:auto increment id" json:"id"`
	CreatedAt      time.Time `gorm:"comment:create time" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"comment:update time" json:"updatedAt"`
	DeletedAt      time.Time `gorm:"index:idx_deleted_at;comment:soft delete time" json:"deletedAt"`
	ModifiedUserId uint      `gorm:"comment:modified user id" json:"modifiedUserId"`
}
