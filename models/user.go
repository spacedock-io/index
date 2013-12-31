package models

import (
  "github.com/yawnt/index.spacedock/couch"
  "github.com/fjl/go-couchdb"
)

type User struct {

}

func NewUser() *User {
  return &User{}
}

func GetUser(name string) (*User, error) {
  ret := &User{}
  err := couch.Couch.Get("user/" + name, nil, ret)

  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 404 {
      err = nil
      ret = nil
    }
  }
  return ret, err
}
