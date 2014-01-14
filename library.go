package main

import (
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func CreateLibrary(req *f.Request, res *f.Response, next func()) {
  // @TODO: Make this smarter, and maybe a middleware
  if req.Map["_admin"] != true {
    res.Send("Not Authorized", 401)
    return
  }

  images := req.Map["json"]

  repo := req.Params["repo"]

  r := models.Repo{}


  ts, err := r.Create(repo, "", "1", req.Map["_uid"].(int64), images.([]interface{}))
  if err != nil {
    res.Send(err.Error(), 400)
  }

  res.Set("X-Docker-Token", ts)
  res.Set("WWW-Authenticate", "Token " + ts)
  res.Set("X-Docker-Endpoints", "reg22.spacedock.io, reg41.spacedock.io")

  res.Send("Created", 200)
}

func DeleteLibrary(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
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
