CREATE TABLE IF NOT EXISTS owners (
    id SERIAL PRIMARY KEY,
    photo TEXT,
    name TEXT NOT NULL,
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pets (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    photo TEXT,
    species TEXT NOT NULL,
    breed TEXT,
    age INT,
    owner_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS veterinarians (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    photo TEXT,
    specialty TEXT,
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    vet_id INT NOT NULL,
    service TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS medical_histories (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
    date TIMESTAMP NOT NULL,
    description TEXT,
    vet_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    appointment_id INT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    payment_method TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS reminders (
    id SERIAL PRIMARY KEY,
    appointment_id INT NOT NULL,
    method TEXT NOT NULL,
    time_before TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,    
    phone_number TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add Relations
ALTER TABLE pets
ADD CONSTRAINT fk_owner
FOREIGN KEY (owner_id) REFERENCES owners(id);

ALTER TABLE appointments
ADD CONSTRAINT fk_pet
FOREIGN KEY (pet_id) REFERENCES pets(id),
ADD CONSTRAINT fk_vet
FOREIGN KEY (vet_id) REFERENCES veterinarians(id);

ALTER TABLE medical_histories
ADD CONSTRAINT fk_pet
FOREIGN KEY (pet_id) REFERENCES pets(id),
ADD CONSTRAINT fk_vet
FOREIGN KEY (vet_id) REFERENCES veterinarians(id);

ALTER TABLE payments
ADD CONSTRAINT fk_appointment
FOREIGN KEY (appointment_id) REFERENCES appointments(id);

ALTER TABLE reminders
ADD CONSTRAINT fk_appointment
FOREIGN KEY (appointment_id) REFERENCES appointments(id);

ALTER TABLE owners
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE veterinarians
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id) REFERENCES users(id);
