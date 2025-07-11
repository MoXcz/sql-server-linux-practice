-- SET XACT_ABORT ON will render the transaction uncommittable
-- when the constraint violation occurs.
SET XACT_ABORT ON;

BEGIN TRY
    BEGIN DISTRIBUTED TRANSACTION;

    UPDATE [BankDB1].[dbo].[Account] SET Balance = Balance - 1000 WHERE AccountID = 1;
    UPDATE [BankDB2].[dbo].[Account] SET Balance = Balance + 1000 WHERE AccountID = 2;

    COMMIT TRANSACTION;
    PRINT 'Transaction committed successfully.';
END TRY
BEGIN CATCH
    -- Test XACT_STATE for 0, 1, or -1.
    -- If 1, the transaction is committable.
    -- If -1, the transaction is uncommittable and should
    --     be rolled back.
    -- XACT_STATE = 0 means there is no transaction and
    --     a commit or rollback operation would generate an error.

    -- Test whether the transaction is uncommittable.
    IF (XACT_STATE()) = -1 BEGIN
        PRINT 'The transaction is in an uncommittable state.' +
              ' Rolling back transaction.'
        ROLLBACK TRANSACTION;
    END;

    -- Test whether the transaction is active and valid.
    IF (XACT_STATE()) = 1 BEGIN
        PRINT 'The transaction is committable.' +
              ' Committing transaction.'
        COMMIT TRANSACTION;
    END;

    PRINT 'Error: ' + ERROR_MESSAGE();
END CATCH
