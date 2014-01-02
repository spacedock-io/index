package couch

import (
  "fmt"
  "github.com/fjl/go-couchdb"
  "github.com/yawnt/index.spacedock/config"
)

var Couch *couchdb.Database

func New() *couchdb.Database {
  var err error

  proto, host, port, db := config.Global.Get("couch.protocol").Str("http"),
    config.Global.Get("couch.host").Str("localhost"),
    int(config.Global.Get("couch.port").Float64(5984)),
    config.Global.Get("couch.database").Str()

  if len(db) == 0 {
    config.Logger.Log("c", "There was no database specified in the config.")
  }

  url := fmt.Sprintf("%s://%s:%d/", proto, host, port)

  server := couchdb.NewServer(url, nil)
  Couch = server.Db(db)

  if err != nil {
    config.Logger.Log("c", err.Error())
  }

  return Couch
}
