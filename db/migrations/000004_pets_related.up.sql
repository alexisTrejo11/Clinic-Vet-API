-- 000004_pets_related.up.sql
-- Pets and related tables

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
    allergies TEXT,
    current_medications TEXT,
    special_needs TEXT,
    feeding_instructions TEXT,
    behavioral_notes TEXT,
    veterinary_contact TEXT,
    emergency_contact_name VARCHAR(255),
    emergency_contact_phone VARCHAR(20),
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
    noted_by INT,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS pet_vaccinations (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    vaccine_name VARCHAR(255) NOT NULL,
    administered_date DATE NOT NULL,
    next_due_date DATE,
    administered_by INT,
    batch_number VARCHAR(100),
    vaccine_type VARCHAR(100) NOT NULL DEFAULT 'core',
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
    administered_by INT,
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
    implanted_by INT,
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (implanted_by) REFERENCES employees(id) ON DELETE SET NULL
);
