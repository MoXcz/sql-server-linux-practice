# SQL Server in Linux

1. [Install Docker](https://docs.docker.com/engine/install/)
2. Create container:
```sh
docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=<password>" \
   -p 1433:1433 --name sql1 --hostname sql1 \
   -d \
   mcr.microsoft.com/mssql/server:2025-latest
```
3. Connect to Microsoft SQL Server Management Studio or `sqlcmd`:
```sh
docker exec -it sql1 bash
/opt/mssql-tools18/bin/sqlcmd -S localhost -U "<user>" -P "<password>" -C
```

## Start application

> [!WARNING]
> Before doing any of this it's necessary to have an `.env` file. Just rename local.env to .env and adjust the values accordingly

```sh
# with Go installed
go run .
# with Docker
docker compose up
```

## Some `.env` variables

```sh
SA_PASSWORD=Password123$                       # deprecated
MSSQL_PID=Developer                            # version
MSSQL_SA_PASSWORD=Password123$                 # password
ACCEPT_EULA=Y                                  # accept terms
MSSQL_AGENT_ENABLED=True
MSSQL_DATA_DIR=/var/opt/sqlserver/sqldata
MSSQL_LOG_DIR=/var/opt/sqlserver/sqllog
MSSQL_BACKUP_DIR=/var/opt/sqlserver/sqlbackups
```

## Available value for `MSSQL_PID`:

- Evaluation
- Developer
- Express
- Web
- Standard
- Enterprise
- EnterpriseCore
- A product key

## Inside `sqlcmd` (or SSMS)

```sql
CREATE DATABASE testDB;
SELECT name FROM sys.databases;
GO

USE testDB;

CREATE TABLE Account (
    AccountID INT,
    Balance INT
);

INSERT INTO Account (AccountID, Balance) VALUES
(1, 1500),
(2, 3000);

GO

SELECT AccountID, Balance FROM Account;
GO
```

## References

1. Blog env vars: https://www.sqlservercentral.com/blogs/sql-server-on-docker-with-new-environment-variables
2. Blog env vars: https://dbafromthecold.com/2021/09/24/using-environment-variable-files-for-sql-server-in-containers/
3. Docs env vars: https://learn.microsoft.com/en-us/sql/linux/sql-server-linux-configure-environment-variables?view=sql-server-ver15
4. Tags container: https://hub.docker.com/r/microsoft/mssql-server
5. Installation: https://learn.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker?view=sql-server-linux-ver17&preserve-view=true&tabs=cli&pivots=cs1-bash
6. MSDTC containers: https://learn.microsoft.com/en-us/sql/linux/sql-server-linux-configure-msdtc-docker?view=sql-server-ver16
7. `mssql-conf`: https://learn.microsoft.com/en-us/sql/linux/sql-server-linux-configure-mssql-conf?view=sql-server-ver16
8. MSDTC Linux: https://learn.microsoft.com/en-us/sql/linux/sql-server-linux-configure-msdtc?view=sql-server-ver16
9. SQL Server and containers guide: https://github.com/dbafromthecold/SqlServerAndContainersGuide/wiki
10. https://github.com/vishnubob/wait-for-it
11. Vendoring: https://htmx.org/

