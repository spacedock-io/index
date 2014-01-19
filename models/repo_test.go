package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateRepo(t *testing.T) {
  token, err := SpacedockFooRepo.Create("0.spcdck.net", SpacedockUser, SpacedockFooImages)
  assert.Nil(t, err, "Error should be `nil`")
  assert.IsType(t, token, Token{})
}

func TestRepoAlreadyExists(t *testing.T) {
  token, err := SpacedockFooRepo.Create("0.spcdck.net", SpacedockUser, SpacedockFooImages)
  assert.NotNil(t, err, "Error should not be `nil`")
  assert.Equal(t, token, "", "Token should be empty")
}

func TestGetRepo(t *testing.T) {
  repo, err := GetRepo(SpacedockFooRepo.Namespace, SpacedockFooRepo.Name)
  assert.Nil(t, err, "Error should be `nil`")
  assert.NotNil(t, repo, "Repo should not be `nil`")
  assert.Equal(t, repo.Namespace, SpacedockFooRepo.Namespace, "Namespace should be correct")
  assert.Equal(t, repo.Name, SpacedockFooRepo.Name, "Name should be correct")
}

func TestDeleteRepo(t *testing.T) {
  repo, err := GetRepo(SpacedockFooRepo.Namespace, SpacedockFooRepo.Name)
  assert.Nil(t, err, "Get error should be `nil`")
  err = repo.Delete()
  assert.Nil(t, err, "Delete error should be `nil`")
}

func TestNoSuchRepo(t *testing.T) {
  repo, _ := GetRepo(SpacedockFooRepo.Namespace, SpacedockFooRepo.Name)
  assert.Nil(t, repo, "Repo shouldn't exists after being deleted")
}
