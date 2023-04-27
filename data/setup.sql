DROP TABLE replays;
DROP TABLE forums;
DROP TABLE messages;
DROP TABLE users;

CREATE TABLE users (
    id  SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    email   VARCHAR(255) NOT NULL ,
    password    VARCHAR(20) NOT NULL ,
    created_at  TIMESTAMP NOT NULL
);
CREATE TABLE forums (
    id  SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content   text NOT NULL ,
    author_id INTEGER NOT NULL REFERENCES users(id),
    created_at  TIMESTAMP NOT NULL
);
CREATE TABLE replays (
    id  SERIAL PRIMARY KEY,
    content   text NOT NULL ,
    author_id INTEGER NOT NULL REFERENCES users(id),
    forum_id INTEGER NOT NULL REFERENCES forums(id) ,
    created_at  TIMESTAMP NOT NULL
);
CREATE TABLE messages (
    id  SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL REFERENCES users(id),
    recipient_id INTEGER NOT NULL REFERENCES users(id) ,
    content   text NOT NULL ,
    created_at  TIMESTAMP NOT NULL
);