CREATE TABLE champions (
    id INT,
    name VARCHAR(255),
    class VARCHAR(255),
    price INT
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    cost INT
);

CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    duration INTERVAL,
    winner_team_id INT
);

CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    summoner_id INT,
    champion_id INT,
    role VARCHAR(50) NOT NULL
);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    captain_id INT
);
