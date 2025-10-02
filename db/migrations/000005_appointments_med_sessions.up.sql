-- 000005_appointments_med_sessions.up.sql
-- Appointments and medical sessions

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
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    FOREIGN KEY (pet_id) REFERENCES pets(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
);

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
    weight DECIMAL(4,2),
    temperature DECIMAL(3,1),
    heart_rate INT,
    respiratory_rate INT,
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
