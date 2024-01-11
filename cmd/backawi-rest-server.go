package main

import (
	"encoding/json"
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

	app.handleRoutes()
	app.Run(":8000")

}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProductById).Methods("GET")
	app.Router.HandleFunc("/product/", app.createProduct).Methods("POST", "PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProductById).Methods("DELETE")

}

func sendResponse(rw http.ResponseWriter, statusCode int, payload interface{}) {

	response, _ := json.Marshal(payload)
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(statusCode)
	rw.Write(response)

}

func sendError(rw http.ResponseWriter, statusCode int, err string) {

	error_message := map[string]string{"error": err}
	sendResponse(rw, statusCode, error_message)

}

func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.DBManager.GetAllItemsInDB()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	} else {
		sendResponse(w, http.StatusOK, products)

	}
}

func (app *App) getProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	product, err := app.DBManager.GetProductWithID(vars["id"])
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	} else if len(product) == 0 {
		sendResponse(w, http.StatusNotFound, map[string]string{"Message": "Product not found"})
	} else {
		sendResponse(w, http.StatusOK, product)
	}
}

func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {

	var p db.ProductItem

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request body")
	}

	err = app.DBManager.CreateItemInDB(p)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}
	sendResponse(w, http.StatusOK, map[string]string{"Message": "Item added successfuly"})

}

func (app *App) deleteProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	err := app.DBManager.DeleteProductWithID(vars["id"])
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	} else {
		sendResponse(w, http.StatusOK, map[string]string{"Message": "Item deleted successfuly"})
	}
}
