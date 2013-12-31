package models_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/yawnt/index.spacedock/config"
  "github.com/yawnt/index.spacedock/couch"
  "github.com/yawnt/index.spacedock/models"
)

func init() {
  config.Global = config.Load("test")
  couch.Couch = couch.New()
}

func TestUserGet(t *testing.T) {
  user, err := models.GetUser("mmalecki")
  assert.Nil(t, err, "Error should be `nil`")
  assert.NotNil(t, user, "User should not be `nil`")
}

func TestUserGetNoSuchUser(t *testing.T) {
  user, err := models.GetUser("404")
  assert.Nil(t, err, "Error should be `nil`")
  assert.Nil(t, user, "User should be `nil`")
}
