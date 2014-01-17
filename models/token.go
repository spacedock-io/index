package models

import (
  "fmt"
  "github.com/gokyle/uuid"
  "github.com/spacedock-io/registry/db"
  "strings"
)

type Token struct {
  Id        int64
  Signature string
  Access    string
  UserId    int64
  RepoId    int64
  Repo      string
}

func CreateToken(access string, uid int64, repo string) (Token, bool) {
  token := Token{}
  // @TODO: Validate access string
  token.Access = access
  sig, err := uuid.GenerateV4String()
  if err != nil {
    return Token{}, false
  }
  token.Signature = sig
  token.UserId = uid
  token.Repo = repo
  return token, true
}

func GetToken(user *User, repo, access string) (Token, error) {
  t := Token{
    UserId: user.Id,
    Access: access,
    Repo: repo,
  }

  q := db.DB.Where(&t).Find(&t)
  if q.RecordNotFound() {
    return Token{}, TokenNotFound
  } else if q.Error != nil {
    return Token{}, q.Error
  }

  return t, nil
}

func GetTokenString(token string) (Token, error) {
  t := Token{}

  split := strings.Split(token, ",")
  for _, v := range split {
    v := strings.Split(v, "=")
    switch v[0] {
      case "signature": t.Signature = v[1]
      case "repository": t.Repo = v[1]
      case "access": t.Access = v[1]
    }
  }

  q := db.DB.Where(&t).Find(&t)
  if q.RecordNotFound() {
    return Token{}, TokenNotFound
  } else if q.Error != nil {
    return Token{}, q.Error
  }
  return t, nil
}

func (token *Token) String() string {
  return fmt.Sprintf("signature=%s,repository=%s,access=%s", token.Signature,
    token.Repo, token.Access)
}
