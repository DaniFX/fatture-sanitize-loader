-- Database Schema per fatture-sanitize-loader
-- SQL Server Database

-- Tabella per le fatture attive
CREATE TABLE FattureAttive (
    Id              BIGINT PRIMARY KEY IDENTITY(1,1),
    Numero          NVARCHAR(50) NOT NULL,
    Data            NVARCHAR(50) NOT NULL,  -- Formato data dalla fattura XML
    CedenteCF       NVARCHAR(16),           -- Codice Fiscale del Cedente/Prestatore
    CessionarioCF   NVARCHAR(16),           -- Codice Fiscale del Cessionario/Committente
    FileXml         VARBINARY(MAX),         -- File XML della fattura
    FileP7M         VARBINARY(MAX),         -- File P7M firmato digitalmente (opzionale)
    IsP7M           BIT DEFAULT 0,          -- Flag: 1 se il file è in formato P7M
    DaElaborare     BIT DEFAULT 1,          -- Flag: 1 se la fattura deve essere elaborata
    DataInserimento DATETIME DEFAULT GETDATE(),
    DataElaborazione DATETIME NULL
);

-- Indici per performance
CREATE INDEX IX_FattureAttive_DaElaborare ON FattureAttive(DaElaborare, Id);
CREATE INDEX IX_FattureAttive_Data ON FattureAttive(Data);

-- Tabella per le fatture passive
CREATE TABLE FatturePassive (
    Id              BIGINT PRIMARY KEY IDENTITY(1,1),
    Numero          NVARCHAR(50) NOT NULL,
    Data            NVARCHAR(50) NOT NULL,  -- Formato data dalla fattura XML
    CedenteCF       NVARCHAR(16),           -- Codice Fiscale del Cedente/Prestatore
    CessionarioCF   NVARCHAR(16),           -- Codice Fiscale del Cessionario/Committente
    FileXml         VARBINARY(MAX),         -- File XML della fattura
    FileP7M         VARBINARY(MAX),         -- File P7M firmato digitalmente (opzionale)
    IsP7M           BIT DEFAULT 0,          -- Flag: 1 se il file è in formato P7M
    DaElaborare     BIT DEFAULT 1,          -- Flag: 1 se la fattura deve essere elaborata
    DataInserimento DATETIME DEFAULT GETDATE(),
    DataElaborazione DATETIME NULL
);

-- Indici per performance
CREATE INDEX IX_FatturePassive_DaElaborare ON FatturePassive(DaElaborare, Id);
CREATE INDEX IX_FatturePassive_Data ON FatturePassive(Data);

-- Note:
-- 1. I campi FileXml e FileP7M contengono i dati binari delle fatture
-- 2. IsP7M indica quale dei due file utilizzare (P7M ha priorità se presente)
-- 3. DaElaborare viene usato dal loader per processare solo le fatture nuove
-- 4. Dopo l'elaborazione, DaElaborare può essere impostato a 0 e DataElaborazione aggiornata
-- 5. Gli indici su DaElaborare e Id ottimizzano la query "WHERE DaElaborare = 1 ORDER BY Id"
