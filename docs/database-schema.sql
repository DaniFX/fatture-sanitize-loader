-- Database Schema per fatture-sanitize-loader
-- SQL Server Database: Evo4WebSviluppo

-- ============================================================================
-- Tabella PA_Storico - Storico Fatture Elettroniche ATTIVE
-- ============================================================================

USE [Evo4WebSviluppo]
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[PA_Storico](
	[Chiave] [int] IDENTITY(1,1) NOT NULL,
	[NomeAzienda] [nvarchar](max) NULL,
	[AziendaID] [int] NULL,
	[TipoDoc] [nvarchar](2) NULL,
	[StatoPosta] [nvarchar](20) NULL,
	[DataOraCreazione] [datetime] NULL,
	[NomeFilePDF] [nvarchar](max) NULL,
	[FilePDF] [image] NULL,
	[NomeFileDati] [nvarchar](max) NULL,
	[FileDati] [image] NULL,
	[FormatoFileDati] [nvarchar](max) NULL,
	[NomeFileXML] [nvarchar](max) NULL,
	[FileXML] [image] NULL,
	[Oggetto] [nvarchar](max) NULL,
	[Testo] [nvarchar](max) NULL,
	[UtenteAS] [nvarchar](10) NULL,
	[UtenteMail] [nvarchar](100) NULL,
	[EmailTo] [nvarchar](255) NULL,
	[EmailCC] [nvarchar](255) NULL,
	[EmailBCC] [nvarchar](255) NULL,
	[PswMail] [nvarchar](50) NULL,
	[Mittente] [nvarchar](255) NULL,
	[Note] [nvarchar](max) NULL,
	[DataOraInvio] [datetime] NULL,
	[DataInvioRich] [datetime] NULL,
	[OraInvioRich] [nvarchar](50) NULL,
	[FlagHold] [bit] NULL,
	[DataStampaAvvenuta] [datetime] NULL,
	[FlagCopiaMitt] [bit] NULL,
	[FlagRicevuta] [bit] NULL,
	[FormatoMail] [nvarchar](50) NULL,
	[PrioritaEmail] [nvarchar](50) NULL,
	[FlagOpzServerSmtp] [bit] NULL,
	[FlagLogin] [bit] NULL,
	[FlagSsl] [bit] NULL,
	[SmtpServer] [nvarchar](255) NULL,
	[SmtpPorta] [nvarchar](50) NULL,
	[LoginUser] [nvarchar](255) NULL,
	[LoginPsw] [nvarchar](255) NULL,
	[SmtpTimeOut] [int] NULL,
	[SmtpTentativi] [int] NULL,
	[FlagPec] [bit] NULL,
	[FlagBccToCC] [bit] NULL,
	[FlagBccToMail] [bit] NULL,
	[FlabElabPdfMan] [bit] NULL,
	[FileTXT] [nvarchar](max) NULL,
	[NomeFileTxt] [nvarchar](max) NULL,
	[FileXMLFirmato] [image] NULL,
	[NomeFileXMLFirmato] [nvarchar](max) NULL,
	[DaFirmare] [bit] NULL,
	[FirmaAutomatica] [bit] NULL,
	[DataApposizioneFirma] [datetime] NULL,
	[UtenteFirma] [nvarchar](max) NULL,
	[idUtenteFirma] [int] NULL,
	[CodAmmDest] [nvarchar](max) NULL,
	[Committente] [nvarchar](max) NULL,
	[DataDoc] [datetime] NULL,
	[NrDoc] [nvarchar](max) NULL,
	[PivaCommittente] [nvarchar](max) NULL,
	[UltimaNotifica] [nvarchar](max) NULL,
	[XMLEstratto] [bit] NULL,
	[FormatoTrasmissione] [nvarchar](max) NULL,
	[PECDestinatario] [nvarchar](max) NULL,
	[InviaPECDestinatario] [bit] NULL,
	[Copiato] [bit] NULL,
	[DataCopiato] [datetime] NULL,
	[TipoFlusso] [nvarchar](max) NULL,
	[FiltroTesto] [nvarchar](max) NULL,
	[Imponibile] [float] NULL,
	[Imposta] [float] NULL,
	[StatoPrecedente] [nvarchar](max) NULL,
	[GlobeRestApiExported] [bit] NOT NULL,
	[DataEstrazione] [datetime] NULL,
	[TipoDocSDI] [nvarchar](max) NULL,
	[NomeFileXMLArchivio] [nvarchar](max) NULL,
	[IdentificativoSdI] [nvarchar](max) NULL,
	[Bollo] [bit] NOT NULL,
	[ImportoBollo] [nvarchar](max) NULL,
	[ArchiviatoCMT] [bit] NOT NULL,
	[DataArchiviatoCMT] [datetime] NULL,
	[ConservatoCMT] [bit] NOT NULL,
	[DataConservazione] [datetime] NULL,
	[DocEsportato] [bit] NOT NULL,
	[DataDocEsportato] [datetime] NULL,
	[DataConservazioneCMT] [datetime] NULL,
	[AlertInviato] [bit] NOT NULL,
	[VisualizzaRecord] [nvarchar](max) NULL,
	[DataConsegnaNotifica] [datetime] NULL,
	[CopiaCortesiaInviata] [bit] NOT NULL,
	[Divisa] [nvarchar](max) NULL,
 CONSTRAINT [PK_PA_Storico] PRIMARY KEY CLUSTERED 
(
	[Chiave] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

ALTER TABLE [dbo].[PA_Storico] ADD  CONSTRAINT [DF_PA_Storico_GlobeRestApiExported]  DEFAULT ((1)) FOR [GlobeRestApiExported]
GO

ALTER TABLE [dbo].[PA_Storico] ADD  CONSTRAINT [DF_PA_Storico_Bollo]  DEFAULT ((0)) FOR [Bollo]
GO

ALTER TABLE [dbo].[PA_Storico] ADD  CONSTRAINT [DF_PA_Storico_ArchiviatoCMT]  DEFAULT ((0)) FOR [ArchiviatoCMT]
GO

ALTER TABLE [dbo].[PA_Storico] ADD  CONSTRAINT [DF_PA_Storico_ConservatoCMT]  DEFAULT ((0)) FOR [ConservatoCMT]
GO

ALTER TABLE [dbo].[PA_Storico] ADD  CONSTRAINT [DF_PA_Storico_DocEsportato]  DEFAULT ((0)) FOR [DocEsportato]
GO

ALTER TABLE [dbo].[PA_Storico] ADD  CONSTRAINT [DF_PA_Storico_AlertInviato]  DEFAULT ((0)) FOR [AlertInviato]
GO

ALTER TABLE [dbo].[PA_Storico] ADD  CONSTRAINT [DF_PA_Storico_CopiaCortesiaInviata]  DEFAULT ((0)) FOR [CopiaCortesiaInviata]
GO

ALTER TABLE [dbo].[PA_Storico]  WITH CHECK ADD  CONSTRAINT [FK_PA_Storico_PA_Aziende] FOREIGN KEY([AziendaID])
REFERENCES [dbo].[PA_Aziende] ([PAAziendeID])
GO

ALTER TABLE [dbo].[PA_Storico] CHECK CONSTRAINT [FK_PA_Storico_PA_Aziende]
GO

-- ============================================================================
-- Tabella B2B_Storico_RCV - Storico Fatture Elettroniche PASSIVE
-- ============================================================================

CREATE TABLE [dbo].[B2B_Storico_RCV](
	[Id] [int] IDENTITY(1,1) NOT NULL,
	[TipoDoc] [nvarchar](max) NULL,
	[IdAzienda] [int] NULL,
	[NomeFile] [nvarchar](max) NULL,
	[NomeFileArchivio] [nvarchar](max) NULL,
	[TipoFirma] [nvarchar](max) NULL,
	[NumeroDocumento] [nvarchar](max) NULL,
	[DataDocumento] [datetime] NULL,
	[DataCreazione] [datetime] NULL,
	[Fornitore] [nvarchar](max) NULL,
	[ImportoTotale] [float] NULL,
	[ImportoPagamento] [float] NULL,
	[PIVAFornitore] [nvarchar](max) NULL,
	[TipoDocSDI] [nvarchar](max) NULL,
	[CodicefiscaleFornitore] [nvarchar](max) NULL,
	[FormatoTrasmissione] [nvarchar](max) NULL,
	[IdentificativoSdI] [nvarchar](max) NULL,
	[TentativiInvio] [nvarchar](max) NULL,
	[NomeFileNotificaMT] [nvarchar](max) NULL,
	[NomeFileNotificaArchivioMT] [nvarchar](max) NULL,
	[NoteMT] [nvarchar](max) NULL,
	[Stato] [nvarchar](max) NULL,
	[DataRicezione] [datetime] NULL,
	[Azienda] [nvarchar](max) NULL,
	[Imponibile] [float] NULL,
	[Imposta] [float] NULL,
	[Iva] [float] NULL,
	[Allegati] [nvarchar](max) NULL,
	[FileSystemExported] [bit] NOT NULL,
	[GlobeRestApiExported] [bit] NOT NULL,
	[StatoPrecedente] [nvarchar](max) NULL,
	[DocumentoEstratto] [bit] NOT NULL,
	[DataEstrazione] [datetime] NULL,
	[Stampato] [bit] NOT NULL,
	[DataStampa] [datetime] NULL,
	[DataEvaso] [datetime] NULL,
	[DocEsportato] [bit] NOT NULL,
	[DataDocEsportato] [datetime] NULL,
	[ExpAS] [bit] NOT NULL,
	[Divisa] [nvarchar](max) NULL,
 CONSTRAINT [PK_B2B_Storico] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

ALTER TABLE [dbo].[B2B_Storico_RCV] ADD  CONSTRAINT [DF_B2B_Storico_RCV_FileSystemExported]  DEFAULT ((0)) FOR [FileSystemExported]
GO

ALTER TABLE [dbo].[B2B_Storico_RCV] ADD  CONSTRAINT [DF_B2B_Storico_RCV_GlobeRestApiExported]  DEFAULT ((0)) FOR [GlobeRestApiExported]
GO

ALTER TABLE [dbo].[B2B_Storico_RCV] ADD  CONSTRAINT [DF_B2B_Storico_RCV_DocumentoEstratto]  DEFAULT ((0)) FOR [DocumentoEstratto]
GO

ALTER TABLE [dbo].[B2B_Storico_RCV] ADD  CONSTRAINT [DF_B2B_Storico_RCV_Stampato]  DEFAULT ((0)) FOR [Stampato]
GO

ALTER TABLE [dbo].[B2B_Storico_RCV] ADD  CONSTRAINT [DF_B2B_Storico_RCV_DocEsportato]  DEFAULT ((0)) FOR [DocEsportato]
GO

ALTER TABLE [dbo].[B2B_Storico_RCV] ADD  CONSTRAINT [DF_B2B_Storico_RCV_ExpAS]  DEFAULT ((0)) FOR [ExpAS]
GO

ALTER TABLE [dbo].[B2B_Storico_RCV]  WITH CHECK ADD  CONSTRAINT [FK_B2B_Storico_RCV_PA_Aziende] FOREIGN KEY([IdAzienda])
REFERENCES [dbo].[PA_Aziende] ([PAAziendeID])
GO

ALTER TABLE [dbo].[B2B_Storico_RCV] CHECK CONSTRAINT [FK_B2B_Storico_RCV_PA_Aziende]
GO

-- ============================================================================
-- CAMPI RILEVANTI PER IL LOADER - PA_Storico (ATTIVE)
-- ============================================================================

-- Chiave: ID univoco del record
-- TipoDoc: Tipo documento (es: TD01=Fattura, TD04=Nota Credito, etc.)
-- TipoFlusso: Discrimina tra fatture attive e passive ('ATTIVO')
-- NrDoc: Numero del documento
-- DataDoc: Data del documento
-- FileXML: File XML della fattura (tipo IMAGE/VARBINARY)
-- FileXMLFirmato: File XML con firma digitale P7M (tipo IMAGE/VARBINARY)
-- NomeFileXML: Nome originale del file XML
-- NomeFileXMLFirmato: Nome del file P7M se presente
-- Committente: Ragione sociale committente/cessionario
-- PivaCommittente: P.IVA del committente
-- GlobeRestApiExported: Flag per indicare se già esportato (default 1, usare 0 per "da elaborare")
-- DataEstrazione: Data di estrazione/elaborazione

-- ============================================================================
-- CAMPI RILEVANTI PER IL LOADER - B2B_Storico_RCV (PASSIVE)
-- ============================================================================

-- Id: ID univoco del record
-- TipoDoc: Tipo documento
-- NumeroDocumento: Numero del documento
-- DataDocumento: Data del documento
-- NomeFile: Nome originale del file
-- NomeFileArchivio: Nome del file archiviato (contiene il file XML o P7M)
-- TipoFirma: Indica se il file è firmato (es: "P7M")
-- Fornitore: Ragione sociale del fornitore (cedente)
-- PIVAFornitore: P.IVA del fornitore
-- CodicefiscaleFornitore: CF del fornitore
-- GlobeRestApiExported: Flag per indicare se già esportato (default 0, usare 0 per "da elaborare")
-- DataEstrazione: Data di estrazione/elaborazione
-- FileSystemExported: Flag per export su filesystem

-- NOTE: Per le passive, i file XML/P7M sono salvati su filesystem nella cartella
--       indicata da NomeFileArchivio, non come BLOB nel database

-- ============================================================================
-- NOTE OPERATIVE
-- ============================================================================

-- Per identificare le fatture ATTIVE da elaborare:
-- SELECT * FROM PA_Storico 
-- WHERE GlobeRestApiExported = 0 (o IS NULL)
-- AND TipoFlusso = 'ATTIVO'

-- Per identificare le fatture PASSIVE da elaborare:
-- SELECT * FROM B2B_Storico_RCV
-- WHERE GlobeRestApiExported = 0

-- Dopo l'elaborazione ATTIVE:
-- UPDATE PA_Storico SET GlobeRestApiExported = 1, DataEstrazione = GETDATE() WHERE Chiave = ?

-- Dopo l'elaborazione PASSIVE:
-- UPDATE B2B_Storico_RCV SET GlobeRestApiExported = 1, DataEstrazione = GETDATE() WHERE Id = ?
