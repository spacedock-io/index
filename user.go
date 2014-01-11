package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
  "github.com/spacedock-io/registry/db"
)

func CreateUser(req *f.Request, res *f.Response, next func()) {
  var username, email, password string

  if len(req.Body) > 0 {
    username, email, password = req.Body["username"], req.Body["email"],
      req.Body["password"]
  } else if len(req.Request.Map) > 0 {
    json := req.Map["json"].(map[string]string)
    username, _ = json["username"]
    password, _ = json["password"]
    email, _ = json["email"]
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
    u.Emails = append(u.Emails, email)

    q := u.Create(password)
    if (q.Error != nil) {
      // @TODO: Don't just send the whole error here
      res.Send(q.Error(), 400)
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
  var u models.User

  username = req.Params["username"]

  if len(req.Body) > 0 {
    email = req.Body["email"]
    newPass = req.Body["password"]
  } else if len(req.Map) > 0 {
    email = req.Map["email"].(string)
    newPass = req.Map["password"].(string)
  }

  query := db.DB.Where("Username = ?", username).Find(&u)
  if query.Error != nil {
    res.Send(query.Error, 400)
    return
  }

  if len(newPass) > 5 {
    u.SetPassword(newPass)
  } else if len(newPass) > 0 {
    res.Send("Password too short", 400)
    return
  }

  if len(email) > 0 { u.Emails = append(u.Emails, email) }

  query = db.DB.Save(&u)
  if query.Error != nil {
    res.Send(query.Error, 400)
    return
  }

  res.Send("User Updated", 204)
}
