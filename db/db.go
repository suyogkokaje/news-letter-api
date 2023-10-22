package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    DB  *gorm.DB
)

func Initialize(connection string) {
    db, err := gorm.Open("postgres", connection)
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }
    DB = db
}
