package models

import "github.com/yawnt/index.spacedock/couch"

type User struct {

}

func NewUser() *User {
  return &User{}
}

func GetUser(name string) (error, *User) {
  ret := &User{}
  err := couch.Couch.Get("user/" + name, nil, ret)
  if err != nil && err.StatusCode == 404 {
    err = nil
    ret = nil
  }
  return ret, err
}
