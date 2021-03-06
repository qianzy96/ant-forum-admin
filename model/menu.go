package model

import (
	"ant-forum/pkg/constvar"
)

type MenuModel struct {
	BaseModel
	Name   string `form:"name" json:"name" xml:"name" gorm:"column:name;not null" binding:"required"`
	Path   string `form:"path" json:"path" xml:"path" gorm:"column:path;not null" binding:"required"`
	Method string `form:"method" json:"method" xml:"method" gorm:"column:method;not null" binding:"required"`
}

func (menu *MenuModel) TableName() string {
	return "menu"
}

type MenuInfo struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

// 创建新菜单
func (menu *MenuModel) Create() error {
	return DB.Self.Create(&menu).Error
}

// 获取菜单列表
func (menu *MenuModel) ListMenu(offset, limit int) ([]*MenuModel, uint64, error) {
	t := MenuModel{}
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	list := make([]*MenuModel, 0)
	var count uint64
	if err := DB.Self.Model(&t).Count(&count).Error; err != nil {
		return list, count, err
	}
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&list).Error; err != nil {
		return list, count, err
	}

	return list, count, nil
}

// 根据标签id获取菜单数据.
func (menu *MenuModel) GetMenuById(id uint64) (*MenuModel, error) {
	d := DB.Self.First(&menu, id)
	return menu, d.Error
}

// 根据标签id删除菜单
func (menu *MenuModel) DeleteMenu(id uint64) error {
	menu.BaseModel.Id = id
	return DB.Self.Delete(&menu).Error
}

// 更新菜单
func (menu *MenuModel) Update() error {
	return DB.Self.Omit("created_at").Save(menu).Error
}
