/// {{$n := .Name}}
/// {{$p := .Pointer}}
/// {{$r := .Receiver}}

package __FIELD_Package__

/// {{if .Interface.Package}}import "{{.Interface.Package}}"{{end}}

/// {{if false}}
type __FIELDS_Interface_Qualifier_ENDFIELDS____X_PERIOD____FIELDS_Interface_Name_ENDFIELDS__ interface{}
type __FIELDS_Interface_Name_ENDFIELDS__ interface{}
type __FIELD_Name__ struct{}
type __X_IF_FIELD_Pointer_THEN_AMPERSAND_ENDIF_ACTION_FIELD_Name_ENDACTION__ struct{}
type __X_IF_VAR_p_THEN_ASTERISK_ENDIF____VAR_n__ struct{} /// {{end}}

/// {{if .Name}}
/// {{if .Interface.Qualifier}}
var _ __FIELDS_Interface_Qualifier_ENDFIELDS____X_PERIOD____FIELDS_Interface_Name_ENDFIELDS__ = __X_IF_FIELD_Pointer_THEN_AMPERSAND_ENDIF_ACTION_FIELD_Name_ENDACTION__{} /// {{else}}
var _ __FIELDS_Interface_Name_ENDFIELDS__ = &__FIELD_Name__{}                                                                                                             /// {{end}} {{end}}

/// {{range .Interface.Methods}}
/// {{range .Doc}}
/// {{.}}{{end}}
func (__VAR_r__ __X_IF_VAR_p_THEN_ASTERISK_ENDIF____VAR_n__) __FIELD_Name__( /** {{range $i, $p := .Params}}{{if $i}}, {{end}}{{range $j, $n := $p.Names}}{{if $j}}, {{end}}{{$n}}{{end}} {{$p.Type}}{{end}} **/ ) /** {{if .Results}}{{if or (gt (len .Results) 1) (index .Results 0).Names}}({{end}}{{end}}{{range $i, $r := .Results}}{{if $i}}, {{end}}{{range $j, $n := $r.Names}}{{if $j}}, {{end}}{{$n}}{{end}} {{$r.Type}}{{end}}{{if .Results}}{{if or (gt (len .Results) 1) (index .Results 0).Names}}){{end}}{{end}} **/ {
	panic("not implemented")
}

/// {{end}}
