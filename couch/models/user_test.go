package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

var (
  username = "foo"
  email = "foo@bar.com"
  pass = "4321"
  user = &User{
    Username: username,
    Email: email,
    Pass: []byte(pass),
  }
)

func TestUserCreate(t *testing.T) {
  err := user.Create()
  assert.Nil(t, err, "Error should be `nil`")
}

func TestUserAlreadyExists(t *testing.T) {
  err := user.Create()
  assert.NotNil(t, err, "Error should not be `nil`")
  assert.IsType(t, AlreadyExistsError{}, err)
}

func TestUserGet(t *testing.T) {
  getUser, getError := GetUser(username)
  assert.Nil(t, getError, "Get error should be `nil`")
  assert.True(t, getUser.MatchPassword(pass), "Password should match")
  assert.Equal(t, getUser.Username, username)
}

func TestUserDelete(t *testing.T) {
  user, err := GetUser(username)
  assert.Nil(t, err, "Get error should be `nil`")
  err = user.Delete()
  assert.Nil(t, err, "Delete error should be `nil`")
}

func TestNoSuchUser(t *testing.T) {
  getUser, _ := GetUser(username)
  assert.Nil(t, getUser, "User shouldn't exists after being deleted")
}
