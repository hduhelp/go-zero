package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type (

	// A GinServer is a http server with gin.
	GinServer struct {
		ngin *gin.Engine
		conf RestConf
	}
)

// MustNewGinServer returns a server with given config of c and options defined in opts.
// Be aware that later RunOption might overwrite previous one that write the same option.
// The process will exit if error occurs.
func MustNewGinServer(c RestConf, opts ...RunOption) *GinServer {
	server, err := NewGinServer(c, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return server
}

// NewGinServer returns a server with given config of c and options defined in opts.
// Be aware that later RunOption might overwrite previous one that write the same option.
func NewGinServer(c RestConf, opts ...RunOption) (*GinServer, error) {
	if err := c.SetUp(); err != nil {
		return nil, err
	}

	server := &GinServer{
		ngin: gin.New(),
		conf: c,
	}

	return server, nil
}

// Start starts the GinServer.
// Graceful shutdown is enabled by default.
// Use proc.SetTimeToForceQuit to customize the graceful shutdown period.
func (s *GinServer) Start() {
	handleError(s.ngin.Run(fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port)))
}

// Stop stops the GinServer.
func (s *GinServer) Stop() {
	logx.Close()
}

// Use adds the given middleware in the Server.
func (s *GinServer) Use(middleware gin.HandlerFunc) {
	s.ngin.Use(middleware)
}

func (s *GinServer) GET(relativePath string, handler ...gin.HandlerFunc) {
	s.ngin.GET(relativePath, handler...)
}

func (s *GinServer) POST(relativePath string, handler ...gin.HandlerFunc) {
	s.ngin.POST(relativePath, handler...)
}

func (s *GinServer) DELETE(relativePath string, handler ...gin.HandlerFunc) {
	s.ngin.DELETE(relativePath, handler...)
}

func (s *GinServer) PATCH(relativePath string, handler ...gin.HandlerFunc) {
	s.ngin.PATCH(relativePath, handler...)
}

func (s *GinServer) PUT(relativePath string, handler ...gin.HandlerFunc) {
	s.ngin.PUT(relativePath, handler...)
}

func (s *GinServer) OPTIONS(relativePath string, handler ...gin.HandlerFunc) {
	s.ngin.OPTIONS(relativePath, handler...)
}

func (s *GinServer) HEAD(relativePath string, handler ...gin.HandlerFunc) {
	s.ngin.HEAD(relativePath, handler...)
}
