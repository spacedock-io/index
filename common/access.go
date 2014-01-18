package common

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
  "strings"
)

var AccessMap = map[string]int{
  "none": 0,
  "read": 1,
  "write": 2,
  "delete": 3,
}

func sendToken(req *f.Request, res *f.Response, access string) {
  user := req.Map["_user"].(*models.User)
  ns := req.Params["namespace"]
  repo := req.Params["repo"]

  ok := hasAccess(user, ns, repo, access)
  if !ok {
    res.Send("You do not have access to perform this action.", 400)
  }

  if len(ns) == 0 {
    ns = "library"
  }

  repo = ns + "/" + repo

  if wantsToken(req) {
    token, err := models.GetToken(user, repo, access)
    if err != nil {
      res.Send(err.Error(), 400)
      return
    }
    res.Set("x-docker-token", token.String())
    res.Set("www-authenticate", "Token " + token.String())
  }
}

func wantsToken(req *f.Request) bool {
  return strings.ToLower(strings.Trim(req.Get("x-docker-token"), " ")) == "true"
}

func hasAccess(user *models.User, ns, repo, access string) bool {
  return AccessMap[access] <= AccessMap[user.GetAccess(ns, repo)]
}

func Access(access string) func(*f.Request, *f.Response, func()) {
  return func(req *f.Request, res *f.Response, next func()) {
    sendToken(req, res, access)
  }
}
