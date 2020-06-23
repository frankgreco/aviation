CREATE SCHEMA IF NOT EXISTS aviation;
ALTER SCHEMA aviation OWNER TO "frank.greco";

DROP TABLE IF EXISTS aviation.registration;
CREATE TABLE IF NOT EXISTS aviation.registration (
    unique_id           text PRIMARY KEY NOT NULL,
    id                  text,
    serial_number       text,
    year_manufactured   text,
    manufacturer        text,
    model               text,
    series              text,
    created             timestamp without time zone
);
ALTER TABLE aviation.registration OWNER TO "frank.greco";

DROP TABLE IF EXISTS aviation.aircraft;
CREATE TABLE IF NOT EXISTS aviation.aircraft (
    manufacturer        text NOT NULL,
    model               text,
    series              text,
    manufactuer_name    text,
    model_name          text,
    num_engines         text,
    num_seats           text,
    weight              text,
    cruising_speed      text,
    created             timestamp without time zone,
    PRIMARY KEY(manufacturer, model, series)
);
ALTER TABLE aviation.aircraft OWNER TO "frank.greco";