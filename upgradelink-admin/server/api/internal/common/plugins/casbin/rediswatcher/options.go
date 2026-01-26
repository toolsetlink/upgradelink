package rediswatcher

import (
	"github.com/google/uuid"
	rds "github.com/redis/go-redis/v9"
)

type WatcherOptions struct {
	Options                rds.Options
	ClusterOptions         rds.ClusterOptions
	SubClient              *rds.Client
	PubClient              *rds.Client
	SubClusterClient       *rds.ClusterClient
	PubClusterClient       *rds.ClusterClient
	Channel                string
	IgnoreSelf             bool
	LocalID                string
	OptionalUpdateCallback func(string)
}

func initConfig(option *WatcherOptions) {
	if option.LocalID == "" {
		option.LocalID = uuid.New().String()
	}
	if option.Channel == "" {
		option.Channel = "/casbin"
	}
}
