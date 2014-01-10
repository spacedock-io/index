
package common


import (
  "encoding/base64"
  "strings"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/couch/models"
)

func UnpackAuth(raw string) (creds []string, err error) {
  auth := strings.Split(raw, " ")
  decoded, err := base64.StdEncoding.DecodeString(auth[1])
  if err != nil { return nil, err }

  creds = strings.Split(string(decoded), ":")
  return creds, nil
}

func CheckAuth(req *f.Request, res *f.Response, next func()) {
  auth := req.Get("authorization")

  if len(auth) == 0 {
    res.Send("No authorization provided.", 401)
    return
  }

  creds, err := UnpackAuth(auth)
  if err != nil {
    res.Send("Unauthorized", 401)
    return
  }

  result := models.AuthUser(creds[0], creds[1])
  if result != true {
    res.Send("Unauthorized", 401)
  }
}
