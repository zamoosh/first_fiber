package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	DB   *mongo.Client
	opts = options.Client().
		SetMaxConnIdleTime(time.Minute * 20).
		SetTimeout(time.Second * 10).
		SetMinPoolSize(5).
		SetMaxPoolSize(10).
		SetHeartbeatInterval(time.Second)
)

const (
	ClientLog    string = "client_log"
	ActiveGps    string = "active_gps"
	ActiveGpsLog string = "active_gps_log"
	SampleGpsLog string = "sample_gps_log"
	Carstates    string = "carstates"
	Charge       string = "charge"
)
