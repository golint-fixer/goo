/// {{$r := .r}}
/// {{$s := .s}}

package __p__

/// {{if .i}}import "{{.i}}"{{end}}

/// {{if false}}
type __CUSTOM_IFX_FIELD_a_THEN_LBRACES_FIELD_a_RBRACES_PERIOD_END____t__ interface{}
type __s__ struct{}
type __VAR_s__ struct{} /// {{{end}}}

var _ __CUSTOM_IFX_FIELD_a_THEN_LBRACES_FIELD_a_RBRACES_PERIOD_END____t__ = &__s__{}

// {{.d}}
type __CUSTOM_IFEX_FIELD_s_THEN_LBRACES_FIELD_s_RBRACES_ELSE_LBRACES_FIELD_t_RBRACESStub_END__ struct{}

/// {{range .m}}
func (__VAR_r__ *__VAR_s__) __n__( /** {{range $i, $p := .p}}{{if $i}}, {{end}}{{range $j, $n := $p.n}}{{if $j}}, {{end}}{{$n}}{{end}} {{$p.t}}{{end}} **/ ) /** {{if or (gt (len .r) 1) (and .r (index .r 0).n)}}({{end}}{{range $i, $r := .r}}{{if $i}}, {{end}}{{range $j, $n := $r.n}}{{if $j}}, {{end}}{{$n}}{{end}} {{$r.t}}{{end}}{{if or (gt (len .r) 1) (and .r (index .r 0).n)}}){{end}} **/ {
	panic("not implemented")
}

/// {{end}}
