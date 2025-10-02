This folder contains database migrations split into small, focused files for readability and safer rollbacks.

Apply order (up):
  1. 000001_types.up.sql
  2. 000002_users.up.sql
  3. 000003_customers_employees.up.sql
  4. 000004_pets_related.up.sql
  5. 000005_appointments_med_sessions.up.sql
  6. 000006_payments_indexes.up.sql

Rollback order (down):
  Run the corresponding .down.sql files in reverse order (or use your migration tool which should handle ordering):
  1. 000006_payments_indexes.down.sql
  2. 000005_appointments_med_sessions.down.sql
  3. 000004_pets_related.down.sql
  4. 000003_customers_employees.down.sql
  5. 000002_users.down.sql
  6. 000001_types.down.sql

Notes:
- Each file contains comments and related DDL grouped by domain area.
- Enum types are created first and dropped last to avoid dependency issues.
- Tables are created with IF NOT EXISTS and dropped with IF EXISTS to make migrations idempotent.
- If you use a migration tool (flyway, goose, sql-migrate, etc.) ensure it picks up these files in numeric order.
