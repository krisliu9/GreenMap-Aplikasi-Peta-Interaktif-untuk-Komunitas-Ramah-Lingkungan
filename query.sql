CREATE TABLE reports (
    Id BIGINT PRIMARY KEY,
    user_id BIGINT,
    pinpoint_id BIGINT,
    reasons TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(Id),
    FOREIGN KEY (pinpoint_id) REFERENCES pinpoints(Id)
);

CREATE TABLE users (
    Id BIGINT PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    name VARCHAR(255),
    tier VARCHAR(255),
    role VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT email_unique UNIQUE (email)
);

CREATE TABLE pinpoints (
    Id BIGINT PRIMARY KEY,
    user_id BIGINT,
    pinpoint_type_id BIGINT,
    longitude FLOAT,
    latitude FLOAT,
    image VARCHAR(255),
    name VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(Id),
    FOREIGN KEY (pinpoint_type_id) REFERENCES pinpoint_types(Id)
);

CREATE TABLE pinpoint_types (
    Id BIGINT PRIMARY KEY,
    image_url VARCHAR(255),
    name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE missions (
    Id BIGINT PRIMARY KEY,
    target INT,
    point INT,
    description VARCHAR(255),
    start_at TIMESTAMP,
    end_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE user_missions (
    Id BIGINT PRIMARY KEY,
    user_id BIGINT,
    mission_id BIGINT,
    current_progress INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(Id),
    FOREIGN KEY (mission_id) REFERENCES missions(Id)
);
