CREATE TYPE person_gender AS ENUM ('male', 'female', 'not_specified');
CREATE TYPE veterinarian_speciality AS ENUM ('unknown_specialty', 'general_practice', 'surgery', 'internal_medicine', 'dentistry', 'dermatology', 'oncology', 'cardiology', 'neurology', 'ophthalmology', 'radiology', 'emergency_critical_care', 'anesthesiology', 'pathology', 'preventive_medicine', 'exotic_animal_medicine', 'equine_medicine', 'avian_medicine', 'zoo_animal_medicine', 'food_animal_medicine', 'public_health');

CREATE TABLE IF NOT EXISTS owners (
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
    owner_id INTEGER NOT NULL,
    allergies TEXT,
    current_medications TEXT,
    special_needs TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS veterinarians(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL, 
    last_name  VARCHAR(255) NOT NULL,
    photo VARCHAR(500) NOT NULL,
    license_number VARCHAR(20) NOT NULL,
    speciality veterinarian_speciality NOT NULL,
    years_of_experience INT NOT NULL,
    is_active BOOLEAN,
    user_id int,
    schedule_json JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);


-- Add Relations
ALTER TABLE pets
ADD CONSTRAINT fk_owner
FOREIGN KEY (owner_id) REFERENCES owners(id) ON DELETE CASCADE;

-- Create Medical History Table
CREATE TABLE IF NOT EXISTS medical_history (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    owner_id INT NOT NULL,
    veterinarian_id INT NOT NULL,
    visit_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    diagnosis TEXT,
    notes TEXT,
    treatment TEXT,
    condition VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (owner_id) REFERENCES owners(id) ON DELETE CASCADE,
    FOREIGN KEY (veterinarian_id) REFERENCES veterinarians(id) ON DELETE CASCADE
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_med_hist_pet_id ON medical_history(pet_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_owner_id ON medical_history(owner_id);
CREATE INDEX IF NOT EXISTS idx_med_hist_vet_id ON medical_history(veterinarian_id);
