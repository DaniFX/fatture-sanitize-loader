// internal/dbsource/sqlserver.go
func (s *SQLServerSource) NextBatchAttive(ctx context.Context, limit int) ([]document.Document, error) {
	const q = `
        SELECT TOP (@p1) Id, Numero, Data, CedenteCF, CessionarioCF, FileXml, FileP7M, IsP7M
        FROM   FattureAttive
        WHERE  DaElaborare = 1
        ORDER BY Id
    ` // driver sqlserver usato con database/sql [web:20][web:24]
	// ...
}
