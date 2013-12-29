package main

import (
  "github.com/ricallinson/forgery"
)

func CreateUser(req *f.Request, res *f.Response, next func()) {
  // res.Send("Not implemented yet.")
  res.Send("Unknown error while trying to register user", 400)
}

func Login(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}

func UpdateUser(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}
