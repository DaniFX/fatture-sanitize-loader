// internal/sink/http_sink.go
package sink

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"fatture-sanitize-loader/internal/document"
)

type HttpSink struct {
	baseURL string
	client  *http.Client
}

func NewHttpSink(baseURL string) *HttpSink {
	return &HttpSink{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

type importFatturaReq struct {
	XmlBase64     string `json:"XmlBase64"`
	Tipo          string `json:"Tipo"`
	CedenteCFPI   string `json:"CedenteCFPI"`
	CessionarioCF string `json:"CessionarioCF"`
	Numero        string `json:"Numero"`
	Data          string `json:"Data"`
}

func (s *HttpSink) WriteAttiva(doc document.Document) error {
	return s.postImport("/api/Fatture/ImportAttiva", doc)
}

func (s *HttpSink) WritePassiva(doc document.Document) error {
	return s.postImport("/api/Fatture/ImportPassiva", doc)
}

func (s *HttpSink) postImport(path string, doc document.Document) error {
	req := importFatturaReq{
		XmlBase64:     base64.StdEncoding.EncodeToString(doc.XML),
		Tipo:          string(doc.Tipo),
		CedenteCFPI:   doc.Meta.CedenteCFPI,
		CessionarioCF: doc.Meta.CessionarioCF,
		Numero:        doc.Meta.Numero,
		Data:          doc.Meta.Data,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url := s.baseURL + path
	resp, err := s.client.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("import %s status %d", path, resp.StatusCode)
	}
	return nil
}
