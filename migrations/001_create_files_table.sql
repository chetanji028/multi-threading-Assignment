CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    file_id UUID NOT NULL,
    part_number INT NOT NULL,
    data BYTEA NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_file_id ON files(file_id);
