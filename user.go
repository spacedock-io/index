package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/couch/models"
)

func CreateUser(req *f.Request, res *f.Response, next func()) {
  var username, email, password string

  if len(req.Body) > 0 {
    username, email, password = req.Body["username"], req.Body["email"],
      req.Body["password"]
  } else if len(req.Request.Map) > 0 {
    username, _ = req.Map["username"].(string)
    password, _ = req.Map["password"].(string)
    email, _ = req.Map["email"].(string)
  }

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
    u.Email = append(u.Email, email)

    e := u.Create(password)
    if (e != nil) {
      // @TODO: Don't just send the whole error here
      res.Send(e.Error(), 400)
    }
    res.Send("User created successfully", 201)
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
  var username, email, newPass string

  username = req.Params["username"]

  if len(req.Body) > 0 {
    email = req.Body["email"]
    newPass = req.Body["password"]
  } else if len(req.Map) > 0 {
    email = req.Map["email"].(string)
    newPass = req.Map["password"].(string)
  }

  u, e := models.GetUser(username)
  if e != nil {
    res.Send(e.Error(), 400)
    return
  }

  if len(newPass) > 5 {
    u.SetPassword(newPass)
  } else if len(newPass) > 0 {
    res.Send("Password too short", 400)
    return
  }

  if len(email) > 0 { u.Email = append(u.Email, email) }

  e = u.Save(true)
  if e != nil {
    res.Send(e.Error(), 400)
    return
  }

  res.Send("User Updated", 204)
}
