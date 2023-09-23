BEGIN;
  CREATE TABLE locations (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    visited_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP(0) DEFAULT current_timestamp(0) NOT NULL
  );
COMMIT;