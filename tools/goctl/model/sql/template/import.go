package template

const (
	// Imports defines an import template for model in cache case
	Imports = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	{{if .containsPQ}}"github.com/lib/pq"{{end}}
	"github.com/hduhelp/go-zero/core/stores/builder"
	"github.com/hduhelp/go-zero/core/stores/cache"
	"github.com/hduhelp/go-zero/core/stores/sqlc"
	"github.com/hduhelp/go-zero/core/stores/sqlx"
	"github.com/hduhelp/go-zero/core/stringx"
)
`
	// ImportsNoCache defines an import template for model in normal case
	ImportsNoCache = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

    {{if .containsPQ}}"github.com/lib/pq"{{end}}
	"github.com/hduhelp/go-zero/core/stores/builder"
	"github.com/hduhelp/go-zero/core/stores/sqlc"
	"github.com/hduhelp/go-zero/core/stores/sqlx"
	"github.com/hduhelp/go-zero/core/stringx"
)
`
)
