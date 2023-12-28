CREATE DATABASE authx;
CREATE USER authx WITH PASSWORD 'authx123';
GRANT CONNECT ON DATABASE authx TO authx;

\c authx;

CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS credentials(
	employee_id TEXT NOT NULL,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted_at TIMESTAMPTZ DEFAULT NULL,
	CONSTRAINT pk_credentials PRIMARY KEY(employee_id)
);
GRANT SELECT, UPDATE, INSERT, DELETE ON credentials TO authx;

CREATE TABLE IF NOT EXISTS profiles(
	ehid TEXT NOT NULL,
	employee_id TEXT NOT NULL,
	email_address TEXT,
	name TEXT NOT NULL,
	dob DATE DEFAULT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT pk_profiles PRIMARY KEY(ehid)
);
GRANT SELECT, UPDATE, INSERT, DELETE ON profiles TO authx;
