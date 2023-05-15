package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	//postgres://uuuwxbte:ydMbQltQR9VRY5vSTmzdTm2EoTTmQRmI@arjuna.db.elephantsql.com/uuuwxbte
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falied to connect to DB")
	}

}
