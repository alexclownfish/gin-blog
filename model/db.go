package model

import (
	"fmt"
	"gin-blog/utils"
	"github.com/wonderivan/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//type Test struct {
//	Category Category `gorm:"foreignkey:Cid"`
//	gorm.Model
//	Title        string `gorm:"type:varchar(100);not null" json:"title"`
//	Cid          int    `gorm:"type:int;not null" json:"cid"`
//	Desc         string `gorm:"type:varchar(200)" json:"desc"`
//	Content      string `gorm:"type:longtext" json:"content"`
//	Img          string `gorm:"type:varchar(100)" json:"img"`
//	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
//	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
//}

var db *gorm.DB

var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPasswd,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		//关闭自动生成复数表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logger.Error("连接数据库失败，请检查参数：", err)
		panic(err)
	}
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取通用数据库对象，使用连接池失败，请检查错误：", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
