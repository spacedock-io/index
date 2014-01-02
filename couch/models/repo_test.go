package models

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/spacedock-io/index/config"
  "github.com/spacedock-io/index/couch"
)

var (
  namespace = "spacedock"
  name = "foo"
  newRepo = &Repo{
    Namespace: namespace,
    Name: name,
  }
)

func init() {
  config.Global = config.Load("test")
  couch.Couch = couch.New()
}

func TestCreateRepo(t *testing.T) {
  err := CreateRepo(newRepo)
  assert.Nil(t, err, "Error should be `nil`")
}

func TestRepoAlreadyExists(t *testing.T) {
  err := CreateRepo(newRepo)
  assert.NotNil(t, err, "Error should not be `nil`")
  assert.IsType(t, err, AlreadyExistsError{})
}

func TestGetRepo(t *testing.T) {
  repo, err := GetRepo(namespace, name)
  assert.Nil(t, err, "Error should be `nil`")
  assert.NotNil(t, repo, "Repo should not be `nil`")
  assert.Equal(t, repo.Namespace, namespace, "Namespace should be correct")
  assert.Equal(t, repo.Name, name, "Name should be correct")
}

func TestDeleteRepo(t *testing.T) {
  err := DeleteRepo(namespace, name)
  assert.Nil(t, err, "Delete error should be `nil`")
}

func TestNoSuchRepo(t *testing.T) {
  repo, _ := GetRepo(namespace, name)
  assert.Nil(t, repo, "Repo shouldn't exists after being deleted")
}
