package common

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
  // "github.com/spacedock-io/registry/db"
  "strings"
)

var AccessMap = map[string]int{
  "none": 0,
  "read": 1,
  "write": 2,
  "delete": 3,
}

func wantsToken(req *f.Request) bool {
  return strings.ToLower(strings.Trim(req.Get("x-docker-token"), " ")) == "true"
}

func hasAccess(user *models.User, ns, repo, access string) bool {
  return AccessMap[access] <= AccessMap[user.GetAccess(ns, repo)]
}

func DeleteAccess(req *f.Request, res *f.Response, next func()) {
  user := req.Map["_user"].(*models.User)
  ns := req.Params["namespace"]
  repo := req.Params["repo"]

  ok := hasAccess(user, ns, repo, "delete")
  if !ok {
    res.Send("You do not have access to delete this repository.", 400)
  }

  if wantsToken(req) {
    // Create the token and return it
  }
}

func ReadAccess(req *f.Request, res *f.Response, next func()) {
  user := req.Map["_user"].(*models.User)
  ns := req.Params["namespace"]
  repo := req.Params["repo"]

  ok := hasAccess(user, ns, repo, "read")
  if !ok {
    res.Send("You do not have access to read from this repository.", 400)
  }

  if wantsToken(req) {
    // Create the token and return it
  }
}

func WriteAccess(req *f.Request, res *f.Response, next func()) {
  user := req.Map["_user"].(*models.User)
  ns := req.Params["namespace"]
  repo := req.Params["repo"]

  ok := hasAccess(user, ns, repo, "write")
  if !ok {
    res.Send("You do not have access to write to this repository.", 400)
    return
  }

  if len(ns) > 0 {
    repo = ns + "/" + repo
  }

  if wantsToken(req) {
    // Create the token and return it
  }
}
