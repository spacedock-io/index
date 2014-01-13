package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func CreateRepo(req *f.Request, res *f.Response, next func()) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]
  fullname := namespace + "/" + repo

  r := &models.Repo{}
  r.Namespace = namespace
  r.Name = repo
  r.RegistryId = "221"

  // @TODO: make sure this access level is right
  t, ok := models.CreateToken("write", req.Map["_uid"].(int64), fullname)
  if !ok { res.Send("Token error", 400) }

  r.Tokens = append(r.Tokens, t)

  err := r.Create()
  if err != nil {
    res.Send(err.Error(), 400)
  }

  res.Set("X-Docker-Token", t.String())
  res.Set("WWW-Authenticate", "Token " + t.String())
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
  res.Send("Not implemented yet.")
}

func UpdateUserImage(req *f.Request, res *f.Response, next func()) {
  res.Send("Not implemented yet.")
}
