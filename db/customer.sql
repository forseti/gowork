CREATE TABLE customer(
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
) ENGINE = InnoDb;

ALTER TABLE customer
    ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INT DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES ('eka', 'Eka', 'eka@gmail.com', 1000000, 90.0, '1999-10-10', true),
       ('badu', 'badu', 'badu@gmail.com', 500000, 45.0, '1986-06-07', true),
       ('jaka', 'jaka', 'jaka@gmail.com', 250000, 22.5, '1975-02-04', true);