-- Создаем таблицу для пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создаем таблицу для организаций
CREATE TABLE IF NOT EXISTS organizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создаем таблицу для тендеров
CREATE TABLE IF NOT EXISTS tenders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    service_type VARCHAR(50),
    status VARCHAR(20),
    organization_id INTEGER REFERENCES organizations(id),
    creator_username VARCHAR(50) REFERENCES users(username)
);

-- Создаем таблицу для заявок (bids)
CREATE TABLE IF NOT EXISTS bids (
    id SERIAL PRIMARY KEY,
    tender_id INTEGER REFERENCES tenders(id),
    user_id INTEGER REFERENCES users(id),
    amount DECIMAL(12, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
