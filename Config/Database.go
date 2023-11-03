package Config

import (
    "fmt"
    "os"
    _"time"
    "gorm.io/gorm"
)

var DB *gorm.DB

func DbURL() string {
    USER        := os.Getenv("DB_USER")
    PASSWORD    := os.Getenv("DB_PASSWORD")
    HOST        := os.Getenv("DB_HOST")
    DBNAME      := os.Getenv("DB_NAME")
    DBPORT      := os.Getenv("DB_PORT")

    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD, HOST, DBPORT, DBNAME)
}