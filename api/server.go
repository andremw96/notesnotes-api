package api

import (
	db "andre/notesnotes-api/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes
	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser) // id is URI parameter
	router.GET("users", server.listUser)
	router.POST("/insertnote", server.insertNewNote)

	server.router = router
	return server
}

// Start run http server on specifi address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
