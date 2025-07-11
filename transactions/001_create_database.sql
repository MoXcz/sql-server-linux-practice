CREATE DATABASE BankDB1;
CREATE DATABASE BankDB2;

-- run on both databases
DROP TABLE IF EXISTS Account

CREATE TABLE Account (
    AccountID INT,
    Balance INT
)

INSERT INTO Account (AccountID, Balance) VALUES
(1, 1500),
(2, 3000)
