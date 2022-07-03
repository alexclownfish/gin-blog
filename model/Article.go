package model

import (
	"fmt"
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
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	//maps["CommentCount"] = data.CommentCount
	//maps["ReadCount"] = data.ReadCount
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

func (a *articlemethod) GetArticleList(PageSize int, PageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	//err = db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	err = db.Preload("Category").Limit(PageSize).Offset((PageNum - 1) * PageSize).Count(&total).Find(&articleList).Error
	if err != nil {
		fmt.Println(err)
	}
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return nil, errmsg.ERROR, 0
	//}
	return articleList, errmsg.SUCCESS, total
}

////查询单个文章信息
//
//func GetArticleInfo(id int) (Article, int) {
//	var art Article
//	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
//	if err != nil {
//		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
//	}
//	return art, errmsg.SUCCESS
//}

//查询分类下所有文章
func (a *articlemethod) GetCategoryArticle(cid, PageSize, PageNum int) ([]Article, int, int64) {
	var CateArtList []Article
	var total int64
	err := db.Preload("Category").Limit(PageSize).Offset((PageNum-1)*PageSize).Where("cid =?", cid).Find(&CateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return CateArtList, errmsg.SUCCESS, total
}

// GetArtInfo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var art Article
	err = db.Where("id = ?", id).Preload("Category").First(&art).Error
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

//搜索文章标题
func SearchArticle(title string, PageSize, PageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").
		Where("title LIKE ?", title+"%").Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&articleList).Error
	//单独计数
	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}
