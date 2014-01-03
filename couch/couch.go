package couch

import (
  "fmt"
  "github.com/fjl/go-couchdb"
  "github.com/spacedock-io/index/config"
  "github.com/stretchr/objx"
)

var Global *couchdb.Database

func New(c objx.Map) *couchdb.Database {
  var err error

  proto, host, port, name := c.Get("couch.protocol").Str("http"),
    c.Get("couch.host").Str("localhost"),
    int(c.Get("couch.port").Float64(5984)),
    c.Get("couch.database").Str()

  if len(name) == 0 {
    config.Logger.Log("c", "There was no database specified in the config.")
  }

  url := fmt.Sprintf("%s://%s:%d/", proto, host, port)

  server := couchdb.NewServer(url, nil)
  db := server.Db(name)

  if err != nil {
    config.Logger.Log("c", err.Error())
  }

  return db
}
