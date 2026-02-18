// cmd/attive/main.go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fatture-sanitize-loader/internal/anagrafica"
	"fatture-sanitize-loader/internal/config"
	"fatture-sanitize-loader/internal/dbsource"
	"fatture-sanitize-loader/internal/sanitize"
	"fatture-sanitize-loader/internal/sink"
)

func main() {
	// logging base
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// contesto con cancellazione su CTRL+C
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.Load("configs/config.dev.yaml")
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	src, err := dbsource.NewSQLServerSource(cfg.SourceDB.DSN)
	if err != nil {
		log.Fatalf("source db: %v", err)
	}
	defer src.Close()

	ana := anagrafica.NewClient(cfg.APIDest.BaseURL)
	sinkCli := sink.NewHttpSink(cfg.APIDest.BaseURL)

	batchSize := cfg.Batch.Size
	if batchSize <= 0 {
		batchSize = 100
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("shutdown richiesto, esco dal loop")
			return
		default:
		}

		// timeout per il batch (evita blocchi infiniti)
		batchCtx, cancelBatch := context.WithTimeout(ctx, 2*time.Minute)
		docs, err := src.NextBatchAttive(batchCtx, batchSize)
		cancelBatch()
		if err != nil {
			log.Fatalf("read batch: %v", err)
		}
		if len(docs) == 0 {
			log.Println("nessuna fattura attiva da elaborare, fine")
			break
		}

		for _, d := range docs {
			// opzionale: controlla ctx a ogni documento
			if ctx.Err() != nil {
				log.Println("shutdown richiesto durante il batch")
				return
			}

			docSan, err := sanitize.Attiva(d)
			if err != nil {
				log.Printf("sanitize attiva id=%v: %v", d.Meta.SourceID, err)
				continue
			}

			if _, err := ana.EnsureEmittente(docSan); err != nil {
				log.Printf("anagrafica emittente id=%v: %v", d.Meta.SourceID, err)
				continue
			}

			if err := sinkCli.WriteAttiva(docSan); err != nil {
				log.Printf("sink attiva id=%v: %v", d.Meta.SourceID, err)
				continue
			}
		}
	}
}
