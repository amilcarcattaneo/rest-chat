package check

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Client *gorm.DB
)

func init() {
	var err error
	if Client, err = gorm.Open(
		"sqlite3",
		"challenge.db",
	); err != nil {
		panic(err.Error())
	}
}

func Check() (bool, error) {
	var res int
	err := Client.DB().QueryRow("SELECT 1").Scan(&res)
	if err != nil {
		return false, errors.New("DB connection error")
	}

	if res != 1 {
		return false, errors.New("Unexpected query result")
	}

	return true, nil
}
