package models

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"xorm.io/xorm"
)

//设置一个外部变量用来调用init方法
var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", "resource/cloud-disk.db")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}
