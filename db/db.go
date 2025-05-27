package db

import (
	"fmt"
	"trinity/includes/Log"
	"trinity/includes/cfg"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.RowAccessString()), &gorm.Config{})
	if err != nil {
		Log.WriteLog(err)
	}
	return db
}

func Add(ValuesString []string, tablename string) error {
	var db *gorm.DB = Db()

	if len(ValuesString) == 0 {
		Log.WriteLog("values required")
	}

	data := make(map[string]interface{})
	for i, value := range ValuesString {
		data[fmt.Sprintf("field%d", i+1)] = value
	}

	return db.Table(tablename).Create(data).Error
}

func Update(id string, valueToUpdate string, Updatestring string, tablename string) error {
	var db *gorm.DB = Db()

	return db.Table(tablename).
		Where("id = ?", id).
		Update(valueToUpdate, Updatestring).Error
}

func DeleteRow(id string, tablename string) error {
	var db *gorm.DB = Db()

	return db.Table(tablename).
		Where("id = ?", id).
		Delete(nil).Error
}

func GetRow(id string, tablename string) *gorm.DB {
	var db *gorm.DB = Db()

	return db.Table(tablename).
		Where("id = ?", id).
		Limit(1)
}

func GetRowAll(tablename string) *gorm.DB {
	var db *gorm.DB = Db()

	return db.Table(tablename)
}

func UserExists(nickname, tablename string) (bool, error) {
	var db *gorm.DB = Db()

	var count int64
	err := db.Table(tablename).
		Where("nickname = ?", nickname).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func PasswordMatches(nickname, password, tablename string) (bool, error) {
	var db *gorm.DB = Db()

	var storedPassword string

	err := db.Table(tablename).
		Select("password").
		Where("nickname = ?", nickname).
		Scan(&storedPassword).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return storedPassword == password, nil
}
