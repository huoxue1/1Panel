package dto

import (
	"encoding/json"
	"github.com/1Panel-dev/1Panel/app/model"
)

type AppRes struct {
	Version   string      `json:"version"`
	CanUpdate bool        `json:"canUpdate"`
	Items     []*AppDTO   `json:"items"`
	Tags      []model.Tag `json:"tags"`
	Total     int64       `json:"total"`
}

type AppDTO struct {
	model.App
	Versions []string    `json:"versions"`
	Tags     []model.Tag `json:"tags"`
}

type AppDetailDTO struct {
	model.AppDetail
	Params interface{} `json:"params"`
}

type AppList struct {
	Version string      `json:"version"`
	Tags    []Tag       `json:"tags"`
	Items   []AppDefine `json:"items"`
}

type AppDefine struct {
	Key                string   `json:"key"`
	Name               string   `json:"name"`
	Tags               []string `json:"tags"`
	Versions           []string `json:"versions"`
	Icon               string   `json:"icon"`
	Author             string   `json:"author"`
	Source             string   `json:"source"`
	ShortDesc          string   `json:"short_desc"`
	Type               string   `json:"type"`
	Required           []string `json:"Required"`
	CrossVersionUpdate bool     `json:"crossVersionUpdate"`
}

func (define AppDefine) GetRequired() string {
	by, _ := json.Marshal(define.Required)
	return string(by)
}

type Tag struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type AppForm struct {
	FormFields []AppFormFields `json:"form_fields"`
}

type AppFormFields struct {
	Type     string `json:"type"`
	LabelZh  string `json:"label_zh"`
	LabelEn  string `json:"label_en"`
	Required string `json:"required"`
	Default  string `json:"default"`
	EnvKey   string `json:"env_variable"`
}

type AppRequest struct {
	PageInfo
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type AppInstallRequest struct {
	AppDetailId uint                   `json:"appDetailId" validate:"required"`
	Params      map[string]interface{} `json:"params"`
	Name        string                 `json:"name" validate:"required"`
	Services    map[string]string      `json:"services"`
}

type AppInstalled struct {
	model.AppInstall
	Total   int    `json:"total"`
	Ready   int    `json:"ready"`
	AppName string `json:"appName"`
	Icon    string `json:"icon"`
}

type AppInstalledRequest struct {
	PageInfo
}

type AppOperate string

var (
	Up      AppOperate = "up"
	Down    AppOperate = "down"
	Restart AppOperate = "restart"
	Delete  AppOperate = "delete"
	Sync    AppOperate = "sync"
)

type AppInstallOperate struct {
	InstallId uint       `json:"installId" validate:"required"`
	Operate   AppOperate `json:"operate" validate:"required"`
}

type AppService struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type AppDatabase struct {
	ServiceName string `json:"PANEL_DB_HOST"`
	DbName      string `json:"PANEL_DB_NAME"`
	DbUser      string `json:"PANEL_DB_USER"`
	Password    string `json:"PANEL_DB_USER_PASSWORD"`
}

type AuthParam struct {
	RootPassword string `json:"PANEL_DB_ROOT_PASSWORD"`
}

type ContainerExec struct {
	ContainerName string      `json:"containerName"`
	DbParam       AppDatabase `json:"dbParam"`
	Auth          AuthParam   `json:"auth"`
}