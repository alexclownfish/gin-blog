package model

import (
	"encoding/base64"

	"gin-blog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
	//Avatar   string
}

var UserMethod usermethod

type usermethod struct {
}

//查询用户是否存在

func (u *usermethod) CheckUser(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户

func CreateUser(data *User) (code int) {
	//data.Password = ScryptPassWord(data.Password)

	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.ERROR
}

//编辑用户
func EditUser(id int, data *User) (code int) {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户

func (u *usermethod) DeleteUser(id int) (code int) {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//密码加密 狗子
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPassWord(u.Password)
	return
}

//查询用户列表

func (u *usermethod) GetUserList(PageSize int, PageNum int) []User {
	var users []User
	err = db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

func ScryptPassWord(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPassWord, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPassWord)
	return fpw
}

//登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPassWord(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_PERMISSION
	}
	return errmsg.SUCCESS
}
