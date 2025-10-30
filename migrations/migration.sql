CREATE TABLE IF NOT EXISTS patients (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    address TEXT,
    date_of_birth DATE,
    gender VARCHAR(20),
    blood_type VARCHAR(5),
    registered_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Doctors Table
CREATE TABLE IF NOT EXISTS doctors (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    schedule VARCHAR(50),
    available_times TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Appointments Table
CREATE TABLE IF NOT EXISTS appointments (
    id VARCHAR(50) PRIMARY KEY,
    patient_id VARCHAR(50) REFERENCES patients(id) ON DELETE CASCADE,
    patient_name VARCHAR(255),
    doctor_id VARCHAR(50) REFERENCES doctors(id) ON DELETE CASCADE,
    doctor_name VARCHAR(255),
    poli VARCHAR(50),
    date DATE NOT NULL,
    time VARCHAR(10),
    complaint TEXT,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Medical Records Table
CREATE TABLE IF NOT EXISTS medical_records (
    id VARCHAR(50) PRIMARY KEY,
    appointment_id VARCHAR(50) REFERENCES appointments(id) ON DELETE CASCADE,
    patient_id VARCHAR(50) REFERENCES patients(id) ON DELETE CASCADE,
    patient_name VARCHAR(255),
    poli VARCHAR(50),
    date DATE,
    anamnesa TEXT,
    objective TEXT,
    diagnosis TEXT,
    therapy TEXT,
    prescription TEXT,
    next_visit DATE,
    created_by VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert Default Doctors
INSERT INTO doctors (id, name, email, password, schedule, available_times) VALUES
('doctor1', 'dr. Romadhani Nadri', 'romadhani@klinik.com', '$2a$10$YourHashedPasswordHere', '08:00-12:00', ARRAY['08:00', '09:00', '10:00', '11:00']),
('doctor2', 'dr. Maudina HF Diantoro', 'maudina@klinik.com', '$2a$10$YourHashedPasswordHere', '12:00-16:00', ARRAY['12:00', '13:00', '14:00', '15:00'])
ON CONFLICT (id) DO NOTHING;

-- Create Indexes
CREATE INDEX idx_appointments_patient ON appointments(patient_id);
CREATE INDEX idx_appointments_doctor ON appointments(doctor_id);
CREATE INDEX idx_appointments_date ON appointments(date);
CREATE INDEX idx_medical_records_patient ON medical_records(patient_id);
CREATE INDEX idx_medical_records_date ON medical_records(date);