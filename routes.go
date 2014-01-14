package main

import (
  "github.com/ricallinson/forgery"
  "github.com/spacedock-io/index/common"
)

func Routes(server *f.Server) {
  /*
    Stuff needed for private index
   */
  server.Get("/v1/_ping", func(req *f.Request, res *f.Response, next func()) {
    res.SetHeader("X-Docker-Registry-Standalone", "False")
    res.Send("true")
  })

  /*
    Library repository routes
   */
  server.Put("/v1/repositories/:repo/auth/", common.TokenAuth, LibraryAuth)
  server.Put("/v1/repositories/:repo/", common.BasicAuth, CreateLibrary)
  server.Delete("/v1/repositories/:repo/", common.BasicAuth, DeleteLibrary)
  server.Put("/v1/repositories/:repo/images/", common.BasicAuth, UpdateLibraryImage)
  server.Get("/v1/repositories/:repo/images/", GetLibraryImage)

  /*
    User routes
   */
  server.Get("/v1/users/", common.BasicAuth, Login)
  server.Post("/v1/users/", CreateUser)
  server.Put("/v1/users/:username/", common.BasicAuth, UpdateUser)

  /*
    User repository routes
   */
  server.Put("/v1/repositories/:namespace/:repo/auth/", common.TokenAuth, RepoAuth)
  server.Put("/v1/repositories/:namespace/:repo/", common.BasicAuth, CreateRepo)
  server.Delete("/v1/repositories/:namespace/:repo/", common.BasicAuth, DeleteRepo)
  server.Put("/v1/repositories/:namespace/:repo/images/", common.BasicAuth, UpdateUserImage)
  server.Get("/v1/repositories/:namespace/:repo/images/", GetUserImage)

  // Search route
  server.Get("/v1/search/", Search)
}
