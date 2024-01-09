package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/AbdouTlili/backawi/pkg/config"
	"github.com/AbdouTlili/backawi/pkg/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type App struct {
	Config    *aws.Config
	DBManager *db.DBManager
}

func main() {

	cfg := config.LoadConfig()

	log.Info("Config loaded form Env \nTable name is : ", cfg.DynamoDBTableName, " \nAWS region is : ", cfg.AWSRegion, "\nSecurity Keys loaded")

	app := App{Config: &aws.Config{
		Region:      &cfg.AWSRegion,
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	}, DBManager: &db.DBManager{}}

	app.DBManager.Init(app.Config, cfg.DynamoDBTableName)

	app.DBManager.CreateItemInDB("5585", "productname", "20", "50", "2555.5")

}
