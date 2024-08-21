CREATE TABLE IF NOT EXISTS owners (
    id SERIAL PRIMARY KEY,
    photo TEXT,
    name TEXT NOT NULL,
    phone TEXT,
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
    email TEXT NOT NULL UNIQUE,
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


-- Insert example data into the users table first
INSERT INTO users (name, email, password, user_id, role)
VALUES ('Alice Johnson', 'alice.johnson@example.com', 'securepassword', 1, 'admin');

-- Insert example data into the owners table
INSERT INTO owners (photo, name, phone, user_id)
VALUES ('owner_photo_url', 'John Doe', '555-1234', 1);

-- Insert example data into the pets table
INSERT INTO pets (name, photo, species, breed, age, owner_id)
VALUES ('Rex', 'pet_photo_url', 'Dog', 'Labrador', 5, 1);

-- Insert example data into the veterinarians table
INSERT INTO veterinarians (name, photo, email, specialty, user_id)
VALUES ('Dr. Smith', 'vet_photo_url', 'dr.smith@example.com', 'Surgery', 1); -- Adjusted to match user_id = 1

-- Insert example data into the appointments table
INSERT INTO appointments (pet_id, vet_id, service, date)
VALUES (1, 1, 'Annual Checkup', '2024-08-19 10:00:00');

-- Insert example data into the medical_histories table
INSERT INTO medical_histories (pet_id, date, description, vet_id)
VALUES (1, '2024-08-19 10:00:00', 'Routine checkup, no issues found.', 1);

-- Insert example data into the payments table
INSERT INTO payments (appointment_id, amount, payment_method)
VALUES (1, 100.00, 'Credit Card');

-- Insert example data into the reminders table
INSERT INTO reminders (appointment_id, method, time_before)
VALUES (1, 'Email', '24 hours');
