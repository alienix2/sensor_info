CREATE DATABASE mqtt_users;
USE mqtt_users;

CREATE TABLE account (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    is_superuser TINYINT(1) DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE topics (
    id INT AUTO_INCREMENT PRIMARY KEY,
    topic VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE acl (
    id INT AUTO_INCREMENT PRIMARY KEY,
    topic VARCHAR(255),
    user_id INT NOT NULL,
    rw TINYINT(1), -- 1: read-only, 2: write-only, 3: read-write
    FOREIGN KEY (user_id) REFERENCES account(id),
    FOREIGN KEY (topic) REFERENCES topics(topic)
);

CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    topic VARCHAR(255),
    sent_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    device_data TEXT,
    device_id INT,
    device_name VARCHAR(255),
    device_unit VARCHAR(255),
    payload TEXT,
    FOREIGN KEY (topic) REFERENCES topics(topic)
);

INSERT INTO account (username, password_hash, email, is_superuser)
VALUES
('alice', '$2a$12$daqa1nbpP7uZ2JVrCP5iG.svoU6tIxTVhzFlzjGXCjca8SGswSeNq', 'alice@example.com', 1),
('bob', '$2a$12$.Ty33eQhO4YZiBF70m3Gm.Qrn0cD2g2yepuOPlGAB48.MJ6RDNNPS', 'bob@example.com', 0),
('charlie', '$2a$12$ohkW5.c.EaEwCl9ERJhuF.klr7eNfnowvcTRVKUB7Rkh3b2Oyy31e', 'charlie@example.com', 0);
-- Passwords: password

INSERT INTO topics (topic) VALUES ('home/temperature'), ('home/humidity'), ('office/temperature'), ('office/humidity');

INSERT INTO acl (topic, user_id, rw)
VALUES
  ('home/temperature', (SELECT id FROM account WHERE username = 'alice'), 3), -- Alice can read and write to home/temperature
  ('home/humidity', (SELECT id FROM account WHERE username = 'alice'), 1),    -- Alice can only read home/humidity
  ('home/temperature', (SELECT id FROM account WHERE username = 'bob'), 2),   -- Bob can only write to home/temperature
  ('home/humidity', (SELECT id FROM account WHERE username = 'bob'), 3),      -- Bob can read and write to home/humidity
  ('office/temperature', (SELECT id FROM account WHERE username = 'charlie'), 1), -- Charlie can only read office/temperature
  ('office/humidity', (SELECT id FROM account WHERE username = 'charlie'), 2); -- Charlie can only write office/humidity

INSERT INTO account (username, password_hash, email, is_superuser) VALUES ('omnisub', '$2y$10$G9omfphPp44ydARgfuCzn.2lBSDUUxzoy7pbPh41iEIncyuP8wqUe', 'omnisub@example.com', 1);
