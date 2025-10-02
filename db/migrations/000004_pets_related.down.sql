-- 000004_pets_related.down.sql
-- Drop pet-related tables in reverse order

DROP TABLE IF EXISTS pet_chip_implants CASCADE;
DROP TABLE IF EXISTS pet_deworming CASCADE;
DROP TABLE IF EXISTS pet_vaccinations CASCADE;
DROP TABLE IF EXISTS pet_behavioral_notes CASCADE;
DROP TABLE IF EXISTS pet_feeding_instructions CASCADE;
DROP TABLE IF EXISTS pets CASCADE;
