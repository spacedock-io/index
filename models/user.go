package models

import (
  "crypto/sha256"
  "github.com/gokyle/pbkdf2"
  "github.com/spacedock-io/registry/db"
)

func init() {
  pbkdf2.HashFunc = sha256.New
}

type Email struct {
  Id        int64
  Email     string
  UserId    int64
}

type User struct {
  Id        int64
  Username  string
  Admin     bool
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

