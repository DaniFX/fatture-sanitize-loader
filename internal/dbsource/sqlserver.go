// internal/dbsource/sqlserver.go
package dbsource

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
        SELECT TOP (@p1) 
            Chiave, 
            NrDoc, 
            DataDoc, 
            PivaCommittente,
            Committente,
            FileXML, 
            FileXMLFirmato, 
            NomeFileXML,
            NomeFileXMLFirmato
        FROM   PA_Storico
        WHERE  (GlobeRestApiExported = 0 OR GlobeRestApiExported IS NULL)
        AND    TipoFlusso = 'ATTIVO'
        ORDER BY Chiave
    `

	rows, err := s.db.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, fmt.Errorf("query batch attive: %w", err)
	}
	defer rows.Close()

	var docs []document.Document
	for rows.Next() {
		var (
			chiave              int64
			nrDoc               sql.NullString
			dataDoc             sql.NullTime
			pivaCommittente     sql.NullString
			committente         sql.NullString
			fileXML             []byte
			fileXMLFirmato      []byte
			nomeFileXML         sql.NullString
			nomeFileXMLFirmato  sql.NullString
		)

		err := rows.Scan(
			&chiave,
			&nrDoc,
			&dataDoc,
			&pivaCommittente,
			&committente,
			&fileXML,
			&fileXMLFirmato,
			&nomeFileXML,
			&nomeFileXMLFirmato,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		// Determina quale file utilizzare (P7M ha priorità)
		var xmlData []byte
		var isP7M bool
		var fileName string

		if len(fileXMLFirmato) > 0 {
			xmlData = fileXMLFirmato
			isP7M = true
			if nomeFileXMLFirmato.Valid {
				fileName = nomeFileXMLFirmato.String
			}
		} else {
			xmlData = fileXML
			isP7M = false
			if nomeFileXML.Valid {
				fileName = nomeFileXML.String
			}
		}

		// Formatta la data
		var dataStr string
		if dataDoc.Valid {
			dataStr = dataDoc.Time.Format("2006-01-02")
		}

		doc := document.Document{
			Tipo: document.TipoAttiva,
			Meta: document.Meta{
				SourceID:      fmt.Sprintf("%d", chiave),
				Numero:        nrDoc.String,
				Data:          dataStr,
				CedenteCFPI:   "",                    // Per le attive, il cedente siamo noi (da configurazione)
				CessionarioCF: pivaCommittente.String, // Il committente è il cessionario
				FileNameOrig:  fileName,
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
        SELECT TOP (@p1) 
            Chiave, 
            NrDoc, 
            DataDoc, 
            PivaCommittente,
            Committente,
            FileXML, 
            FileXMLFirmato, 
            NomeFileXML,
            NomeFileXMLFirmato
        FROM   PA_Storico
        WHERE  (GlobeRestApiExported = 0 OR GlobeRestApiExported IS NULL)
        AND    TipoFlusso = 'PASSIVO'
        ORDER BY Chiave
    `

	rows, err := s.db.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, fmt.Errorf("query batch passive: %w", err)
	}
	defer rows.Close()

	var docs []document.Document
	for rows.Next() {
		var (
			chiave              int64
			nrDoc               sql.NullString
			dataDoc             sql.NullTime
			pivaCommittente     sql.NullString
			committente         sql.NullString
			fileXML             []byte
			fileXMLFirmato      []byte
			nomeFileXML         sql.NullString
			nomeFileXMLFirmato  sql.NullString
		)

		err := rows.Scan(
			&chiave,
			&nrDoc,
			&dataDoc,
			&pivaCommittente,
			&committente,
			&fileXML,
			&fileXMLFirmato,
			&nomeFileXML,
			&nomeFileXMLFirmato,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		// Determina quale file utilizzare (P7M ha priorità)
		var xmlData []byte
		var isP7M bool
		var fileName string

		if len(fileXMLFirmato) > 0 {
			xmlData = fileXMLFirmato
			isP7M = true
			if nomeFileXMLFirmato.Valid {
				fileName = nomeFileXMLFirmato.String
			}
		} else {
			xmlData = fileXML
			isP7M = false
			if nomeFileXML.Valid {
				fileName = nomeFileXML.String
			}
		}

		// Formatta la data
		var dataStr string
		if dataDoc.Valid {
			dataStr = dataDoc.Time.Format("2006-01-02")
		}

		doc := document.Document{
			Tipo: document.TipoPassiva,
			Meta: document.Meta{
				SourceID:      fmt.Sprintf("%d", chiave),
				Numero:        nrDoc.String,
				Data:          dataStr,
				CedenteCFPI:   pivaCommittente.String, // Per le passive, il cedente è il fornitore
				CessionarioCF: "",                     // Il cessionario siamo noi (da configurazione)
				FileNameOrig:  fileName,
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

// MarkAsProcessed segna un documento come elaborato nel database
func (s *SQLServerSource) MarkAsProcessed(ctx context.Context, sourceID string) error {
	const q = `
        UPDATE PA_Storico 
        SET GlobeRestApiExported = 1, 
            DataEstrazione = @p1
        WHERE Chiave = @p2
    `

	_, err := s.db.ExecContext(ctx, q, time.Now(), sourceID)
	if err != nil {
		return fmt.Errorf("mark as processed: %w", err)
	}

	return nil
}
