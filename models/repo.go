package models

import (
  "github.com/spacedock-io/registry/db"
)

type Repo struct {
  Id            int64
  RegistryId    string  `sql:"not null"`
  Namespace     string  `sql:"not null"`
  Name          string  `sql:"not null"`
  Tokens        []Token
  Images        []Image
  Deleted       bool
}

type Image struct {
  Id            int64   `json:"-"`
  Uuid          string  `json:"id"`
  Checksum      string  `json:"checksum,omitempty"`
  RepoId        int64   `json:"-"`
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

func (r *Repo) AddToken(access string, user *User) (Token, error) {
  repo := r.Namespace + "/" + r.Name

  t, ok := CreateToken(access, user.Id, repo)
  if !ok {
    return Token{}, TokenErr
  }

  r.Tokens = append(r.Tokens, t)
  return t, nil
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

func (r *Repo) Create(regId string, user *User,
                      images []interface{}) (string, error) {
  r.RegistryId = regId

  if len(r.Namespace) == 0 {
    r.Namespace = "library"
  }

  q := db.DB.Where("namespace = ? and name = ?", r.Namespace, r.Name).Find(&Repo{})
  if !q.RecordNotFound() {
    return "", AlreadyExistsError
  }

  ok := user.SetAccess(r.Namespace, r.Name, "delete")
  if !ok {
    return "", AccessSetError
  }

  t, err := r.AddToken("write", user)
  if err != nil {
    return "", err
  }

  for _, v := range images {
    row := v.(map[string]interface{})
    id := row["id"].(string)
    img := Image{Uuid: id}
    r.Images = append(r.Images, img)
  }

  err = r.Save()
  if err != nil {
    return "", err
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

func (repo *Repo) HasToken(token string) bool {
  for _, v := range repo.Tokens {
    if v.String() == token {
      return true
    }
  }

  return false
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

  err := repo.Save()
  if err != nil {
    return "", err
  }

  return ts, nil
}

func (r *Repo) Save() error {
  q := db.DB.Save(r)
  if q.Error != nil {
    return DBErr
  }
  return nil
}

func (repo *Repo) UpdateImages(updates []interface{}) error {
  var imgs, updated []Image

  q := db.DB.Model(repo).Related(&imgs)
  if q.Error != nil {
    return DBErr
  }

  for _, update := range updates {
    var found bool
    row := update.(map[string]interface{})
    for _, image := range imgs {
      if row["id"].(string) == image.Uuid {
        image.Checksum = row["checksum"].(string)
        updated = append(updated, image)
        found = true
      }
    }
    if !found {
      img := Image{
        Uuid: row["id"].(string),
        Checksum: row["checksum"].(string),
      }
      updated = append(updated, img)
      found = false
    }
  }

  repo.Images = updated

  err := repo.Save()
  if err != nil {
    return err
  }

  return nil
}
