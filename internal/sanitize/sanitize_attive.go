// internal/sanitize/sanitize_attive.go
package sanitize

import (
	"fatture-sanitize-loader/internal/document"
	"fatture-sanitize-loader/internal/p7m"
)

func Attiva(in document.Document) (document.Document, error) {
	out := in

	// 1) Se P7M → estrai XML
	if in.Meta.IsP7M {
		xmlBytes, err := p7m.ExtractXML(in.XML)
		if err != nil {
			return out, err
		}
		out.XML = xmlBytes
	}

	// 2) parse XML e applica regole (da definire nel dettaglio)
	//    es. mascherare CF/PIVA, indirizzi, contatti, IBAN, ecc. [web:29][web:35]

	return out, nil
}
