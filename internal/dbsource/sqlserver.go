// internal/dbsource/sqlserver.go
package dbsource

import (
	"context"
	"database/sql"
	"fmt"

	"fatture-sanitize-loader/internal/document"

	_ "github.com/denisenkom/go-mssqldb" // driver SQL Server
)

// SQLServerSource rappresenta la connessione al database SQL Server sorgente
type SQLServerSource struct {
	db *sql.DB
}

// NewSQLServerSource crea una nuova connessione al database SQL Server
// dsn esempio: "sqlserver://user:password@host:port?database=dbname"
func NewSQLServerSource(dsn string) (*SQLServerSource, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, fmt.Errorf("open connection: %w", err)
	}

	// Verifica la connessione
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &SQLServerSource{db: db}, nil
}

// Close chiude la connessione al database
func (s *SQLServerSource) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// NextBatchAttive legge il prossimo batch di fatture attive da elaborare
func (s *SQLServerSource) NextBatchAttive(ctx context.Context, limit int) ([]document.Document, error) {
	const q = `
        SELECT TOP (@p1) Id, Numero, Data, CedenteCF, CessionarioCF, FileXml, FileP7M, IsP7M
        FROM   FattureAttive
        WHERE  DaElaborare = 1
        ORDER BY Id
    `

	rows, err := s.db.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, fmt.Errorf("query batch: %w", err)
	}
	defer rows.Close()

	var docs []document.Document
	for rows.Next() {
		var (
			id             int64
			numero         string
			data           string
			cedenteCF      string
			cessionarioCF  string
			fileXml        []byte
			fileP7M        []byte
			isP7M          bool
		)

		err := rows.Scan(&id, &numero, &data, &cedenteCF, &cessionarioCF, &fileXml, &fileP7M, &isP7M)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		// Determina quale file utilizzare
		var xmlData []byte
		if isP7M && len(fileP7M) > 0 {
			xmlData = fileP7M
		} else {
			xmlData = fileXml
		}

		doc := document.Document{
			Tipo: document.TipoAttiva,
			Meta: document.Meta{
				SourceID:      fmt.Sprintf("%d", id),
				Numero:        numero,
				Data:          data,
				CedenteCFPI:   cedenteCF,
				CessionarioCF: cessionarioCF,
				IsP7M:         isP7M,
			},
			XML: xmlData,
		}

		docs = append(docs, doc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return docs, nil
}

// NextBatchPassive legge il prossimo batch di fatture passive da elaborare
func (s *SQLServerSource) NextBatchPassive(ctx context.Context, limit int) ([]document.Document, error) {
	const q = `
        SELECT TOP (@p1) Id, Numero, Data, CedenteCF, CessionarioCF, FileXml, FileP7M, IsP7M
        FROM   FatturePassive
        WHERE  DaElaborare = 1
        ORDER BY Id
    `

	rows, err := s.db.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, fmt.Errorf("query batch: %w", err)
	}
	defer rows.Close()

	var docs []document.Document
	for rows.Next() {
		var (
			id             int64
			numero         string
			data           string
			cedenteCF      string
			cessionarioCF  string
			fileXml        []byte
			fileP7M        []byte
			isP7M          bool
		)

		err := rows.Scan(&id, &numero, &data, &cedenteCF, &cessionarioCF, &fileXml, &fileP7M, &isP7M)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		// Determina quale file utilizzare
		var xmlData []byte
		if isP7M && len(fileP7M) > 0 {
			xmlData = fileP7M
		} else {
			xmlData = fileXml
		}

		doc := document.Document{
			Tipo: document.TipoPassiva,
			Meta: document.Meta{
				SourceID:      fmt.Sprintf("%d", id),
				Numero:        numero,
				Data:          data,
				CedenteCFPI:   cedenteCF,
				CessionarioCF: cessionarioCF,
				IsP7M:         isP7M,
			},
			XML: xmlData,
		}

		docs = append(docs, doc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return docs, nil
}
