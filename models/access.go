package models

type Access struct {
	Id          int
	ModuleName  string
	ActionName  string
	Type        int // 节点类型：1-模块 2-菜单 3-操作
	Url         string
	ModuleId    int
	Sort        int
	Description string
	Status      int
	AddTime     int
	AccessItem  []Access `gorm:"foreignKey:ModuleId"`
	Checked     bool     `gorm:"-"` // 数据库忽略这个字段
}

func (Access) TableName() string {
	return "access"
}
