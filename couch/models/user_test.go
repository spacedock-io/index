package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/spacedock-io/index/config"
  "github.com/spacedock-io/index/couch"
)

var (
  username = "foo"
  user = &User{
    Username: username,
  }
)

func init() {
  config.Global = config.Load("test")
  couch.Couch = couch.New()
}

func TestUserCreate(t *testing.T) {
  // Since tests are ran in parallel (in goroutines), we need to create and
  // destroy user in one test.
  err := CreateUser(user)
  assert.Nil(t, err, "Error should be `nil`")
}

func TestAlreadyExists(t *testing.T) {
  err := CreateUser(user)
  assert.NotNil(t, err, "Error should not be `nil`")
  assert.IsType(t, err, AlreadyExistsError{})
}

func TestUserGet(t *testing.T) {
  getUser, getError := GetUser(username)
  assert.Nil(t, getError, "Get error should be `nil`")
  assert.Equal(t, getUser.Username, username)
}

func TestUserDelete(t *testing.T) {
  err := DeleteUser(username)
  assert.Nil(t, err, "Delete error should be `nil`")  
}

func TestNoSuchUser(t *testing.T) {
  getUser, _ := GetUser(username)
  assert.Nil(t, getUser, "User shouldn't exists after being deleted")
}