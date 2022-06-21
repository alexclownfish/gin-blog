package model

import (
	"gin-blog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

var ArticleMethod articlemethod

type articlemethod struct {
}

// 新增文章

func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.ERROR
}

//编辑文章
func EditArticle(id int, data *Article) (code int) {
	var art Article
	var maps = make(map[string]interface{})
	maps["Title"] = data.Title
	maps["Cid"] = data.Cid
	maps["Desc"] = data.Desc
	maps["Content"] = data.Content
	maps["Img"] = data.Img
	maps["CommentCount"] = data.CommentCount
	maps["ReadCount"] = data.ReadCount
	err := db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章

func DeleteArticle(id int) (code int) {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//查询文章列表

func (a *articlemethod) GetArticleList(PageSize int, PageNum int) []Article {
	var art []Article
	err := db.Preload("Category").Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return art
}

//查询单个文章信息

func GetArticleInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

//查询分类下所有文章
func (a *articlemethod) GetCategoryArticle(cid, PageSize, PageNum int) ([]Article, int) {
	var CateArtList []Article
	err := db.Preload("Category").Limit(PageSize).Offset((PageNum-1)*PageSize).Where("cid =?", cid).Find(&CateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return CateArtList, errmsg.SUCCESS
}
