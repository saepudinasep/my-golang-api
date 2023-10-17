package main

import (
	"github.com/saepudinasep/my-golang-api/database"
	"github.com/saepudinasep/my-golang-api/router"
)

func main() {
	database.InitDB()
	database.Migrate() // Jalankan migrasi saat aplikasi dimulai
	r := router.SetupRouter()
	r.Run(":8080")
}
