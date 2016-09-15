package macro

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"strings"
)

//goo:case
//goo:default
//goo:else
//goo:elseif
//goo:elseifnot
//goo:end
//goo:for
//goo:if
//goo:ifnot
//goo:write
//goo:omit
//goo:switch

var (
	ifre    = regexp.MustCompile(`^//goo:if ([\w_]+)$`)
	ifvalre = regexp.MustCompile(`^//goo:if ([\w_]+)( ([\w_]+|"[^"\r\n]"))?$`)
	ifnotre = regexp.MustCompile(`^//goo:ifnot ([\w_]+)$`)
)

/*func omit(cm ast.CommentMap, n ast.Node) bool {
	for _, c := range cm[n].List {
		if c.Text == "goo:omit" {
			return true
		}
	}

	return false
}

func vars(cm ast.CommentMap, n ast.Node, r *regexp.Regexp) []string {
	var cg, ok = cm[n]

	if !ok {
		return nil
	}

	var vars []string

	for _, c := range cg.List {
		if ss := r.FindStringSubmatch(c.Text); len(ss) > 1 {
			vars = append(vars, ss[1])
		}
	}

	return vars
}

func ifvars(cm ast.CommentMap, n ast.Node) []string {
	return vars(cm, n, ifre)
}

func ifnotvars(cm ast.CommentMap, n ast.Node) []string {
	return vars(cm, n, ifnotre)
}*/

func Generate(fs *token.FileSet, f *ast.File, vars map[string][]string) error {
	//var cm = ast.NewCommentMap(fs, f, f.Comments)

	// strip go generate

	// vars

	var reps []string

	for k, vs := range vars {
		if len(vs) == 0 {
			continue
		}

		var v = vs[len(vs)-1]
		var l = fmt.Sprintf("_%v", k)
		var r = fmt.Sprintf("%v_", k)
		var lr = fmt.Sprintf("_%v_", k)

		reps = append(reps, k, v, l, v, r, v, lr, v)
	}

	var repl = strings.NewReplacer(reps...)

	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		switch n := n.(type) {
		case *ast.BasicLit:
			if n.Kind == token.STRING {
				n.Value = repl.Replace(n.Value)
			}

		case *ast.Comment:
			fmt.Println("COM", n.Text, repl.Replace(n.Text))
			n.Text = repl.Replace(n.Text)

		case *ast.Ident:
			n.Name = repl.Replace(n.Name)
		}

		return true
	})

	// decls

	/*for _, n := range f.Decls {

	}*/

	//BlockStmt.List (Stmt)

	return nil
}
