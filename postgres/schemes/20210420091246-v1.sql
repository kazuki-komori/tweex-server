
-- +migrate Up
CREATE TABLE users (
    uid VARCHAR(50) PRIMARY KEY NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    access_token VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP
);
CREATE TABLE tweets (
    tid BIGINT PRIMARY KEY,
    user_id VARCHAR(50) REFERENCES users(uid) ON DELETE SET NULL ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS tweets;
DROP TABLE IF EXISTS users;