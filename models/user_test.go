package models_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/spacedock-io/index/config"
  "github.com/spacedock-io/index/couch"
  "github.com/spacedock-io/index/models"
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

func TestUserCreateAndDestroy(t *testing.T) {
  // Since tests are ran in parallel (in goroutines), we need to create and
  // destroy user in one test.
  username := "foo"

  user := &models.User{
    Username: username,
  }

  err := models.CreateUser(user)
  assert.Nil(t, err, "Error should be `nil`")

  getUser, getError := models.GetUser(username)
  assert.Nil(t, getError, "Get error should be `nil`")
  assert.Equal(t, getUser.Username, username)

  err = models.DeleteUser(username)
  assert.Nil(t, err, "Delete error should be `nil`")

  getUser, getError = models.GetUser(username)
  assert.Nil(t, getUser, "User shouldn't exists after being deleted")
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
