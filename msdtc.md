# MSDTC

To enable MSDTC transactions it's necessary to define:
- `MSSQL_RPC_PORT`. TCP port for RPC endpoint mapper service.
- `MSSQL_DTC_TCP_PORT`. TCP port for MSDTC service.

```sh
docker run \
   -e 'ACCEPT_EULA=Y' -e 'MSSQL_SA_PASSWORD=<password>' \
   -e 'MSSQL_RPC_PORT=135' -e 'MSSQL_DTC_TCP_PORT=51000' \
   -p 51433:1433 -p 1350:135 -p 51000:51000  \
   -d mcr.microsoft.com/mssql/server:2019-GA-ubuntu-20.04
```

Firewall rules:
```sh
sudo ufw allow from any to any port 51433 proto tcp
sudo ufw allow from any to any port 51000 proto tcp
sudo ufw allow from any to any port 135 proto tcp
```

```sql
USE [master];
GO

EXECUTE master.dbo.sp_addlinkedserver
    @server = N'10.88.213.209',
    @srvproduct = N'SQL Server';
GO

EXECUTE master.dbo.sp_addlinkedsrvlogin
    @rmtsrvname = N'10.88.213.209',
    @rmtuser = 'sa',
    @rmtpassword = '<password>',
    @useself = N'False';
GO
```

```sql
SET XACT_ABORT ON;

BEGIN DISTRIBUTED TRANSACTION;

SELECT *
FROM [10.88.213.209].master.dbo.sysprocesses;

COMMIT TRANSACTION;
GO
```
