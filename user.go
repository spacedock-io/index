package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
  "strings"
)

func CreateUser(req *f.Request, res *f.Response, next func()) {
  var username, email, password string

  json := req.Map["json"].(map[string]interface{})

  if len(req.Body) > 0 {
    username, email, password = req.Body["username"], req.Body["email"],
      req.Body["password"]
  } else if len(json) > 0 {
    username, _ = json["username"].(string)
    password, _ = json["password"].(string)
    email, _ = json["email"].(string)
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
    u := models.User{}

    u.Username = username
    u.Emails = append(u.Emails, models.Email{Email: email})

    err := u.Create(password)
    if (err != nil) {
      // @TODO: Don't just send the whole error here
      res.Send(err.Error(), 400)
    }
    res.Send(201)
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

  username = strings.ToLower(req.Params["username"])
  u := req.Map["_user"].(*models.User)

  if strings.ToLower(u.Username) != username && !u.Admin {
    res.Send("You are not authorized to update this user.", 401)
    return
  }

  if len(req.Body) > 0 {
    email = req.Body["email"]
    newPass = req.Body["password"]
  } else if len(req.Map) > 0 {
    email = req.Map["email"].(string)
    newPass = req.Map["password"].(string)
  }

  if len(newPass) < 5 {
    res.Send("Password too short", 400)
    return
  }

  err := u.Update(email, newPass)
  if err != nil {
    res.Send(err.Error(), 400)
    return
  }

  res.Send("User Updated", 204)
}
