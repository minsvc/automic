package model

import "github.com/jinzhu/gorm"

func (t Script) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Title != "" {
		db = db.Where("name = ?", t.Title)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Script) List(db *gorm.DB, pageOffset, pageSize int) ([]*Script, error) {
	var scripts []*Script
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Title != "" {
		db = db.Where("name = ?", t.Title)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&scripts).Error; err != nil {
		return nil, err
	}

	return scripts, nil
}

func (t Script) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Script) Update(db *gorm.DB) error {
	return db.Model(&Script{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(t).Error
}

func (t Script) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

func (t Script) Get(db *gorm.DB) (Script, error) {
	var script Script
	db = db.Where("id = ?", t.ID)
	err := db.First(&script).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return script, err
	}

	return script, nil
}
