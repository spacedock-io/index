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
  user, err := models.GetUser("404")
  if err != nil {
    t.Errorf("Error should be `nil`, got: %s", err)
  }
  if user != nil {
    t.Errorf("User should be `nil`, got: %s", user)
  }
}
