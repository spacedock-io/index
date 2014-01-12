package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func CreateRepo(req *f.Request, res *f.Response, next func()) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]

  r := &models.Repo{}
  r.Namespace = namespace
  r.Name = repo
  r.RegistryId = "221"

  // Token stuff would go here

  err := r.Create()
  if err != nil {
    res.Send(err.Error(), 400)
  }

  res.Set("X-Docker-Token", "token string value")
  res.Set("WWW-Authenticate", "Token " + "token string value")

  res.Set("X-Docker-Endpoints", "reg22.spacedock.io, reg41.spacedock.io")
  res.Send("Not implemented yet.")
}

func DeleteRepo(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}

func GetUserImage(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}

func RepoAuth(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}

func UpdateUserImage(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}
