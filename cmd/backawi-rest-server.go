package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/AbdouTlili/backawi/pkg/config"
	"github.com/AbdouTlili/backawi/pkg/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type App struct {
	Config    *aws.Config
	DBManager *db.DBManager
	Router    *mux.Router
}

func (app *App) Init() {

	cfg := config.LoadConfig()

	log.Info("Config loaded form Env \nTable name is : ", cfg.DynamoDBTableName, " \nAWS region is : ", cfg.AWSRegion, "\nSecurity Keys loaded")

	// app := App{Config: &aws.Config{
	// 	Region:      &cfg.AWSRegion,
	// 	Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	// }, DBManager: &db.DBManager{}}

	app.Config = &aws.Config{
		Region:      &cfg.AWSRegion,
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	}

	app.DBManager = &db.DBManager{}

	app.DBManager.Init(app.Config, cfg.DynamoDBTableName)

	app.Router = mux.NewRouter().StrictSlash(true)

}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func main() {

	app := App{}

	app.Init()

	app.DBManager.GetAllItemsInDB()
	app.DBManager.GetProductWithID("5585")
	app.DBManager.DeleteProductWithID("5585")
	app.DBManager.GetAllItemsInDB()

	// app.Run(":8000")

}

// func (app *App) handleRoutes(){
// 	app.Router.HandleFunc("/products",getProducts).Methods("GET")

// }
