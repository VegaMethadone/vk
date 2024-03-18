CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(250),
    password VARCHAR(250),
    access INTEGER
);

CREATE TABLE IF NOT EXISTS acters (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(250),
    sex   VARCHAR(50),
    dateOfBirth DATE
);

CREATE TABLE IF NOT EXISTS films (
    id SERIAL PRIMARY KEY,
    name VARCHAR(250),
    description VARCHAR(1000),
    enterDate DATE,
    rate FLOAT,
    score INTEGER
);



CREATE TABLE IF NOT EXISTS film_acters  (
    id SERIAL PRIMARY KEY,
    film_id INTEGER REFERENCES films(id),
    acter_id INTEGER REFERENCES acters(id)
);