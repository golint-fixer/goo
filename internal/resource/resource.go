package __FIELD_PackageName__

import "github.com/willfaught/goo"

/// {{if false}}
var __FIELD_Compressed__ bool
var __FIELD_Data__ []byte /// {{end}}

func init() {
	goo.AddResource("{{.PackagePath}}", "{{.Name}}", __FIELD_Data__, __FIELD_Compressed__)
}
