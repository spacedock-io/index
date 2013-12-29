/*
    Spacedex, the new and improved Docker index.
 */
package main

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/ricallinson/forgery"
  "github.com/southern/logger"
  "github.com/yawnt/index.spacedock/config"
  "github.com/yawnt/index.spacedock/couch"
  "os"
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

  app.Action = func (c *cli.Context) {
    env := c.String("env")
    dir := c.String("config")

    // Default to dev if someone enters a blank string
    if len(env) == 0 {
      env = "dev"
    }
    if len(dir) > 0 {
      config.Dir = dir
    }

    config.Global = config.Load(env)
    config.Logger = logger.New()

    port := c.Int("port")
    if port == 0 {
      // Bug(Colton): Not quite sure why port is being picked up as Float64 at
      // the moment. Still looking into this. It may be intended functionality.
      port = int(config.Global.Get("port").Float64())
    }

    config.Logger.Debug = config.Global.Get("logger.debug").Bool(false)
    config.Logger.Exit = config.Global.Get("logger.exit").Bool(false)

    couch.Couch = couch.New()
    Routes(server)

    config.Logger.Log("Index listening on port " + fmt.Sprint(port))
    server.Listen(port)
  }

  app.Run(os.Args)
}
