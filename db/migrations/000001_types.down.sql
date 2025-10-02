-- 000001_types.down.sql
-- Drop custom enum types (reverse of 000001_types.up.sql)

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'appointment_status') THEN
        DROP TYPE appointment_status;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'clinic_service') THEN
        DROP TYPE clinic_service;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_method') THEN
        DROP TYPE payment_method;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status') THEN
        DROP TYPE payment_status;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'currency') THEN
        DROP TYPE currency;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        DROP TYPE user_role;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
        DROP TYPE user_status;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'veterinarian_speciality') THEN
        DROP TYPE veterinarian_speciality;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'person_gender') THEN
        DROP TYPE person_gender;
    END IF;
END
$$;
