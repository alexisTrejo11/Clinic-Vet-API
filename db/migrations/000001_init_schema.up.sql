DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'person_gender') THEN
        CREATE TYPE person_gender AS ENUM ('male', 'female', 'not_specified');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'employee_speciality') THEN
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
        CREATE TYPE user_role AS ENUM('veterinarian', 'receptionist', 'manager' ,'customer', 'admin', 'superadmin');
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

-- Create User Table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    status user_status NOT NULL,
    role user_role NOT NULL,
    last_login TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);


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

-- Create User Table
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

CREATE TABLE IF NOT EXISTS pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    photo TEXT,
    species VARCHAR(100) NOT NULL,
    breed VARCHAR(100),
    age SMALLINT,
    gender VARCHAR(20) CHECK (gender IN ('male', 'female', 'unknown', 'other')),
    color VARCHAR(50),
    microchip VARCHAR(50) UNIQUE,
    tattoo VARCHAR(50),
    blood_type VARCHAR(10),
    is_neutered BOOLEAN DEFAULT FALSE,
    customer_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    
    -- Campos de información general
    allergies TEXT,
    current_medications TEXT,
    special_needs TEXT,
    feeding_instructions TEXT,
    behavioral_notes TEXT,
    veterinary_contact TEXT,
    emergency_contact_name VARCHAR(255),
    emergency_contact_phone VARCHAR(20),
    
    -- Auditoría
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL,
    
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    
    CONSTRAINT chk_gender_valid CHECK (gender IN ('male', 'female', 'unknown', 'other'))
);

CREATE TABLE IF NOT EXISTS pet_feeding_instructions (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    food_brand VARCHAR(255),
    food_type VARCHAR(100), 
    amount_per_serving VARCHAR(100), 
    frequency VARCHAR(100), 
    special_instructions TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS pet_behavioral_notes (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    note TEXT NOT NULL,
    noted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    noted_by INT, -- employee_id or customer_id
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS pet_vaccinations (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    vaccine_name VARCHAR(255) NOT NULL,
    administered_date DATE NOT NULL,
    next_due_date DATE,
    administered_by INT, -- employee_id
    batch_number VARCHAR(100),
    vaccine_type VARCHAR(100) NOT NULL DEFAULT 'core', -- core, non-core, optional
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (administered_by) REFERENCES employees(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS pet_deworming (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    medication_name VARCHAR(255) NOT NULL,
    administered_date DATE NOT NULL,
    next_due_date DATE,
    administered_by INT, -- employee_id
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (administered_by) REFERENCES employees(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS pet_chip_implants (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    implant_date DATE NOT NULL,
    implant_location VARCHAR(100),
    chip_number VARCHAR(50) UNIQUE,
    implanted_by INT, -- employee_id
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (implanted_by) REFERENCES employees(id) ON DELETE SET NULL
);

-- Create Appointments Table
CREATE TABLE IF NOT EXISTS appointments(
    id SERIAL PRIMARY KEY,
    clinic_service clinic_service NOT NULL,
    scheduled_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status appointment_status NOT NULL,
    notes TEXT DEFAULT '',
    customer_id INT NOT NULL,
    pet_id INT NOT NULL,
    employee_id INT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL, FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
);



-- Create Medical Session Table
CREATE TABLE IF NOT EXISTS medical_sessions (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    customer_id INT NOT NULL,
    employee_id INT NOT NULL,
    appointment_id INT,
    clinic_service clinic_service NOT NULL,
    visit_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    visit_type VARCHAR(50) NOT NULL DEFAULT '',
    diagnosis TEXT,
    notes TEXT,
    treatment TEXT,
    condition VARCHAR(254),
    weight DECIMAL(4,2),               -- (ej: 25.5 kg)
    temperature DECIMAL(3,1),          --  (ej: 38.5 °C)
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
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (appointment_id) REFERENCES appointments(id) ON DELETE SET NULL
);

-- Create Payment Table
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(10) NOT NULL DEFAULT 'MXN',
    status payment_status NOT NULL DEFAULT 'pending',
    method payment_method NOT NULL DEFAULT 'cash',
    med_session_id INT,
    transaction_id VARCHAR(255) UNIQUE,
    description TEXT,
    due_date TIMESTAMP WITH TIME ZONE,
    paid_at TIMESTAMP WITH TIME ZONE,
    refunded_at TIMESTAMP WITH TIME ZONE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    paid_by_customer_id INT,
    invoice_id VARCHAR(100), 
    refund_amount NUMERIC(10, 2) CHECK (refund_amount >= 0),
    failure_reason TEXT,
    
    FOREIGN KEY (paid_by_customer_id) REFERENCES customers(id) ON DELETE SET NULL,
    FOREIGN KEY (paid_to_employee) REFERENCES employees(id) ON DELETE SET NULL,
    FOREIGN KEY (med_session_id) REFERENCES medical_sessions(id) ON DELETE SET NULL,
    
    CONSTRAINT chk_refund_amount CHECK (refund_amount IS NULL OR refund_amount <= amount),
    CONSTRAINT chk_paid_date CHECK (paid_at IS NULL OR status = 'completed' OR status = 'refunded'),
    CONSTRAINT chk_refund_date CHECK (refunded_at IS NULL OR status = 'refunded')
);

-- Indexes for Medical Sessions
CREATE INDEX IF NOT EXISTS idx_med_hist_pet_id ON medical_sessions(pet_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_customer_id ON medical_sessions(customer_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_employee_id ON medical_sessions(employee_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_visit_date ON medical_sessions(visit_date);
CREATE INDEX IF NOT EXISTS idx_med_hist_visit_type ON medical_sessions(visit_type);
CREATE INDEX IF NOT EXISTS idx_med_hist_condition ON medical_sessions(condition);
CREATE INDEX IF NOT EXISTS idx_med_hist_is_emergency ON medical_sessions(is_emergency);
CREATE INDEX IF NOT EXISTS idx_med_hist_follow_up_date ON medical_sessions(follow_up_date);


-- Indexes for Payments
CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(status);
CREATE INDEX IF NOT EXISTS idx_payments_customer_id ON payments(paid_by_customer_id);
CREATE INDEX IF NOT EXISTS idx_payments_method ON payments(method);
CREATE INDEX IF NOT EXISTS idx_payments_created_at ON payments(created_at);
CREATE INDEX IF NOT EXISTS idx_payments_due_date ON payments(due_date);
CREATE INDEX IF NOT EXISTS idx_payments_paid_at ON payments(paid_at);
CREATE INDEX IF NOT EXISTS idx_payments_transaction_id ON payments(transaction_id);
CREATE INDEX IF NOT EXISTS idx_payments_appointment_id ON payments(appointment_id);
CREATE INDEX IF NOT EXISTS idx_payments_invoice_id ON payments(invoice_id);
CREATE INDEX IF NOT EXISTS idx_payments_is_active ON payments(is_active) WHERE is_active = TRUE;
CREATE INDEX IF NOT EXISTS idx_payments_deleted_at ON payments(deleted_at) WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_payments_customer_status ON payments(paid_by_customer_id, status);
CREATE INDEX IF NOT EXISTS idx_payments_status_date ON payments(status, created_at);
CREATE INDEX IF NOT EXISTS idx_payments_customer_date ON payments(paid_by_customer_id, created_at);
CREATE INDEX IF NOT EXISTS idx_payments_due_status ON payments(due_date, status) WHERE status != 'completed';


-- Indexes for Appointments
CREATE INDEX IF NOT EXISTS idx_appointment_customer_id ON appointments(customer_id);
CREATE INDEX IF NOT EXISTS idx_appointment_vet_id ON appointments(employee_id);
CREATE INDEX IF NOT EXISTS idx_appointment_pet_id ON appointments(pet_id);
CREATE INDEX IF NOT EXISTS idx_appointment_status ON appointments(status);


CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_user_phone_number ON users(phone_number);


-- Indexes
CREATE INDEX IF NOT EXISTS idx_pets_customer_id ON pets(customer_id);
CREATE INDEX IF NOT EXISTS idx_pets_is_active ON pets(is_active) WHERE is_active = TRUE;
CREATE INDEX IF NOT EXISTS idx_pets_microchip ON pets(microchip) WHERE microchip IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_pets_species_breed ON pets(species, breed);
CREATE INDEX IF NOT EXISTS idx_pets_active_customer ON pets(customer_id, is_active) WHERE is_active = TRUE;

CREATE INDEX IF NOT EXISTS idx_pets_name ON pets(name);
CREATE INDEX IF NOT EXISTS idx_pets_species ON pets(species);
CREATE INDEX IF NOT EXISTS idx_pets_breed ON pets(breed) WHERE breed IS NOT NULL;

CREATE INDEX IF NOT EXISTS idx_pet_vaccinations_pet_id ON pet_vaccinations(pet_id);
CREATE INDEX IF NOT EXISTS idx_pet_vaccinations_dates ON pet_vaccinations(administered_date, next_due_date);
CREATE INDEX IF NOT EXISTS idx_pet_deworming_pet_id ON pet_deworming(pet_id);
CREATE INDEX IF NOT EXISTS idx_pet_deworming_dates ON pet_deworming(administered_date, next_due_date);
CREATE INDEX IF NOT EXISTS idx_pet_chip_implants_pet_id ON pet_chip_implants(pet_id);

CREATE UNIQUE INDEX IF NOT EXISTS uq_pets_microchip ON pets(microchip) WHERE microchip IS NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS uq_pet_chip_implants_chip ON pet_chip_implants(chip_number) WHERE chip_number IS NOT NULL;