// internal/sanitize/sanitize_passive.go
package sanitize

import "fatture-sanitize-loader/internal/document"

func Passiva(in document.Document) (document.Document, error) {
	out := in
	// analogo a Attiva ma con regole centrali sul fornitore (CedentePrestatore) [web:29]
	return out, nil
}
