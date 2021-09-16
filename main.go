package main

import (
	_ "database/sql"
	"fmt"
	"os"

	"github.com/BoucherR/aba_server/db"
	"github.com/BoucherR/aba_server/router"

	_ "github.com/lib/pq"
)

func init() {
	db.Connect()
}

func main() {
	r := router.SetupRouter()
	// Listen and Serve in 0.0.0.0:8081
	// r.Run(":8081")
	fmt.Printf("STARTING ON PORT: %v", os.Getenv("PORT"))
	r.Run(os.Getenv("PORT"))
}
