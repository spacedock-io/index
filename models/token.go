package models

import (
  //"github.com/spacedock-io/registry/db"
  "github.com/gokyle/uuid"
)

type Token struct {
  Id        int64
  Signature string
  Access    string
  UserId    int64
  Repo      string
}

func CreateToken(access, repo string, uid int64) (Token, bool) {
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
