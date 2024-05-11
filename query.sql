CREATE TABLE pinpoint_types (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    image_url VARCHAR(255),
    name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE missions (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    target INT,
    point INT,
    description VARCHAR(255),
    start_at TIMESTAMP,
    end_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE tiers (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    tier_name VARCHAR(255),
    minimal_point INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    tier_id BIGINT,
    email VARCHAR(255),
    password VARCHAR(255),
    name VARCHAR(255),
    current_point INT,
    role VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT email_unique UNIQUE (email),
    FOREIGN KEY (tier_id) REFERENCES tiers(id) ON DELETE SET NULL
);


CREATE TABLE user_missions (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT,
    mission_id BIGINT,
    current_progress INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (mission_id) REFERENCES missions(id) ON DELETE SET NULL
);

CREATE TABLE pinpoints (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT,
    pinpoint_type_id BIGINT,
    longitude FLOAT,
    latitude FLOAT,
    image VARCHAR(255),
    name VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (pinpoint_type_id) REFERENCES pinpoint_types(id) ON DELETE SET NULL
);

CREATE TABLE reports (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT,
    pinpoint_id BIGINT,
    reason TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (pinpoint_id) REFERENCES pinpoints(id) ON DELETE SET NULL
);