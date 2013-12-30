package models_test

import (
  "testing"
  "github.com/yawnt/index.spacedock/config"
  "github.com/yawnt/index.spacedock/couch"
  "github.com/yawnt/index.spacedock/models"
)

func init() {
  config.Global = config.Load("test")
  couch.Couch = couch.New()
}

func TestUserGetNoSuchUser(t *testing.T) {
  err, user := models.GetUser("404")
  if err != nil {
    t.Error("Error while getting user")
  }
  if user != nil {
    t.Error("User should be `nil`")
  }
}
