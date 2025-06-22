CREATE TABLE IF NOT EXISTS owners (
    id SERIAL PRIMARY KEY,
    photo TEXT,
    name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    user_id INT,
    birthday DATE,
    genre TEXT CHECK (genre IN ('male', 'female', 'other')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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


-- Add Relations
ALTER TABLE pets
ADD CONSTRAINT fk_owner
FOREIGN KEY (owner_id) REFERENCES owners(id) ON DELETE CASCADE;


INSERT INTO owners (photo, name, last_name, user_id)
VALUES ('owner_photo_url', 'John', 'Doe', 1);

