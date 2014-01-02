package models

import (
  "github.com/spacedock-io/index/couch"
  "github.com/fjl/go-couchdb"
)

var repoSeparator string = ":"
var repoPrefix string = "repo" + repoSeparator

type Repo struct {
  Id string `json:"_id,omitempty"`
  Rev string `json:"_rev,omitempty"`
  RegistryId string `json:"registryId"`
  Namespace string `json:"namespace"`
  Name string `json:"name"`
}

func NewRepo() *Repo {
  return &Repo{}
}

func GetRepo(namespace string, repo string) (*Repo, error) {
  ret := &Repo{}
  err := couch.Couch.Get(repoPrefix + namespace + repoSeparator + repo, nil, ret)

  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 404 {
      err = nil
      ret = nil
    }
  }
  return ret, err
}

func CreateRepo(repo *Repo) error {
  _, err := couch.Couch.Put(repoPrefix + repo.Namespace + repoSeparator + repo.Name, repo)
  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 409 {
      err = AlreadyExistsError{}
    }
  }

  return err
}

func DeleteRepo(namespace string, name string) error {
  repo, err := GetRepo(namespace, name)
  if err != nil {
    return err
  }

  _, err = couch.Couch.Delete(repo.Id, repo.Rev)
  return err
}
