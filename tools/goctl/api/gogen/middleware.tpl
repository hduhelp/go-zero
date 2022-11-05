package middleware

import (
    {{if .useGin}}"github.com/gin-gonic/gin"{{else}}"net/http"{{end}}
)

type {{.name}} struct {
}

func New{{.name}}() *{{.name}} {
	return &{{.name}}{}
}

{{if .useGin}}func (m *{{.name}})Handle(c *gin.Context) {
    // TODO generate middleware implement function, delete after code implementation

    // Passthrough to next handler if need
    c.Next()
}{{else}}func (m *{{.name}})Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}{{end}}
