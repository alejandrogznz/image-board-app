
package main

import (
	"os"
	"log/slog"
	//"database/sql"
	"net/http"
//	"encoding/json"
	_ "github.com/go-sql-driver/mysql"	
	"github.com/joho/godotenv"
	//"image-board/model"
//	"image-board/handlers"
	"image-board/middleware"
	"image-board/sqldb"
	"image-board/controllers"

)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	genv = godotenv.Load()
)



func main() {
	logger.Info("Hello Server")		
	logger.Info("Opening database connection...")
	
	// Open the connection to the database
	db := sqldb.ConnectDB()
	h := controllers.NewBaseHandler(db)
	logger.Info("Successfully connected to database")
	// Open the server
	router := http.NewServeMux()
	router.HandleFunc("GET /", h.HelloWorld)
	server := http.Server{
		Addr:	 ":8000",
		Handler: middleware.Logging(router),
	}

	logger.Error("Server Closed", server.ListenAndServe())
}
