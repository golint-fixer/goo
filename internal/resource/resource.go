package __FIELD_Package__

import "github.com/willfaught/goo"

/// {{if false}}
const __FIELD_Compressed__ = false /// {{end}}

func init() {
	goo.SetResource("{{.Name}}", []byte{ /** {{.Data}} **/ }, __FIELD_Compressed__)
}
