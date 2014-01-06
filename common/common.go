
package common


import (
  "encoding/base64"
  "strings"
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/couch/models"
)

func UnpackAuth(raw []string) (creds []string, err error) {
  auth := strings.Split(raw[0], " ")
  decoded, err := base64.StdEncoding.DecodeString(auth[1])
  if err != nil { return nil, err }

  creds = strings.Split(string(decoded), ":")
  return creds, nil
}

func CheckAuth(req *f.Request, res *f.Response, next func()) {
  creds, err := UnpackAuth(req.Header["Authorization"])
  if err != nil {
    res.Send("Unauthorized", 401)
    return
  }

  result := models.AuthUser(creds[0], creds[1])
  if result != true {
    res.Send("Unauthorized", 401)
  }
}
