package httpg

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/internal/header"
)

var (
	errorHandler func(*gin.Context, error)
	lock         sync.RWMutex
)

// Error writes err into w.
func Error(c *gin.Context, err error) {
	lock.RLock()
	handler := errorHandler
	lock.RUnlock()

	handler(c, err)
	if c.Writer.Written() == false {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

// OkJsonFormat writes v into w with 200 OK.
func OkJsonFormat(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: 0,
		Msg:  "ok",
		Data: v,
	})
}

// SetErrorHandler sets the error handler, which is called on calling Error.
func SetErrorHandler(handler func(*gin.Context, error)) {
	lock.Lock()
	defer lock.Unlock()
	errorHandler = handler
}

// WriteJson writes v as json string into w with code.
func WriteJson(w http.ResponseWriter, code int, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, header.JsonContentType)
	w.WriteHeader(code)

	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
}
