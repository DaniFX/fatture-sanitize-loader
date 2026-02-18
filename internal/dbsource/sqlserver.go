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
// Le fatture attive sono nella tabella PA_Storico con TipoFlusso = 'ATTIVO'
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
// Le fatture passive sono nella tabella B2B_Storico_RCV
// NOTA: I file XML/P7M sono salvati su filesystem, non come BLOB nel database
func (s *SQLServerSource) NextBatchPassive(ctx context.Context, limit int) ([]document.Document, error) {
	const q = `
        SELECT TOP (@p1) 
            Id, 
            NumeroDocumento, 
            DataDocumento, 
            PIVAFornitore,
            CodicefiscaleFornitore,
            Fornitore,
            NomeFile,
            NomeFileArchivio,
            TipoFirma
        FROM   B2B_Storico_RCV
        WHERE  GlobeRestApiExported = 0
        ORDER BY Id
    `

	rows, err := s.db.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, fmt.Errorf("query batch passive: %w", err)
	}
	defer rows.Close()

	var docs []document.Document
	for rows.Next() {
		var (
			id                      int64
			numeroDocumento         sql.NullString
			dataDocumento           sql.NullTime
			pivaFornitore           sql.NullString
			codicefiscaleFornitore  sql.NullString
			fornitore               sql.NullString
			nomeFile                sql.NullString
			nomeFileArchivio        sql.NullString
			tipoFirma               sql.NullString
		)

		err := rows.Scan(
			&id,
			&numeroDocumento,
			&dataDocumento,
			&pivaFornitore,
			&codicefiscaleFornitore,
			&fornitore,
			&nomeFile,
			&nomeFileArchivio,
			&tipoFirma,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		// Per le passive, i file sono su filesystem
		// NomeFileArchivio contiene il path del file
		var fileName string
		if nomeFileArchivio.Valid && nomeFileArchivio.String != "" {
			fileName = nomeFileArchivio.String
		} else if nomeFile.Valid {
			fileName = nomeFile.String
		}

		// Determina se è un file P7M
		isP7M := false
		if tipoFirma.Valid && tipoFirma.String == "P7M" {
			isP7M = true
		}

		// Formatta la data
		var dataStr string
		if dataDocumento.Valid {
			dataStr = dataDocumento.Time.Format("2006-01-02")
		}

		// Per le passive, usa PIVA se disponibile, altrimenti CF
		cedenteCFPI := pivaFornitore.String
		if cedenteCFPI == "" {
			cedenteCFPI = codicefiscaleFornitore.String
		}

		doc := document.Document{
			Tipo: document.TipoPassiva,
			Meta: document.Meta{
				SourceID:      fmt.Sprintf("%d", id),
				Numero:        numeroDocumento.String,
				Data:          dataStr,
				CedenteCFPI:   cedenteCFPI,            // Per le passive, il cedente è il fornitore
				CessionarioCF: "",                     // Il cessionario siamo noi (da configurazione)
				FileNameOrig:  fileName,
				IsP7M:         isP7M,
			},
			XML: nil, // Per le passive, XML deve essere letto da filesystem
		}

		docs = append(docs, doc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return docs, nil
}

// MarkAsProcessedAttive segna una fattura attiva come elaborata
func (s *SQLServerSource) MarkAsProcessedAttive(ctx context.Context, sourceID string) error {
	const q = `
        UPDATE PA_Storico 
        SET GlobeRestApiExported = 1, 
            DataEstrazione = @p1
        WHERE Chiave = @p2
    `

	_, err := s.db.ExecContext(ctx, q, time.Now(), sourceID)
	if err != nil {
		return fmt.Errorf("mark attiva as processed: %w", err)
	}

	return nil
}

// MarkAsProcessedPassive segna una fattura passiva come elaborata
func (s *SQLServerSource) MarkAsProcessedPassive(ctx context.Context, sourceID string) error {
	const q = `
        UPDATE B2B_Storico_RCV 
        SET GlobeRestApiExported = 1, 
            DataEstrazione = @p1
        WHERE Id = @p2
    `

	_, err := s.db.ExecContext(ctx, q, time.Now(), sourceID)
	if err != nil {
		return fmt.Errorf("mark passiva as processed: %w", err)
	}

	return nil
}
