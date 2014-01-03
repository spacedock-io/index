/*
  Setup global test variables
 */
package models

import (
  "github.com/spacedock-io/index/config"
  "github.com/spacedock-io/index/couch"
)

func init() {
  couch.Global = couch.New(config.Load("test"))
}
