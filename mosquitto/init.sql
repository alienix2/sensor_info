DROP DATABASE IF EXISTS mqtt_users;

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
    rw TINYINT(1),
    FOREIGN KEY (user_id) REFERENCES account(id),
    FOREIGN KEY (topic) REFERENCES topics(topic)
);

CREATE TABLE message_data(
    id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    sent_at DATETIME(3) NOT NULL,
    created_at DATETIME(3) DEFAULT NULL,
    topic VARCHAR(255) DEFAULT NULL,
    device_name VARCHAR(255) DEFAULT NULL,
    device_unit VARCHAR(255) DEFAULT NULL,
    device_id LONGTEXT NOT NULL,
    device_data DOUBLE NOT NULL,
    control_data VARCHAR(255) DEFAULT NULL,
    command VARCHAR(255) DEFAULT NULL,
    notes VARCHAR(255) DEFAULT NULL
);


INSERT INTO account (username, password_hash, email, is_superuser)
VALUES
('alice', '$2a$12$daqa1nbpP7uZ2JVrCP5iG.svoU6tIxTVhzFlzjGXCjca8SGswSeNq', 'alice@example.com', 1),
('bob', '$2a$12$.Ty33eQhO4YZiBF70m3Gm.Qrn0cD2g2yepuOPlGAB48.MJ6RDNNPS', 'bob@example.com', 0),
('charlie', '$2a$12$ohkW5.c.EaEwCl9ERJhuF.klr7eNfnowvcTRVKUB7Rkh3b2Oyy31e', 'charlie@example.com', 0);
-- Passwords: password

INSERT INTO topics (topic) VALUES ('home/temperature'), ('home/humidity'), ('office/temperature'), ('office/humidity'), ('#'), ('command/home/temperature'), ('command/home/humidity'), ('command/office/temperature'), ('command/office/humidity');

INSERT INTO acl (topic, user_id, rw)
VALUES
  ('home/temperature', (SELECT id FROM account WHERE username = 'alice'), 3), -- Alice can read and write to home/temperature
  ('command/home/temperature', (SELECT id FROM account WHERE username = 'alice'), 3), -- Alice can read and write to home/temperature
  ('home/humidity', (SELECT id FROM account WHERE username = 'alice'), 1),    -- Alice can only read home/humidity
  ('command/home/humidity', (SELECT id FROM account WHERE username = 'alice'), 3),    -- Alice can only read home/humidity
  ('home/temperature', (SELECT id FROM account WHERE username = 'bob'), 2),   -- Bob can only write to home/temperature
  ('command/home/temperature', (SELECT id FROM account WHERE username = 'bob'), 3),   -- Bob can only write to home/temperature
  ('home/humidity', (SELECT id FROM account WHERE username = 'bob'), 3),      -- Bob can read and write to home/humidity
  ('command/home/humidity', (SELECT id FROM account WHERE username = 'bob'), 3),      -- Bob can read and write to home/humidity
  ('office/temperature', (SELECT id FROM account WHERE username = 'charlie'), 1), -- Charlie can only read office/temperature
  ('command/office/temperature', (SELECT id FROM account WHERE username = 'charlie'), 3), -- Charlie can only read office/temperature
  ('office/humidity', (SELECT id FROM account WHERE username = 'charlie'), 2), -- Charlie can only write office/humiditY
  ('command/office/humidity', (SELECT id FROM account WHERE username = 'charlie'), 3); -- Charlie can only write office/humidity

INSERT INTO account (username, password_hash, email, is_superuser) VALUES ('omnisub', '$2a$12$ohkW5.c.EaEwCl9ERJhuF.klr7eNfnowvcTRVKUB7Rkh3b2Oyy31e', 'omnisub@example.com', 1);
INSERT INTO acl (topic, user_id, rw)
VALUES
  ('#', (SELECT id FROM account WHERE username = 'omnisub'), 3); -- Omnisub can read-write everything
