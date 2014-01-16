package models

import (
  "crypto/sha256"
  "github.com/gokyle/pbkdf2"
  "github.com/spacedock-io/registry/db"
  "strings"
)

func init() {
  pbkdf2.HashFunc = sha256.New
}

type Access struct {
  Id        int64
  UserId    int64
  Repo      string  `sql:"not null;unique"`
  Access    string  `sql:not null`
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

func (user *User) SetAccess(repo string, access string) bool {
  if len(repo) == 0 {
    return false
  }

  a := Access{
    Repo: repo,
    Access: access,
  }
  user.Access = append(user.Access, a)

  q := db.DB.Save(user)
  if q.Error != nil {
    return false
  }

  return true
}

func (user *User) GetAccess(repo string) string {
  for _, v := range user.Access {
    if strings.ToLower(repo) == strings.ToLower(v.Repo) {
      return v.Access
    }
  }
  return ""
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

// @TODO: Make this not return a bool
func AuthUser(user string, pass string) (*User, bool) {
  u, err := GetUser(user)
  if err != nil { return nil, false }
  res := u.MatchPassword(pass)
  if !res { return nil, false }
  return u, true
}

