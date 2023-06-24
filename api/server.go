package api

import (
	"fmt"

	db "github.com/foyez/simplebank/db/sqlc"
	"github.com/foyez/simplebank/token"
	"github.com/foyez/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

// setupRouter setups the routers
func (server *Server) setupRouter() {
	router := gin.Default()

	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	authRouter.PATCH("/users/:username", server.updateUser)

	authRouter.POST("/accounts", server.createAccount)
	authRouter.GET("/accounts/:id", server.getAccount)
	authRouter.GET("/accounts", server.listAccounts)
	authRouter.GET("/accountsWithCursor", server.listAccountsWithCursor)
	authRouter.PUT("/accounts/:id", server.updateAccount)
	authRouter.DELETE("/accounts/:id", server.deleteAccount)

	authRouter.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
