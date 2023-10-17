package database

import (
	"github.com/saepudinasep/my-golang-api/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})

	// Definisikan relasi antar tabel di sini jika diperlukan
	// DB.Model(&models.Photo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// DB.Model(&models.Photo{}).
}

func Rollback() {
	// Implementasikan rollback migrasi jika diperlukan
}
