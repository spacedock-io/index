package models

import (
  "fmt"
  "github.com/gokyle/uuid"
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

func (token *Token) String() string {
  return fmt.Sprintf("signature=%s,repository=%s,access=%s", token.Signature,
    token.Repo, token.Access)
}
