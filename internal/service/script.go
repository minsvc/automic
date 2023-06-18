package service

import (
	"automic/pkg/utils"
	"fmt"
	"strconv"
)

type CountScriptRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ScriptListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateScriptRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetScriptRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	//Name       string `form:"name" binding:"min=3,max=100"`
	//State      uint8  `form:"state" binding:"oneof=0 1"`
	//ModifiedBy string `form:"modified_by" binding:"min=3,max=100"`
}

type ExecScriptRequest struct {
	ID       uint32 `form:"id" binding:"required,gte=1"`
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
	Ip       string `form:"ip" binding:"required"`
	Port     string `form:"port" binding:"required"`
}

type DeleteScriptRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CreateScript(versionid string, filename string, desc string, user string, suffix string) error {
	return svc.dao.CreateScript(versionid, filename, desc, user, suffix)
}

type Script struct {
	ID       uint32 `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Version  string `json:"version"`
	Language string `json:"language"`
}

func (svc *Service) GetScript(param *GetScriptRequest) (*Script, error) {

	script, err := svc.dao.GetScript(param.ID)
	if err != nil {
		return nil, err
	}

	return &Script{
		ID:       script.ID,
		Title:    script.Title,
		Desc:     script.Desc,
		Version:  script.Version,
		Language: script.Language,
	}, nil
}

func (svc *Service) ExecScript(param *ExecScriptRequest) (string, error) {

	intVar, _ := strconv.Atoi(param.Port)
	script, _ := svc.GetScript(&GetScriptRequest{ID: param.ID})

	data := utils.OssDownload("mybucket", script.Title, script.Version)
	fmt.Printf("%v", param)
	fmt.Printf(script.Language)
	switch script.Language {
	case "bat":
		result, err := utils.BatCommand(data, param.Ip, intVar, param.User, param.Password)
		return result, err
	default:
		result, err := utils.SshCommand(data, param.Ip, param.Port, param.User, param.Password)
		return result, err
	}

}
