-- +goose Up
CREATE TABLE Account (
    ID                  INT             IDENTITY(1,1) PRIMARY KEY,
    name                NVARCHAR(50)    NOT NULL,
    balance             MONEY           NOT NULL,
    created             DATETIME2       NOT NULL DEFAULT SYSUTCDATETIME(),
);

-- +goose Down
DROP TABLE Account;
