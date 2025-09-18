DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'person_gender') THEN
        CREATE TYPE person_gender AS ENUM ('male', 'female', 'not_specified');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'customer_speciality') THEN
        CREATE TYPE customer_speciality AS ENUM (
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
        CREATE TYPE user_role AS ENUM('owner', 'receptionist', 'customer', 'admin');
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
        );
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
    photo VARCHAR(255) NOT NULL DEFAULT 'https://www.gravatar.com/avatar/', 
    date_of_birth DATE NOT NULL,
    gender person_gender NOT NULL,
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
    gender VARCHAR(20) CHECK (gender IN ('male', 'female', 'unknown')),
    weight NUMERIC(5, 2),
    color VARCHAR(50),
    microchip VARCHAR(50) UNIQUE,
    is_neutered BOOLEAN DEFAULT FALSE,
    customer_id INTEGER NOT NULL,
    allergies TEXT,
    current_medications TEXT,
    special_needs TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    
    date_of_birth DATE,
    insurance_info TEXT,
    veterinary_contact TEXT,
    feeding_instructions TEXT,
    behavioral_notes TEXT,
    tattoo VARCHAR(50),
    last_vaccination_date DATE,
    next_vaccination_date DATE,
    last_deworming_date DATE,
    next_deworming_date DATE,
    last_vet_visit DATE,
    next_vet_visit DATE,
    blood_type VARCHAR(10),
    chip_implant_date DATE,
    chip_implant_location VARCHAR(100),
    insurance_policy_number VARCHAR(100),
    insurance_company VARCHAR(100),
    emergency_contact_name VARCHAR(255),
    emergency_contact_phone VARCHAR(20),
    
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    
    CONSTRAINT chk_weight_positive CHECK (weight IS NULL OR weight > 0),
    CONSTRAINT chk_age_positive CHECK (age IS NULL OR age >= 0),
    CONSTRAINT chk_gender_valid CHECK (gender IN ('male', 'female', 'unknown', 'other'))
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_pets_customer_id ON pets(customer_id);
CREATE INDEX IF NOT EXISTS idx_pets_species ON pets(species);
CREATE INDEX IF NOT EXISTS idx_pets_breed ON pets(breed);
CREATE INDEX IF NOT EXISTS idx_pets_is_active ON pets(is_active);
CREATE INDEX IF NOT EXISTS idx_pets_is_neutered ON pets(is_neutered);
CREATE INDEX IF NOT EXISTS idx_pets_gender ON pets(gender);
CREATE INDEX IF NOT EXISTS idx_pets_name ON pets(name);
CREATE INDEX IF NOT EXISTS idx_pets_microchip ON pets(microchip);
CREATE INDEX IF NOT EXISTS idx_pets_date_of_birth ON pets(date_of_birth);
CREATE INDEX IF NOT EXISTS idx_pets_last_vet_visit ON pets(last_vet_visit);
CREATE INDEX IF NOT EXISTS idx_pets_next_vet_visit ON pets(next_vet_visit);
CREATE INDEX IF NOT EXISTS idx_pets_created_at ON pets(created_at);
CREATE INDEX IF NOT EXISTS idx_pets_deleted_at ON pets(deleted_at) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_pets_species_breed ON pets(species, breed);
CREATE INDEX IF NOT EXISTS idx_pets_customer_active ON pets(customer_id, is_active);
CREATE INDEX IF NOT EXISTS idx_pets_vaccination_dates ON pets(last_vaccination_date, next_vaccination_date);
CREATE INDEX IF NOT EXISTS idx_pets_deworming_dates ON pets(last_deworming_date, next_deworming_date);

CREATE UNIQUE INDEX IF NOT EXISTS uq_pets_microchip ON pets(microchip) WHERE microchip IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_pets_active_filter ON pets(id) WHERE is_active = TRUE AND deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_pets_vet_visit_dates ON pets(last_vet_visit, next_vet_visit) WHERE last_vet_visit IS NOT NULL OR next_vet_visit IS NOT NULL;


CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL, 
    last_name  VARCHAR(255) NOT NULL,
    photo VARCHAR(500) NOT NULL,
    license_number VARCHAR(20) NOT NULL,
    speciality customer_speciality NOT NULL,
    years_of_experience INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    user_id int,
    schedule_json JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);


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
    weight DECIMAL(5,2),               -- (ej: 25.5 kg)
    temperature DECIMAL(4,1),          --  (ej: 38.5 °C)
    heart_rate INT,                    --  (latidos por minuto)
    respiratory_rate INT,              -- (respiraciones por minuto)
    symptoms TEXT,                   
    medications TEXT,                 
    follow_up_date TIMESTAMP WITH TIME ZONE,
    is_emergency BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_med_hist_pet_id ON medical_history(pet_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_customer_id ON medical_history(customer_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_employee_id ON medical_history(employee_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_visit_date ON medical_history(visit_date);
CREATE INDEX IF NOT EXISTS idx_med_hist_visit_type ON medical_history(visit_type);
CREATE INDEX IF NOT EXISTS idx_med_hist_condition ON medical_history(condition);
CREATE INDEX IF NOT EXISTS idx_med_hist_is_emergency ON medical_history(is_emergency);
CREATE INDEX IF NOT EXISTS idx_med_hist_follow_up_date ON medical_history(follow_up_date);



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
    last_login TIMESTAMP WITH TIME ZONE,
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
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE;

ALTER TABLE users
ADD CONSTRAINT fk_users_employee
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE;

-- Foreign Keys for profiles table
ALTER TABLE profiles
ADD CONSTRAINT fk_profiles_user
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

CREATE INDEX IF NOT EXISTS idx_profile_user_id ON profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_user_profile_id ON users(profile_id);
CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_user_phone_number ON users(phone_number);


-- Create Appointments Table
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
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_appointment_customer_id ON appointments(customer_id);
CREATE INDEX IF NOT EXISTS idx_appointment_vet_id ON appointments(employee_id);
CREATE INDEX IF NOT EXISTS idx_appointment_pet_id ON appointments(pet_id);
CREATE INDEX IF NOT EXISTS idx_appointment_status ON appointments(status);

-- Create Payment Table
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(10) NOT NULL DEFAULT 'MXN',
    status payment_status NOT NULL DEFAULT 'pending',
    method payment_method NOT NULL DEFAULT 'cash',
    transaction_id VARCHAR(255) UNIQUE,
    description TEXT,
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    paid_at TIMESTAMP WITH TIME ZONE,
    refunded_at TIMESTAMP WITH TIME ZONE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    paid_from_customer INT,
    paid_to_employee INT,
    appointment_id INT, 
    invoice_id VARCHAR(100), 
    refund_amount NUMERIC(10, 2) CHECK (refund_amount >= 0),
    failure_reason TEXT,
    
    FOREIGN KEY (paid_from_customer) REFERENCES customers(id) ON DELETE SET NULL,
    FOREIGN KEY (paid_to_employee) REFERENCES employees(id) ON DELETE SET NULL,
    FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE SET NULL,
    
    CONSTRAINT chk_refund_amount CHECK (refund_amount IS NULL OR refund_amount <= amount),
    CONSTRAINT chk_paid_date CHECK (paid_at IS NULL OR status = 'completed' OR status = 'refunded'),
    CONSTRAINT chk_refund_date CHECK (refunded_at IS NULL OR status = 'refunded')
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(status);
CREATE INDEX IF NOT EXISTS idx_payments_customer_id ON payments(paid_from_customer);
CREATE INDEX IF NOT EXISTS idx_payments_employee_id ON payments(paid_to_employee);
CREATE INDEX IF NOT EXISTS idx_payments_method ON payments(method);
CREATE INDEX IF NOT EXISTS idx_payments_created_at ON payments(created_at);
CREATE INDEX IF NOT EXISTS idx_payments_due_date ON payments(due_date);
CREATE INDEX IF NOT EXISTS idx_payments_paid_at ON payments(paid_at);
CREATE INDEX IF NOT EXISTS idx_payments_transaction_id ON payments(transaction_id);
CREATE INDEX IF NOT EXISTS idx_payments_appointment_id ON payments(appointment_id);
CREATE INDEX IF NOT EXISTS idx_payments_invoice_id ON payments(invoice_id);
CREATE INDEX IF NOT EXISTS idx_payments_is_active ON payments(is_active) WHERE is_active = TRUE;
CREATE INDEX IF NOT EXISTS idx_payments_deleted_at ON payments(deleted_at) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_payments_customer_status ON payments(paid_from_customer, status);
CREATE INDEX IF NOT EXISTS idx_payments_status_date ON payments(status, created_at);
CREATE INDEX IF NOT EXISTS idx_payments_customer_date ON payments(paid_from_customer, created_at);
CREATE INDEX IF NOT EXISTS idx_payments_due_status ON payments(due_date, status) WHERE status != 'completed';


