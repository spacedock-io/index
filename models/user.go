package models

import (
  "crypto/sha256"
  "github.com/gokyle/pbkdf2"
  "github.com/spacedock-io/registry/db"
)

func init() {
  pbkdf2.HashFunc = sha256.New
}

type Access struct {
  Id        int64
  UserId    int64
  Repo      string
  Access    string  `sql:"not null"`
}

type Email struct {
  Id        int64
  Email     string  `sql:"not null;unique"`
  UserId    int64
}

type User struct {
  Id        int64
  Username  string
  Admin     bool
  Access    []Access
  Emails    []Email
  Hash      []byte
  Salt      []byte
}


// @TODO: Make this not return a bool
func AuthUser(user string, pass string) (*User, bool) {
  u, err := GetUser(user)
  if err != nil { return nil, false }
  res := u.MatchPassword(pass)
  if !res { return nil, false }
  return u, true
}

func GetUser(username string) (*User, error) {
  u := &User{}
  q := db.DB.Where("Username = ?", username).Find(u)
  if q.Error != nil {
    if q.RecordNotFound() {
      return nil, NotFoundErr
    } else {
      return nil, q.Error
    }
  }
  return u, nil
}

func (user *User) Create(password string) error {
  user.SetPassword(password)
  q := db.DB.Save(user)
  if q.Error != nil {
    return SaveErr
  }
  return nil
}

func (user *User) GetAccess(ns, repo string) string {
  if len(ns) > 0 {
    repo = ns + "/" + repo
  }

  a :=  Access{}
  q := db.DB.Where(&Access{UserId: user.Id, Repo: repo}).First(&a)
  if q.Error != nil {
    return ""
  }
  return a.Access
}

func (user *User) SetAccess(ns, repo, access string) bool {
  if len(ns) > 0 {
    repo = ns + "/" + repo
  }

  if len(repo) == 0 {
    return false
  }

  var perms []Access

  a := Access{
    Repo: repo,
    Access: access,
  }

  q := db.DB.Model(user).Related(&perms)
  if q.Error != nil {
    return false
  }

  perms = append(perms, a)
  user.Access = perms

  q = db.DB.Save(user)
  if q.Error != nil {
    return false
  }

  return true
}

func (user *User) SetPassword(pass string) {
  ph := pbkdf2.HashPassword(pass)
  user.Salt = ph.Salt
  user.Hash = ph.Hash
}

func (user *User) MatchPassword(pass string) bool {
  ph := &pbkdf2.PasswordHash{ Hash: user.Hash, Salt: user.Salt}
  return pbkdf2.MatchPassword(pass, ph)
}

func (user *User) Update(email, password string) error {
  if len(password) > 5 {
    user.SetPassword(password)
  }

  if len(email) > 0 {
    user.Emails = append(user.Emails, Email{Email: email})
  }

  q := db.DB.Save(user)
  if q.Error != nil {
    return DBErr
  }
  return nil
}
