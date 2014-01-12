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
  Emails    []Email
  Hash      []byte
  Salt      []byte
}

type UserErr struct{}
type NotFoundErr struct{}

func (err UserErr) Error() string {
  return "There was an error saving to users"
}

func (err NotFoundErr) Error() string {
  return "User not found"
}

func GetUser(username string) (*User, error) {
  u := &User{}
  q := db.DB.Where("Username = ?", username).Find(u)
  if q.Error != nil {
    if q.RecordNotFound() {
      return nil, NotFoundErr{}
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
    return UserErr{}
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
func AuthUser(user string, pass string) bool {
  u, err := GetUser(user)
  if err != nil { return false }
  return u.MatchPassword(pass)
}

