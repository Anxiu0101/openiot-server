package model

import (
	"gorm.io/gorm"
)

type (
	// Notification 警告信息
	Notification struct {
		gorm.Model
		ObDevice  uint64
		Name      string
		Desc      string
		Content   string
		Issuer    uint64
		Level     uint // 警告等级，0-Normal, 1-Warming, 2-Dangerous
		Handled   bool // 是否处理，false 未处理，true 已处理
		CloseRole uint64
	}

	// Noticeboard 常态公告模块
	Noticeboard struct {
	}

	NoticeCloser interface {
	}
)
