/// {{$r := .Receiver}}
/// {{$s := .Name}}

package __FIELD_Package__

/// {{if .Interface.Package}}import "{{.Interface.Package}}"{{end}}

/// {{if false}}
type __X_IF_FIELDS_Interface_Qualifier_ENDFIELDS_THEN_ACTION_FIELDS_Interface_Qualifier_ENDFIELDS_ENDACTION_PERIOD_ENDIF_ACTION_FIELDS_Interface_Name_ENDFIELDS_ENDACTION__ interface{}
type __X_IF_VAR_r_THEN_ACTION_VAR_r_ENDACTION_ENDIF_ASTERISK_ACTION_VAR_s_ENDACTION__ struct{} /// {{end}}

var _ __X_IF_FIELDS_Interface_Qualifier_ENDFIELDS_THEN_ACTION_FIELDS_Interface_Qualifier_ENDFIELDS_ENDACTION_PERIOD_ENDIF_ACTION_FIELDS_Interface_Name_ENDFIELDS_ENDACTION__ = &__FIELD_Name__{}

/// {{range .Doc}}
/// {{.}}{{end}}
type __FIELD_Name__ struct{}

/// {{range .Interface.Methods}}
/// {{range .Doc}}
/// {{.}}{{end}}
func (__X_IF_VAR_r_THEN_ACTION_VAR_r_ENDACTION_ENDIF_ASTERISK_ACTION_VAR_s_ENDACTION__) __FIELD_Name__( /** {{range $i, $p := .Params}}{{if $i}}, {{end}}{{range $j, $n := $p.Names}}{{if $j}}, {{end}}{{$n}}{{end}} {{$p.Type}}{{end}} **/ ) /** {{if or (gt (len .Results) 1) (and .Results (index .Results 0).Names)}}({{end}}{{range $i, $r := .Results}}{{if $i}}, {{end}}{{range $j, $n := $r.Names}}{{if $j}}, {{end}}{{$n}}{{end}} {{$r.Type}}{{end}}{{if or (gt (len .Results) 1) (and .Results (index .Results 0).Names)}}){{end}} **/ {
	panic("not implemented")
}

/// {{end}}
