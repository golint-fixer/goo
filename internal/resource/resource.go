package __FIELD_Package__

import "github.com/willfaught/goo"

/// {{if false}}
var __FIELD_Compressed__ bool
var __FIELD_Data__ []byte /// {{end}}

func init() {
	goo.SetResource("{{.Name}}", __FIELD_Data__, __FIELD_Compressed__)
}
