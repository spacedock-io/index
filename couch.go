package main

import (
  "fmt"
  "github.com/lancecarlson/couchgo"
  "github.com/southern/logger"
)

var Couch *couch.Client

func SetupCouch() {
  var err error

  proto, host, port, db := global.Get("couch.protocol").Str("http"),
    global.Get("couch.host").Str("localhost"),
    string(int(global.Get("couch.port").Float64(5984))),
    global.Get("couch.database").Str()

  if len(db) == 0 {
    globalLog.Log(logger.CRIT,
      "There was no database specified in the config.")
  }

  url := fmt.Sprintf("%s://%s:%s/%s", proto, host, port, db)

  Couch, err = couch.NewClientURL(url)
  if err != nil {
    globalLog.Log(logger.CRIT, err.Error())
  }
}
