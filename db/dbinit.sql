-- Description: This file contains the SQL statements to create the database 
-- tables and load sample data.
--
-- Usage: psql -q -h localhost -U keepm -d keepmdb -f dbinit.sql



-- Drop tables if they exist
\echo Dropping tables...
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "contact";
DROP TABLE IF EXISTS "account" CASCADE;

-- Drop types if they exist
DO $$ BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
        DROP TYPE user_role;
    END IF;
END $$;

-- Create types
CREATE TYPE user_role AS ENUM ('employee', 'admin', 'superadmin');


-- Create tables
\echo Creating tables...
\echo Creating table account...
CREATE TABLE "account" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    url VARCHAR(256),
    address VARCHAR(256),
    city VARCHAR(128),
    state VARCHAR(128),
    zip VARCHAR(10),
    country VARCHAR(128),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

\echo Creating table user...
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    first_name VARCHAR(40) NOT NULL,
    role user_role NOT NULL DEFAULT 'employee',
    email VARCHAR(128) NOT NULL,
    hash TEXT NOT NULL,
    phone VARCHAR(20),
    title VARCHAR(128),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(id)
);

\echo Creating table contact...
CREATE TABLE "contact" (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    first_name VARCHAR(40) NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    company VARCHAR(256) NOT NULL,
    title VARCHAR(128),
    phone VARCHAR(20),
    email VARCHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (account_id) REFERENCES account(id)
);


-- Load sample data
-- account
\echo Loading sample data...
\echo Loading sample data for table account...
INSERT INTO account (name, url, address, city, state, zip, country, created_at, updated_at) VALUES
('RemarQable Software', 'http://www.remarqable.io', '121 South 8th Street, Suite 1320', 
'Minneapolis', 'MN', '55402', 'USA', NOW(), NOW());

-- user
\echo Loading sample data for table user...
INSERT INTO "user" (account_id, email, hash, last_name, first_name, role, phone, title, created_at, updated_at) VALUES 
(1, 'asim@remarqable.io', '$2y$10$GHDHgqrCPap7Ej9vK25aZugZSh.ipt3D5d7GI1vFdycXERZpBdyri%', 'Baig', 'Asim', 'admin', '123-456-7890', 'Software Engineer', NOW(), NOW()),
(1, 'sophie@remarqable.io', '$2y$10$GHDHgqrCPap7Ej9vK25aZugZSh.ipt3D5d7GI1vFdycXERZpBdyri%', 'Foley', 'Sophie', 'employee', '098-765-4321', 'Product Manager', NOW(), NOW()),
(1, 'john@remarqable.io', '$2y$10$GHDHgqrCPap7Ej9vK25aZugZSh.ipt3D5d7GI1vFdycXERZpBdyri%', 'Doe', 'John', 'employee', '123-456-7891', 'Software Developer', NOW(), NOW()),
(1, 'jane@remarqable.io', '$2y$10$GHDHgqrCPap7Ej9vK25aZugZSh.ipt3D5d7GI1vFdycXERZpBdyri%', 'Doe', 'Jane', 'employee', '123-456-7892', 'Product Designer', NOW(), NOW()),
(1, 'mike@remarqable.io', '$2y$10$GHDHgqrCPap7Ej9vK25aZugZSh.ipt3D5d7GI1vFdycXERZpBdyri%', 'Smith', 'Mike', 'employee', '123-456-7893', 'Data Analyst', NOW(), NOW()),
(1, 'susan@remarqable.io', '$2y$10$GHDHgqrCPap7Ej9vK25aZugZSh.ipt3D5d7GI1vFdycXERZpBdyri%', 'Johnson', 'Susan', 'employee', '123-456-7894', 'HR Manager', NOW(), NOW()),
(1, 'bob@remarqable.io', '$2y$10$GHDHgqrCPap7Ej9vK25aZugZSh.ipt3D5d7GI1vFdycXERZpBdyri%', 'Williams', 'Bob', 'employee', '123-456-7895', 'Marketing Specialist', NOW(), NOW());

-- contact
\echo Loading sample data for table contact...
INSERT INTO "contact" (account_id, first_name, last_name, company, title, phone, email, created_at, updated_at) VALUES
(1, 'John', 'Doe', 'Company A', 'Manager', '1234567890', 'john.doe@companya.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 'Jane', 'Doe', 'Company B', 'CEO', '0987654321', 'jane.doe@companyb.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 'Jim', 'Beam', 'Company C', 'CTO', '1122334455', 'jim.beam@companyc.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 'Jack', 'Daniels', 'Company D', 'CFO', '5566778899', 'jack.daniels@companyd.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 'Jill', 'Valentine', 'Company E', 'COO', '7788991122', 'jill.valentine@companye.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 'June', 'Summer', 'Company F', 'CMO', '9911223344', 'june.summer@companyf.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

\echo Done.
