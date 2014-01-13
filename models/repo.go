package models

import (
  "github.com/spacedock-io/registry/db"
)

type Repo struct {
  Id            int64
  RegistryId    string `sql:"not null"`
  Namespace     string `sql:"not null"`
  Name          string `sql:"not null;unique"`
  Tokens        []Token
}

func GetRepo(namespace string, repo string) (*Repo, error) {
  r := &Repo{}
  q := db.DB.Where("namespace = ? and name = ?", namespace, repo).Find(r)

  if q.Error != nil {
    if q.RecordNotFound() {
      return nil, NotFoundErr{}
    } else { return nil, DBErr{} }
  }
  return r, nil
}

func (repo *Repo) Create() error {
  q := db.DB.Save(repo)
  if q.Error != nil {
    return DBErr{}
  }
  return nil
}

func (repo *Repo) Delete() error {
  q := db.DB.Delete(repo)
  if q.Error != nil {
    return DBErr{}
  }
  return nil
}

func (repo *Repo) HasToken(token string) bool {
  for _, v := range repo.Tokens {
    if v.String() == token {
      return true
    }
  }

  return false
}
