package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/couch/models"
)

func CreateUser(req *f.Request, res *f.Response, next func()) {
  username, email, password := req.Body["username"], req.Body["email"], req.Body["password"]

  // @TODO: Validate email format

  if len(password) < 5 {
    res.Send("Password too short", 400)
  } else if len(username) < 4 {
    res.Send("Username too short", 400)
  } else if len(username) > 30 {
    res.Send("Username too long", 400)
  } else {
    // put user in couch, send confirm email
    u := models.NewUser()

    u.Username = username
    u.Email = email
    u.Pass = password

    e := u.Create()
    if (e != nil) {
      // @TODO: Don't just send the whole error here
      res.Send(e, 400)
    }
    res.Send("User created successfully", 200)
    // later on, send an async email
    //go ConfirmEmail()
  }

  res.Send("Unknown error while trying to register user", 400)
}

func Login(req *f.Request, res *f.Response, next func()) {
  // Because of middleware, execution only gets here on success.
  res.Send("OK", 200)
}

func UpdateUser(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}
