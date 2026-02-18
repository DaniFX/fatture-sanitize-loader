// internal/document/document.go
package document

type Tipo string

const (
	TipoAttiva  Tipo = "ATTIVA"
	TipoPassiva Tipo = "PASSIVA"
)

type Meta struct {
	SourceID      string
	Numero        string
	Data          string
	CedenteCFPI   string
	CessionarioCF string
	FileNameOrig  string
	IsP7M         bool
}

type Document struct {
	Tipo Tipo
	Meta Meta
	XML  []byte // XML grezzo in ingresso, sanificato in uscita
}
