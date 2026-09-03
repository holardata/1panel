// 此文件在GPL-3.0协议下开源
// 修改者：bobwu 2026-09-01
package model

type AppDetail struct {
	BaseModel
	AppId               uint   `json:"appId" gorm:"type:integer;not null"`
	Version             string `json:"version" gorm:"type:varchar(64);not null"`
	Params              string `json:"-" gorm:"type:longtext;"`
	DockerCompose       string `json:"dockerCompose"  gorm:"type:longtext;"`
	Status              string `json:"status" gorm:"type:varchar(64);not null"`
	LastVersion         string `json:"lastVersion" gorm:"type:varchar(64);"`
	LastModified        int    `json:"lastModified" gorm:"type:integer;"`
	DownloadUrl         string `json:"downloadUrl"  gorm:"type:varchar;"`
	DownloadCallBackUrl string `json:"downloadCallBackUrl" gorm:"type:longtext;"`
	Update              bool   `json:"update"`
	IgnoreUpgrade       bool   `json:"ignoreUpgrade"`
	ChangeLog           string `json:"changeLog" gorm:"type:longtext;default:''"`
}
