package models

import (
  "crypto/sha256"
  "github.com/spacedock-io/index/couch"
  "github.com/fjl/go-couchdb"
  "github.com/gokyle/pbkdf2"
)

var prefix string = "user:"

type User struct {
  Username string `json:"username"`
  Email string `json:"email"`
  Salt []byte `json:"salt"`
  Pass []byte `json:"pass"`
  Rev string `json:"_rev,omitempty"`
}

func init() {
  pbkdf2.HashFunc = sha256.New
}

func NewUser() *User {
  return &User{}
}

func GetUser(name string) (*User, error) {
  ret := &User{}
  err := couch.Global.Get(prefix + name, nil, ret)

  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 404 {
      err = nil
      ret = nil
    }
  }
  return ret, err
}

func (user *User) Create() error {
  ph := pbkdf2.HashPassword(string(user.Pass))
  user.Salt = ph.Salt
  user.Pass = ph.Hash

  _, err := couch.Global.Put(prefix + user.Username, user)
  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 409 {
      err = AlreadyExistsError{}
    }
  }

  return err
}

func (user *User) Delete() error {
  _, err := couch.Global.Delete(prefix + user.Username, user.Rev)
  return err
}

func (user *User) MatchPassword(pass string) bool {
  ph := &pbkdf2.PasswordHash{ Hash: user.Pass, Salt: user.Salt}
  return pbkdf2.MatchPassword(pass, ph)
}

func AuthUser(user string, pass string) bool {
  ret, err := GetUser(user)
  if (err != nil) {
    return false;
  }
  return ret.MatchPassword(pass)
}
