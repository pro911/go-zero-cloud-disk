package test

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pro911/go-zero-cloud-disk/models"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "/resource/cloud-disk.db")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
}
