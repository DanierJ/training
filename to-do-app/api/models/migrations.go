package models

func AutoMigrations() {
	db := Connect()
	defer db.Close()
	db.Debug().DropTableIfExists(&Todo{})
	db.Debug().AutoMigrate(&Todo{})
}
