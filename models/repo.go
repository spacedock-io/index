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
  Images        []Image
  Deleted       bool
}

func GetRepo(namespace string, repo string) (*Repo, error) {
  r := &Repo{}
  q := db.DB.Where("namespace = ? and name = ?", namespace, repo).Find(r)

  if q.Error != nil {
    if q.RecordNotFound() {
      return nil, NotFoundErr
    } else { return nil, DBErr }
  }
  return r, nil
}

func (r *Repo) GetImages() ([]Image, error) {
  var i []Image

  q := db.DB.Model(r).Related(&i)
  if q.Error != nil {
    if q.RecordNotFound() {
      return nil, NotFoundErr
    } else { return nil, DBErr }
  }

  return i, nil
}

func (r *Repo) Create(repo, ns, regId string, uid int64,
                      images []interface{}) (string, error) {
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
  if !ok { return "", TokenErr }

  r.Tokens = append(r.Tokens, t)

  for _, v := range images {
    row := v.(map[string]interface{})
    img := Image{}
    id := row["id"].(string)
    img.Create(id)
    r.Images = append(r.Images, img)
  }

  q := db.DB.Save(r)
  if q.Error != nil {
    return "", DBErr
  }
  return t.String(), nil
}

func (repo *Repo) Delete() error {
  q := db.DB.Delete(repo)
  if q.Error != nil {
    return DBErr
  }
  return nil
}

func (repo *Repo) MarkAsDeleted(uid int64) (string, error) {
  var fullname, ts string
  if len(repo.Namespace) == 0 {
    fullname = "library/" + repo.Name
  } else {
    fullname = repo.Namespace + repo.Name
  }
  t, ok := CreateToken("delete", uid, fullname)
  if !ok {
    return "", TokenErr
  }

  repo.Tokens = append(repo.Tokens, t)
  ts = t.String()

  repo.Deleted = true
  q := db.DB.Save(repo)
  if q.Error != nil {
    return "", DBErr
  }
  return ts, nil
}

func (repo *Repo) HasToken(token string) bool {
  for _, v := range repo.Tokens {
    if v.String() == token {
      return true
    }
  }

  return false
}
