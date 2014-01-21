
package common


import (
  "encoding/base64"
  "strings"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/models"
)

func BasicAuth(req *f.Request, res *f.Response, next func()) {
  auth := req.Get("authorization")
  req.Map["_user"] = nil

  if len(auth) == 0 {
    res.Send("No authorization provided.", 401)
    return
  }

  u, err := HandleBasic(auth)
  if err != nil {
    res.Send("Unauthorized", 401)
    return
  }

  req.Map["_user"] = u
}

func TokenAuth(req *f.Request, res *f.Response, next func()) {
  auth := req.Get("authorization")
  if len(auth) == 0 {
    res.Send("No authorization provided.", 403)
    return
  }

  _, err := HandleToken(auth)
  if err != nil {
    if err == models.TokenNotFound {
      res.Send(err.Error(), 404)
      return
    }
    res.Send(err.Error(), 403)
    return
  }
}

func HandleBasic(raw string) ((*models.User), error) {
  auth := strings.Split(raw, " ")
  decoded, err := base64.StdEncoding.DecodeString(auth[1])
  if err != nil { return nil, err }

  creds := strings.Split(string(decoded), ":")
  u, ok := models.AuthUser(creds[0], creds[1])
  if !ok {
    return nil, models.AuthErr
  }
  return u, nil
}

func HandleToken(raw string) (models.Token, error) {
  auth := strings.Split(raw, " ")
  return models.UseTokenString(auth[1])
}
