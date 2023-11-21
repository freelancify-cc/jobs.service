CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE 
    IF NOT EXISTS JobStatus (
        id INT PRIMARY KEY NOT NULL IDENTITY,
        status VARCHAR(20) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS PaymentPlans (
        id INT PRIMARY KEY IDENTITY NOT NULL,
        plan_name VARCHAR(30) NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS Jobs (
        id UUID PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
        job_name VARCHAR(50) NOT NULL,
        job_description(300) NOT NULL,
        job_status_id INT NOT NULL,
        CONSTRAINT fk_job_status_id FOREIGN KEY (job_status_id) REFERENCES JobStatus(id),
        posted_employer_id INT NOT NULL,
        CONSTRAINT fk_posted_employer_id INT NOT NULL,
        payment_plan_id INT NOT NULL,
        CONSTRAINT fk_payment_plan_id FOREIGN KEY (payment_plan_id) REFERENCES PaymentPlans(id)
    );
