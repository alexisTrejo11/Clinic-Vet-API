-- 000001_types.up.sql
-- Create custom enum types used by the schema

DO $$
BEGIN
    -- Person gender
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'person_gender') THEN
        CREATE TYPE person_gender AS ENUM ('male', 'female', 'not_specified');
    END IF;

    -- Veterinarian speciality
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'veterinarian_speciality') THEN
        CREATE TYPE veterinarian_speciality AS ENUM (
            'unknown_specialty', 'general_practice','surgery','internal_medicine','dentistry',
            'dermatology','oncology','cardiology','neurology','ophthalmology','radiology',
            'emergency_critical_care','anesthesiology','pathology','preventive_medicine',
            'exotic_animal_medicine','equine_medicine','avian_medicine','zoo_animal_medicine',
            'food_animal_medicine','public_health'
        );
    END IF;

    -- User status and role
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
        CREATE TYPE user_status AS ENUM('active', 'inactive', 'pending', 'banned', 'deleted');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        CREATE TYPE user_role AS ENUM('veterinarian', 'receptionist', 'manager' ,'customer', 'admin', 'superadmin');
    END IF;

    -- Currency and payments
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'currency') THEN
        CREATE TYPE currency AS ENUM('USD', 'MXN', 'EUR', 'GBP', 'JPY', 'AUD', 'CAD', 'CHF', 'CNY', 'SEK', 'NZD');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status') THEN
        CREATE TYPE payment_status AS ENUM('pending', 'completed', 'failed', 'refunded');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_method') THEN
        CREATE TYPE payment_method AS ENUM(
            'cash', 'credit_card', 'debit_card', 'bank_transfer', 'paypal', 'stripe', 'check'
        );
    END IF;

    -- Clinic services and appointment status
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'clinic_service') THEN
        CREATE TYPE clinic_service AS ENUM(
            'general_consultation', 'vaccination', 'surgery', 'dental_care', 'emergency_care',
            'grooming', 'nutrition_consult', 'behavior_consult', 'wellness_exam', 'other'
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'appointment_status') THEN
        CREATE TYPE appointment_status AS ENUM('pending', 'cancelled', 'confirmed', 'rescheduled' ,'completed', 'not_presented');
    END IF;
END
$$;
