package main

import (
  "encoding/json"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func CreateRepo(req *f.Request, res *f.Response, next func()) {
  images := req.Map["json"]
  u := req.Map["_user"].(*models.User)

  r := &models.Repo{
    Namespace: req.Params["namespace"],
    Name: req.Params["repo"],
  }

  ts, err := r.Create("1", u, images.([]interface{}))
  if err != nil {
    res.Send(err.Error(), 400)
  }

  res.Set("X-Docker-Token", ts)
  res.Set("WWW-Authenticate", "Token " + ts)
  res.Set("X-Docker-Endpoints", "reg22.spacedock.io, reg41.spacedock.io")

  res.Send("Created", 200)
}

func DeleteRepo(req *f.Request, res *f.Response, next func()) {
  namespace := req.Params["namespace"]
  repo := req.Params["repo"]

  r, err := models.GetRepo(namespace, repo)
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

func GetUserImage(req *f.Request, res *f.Response, next func()) {
  repo := req.Params["repo"]
  ns := req.Params["ns"]

  r, err := models.GetRepo(ns, repo)
  if err != nil {
    res.Send(err.Error(), 400)
  }

  images, e := r.GetImages()
  if e != nil {
    res.Send(e.Error(), 400)
  }

  j, jsonErr := json.Marshal(images)
  if jsonErr != nil {
    res.Send("Error returning data", 400)
  }
  res.Send(string(j), 200)
}

func RepoAuth(req *f.Request, res *f.Response, next func()) {
  res.Send(200)
}

func UpdateUserImage(req *f.Request, res *f.Response, next func()) {
  repo := req.Params["repo"]
  ns := req.Params["namespace"]
  json := req.Map["json"].([]interface{})

  r, err := models.GetRepo(ns, repo)
  if err != nil {
    res.Send(err.Error(), 400)
  }

  er := r.UpdateImages(json)
  if er != nil {
    res.Send(err.Error(), 400)
  }

  res.Send("Created", 204)
}
