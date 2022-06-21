package model

import (
	"gin-blog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

var CategoryMethod categorymethod

type categorymethod struct {
}

//查询分类是否存在

func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCESS
}

// 新增分类

func CreateCategory(data *Category) int {
	//data.Password = ScryptPassWord(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.ERROR
}

//编辑分类
func EditCategory(id int, data *Category) (code int) {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类

func DeleteCategory(id int) (code int) {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//查询分类列表

func (c *categorymethod) GetCategoryList(PageSize int, PageNum int) []Category {
	var cate []Category
	err = db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

//查询分类下的所有文章
