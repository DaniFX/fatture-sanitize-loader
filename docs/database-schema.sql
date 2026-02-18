-- Database Schema per fatture-sanitize-loader
-- SQL Server Database: Evo4WebSviluppo

-- ============================================================================
-- Tabella PA_Storico - Storico Fatture Elettroniche (Attive e Passive)
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
-- CAMPI RILEVANTI PER IL LOADER
-- ============================================================================

-- Chiave: ID univoco del record
-- TipoDoc: Tipo documento (es: TD01=Fattura, TD04=Nota Credito, etc.)
-- TipoFlusso: Discrimina tra fatture attive e passive
-- NrDoc: Numero del documento
-- DataDoc: Data del documento
-- FileXML: File XML della fattura (tipo IMAGE/VARBINARY)
-- FileXMLFirmato: File XML con firma digitale P7M (tipo IMAGE/VARBINARY)
-- NomeFileXML: Nome originale del file XML
-- NomeFileXMLFirmato: Nome del file P7M se presente
-- Committente: Ragione sociale committente/cessionario
-- PivaCommittente: P.IVA del committente
-- XMLEstratto: Flag che indica se l'XML è già stato estratto
-- GlobeRestApiExported: Flag per indicare se già esportato (default 1, usare 0 per "da elaborare")
-- DataEstrazione: Data di estrazione/elaborazione

-- ============================================================================
-- NOTE OPERATIVE
-- ============================================================================

-- Per identificare le fatture da elaborare, usare:
-- WHERE GlobeRestApiExported = 0 (o IS NULL)
-- AND TipoFlusso = 'ATTIVO' (per attive) o 'PASSIVO' (per passive)

-- Dopo l'elaborazione, aggiornare:
-- UPDATE PA_Storico SET GlobeRestApiExported = 1, DataEstrazione = GETDATE() WHERE Chiave = ?
