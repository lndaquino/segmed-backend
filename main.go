package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/lndaquino/segmed-backend/api/controllers"
	"github.com/lndaquino/segmed-backend/api/routers"
	"github.com/lndaquino/segmed-backend/datastore"
	"github.com/lndaquino/segmed-backend/datastore/repos"
	"github.com/lndaquino/segmed-backend/pkg/domain/file"
)

func init() {
	godotenv.Load()
}

func main() {
	session, err := setupPostgres()
	if err != nil {
		panic("Error starting database:" + err.Error())
	}
	defer session.Close()
	log.Println("==> Sucessfully connected to database")

	app, err := setupApplication(session)
	if err != nil {
		panic("Error starting application:" + err.Error())
	}
	e := app.MakeControllers()

	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// e.Use(middleware.Logger())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
	// }))

	e.Logger.Fatal(e.StartServer(server))
}

func setupPostgres() (*gorm.DB, error) {
	db, err := datastore.NewPostgres()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func setupApplication(db *gorm.DB) (*routers.SystemRoutes, error) {
	fileRepo := repos.NewFileRepo(db)
	fileUsecase := file.NewFileUsecase(fileRepo)
	fileController := controllers.NewFileController(fileUsecase)
	systemRoutes := routers.NewSystemRoutes(fileController)

	return systemRoutes, nil
}
