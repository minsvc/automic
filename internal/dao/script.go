package dao

import (
	"automic/internal/model"
	"automic/pkg/app"
)

func (d *Dao) CountTag(title string, state uint8) (int, error) {
	script := model.Script{Title: title, State: state}
	return script.Count(d.engine)
}

func (d *Dao) GetScriptList(name string, state uint8, page, pageSize int) ([]*model.Script, error) {
	script := model.Script{Title: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return script.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateScript(versionid string, filename string, desc string, user string, suffix string) error {
	script := model.Script{
		Title:    filename,
		State:    1,
		Model:    &model.Model{CreatedBy: user},
		Desc:     desc,
		Version:  versionid,
		Language: suffix,
	}

	return script.Create(d.engine)
}

func (d *Dao) UpdateScript(id uint32, name string, state uint8, modifiedBy string) error {
	script := model.Script{
		Title: name,
		State: state,
		Model: &model.Model{ID: id, ModifiedBy: modifiedBy},
	}

	return script.Update(d.engine)
}

func (d *Dao) DeleteScript(id uint32) error {
	script := model.Script{Model: &model.Model{ID: id}}
	return script.Delete(d.engine)
}

func (d *Dao) GetScript(id uint32) (model.Script, error) {
	script := model.Script{Model: &model.Model{ID: id}}
	return script.Get(d.engine)
}
