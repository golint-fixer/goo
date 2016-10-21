/// {{$n := .Name}}
/// {{$p := .Pointer}}
/// {{$r := .Receiver}}

package __FIELD_Package__

/// {{if .Interface.Package}}import "{{.Interface.Package}}"{{end}}

/// {{if false}}
type __FIELDS_Interface_Qualifier_ENDFIELDS____X_PERIOD____FIELDS_Interface_Name_ENDFIELDS__ interface{}
type __FIELDS_Interface_Name_ENDFIELDS__ interface{}
type __X_IF_FIELD_Pointer_THEN_AMPERSAND_ENDIF_ACTION_FIELD_Name_ENDACTION__ struct{}
type __X_SPACE_ACTION_FIELDS_Interface_Name_ENDFIELDS_ENDACTION__ struct{}
type __FIELD_Name__ struct {
	__X_SPACE_ACTION_FIELDS_Interface_Name_ENDFIELDS_ENDACTION__
} /// {{end}}

/// {{if false}}
type __X_IF_VAR_p_THEN_ASTERISK_ENDIF____VAR_n__ struct{}

func (__X_IF_VAR_p_THEN_ASTERISK_ENDIF____VAR_n__) Called() int { return 0 } /// {{end}}

/// {{if .Name}}
/// {{if .Interface.Qualifier}}
var _ __FIELDS_Interface_Qualifier_ENDFIELDS____X_PERIOD____FIELDS_Interface_Name_ENDFIELDS__ = __X_IF_FIELD_Pointer_THEN_AMPERSAND_ENDIF_ACTION_FIELD_Name_ENDACTION__{} /// {{else}}
var _ __FIELDS_Interface_Name_ENDFIELDS__ = &__FIELD_Name__{}                                                                                                             /// {{end}} {{end}}

func New__FIELD_Name__(__FIELDS_Interface_Initial_ENDFIELDS__ __X_SPACE_ACTION_FIELDS_Interface_Name_ENDFIELDS_ENDACTION__) __FIELDS_Interface_Name_ENDFIELDS__ {
	return &__FIELD_Name__{__X_SPACE_ACTION_FIELDS_Interface_Name_ENDFIELDS_ENDACTION__: __FIELDS_Interface_Initial_ENDFIELDS__}
}

/// {{define "params"}} {{range $i, $p := .ParamsGrouped}} {{if $i}}, {{end}} {{range $j, $n := $p.Names}} {{if $j}} , {{end}} {{$n}} {{end}} {{$p.Type}} {{end}} {{end}}
/// {{define "results"}} {{if .ResultsGrouped}} {{if or (gt (len .ResultsGrouped) 1) (index .ResultsGrouped 0).Names}} ( {{end}} {{end}} {{range $i, $r := .ResultsGrouped}} {{if $i}}, {{end}} {{range $j, $n := $r.Names}} {{if $j}} , {{end}} {{$n}} {{end}} {{$r.Type}} {{end}} {{if .ResultsGrouped}} {{if or (gt (len .ResultsGrouped) 1) (index .ResultsGrouped 0).Names}} ) {{end}} {{end}} {{end}}
/// {{define "args"}} {{range $i, $p := .ParamsFlat}} {{if $i}} , {{end}} {{$p.Name}} {{end}} {{end}}
/// {{define "body-none"}} {{.Receiver}} . {{.Interface.Name}} . {{.Name}} ( {{template "args" .}} ) {{end}}
/// {{define "body-many"}} return {{.Receiver}} . {{.Interface.Name}} . {{.Name}} ( {{template "args" .}} ) {{end}}
/// {{define "body"}} {{if .ResultsFlat}} {{template "body-many" .}} {{else}} {{template "body-none" .}} {{end}} {{end}}

/// {{range .Interface.Methods}}
/// {{range .Doc}}
/// {{.}}{{end}}
func (__VAR_r__ __X_IF_VAR_p_THEN_ASTERISK_ENDIF____VAR_n__) __FIELD_Name__( /** {{template "params" .}} **/ ) /** {{template "results" .}} **/ {
	/// {{template "body" .}}
}

/// {{end}}
