package models

import (
  "github.com/yawnt/index.spacedock/couch"
  "github.com/fjl/go-couchdb"
)

var prefix string = "user:"

type User struct {
  Username string `json:"username"`
}

func NewUser() *User {
  return &User{}
}

func GetUser(name string) (*User, error) {
  ret := &User{}
  err := couch.Couch.Get(prefix + name, nil, ret)

  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 404 {
      err = nil
      ret = nil
    }
  }
  return ret, err
}

func CreateUser(user User) error {
  _, err := couch.Couch.Put(prefix + user.Username, user)
  return err
}
