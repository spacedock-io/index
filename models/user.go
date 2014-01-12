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
  /* UserId    int64 */
}

type User struct {
  Id        int64
  Username  string
  Emails    []Email
  Hash      []byte
  Salt      []byte
}

type LolErr struct{}

func (err LolErr) Error() string {
  return "Lol shit be fucked up dawg"
}



func (user *User) Create(password string) error {
  user.SetPassword(password)
  q := db.DB.Save(&user)
  if q.Error != nil {
    return LolErr{}
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

func AuthUser(user string, pass string) bool {
  var u User
  query := db.DB.Where("Username = ?", user).Find(&u)
  if query.Error != nil {
    return false;
  }
  return u.MatchPassword(pass)
}

