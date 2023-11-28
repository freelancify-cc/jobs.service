CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE 
    IF NOT EXISTS JobStatus (
        id SERIAL PRIMARY KEY NOT NULL,
        status VARCHAR(20) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS PaymentPlans (
        id SERIAL PRIMARY KEY NOT NULL,
        plan_name VARCHAR(30) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS Jobs (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
        job_name VARCHAR(50) NOT NULL,
        job_description VARCHAR(300) NOT NULL,
        posted_employer_id UUID NOT NULL,
        job_status_id INT NOT NULL,
        CONSTRAINT fk_job_status_id FOREIGN KEY (job_status_id) REFERENCES JobStatus(id),
        payment_plan_id INT NOT NULL,
        CONSTRAINT fk_payment_plan_id FOREIGN KEY (payment_plan_id) REFERENCES PaymentPlans(id)
    );
