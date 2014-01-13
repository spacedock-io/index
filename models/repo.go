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

func (r *Repo) Create(repo, ns, regId string, uid int64) (string, error) {
  var fullname string
  r.Name = repo
  r.RegistryId = regId

  if len(ns) == 0 {
    fullname = "library/" + repo
    r.Namespace = ""
  } else {
    fullname = ns + "/" + repo
    r.Namespace = ns
  }

  // @TODO: make sure this access level is right
  t, ok := CreateToken("write", uid, fullname)
  if !ok { return "", TokenErr{} }

  r.Tokens = append(r.Tokens, t)

  q := db.DB.Save(r)
  if q.Error != nil {
    return "", DBErr{}
  }
  return t.String(), nil
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
