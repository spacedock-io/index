package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

var (
  username = "foo"
  email = "foo@bar.com"
  newEmail = "bar@foo.com"
  pass = "4321"
  user = &User{
    Username: username,
    Email: []string{email},
    Pass: HexString(pass),
  }
)

func TestUserCreate(t *testing.T) {
  err := user.Create(pass)
  assert.Nil(t, err, "Error should be `nil`")
}

func TestUserAlreadyExists(t *testing.T) {
  err := user.Create(pass)
  assert.NotNil(t, err, "Error should not be `nil`")
  assert.IsType(t, AlreadyExistsError{}, err)
}

func TestUserAddEmail(t *testing.T) {
  err := user.AddEmail(newEmail)
  assert.Nil(t, err, "Add email error should be `nil`")
}

func TestUserGet(t *testing.T) {
  user, err := GetUser(username)
  assert.Nil(t, err, "Get error should be `nil`")
  assert.True(t, user.MatchPassword(pass), "Password should match")
  assert.Equal(t, user.Username, username)
  assert.Equal(t, user.Email[0], email)
  assert.Equal(t, user.Email[1], newEmail)
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
