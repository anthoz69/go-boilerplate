package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

func NewPostgresDB(v *viper.Viper) *gorm.DB {
	var sslmode string
	if v.GetBool("DB_TLS") {
		sslmode = "require"
	} else {
		sslmode = "disable"
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		v.GetString("db.host"),
		v.GetString("db.user"),
		v.GetString("db.password"),
		v.GetString("db.name"),
		v.GetString("db.port"),
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}
