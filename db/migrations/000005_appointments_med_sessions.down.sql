-- 000005_appointments_med_sessions.down.sql
-- Drop medical_sessions and appointments (reverse order)

DROP TABLE IF EXISTS medical_sessions CASCADE;
DROP TABLE IF EXISTS appointments CASCADE;
