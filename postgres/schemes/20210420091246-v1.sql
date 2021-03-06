
-- +migrate Up
CREATE TABLE users (
    uid VARCHAR(50) PRIMARY KEY NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    access_token VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP
);
CREATE TABLE tags (
    tag_id VARCHAR(50) PRIMARY KEY NOT NULL,
    user_id VARCHAR(50) REFERENCES users(uid) ON DELETE SET NULL ON UPDATE CASCADE
);
CREATE TABLE tweets (
    tweet_id BIGINT NOT NULL,
    user_id VARCHAR(50) REFERENCES users(uid) ON DELETE SET NULL ON UPDATE CASCADE,
    tag_id VARCHAR(50) REFERENCES tags(tag_id) ON DELETE SET NULL ON UPDATE CASCADE,
    PRIMARY KEY(tweet_id, user_id)
);


-- +migrate Down
DROP TABLE IF EXISTS tweets;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tags;