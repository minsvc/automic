package model

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

type Script struct {
	*Model
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Version  string `json:"version"`
	Language string `json:"language"`
	State    uint8  `json:"state"`
}

func (a Script) TableName() string {
	return "automic_script"
}
