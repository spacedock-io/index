package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
  err := SpacedockUser.Create(SpacedockUserPassword)
  assert.Nil(t, err, "Error should be `nil`")
}

func TestUserAlreadyExists(t *testing.T) {
  err := SpacedockUser.Create(SpacedockUserPassword)
  assert.NotNil(t, err, "Error should not be `nil`")
}

func TestUserGet(t *testing.T) {
  getUser, getError := GetUser(SpacedockUser.Username)
  assert.Nil(t, getError, "Get error should be `nil`")
  assert.True(t, getUser.MatchPassword(SpacedockUserPassword), "Password should match")
  assert.Equal(t, getUser.Username, SpacedockUser.Username)
}

func TestUserDelete(t *testing.T) {
  user, err := GetUser(SpacedockUser.Username)
  assert.Nil(t, err, "Get error should be `nil`")
  err = user.Delete()
  assert.Nil(t, err, "Delete error should be `nil`")
}

func TestNoSuchUser(t *testing.T) {
  getUser, _ := GetUser(SpacedockUser.Username)
  assert.Nil(t, getUser, "User shouldn't exists after being deleted")
}
