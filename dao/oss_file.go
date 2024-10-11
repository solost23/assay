package dao

import "gorm.io/gorm"

type OSSFile struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name;type:varchar(500);comment: 文件名称"`
	Path      string `json:"path" gorm:"column:path;type:varchar(500);comment: 文件路径"`
	Type      string `json:"type" gorm:"column:type;type:varchar(20);comment: 文件类型"`
	Size      int64  `json:"size" gorm:"column:size;type:bigint unsigned;comment: 文件大小"`
	CreatorId uint   `json:"creatorId" gorm:"column:creator_id;type:bigint unsigned;comment: 创建人"`
}
