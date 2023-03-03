package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBHOST = "127.0.0.1"
var PORT = "3306"
var DATABASE = "catering"
var USERNAME = "root"
var PASSWORD = ""

var dsn = USERNAME + ":" + PASSWORD + "@tcp(" + DBHOST + ":" + PORT + ")/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
var DataBase, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
