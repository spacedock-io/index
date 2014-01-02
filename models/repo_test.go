// I think I'm starting to understand the Go way.
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

func TestGetRepoNoSuchRepo(t *testing.T) {
  repo, err := models.GetRepo("404", "404")
  assert.Nil(t, err, "Error should be `nil`")
  assert.Nil(t, repo, "Repo should be `nil`")
}

func TestCreateGetDeleteRepo(t *testing.T) {
  namespace := "spacedock"
  name := "foo"

  newRepo := &models.Repo{
    Namespace: namespace,
    Name: name,
  }

  err := models.CreateRepo(newRepo)
  assert.Nil(t, err, "Error should be `nil`")

  repo, err := models.GetRepo(namespace, name)
  assert.Nil(t, err, "Error should be `nil`")
  assert.NotNil(t, repo, "Repo should not be `nil`")
  assert.Equal(t, repo.Namespace, namespace, "Namespace should be correct")
  assert.Equal(t, repo.Name, name, "Name should be correct")

  err = models.DeleteRepo(namespace, name)
  assert.Nil(t, err, "Delete error should be `nil`")

  repo, err = models.GetRepo(namespace, name)
  assert.Nil(t, repo, "Repo shouldn't exists after being deleted")
}

func TestCreateAlreadyExistsRepo(t *testing.T) {
  namespace := "spacedock"
  name := "already-exists"

  repo := &models.Repo{
    Namespace: namespace,
    Name: name,
  }

  err := models.CreateRepo(repo)
  assert.Nil(t, err, "Error should be `nil` when creating the user for the first time")
  err = models.CreateRepo(repo)
  assert.NotNil(t, err, "Error should not be `nil`")
  assert.IsType(t, err, models.AlreadyExistsError{})
}
