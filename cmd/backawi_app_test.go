package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/AbdouTlili/backawi/pkg/db"
)

var a App

func TestMain(m *testing.M) {

	a = App{}

	a.Init()

	a.handleRoutes()
	// a.Run(":8000")

	m.Run()

}

func TestGetProducts(t *testing.T) {

	var p db.ProductItem
	p.Discount = 55
	p.Name = "Test Product"
	p.ProductID = "pid"
	p.Price = 2.55
	p.Quantity = 55
	err := a.DBManager.CreateItemInDB(p)

	if err != nil {
		t.Failed()
	}

	request, _ := http.NewRequest("GET", "/product/pid", nil)
	response := sendRequest(request)
	log.Info(request.Host)
	checkStatusCode(t, http.StatusOK, response.Code)

}
func checkStatusCode(t *testing.T, expectedStatusCode, receivedStatusCode int) {
	if expectedStatusCode != receivedStatusCode {
		t.Errorf("Expected Status code and received did not match, expected %v , got %v", expectedStatusCode, receivedStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder
}
