// internal/anagrafica/http_client.go
package anagrafica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"fatture-sanitize-loader/internal/document"
)

type Client struct {
	baseURL string
	http    *http.Client
	cache   map[string]int64 // CF/PIVA -> ID interno
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		http:    &http.Client{},
		cache:   make(map[string]int64),
	}
}

type ensureResp struct {
	Id    int64  `json:"Id"`
	Stato string `json:"Stato"`
}

type emittenteReq struct {
	CodiceFiscale  string `json:"CodiceFiscale"`
	PartitaIVA     string `json:"PartitaIVA"`
	RagioneSociale string `json:"RagioneSociale"`
}

func (c *Client) EnsureEmittente(doc document.Document) (int64, error) {
	key := doc.Meta.CedenteCFPI
	if id, ok := c.cache[key]; ok {
		return id, nil
	}

	reqBody := emittenteReq{
		CodiceFiscale:  key,
		PartitaIVA:     key, // da raffinare
		RagioneSociale: "",  // da estrarre da XML se ti serve
	}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("%s/api/Anagrafica/EnsureEmittente", c.baseURL)
	resp, err := c.http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return 0, fmt.Errorf("EnsureEmittente status %d", resp.StatusCode)
	}

	var r ensureResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}
	c.cache[key] = r.Id
	return r.Id, nil
}

func (c *Client) EnsureFornitore(doc document.Document) (int64, error) {
	key := doc.Meta.CedenteCFPI
	if id, ok := c.cache[key]; ok {
		return id, nil
	}
	// stesso pattern, chiamando /api/Anagrafica/EnsureFornitore
	return 0, nil
}
