package v1

import (
	"automic/global"
	"automic/internal/service"
	"automic/pkg/app"
	"automic/pkg/convert"
	"automic/pkg/errcode"
	"automic/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type Script struct{}

func NewScript() Script {
	return Script{}
}

// @Summary 获取单个脚本
// @Produce  json
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Script "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/scripts/{id} [get]
func (s Script) Get(c *gin.Context) {

	param := service.GetScriptRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	script, err := svc.GetScript(&param)

	data := utils.OssDownload("mybucket", script.Title, script.Version)
	if err != nil {
		global.Logger.Errorf(c, "Get script err: %v", err)
		response.ToErrorResponse(errcode.ErrorDownloadFileFail.WithDetails(err.Error()))
	}

	response.ToResponse(gin.H{
		"info": data,
	})
	return
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Script "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/scripts [get]
func (s Script) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Script "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/scripts [post]
func (s Script) Create(c *gin.Context) {

	response := app.NewResponse(c)

	file, script, err := c.Request.FormFile("script")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	suffix := strings.Split(script.Filename, ".")[1]

	versionid, err := utils.OssUpload(file, script.Filename, script.Size, suffix)
	if err != nil {
		global.Logger.Errorf(c, "Upload script err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
	}

	svc := service.New(c.Request.Context())

	//models.Add(versionid, script.Filename, "test", claims.Username, suffix)
	err = svc.CreateScript(versionid, script.Filename, "test", "user1", suffix)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{
		"info": "upload success",
	})
	return
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Script "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/scripts/{id} [put]
func (s Script) Update(c *gin.Context) {}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/scripts/{id} [delete]
func (s Script) Delete(c *gin.Context) {}

func (s Script) Exec(c *gin.Context) {
	param := service.ExecScriptRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	result, err := svc.ExecScript(&param)

	if err != nil {

		response.ToErrorResponse(errcode.ErrorSSHConnectFail.WithDetails(err.Error()))
	}

	response.ToResponse(gin.H{
		"info":  result,
		"error": err,
	})
	return
}
