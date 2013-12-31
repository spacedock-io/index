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
  username := "mmalecki"
  user, err := models.GetUser(username)
  assert.Nil(t, err, "Error should be `nil`")
  assert.NotNil(t, user, "User should not be `nil`")
  assert.Equal(t, user.Username, username, "Username should be correct")
}

func TestUserGetNoSuchUser(t *testing.T) {
  user, err := models.GetUser("404")
  assert.Nil(t, err, "Error should be `nil`")
  assert.Nil(t, user, "User should be `nil`")
}

func TestUserCreate(t *testing.T) {
  username := "foo"

  user := &models.User{
    Username: username,
  }

  err := models.CreateUser(user)
  assert.Nil(t, err, "Error should be `nil`")

  getUser, getError := models.GetUser(username)
  assert.Nil(t, getError, "Get error should be `nil`")
  assert.Equal(t, getUser.Username, username)
}

func TestUserCreateAlreadyExists(t *testing.T) {
  username := "mmalecki"

  user := &models.User{
    Username: username,
  }

  err := models.CreateUser(user)
  assert.NotNil(t, err, "Error should not be `nil`")
  assert.IsType(t, err, models.AlreadyExistsError{})
}
