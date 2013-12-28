/*
    Quick JSON configuration parser
*/
package config

import (
  "github.com/stretchr/objx"
  "io/ioutil"
  "log"
  "os"
  "path"
)

var config objx.Map
var gopath = os.Getenv("GOPATH")
var Dir string

func init() {
  // Setup Dir if it's not already set
  if len(Dir) == 0 && len(gopath) > 0 {
    Dir = path.Join(
      gopath,
      "src",
      "github.com",
      "yawnt",
      os.Args[0],
      "config",
    )
  }
}

// Load takes an environment, loads the JSON file associated with the
// environment, and returns an instance of objx.Map for accessing the
// properties.
func Load(env string) (config objx.Map) {
  // Get current directory
  pwd, err := os.Getwd()
  if err != nil {
    log.Fatalln(err)
  }

  // Try reading locally first
  data, localerr := ioutil.ReadFile(path.Join(pwd, "config",
    env + ".config.json"))
  if localerr != nil {
    // Try reading GOPATH next.
    if len(Dir) > 0 {
      data, err = ioutil.ReadFile(path.Join(
        Dir,
        env + ".config.json",
      ))
      if err != nil {
        log.Fatalln(err.Error() + "; " + localerr.Error())
      }
    } else {
      log.Fatalln(localerr.Error() +
        ", and $GOPATH is not defined to determine the config directory.")
    }
  }

  // Convert from JSON to objx.Map
  config, err = objx.FromJSON(string(data))
  if err != nil {
    log.Fatalln(err)
  }

  return
}
