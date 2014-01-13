package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func CreateLibrary(req *f.Request, res *f.Response, next func()) {
  // @TODO: Make this smarter, and maybe a middleware
  if req.Map["_admin"] != true {
    res.Send("Not Authorized", 401)
    return
  }

  images := req.Map["json"].([]map[string]interface{})

  repo := req.Params["repo"]

  r := models.Repo{}

  ts, err := r.Create(repo, "", "1", req.Map["_uid"].(int64), images)
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
  res.Send("Not implemented yet.")
}

func LibraryAuth(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}

func UpdateLibraryImage(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}
