package controllers

import (
	"todolist/api/middlewares"
)

// initializeRoutes sets up the routes for the server
func (s *Server) initializeRoutes() {
 // Create a group for the version 1 of the API
 v1 := s.Router.Group("/api/v1")
 {
  //Users routes
  v1.POST("/login", s.Login)
  v1.POST("/signup", s.CreateUser)
  
  //ToDos routes
  v1.POST("/todos", middlewares.TokenAuthMiddleware(), s.CreateToDo)
  v1.PUT("/todos/:id", middlewares.TokenAuthMiddleware(), s.UpdateToDo)
  v1.DELETE("/todos/:id", middlewares.TokenAuthMiddleware(), s.DeleteToDo)
  v1.GET("/user_todos/:id", middlewares.TokenAuthMiddleware(),s.GetUserToDos)
 }
}