package main

import (
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func CreateLibrary(req *f.Request, res *f.Response, next func()) {
  images := req.Map["json"]
  u := req.Map["_user"].(models.User)

  // @TODO: Make this smarter, and maybe a middleware
  if !u.Admin {
    res.Send("Not Authorized", 401)
    return
  }

  repo := req.Params["repo"]

  r := models.Repo{}

  ts, err := r.Create(repo, "", "1", u.Id, images.([]interface{}))
  if err != nil {
    res.Send(err.Error(), 400)
  }

  res.Set("X-Docker-Token", ts)
  res.Set("WWW-Authenticate", "Token " + ts)
  res.Set("X-Docker-Endpoints", "reg22.spacedock.io, reg41.spacedock.io")

  res.Send("Created", 200)
}

func DeleteLibrary(req *f.Request, res *f.Response, next func()) {
  repo := req.Params["repo"]

  r, err := models.GetRepo("", repo)
  if err != nil {
    res.Send(err.Error(), 400)
    return
  }

  if !r.Deleted {
    ts, err := r.MarkAsDeleted(req.Map["_uid"].(int64))
    if err != nil {
      res.Send(err.Error(), 400)
      return
    }

    res.Set("X-Docker-Token", ts)
    res.Set("WWW-Authenticate", "Token " + ts)
    res.Set("X-Docker-Endpoints", "reg22.spacedock.io, reg41.spacedock.io")

    res.Send(202)
    return
  }

  err = r.Delete()
  if err != nil {
    res.Send(err.Error(), 400)
    return
  }

  res.Send(200)
}

func GetLibraryImage(req *f.Request, res *f.Response, next func()) {
  repo := req.Params["repo"]
  
  r, err := models.GetRepo("", repo)
  if err != nil {
    res.Send(err.Error(), 400)
  }

  images, er := r.GetImages()
  if er != nil {
    res.Send(er.Error(), 400)
  }

  j, e := json.Marshal(images)
  if e != nil {
    res.Send("Error returning data", 400)
  }

  res.Send(string(j), 200)
}

func LibraryAuth(req *f.Request, res *f.Response, next func()) {
  res.Send(200)
}

func UpdateLibraryImage(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}
