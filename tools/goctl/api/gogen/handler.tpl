package {{.PkgName}}

import (
	{{if .UseGin}}"github.com/gin-gonic/gin"{{else}}"net/http"{{end}}

	"github.com/hduhelp/go-zero/rest/{{if .UseGin}}httpg{{else}}httpx{{end}}"
	{{.ImportPackages}}
)

{{if .UseGin}}func {{.HandlerName}}(svcCtx *svc.ServiceContext) gin.HandlerFunc {
    return func(c *gin.Context) {
        {{if .HasRequest}}var req types.{{.RequestType}}
        if err := c.ShouldBind(&req); err != nil {
            httpg.Error(c, err)
            return
        }

        {{end}}l := {{.LogicName}}.New{{.LogicType}}(c, svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        if err != nil {
            httpg.Error(c, err)
        } else {
            {{if .HasResp}}httpg.OkJsonFormat(c, resp){{else}}httpg.OkJsonFormat(c, nil){{end}}
        }
    }
}{{else}}func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        {{if .HasRequest}}var req types.{{.RequestType}}
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }

        {{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        if err != nil {
            httpx.Error(w, err)
        } else {
            {{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
        }
    }
}{{end}}
