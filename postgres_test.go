package parser_test

import (
	"testing"

	"github.com/pingcap/tidb/pkg/parser"
	"github.com/pingcap/tidb/pkg/parser/ast"
)

func TestPgSelect(t *testing.T) {

	p := parser.New()
	src := `select * from A order by IF(ISNULL(a),1,0),a desc`
	st, err := p.ParseOneStmt(src, "", "")
	if err != nil {
		println(err.Error())
		return
	}
	sel, ok := st.(*ast.SelectStmt)
	if !ok {
		println("error")
	} else {
		println(sel.OrderBy)
	}

}
