-- 000006_payments_indexes.down.sql
-- Drop payments indexes and table (reverse order)

DROP INDEX IF EXISTS idx_payments_due_status;
DROP INDEX IF EXISTS idx_payments_customer_date;
DROP INDEX IF EXISTS idx_payments_status_date;
DROP INDEX IF EXISTS idx_payments_customer_status;
DROP INDEX IF EXISTS idx_payments_deleted_at;
DROP INDEX IF EXISTS idx_payments_is_active;
DROP INDEX IF EXISTS idx_payments_invoice_id;
DROP INDEX IF EXISTS idx_payments_transaction_id;
DROP INDEX IF EXISTS idx_payments_paid_at;
DROP INDEX IF EXISTS idx_payments_due_date;
DROP INDEX IF EXISTS idx_payments_created_at;
DROP INDEX IF EXISTS idx_payments_method;
DROP INDEX IF EXISTS idx_payments_customer_id;
DROP INDEX IF EXISTS idx_payments_status;

DROP TABLE IF EXISTS payments CASCADE;
