package repository_test

import (
	"testing"

	"github.com/Edwardz43/mygame/app/db/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestAutiMigration(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/MyGame?parseTime=true")
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=15432 user=admin dbname=postgres password=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.BetRecord{})

	assert.Empty(t, db.Error)
}