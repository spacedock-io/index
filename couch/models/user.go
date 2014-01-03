package models

import (
  "github.com/spacedock-io/index/couch"
  "github.com/fjl/go-couchdb"
)

var prefix string = "user:"

type User struct {
  Username string `json:"username"`
  Rev string `json:"_rev,omitempty"`
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

func CreateUser(user *User) error {
  _, err := couch.Global.Put(prefix + user.Username, user)
  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 409 {
      err = AlreadyExistsError{}
    }
  }

  return err
}

func DeleteUser(name string) error {
  user, err := GetUser(name)
  if err != nil {
    return err
  }

  _, err = couch.Global.Delete(prefix + user.Username, user.Rev)
  return err
}
