-- 000003_customers_employees.down.sql
-- Drop employees and customers tables (reverse order)

DROP TABLE IF EXISTS employees CASCADE;
DROP TABLE IF EXISTS customers CASCADE;
