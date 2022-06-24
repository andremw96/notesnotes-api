package api

import (
	db "andre/notesnotes-api/db/sqlc"
	"andre/notesnotes-api/token"
	"andre/notesnotes-api/util"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	log.Print(config.TokenSymmetricKey)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.initializeRouter()
	return server, nil
}

func (server *Server) initializeRouter() {
	router := gin.Default()

	// add routes
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)
	router.GET("/user/:id", server.getUser) // id is URI parameter
	router.GET("users", server.listUser)
	router.POST("/insertnote", server.insertNewNote)

	server.router = router
}

// Start run http server on specifi address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
