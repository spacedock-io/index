package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func CreateRepo(req *f.Request, res *f.Response, next func()) {
  ns := req.Params["namespace"]
  repo := req.Params["repo"]

  images := req.Map["json"]

  r := &models.Repo{}
  ts, err := r.Create(repo, ns, "1", req.Map["_uid"].(int64), images.([]interface{}))
  if err != nil {
    res.Send(err.Error(), 400)
  }

  res.Set("X-Docker-Token", ts)
  res.Set("WWW-Authenticate", "Token " + ts)
  res.Set("X-Docker-Endpoints", "reg22.spacedock.io, reg41.spacedock.io")

  res.Send("Created", 200)
}

func DeleteRepo(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}

func GetUserImage(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}

func RepoAuth(req *f.Request, res *f.Response, next func()) {
  res.Send(200)
}

func UpdateUserImage(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}
