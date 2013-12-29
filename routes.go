package main

import (
  "github.com/ricallinson/forgery"
)

func Routes(server *f.Server) {
  /*
    Library repository routes
   */
  server.Put("/v1/repositories/:repo/auth", LibraryAuth)
  server.Put("/v1/repositories/:repo", CreateLibrary)
  server.Delete("/v1/repositories/:repo", DeleteLibrary)
  server.Put("/v1/repositories/:repo/images", UpdateLibraryImage)
  server.Get("/v1/repositories/:repo/images", GetLibraryImage)

  /*
    User routes
   */
  server.Get("/v1/users", Login)
  server.Post("/v1/users", CreateUser)
  server.Put("/v1/users", UpdateUser)

  /*
    User repository routes
   */
  server.Put("/v1/repositories/:namespace/:repo/auth", RepoAuth)
  server.Put("/v1/repositories/:namespace/:repo", CreateRepo)
  server.Delete("/v1/repositories/:namespace/:repo", DeleteRepo)
  server.Put("/v1/repositories/:namespace/:repo/images", UpdateUserImage)
  server.Get("/v1/repositories/:namespace/:repo/images", GetUserImage)

  // Search route
  server.Get("/v1/search", Search)
}