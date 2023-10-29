package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(connection string) {
    db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }
    DB = db
}
