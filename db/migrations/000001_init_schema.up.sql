DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'person_gender') THEN
        CREATE TYPE person_gender AS ENUM ('male', 'female', 'not_specified');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'veterinarian_speciality') THEN
        CREATE TYPE veterinarian_speciality AS ENUM (
            'unknown_specialty', 'general_practice','surgery','internal_medicine','dentistry',
            'dermatology','oncology','cardiology','neurology','ophthalmology','radiology',
            'emergency_critical_care','anesthesiology','pathology','preventive_medicine',
            'exotic_animal_medicine','equine_medicine','avian_medicine','zoo_animal_medicine',
            'food_animal_medicine','public_health'
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
        CREATE TYPE user_status AS ENUM('active', 'inactive', 'pending', 'banned', 'deleted');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        CREATE TYPE user_role AS ENUM('owner', 'receptionist', 'veterinarian', 'admin');
    END IF;

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

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'clinic_service') THEN
        CREATE TYPE clinic_service AS ENUM(
            'general_consultation', 'vaccination', 'surgery', 'dental_care', 'emergency_care',
            'grooming', 'nutrition_consult', 'behavior_consult', 'wellness_exam', 'other'
        )
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'appointment_status') THEN
        CREATE TYPE appointment_status AS ENUM('pending', 'cancelled', 'confirmed', 'rescheduled' ,'completed', 'not_presented');
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    photo VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    date_of_birth DATE NOT NULL,
    gender person_gender NOT NULL,
    address VARCHAR(255),
    user_id INT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    photo TEXT,
    species VARCHAR(100) NOT NULL,
    breed VARCHAR(100),
    age SMALLINT,
    gender TEXT,
    weight NUMERIC(5, 2),
    color VARCHAR(50),
    microchip VARCHAR(50) UNIQUE,
    is_neutered BOOLEAN,
    customer_id INTEGER NOT NULL,
    allergies TEXT,
    current_medications TEXT,
    special_needs TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL, 
    last_name  VARCHAR(255) NOT NULL,
    photo VARCHAR(500) NOT NULL,
    license_number VARCHAR(20) NOT NULL,
    speciality veterinarian_speciality NOT NULL,
    years_of_experience INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    user_id int,
    schedule_json JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);


-- Add Relations
ALTER TABLE pets
ADD CONSTRAINT fk_customer_id
FOREIGN KEY (customer_id) REFERENCES owners(id) ON DELETE CASCADE;

-- Create Medical History Table
CREATE TABLE IF NOT EXISTS medical_history (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    customer_id INT NOT NULL,
    employee_id INT NOT NULL,
    visit_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    visit_type VARCHAR(50) NOT NULL,
    diagnosis TEXT,
    notes TEXT,
    treatment TEXT,
    condition VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES owners(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES veterinarians(id) ON DELETE CASCADE
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_med_hist_pet_id ON medical_history(pet_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_customer_id ON medical_history(customer_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_vet_id ON medical_history(employee_id);



-- Create User Table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    status user_status NOT NULL,
    role user_role NOT NULL,
    profile_id INT,
    customer_id INT,
    employee_id INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    user_id INT,
    bio TEXT,
    profile_pic VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    last_update TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Foreign Keys for users table
ALTER TABLE users 
ADD CONSTRAINT fk_users_profile
    FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE;

ALTER TABLE users
ADD CONSTRAINT fk_users_customer
    FOREIGN KEY (customer_id) REFERENCES owners(id) ON DELETE CASCADE;

ALTER TABLE users
ADD CONSTRAINT fk_users_veterinarian
    FOREIGN KEY (employee_id) REFERENCES veterinarians(id) ON DELETE CASCADE;

-- Foreign Keys for profiles table
ALTER TABLE profiles
ADD CONSTRAINT fk_profiles_user
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE profiles
ADD CONSTRAINT fk_profiles_owner
    FOREIGN KEY (customer_id) REFERENCES owners(id) ON DELETE CASCADE;

ALTER TABLE profiles
ADD CONSTRAINT fk_profiles_veterinarian
    FOREIGN KEY (employee_id) REFERENCES veterinarians(id) ON DELETE CASCADE;

CREATE INDEX IF NOT EXISTS idx_profile_user_id ON profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_user_profile_id ON users(profile_id);
CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_user_phone_number ON users(phone_number);


-- Create Payment Table
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(10) NOT NULL DEFAULT 'MXN',
    status payment_status NOT NULL DEFAULT 'pending',
    method payment_method NOT NULL DEFAULT 'cash',
    transaction_id VARCHAR(255) UNIQUE,
    description TEXT,
    duedate TIMESTAMP WITH TIME ZONE NOT NULL,
    paid_at TIMESTAMP WITH TIME ZONE,
    refunded_at TIMESTAMP WITH TIME ZONE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    paid_from_customer INT,
    paid_to_employee INT,
    FOREIGN KEY (padid_from_customer) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (paid_to_employee) REFERENCES employee(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS appointments(
    id SERIAL PRIMARY KEY,
    clinic_service clinic_service NOT NULL,
    schedule_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status appointment_status NOT NULL,
    reason TEXT NOT NULL DEFAULT '',
    notes TEXT DEFAULT '',
    customer_id INT NOT NULL,
    pet_id INT NOT NULL,
    employee_id INT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES owners(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES veterinarians(id) ON DELETE CASCADE
);


CREATE INDEX IF NOT EXISTS idx_appointment_customer_id ON appoinments(customer_id);
CREATE INDEX IF NOT EXISTS idx_appointment_vet_id ON appoinments(employee_id);
CREATE INDEX IF NOT EXISTS idx_appointment_pet_id ON appoinments(pet_id);
CREATE INDEX IF NOT EXISTS idx_appointment_status ON appoinments(status);
