-- Migration 001: Create clothes table

CREATE TABLE IF NOT EXISTS clothes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    photo_url TEXT NOT NULL,
    season_id INT NOT NULL,
    category_id INT NOT NULL,
    status_id INT NOT NULL,
    visible BOOLEAN DEFAULT TRUE,
    creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    notes TEXT
);
