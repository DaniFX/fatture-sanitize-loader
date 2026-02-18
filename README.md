Struttura repository 
fatture-sanitize-loader/
  cmd/
    attive/
      main.go
    passive/
      main.go
  internal/
    config/
      config.go
    dbsource/
      sqlserver.go      # lettura da SQL Server (attive/passive)
    storage/
      files.go          # lettura/scrittura file da/verso filesystem
    document/
      document.go       # modello comune fattura + helpers
    p7m/
      p7m.go            # estrazione XML da .p7m
    sanitize/
      sanitize_attive.go
      sanitize_passive.go
    anagrafica/
      anagrafica.go     # gestione anagrafiche su DB destinazione
    sink/
      dropfolder.go     # scrittura nella cartella monitorata dal gestionale
    log/
      log.go            # wrapper logging
  pkg/
    fatturapa/
      paths.go          # costanti/XPath/tag tipici FatturaPA, utilities XML
  configs/
    config.dev.yaml
    config.prod.yaml
  scripts/
    run_attive.sh
    run_passive.sh
  go.mod
  go.sum
  README.md



## Ruolo delle cartelle
cmd/attive, cmd/passive: entrypoint separati (due binari) che compongono la pipeline read → sanitize → anagrafica → sink.

internal/config: lettura config (path cartelle, connessioni SQL Server e DB destino, regole di batch).

internal/dbsource: tutto ciò che parla con SQL Server (query attive/passive, paging).

internal/storage: accesso a filesystem per le passive (file esterni) e per la dropfolder di destinazione.

internal/document: struct tipo Document con campi comuni (ID, Tipo, XML, Meta…).

internal/p7m: funzione per trasformare []byte .p7m → []byte XML.

internal/sanitize: logica di redazione (due entry, attive/passive, stesso core riusato dove possibile).

internal/anagrafica: funzioni per lookup/creazione azienda nel DB destinazione + cache in memoria.

internal/sink: prende Document sanificato e lo deposita nella cartella giusta con il nome file corretto.

pkg/fatturapa: definisce nomi di tag, sezioni (CedentePrestatore, CessionarioCommittente, ecc.) e piccole utility su XML.