package search

const templateSearch = `//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package {{.Package}}

import ({{if .HasImports}}{{range .Imports}}
	"{{.}}"{{end}}
	{{end}}
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// base filters
type applier func(query *orm.Query) (*orm.Query, error)

// Searcher is interface for every generated filter
type Searcher interface {
	Apply(query *orm.Query) *orm.Query
	Q() applier
}

{{range $model := .Entities}}
type {{.GoName}}Search struct {
	{{range .Columns}}
	{{.GoName}} {{.Type}}{{end}}
}

func (s *{{.GoName}}Search) Apply(query *orm.Query) *orm.Query { {{range .Columns}}{{if .Relaxed}}
	if !reflect.ValueOf(s.{{.GoName}}).IsNil(){ {{else}}
	if s.{{.GoName}} != nil { {{end}}{{if .UseWhereRender}}
		query.Where({{.WhereRender}}){{else}} 
		query.Where("?.? = ?", pg.F(Tables.{{$model.GoName}}.{{if not $model.NoAlias}}Alias{{else}}Name{{end}}), pg.F(Columns.{{$model.GoName}}.{{.GoName}}), s.{{.GoName}}){{end}}
	}{{end}}
	
	return query
}

func (s *{{.GoName}}Search) Q() applier {
	return func(query *orm.Query) (*orm.Query, error) {
		return s.Apply(query), nil
	}
}
{{end}}
`
