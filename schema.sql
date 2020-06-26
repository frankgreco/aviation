CREATE SCHEMA IF NOT EXISTS aviation;

DROP TABLE IF EXISTS aviation.registration;
CREATE TABLE IF NOT EXISTS aviation.registration (
    unique_id               text PRIMARY KEY NOT NULL,
    id                      text,
    serial_number           text,
    year_manufactured       text,
    manufacturer            text,
    model                   text,
    series                  text,
    registrant_type         text,
    registrant_name         text,
    fractional_ownership    text
    created                 timestamp without time zone
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
