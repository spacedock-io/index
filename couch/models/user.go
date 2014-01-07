package models

import (
  "encoding/hex"
  "encoding/json"
  "crypto/sha256"
  "github.com/spacedock-io/index/couch"
  "github.com/fjl/go-couchdb"
  "github.com/gokyle/pbkdf2"
)

type HexString []byte

func (h *HexString) MarshalJSON() ([]byte, error) {
  s := hex.EncodeToString(*h)
  bytes, err := json.Marshal(s)
  return bytes, err
}

func (h *HexString) UnmarshalJSON(data []byte) error {
  var x string
  json.Unmarshal(data, &x)
  s, err := hex.DecodeString(x)
  *h = HexString(s)
  return err
}

var prefix string = "user:"

type User struct {
  Username string `json:"username"`
  Email []string `json:"email"`
  Salt HexString `json:"salt"`
  Pass HexString `json:"pass"`
  Rev string `json:"_rev,omitempty"`
}

func init() {
  pbkdf2.HashFunc = sha256.New
}

func NewUser() *User {
  return &User{}
}

func GetUser(name string) (*User, error) {
  ret := &User{}
  err := couch.Global.Get(prefix + name, nil, ret)

  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && dberr.StatusCode == 404 {
      err = nil
      ret = nil
    }
  }
  return ret, err
}

func (user *User) Create(password string) error {
  user.SetPassword(password)
  return user.Save(false)
}

func (user *User) Save(update bool) error {
  _, err := couch.Global.Put(prefix + user.Username, user)
  if err != nil {
    dberr, ok := err.(couchdb.DatabaseError)
    if ok && !update && dberr.StatusCode == 409 {
      err = AlreadyExistsError{}
    }
  }
  return err
}

func (user *User) Delete() error {
  _, err := couch.Global.Delete(prefix + user.Username, user.Rev)
  return err
}

func (user *User) SetPassword(pass string) {
  ph := pbkdf2.HashPassword(pass)
  user.Salt = ph.Salt
  user.Pass = ph.Hash
}

func (user *User) MatchPassword(pass string) bool {
  ph := &pbkdf2.PasswordHash{ Hash: user.Pass, Salt: user.Salt}
  return pbkdf2.MatchPassword(pass, ph)
}

func AuthUser(user string, pass string) bool {
  ret, err := GetUser(user)
  if err != nil || ret == nil {
    return false;
  }
  return ret.MatchPassword(pass)
}
