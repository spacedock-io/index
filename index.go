/*
    Spacedex, the new and improved Docker index.
 */
package main

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/ricallinson/forgery"
  "github.com/southern/logger"
  "github.com/stretchr/objx"
  "github.com/yawnt/index.spacedock/config"
  "os"
)

var (
  global objx.Map
  globalLog = logger.New()
)

func main() {
  // Initialize new CLI app
  app := cli.NewApp()

  // Initialize new Forgery server
  server := f.CreateServer()

  app.Name = "spacedex"
  app.Author = "Spacedock"
  app.Email = "hello@spacedock.io"
  app.Usage = "Run a standalone Docker index"
  app.Version = "0.0.1"
  app.Flags = []cli.Flag {
    // No default value here, so that our <env>.config.json file will override
    // it.
    cli.StringFlag{"port, p", "", "Port to listen on"},
    cli.StringFlag{"env, e", "dev", "Default environment"},
    cli.StringFlag{"config, c", "", "Configuration directory"},
  }

  app.Action = func (context *cli.Context) {
    env := context.String("env")
    dir := context.String("config")

    // Default to dev if someone enters a blank string
    if len(env) == 0 {
      env = "dev"
    }
    if len(dir) > 0 {
      config.Dir = dir
    }

    global = config.Load(env)

    port := context.Int("port")
    if port == 0 {
      // Bug(Colton): Not quite sure why port is being picked up as Float64 at
      // the moment. Still looking into this. It may be intended functionality.
      port = int(global.Get("port").Float64())
    }

    globalLog.Debug = global.Get("logger.debug").Bool(false)
    globalLog.Exit = global.Get("logger.exit").Bool(false)

    SetupCouch()
    Routes(server)

    globalLog.Log(logger.INFO, "Index listening on port " + fmt.Sprint(port))
    server.Listen(port)
  }

  app.Run(os.Args)
}
