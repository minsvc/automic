package global

import (
	"automic/pkg/logger"
	"automic/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	OssSetting      *setting.OssSettingS
	Logger          *logger.Logger
)
