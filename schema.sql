CREATE SCHEMA IF NOT EXISTS aviation;

DROP TABLE IF EXISTS aviation.registration;
CREATE TABLE IF NOT EXISTS aviation.registration (
    id                      text PRIMARY KEY NOT NULL,
    tail_number             text,
    serial_number           text,
    year_manufactured       text,
    aircraft_id             text,
    registrant              jsonb,
    address                 jsonb,
    last_activity_date      date,
    certificate_issue_date  date,
    classification          text,
    approved_operations     text[],
    type                    text,
    engine_type             text,
    status_code             text,
    model_s_code            text,
    is_fractionally_owned   boolean,
    airworthiness_date      date,
    other_names             text[],
    expiration_date         date,
    kit                     jsonb,
    created                 date
    -- CONSTRAINT aircraft_id_fkey FOREIGN KEY (aircraft_id) REFERENCES aviation.aircraft (id)
);

DROP TABLE IF EXISTS aviation.aircraft;
CREATE TABLE IF NOT EXISTS aviation.aircraft (
    id                          text PRIMARY KEY NOT NULL,
    make                        text,
    model                       text,
    type                        text,
    engine_type                 text,
    category_code               text,
    builder_certification_code  text,
    num_engines                 integer,
    num_seats                   integer,
    weight                      text,
    cruising_speed              integer,
    created                     date
);

DROP TABLE IF EXISTS aviation.engine;
CREATE TABLE IF NOT EXISTS aviation.engine (
    id          text PRIMARY KEY NOT NULL,
    make        text,
    model       text,
    type        text,
    horsepower  integer,
    thrust      integer,
    created     date
);
