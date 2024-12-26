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

CREATE TABLE acl (
    id INT AUTO_INCREMENT PRIMARY KEY,
    topic VARCHAR(255),
    user_id INT NOT NULL,
    rw TINYINT(1), -- 1: read-only, 2: write-only, 3: read-write
    FOREIGN KEY (user_id) REFERENCES account(id)
);

INSERT INTO account (username, password_hash, email, is_superuser)
VALUES 
  ('alice', '$2a$10$8zVowWEmICx7thtNSHse3.j2lYVMeF9k8yCS1Z/Cxq9rG6Q5mvG6y', 'alice@example.com', 1),
  ('bob', '$2a$10$R0py72pgQYv0.Xz5g3w2HeEKlfy0cdb5wll5bXzjZK1pZ9m4z9pVm', 'bob@example.com', 0),
  ('charlie', '$2a$10$yEHeVphV5ysBZmQYN8Kxq/KijPmc/NG5bZfh8e6osgXY0ryZ5a8wu', 'charlie@example.com', 0);

INSERT INTO acl (topic, user_id, rw)
VALUES
  ('home/temperature', (SELECT id FROM account WHERE username = 'alice'), 3), -- Alice can read and write to home/temperature
  ('home/humidity', (SELECT id FROM account WHERE username = 'alice'), 1),    -- Alice can only read home/humidity
  ('home/temperature', (SELECT id FROM account WHERE username = 'bob'), 2),   -- Bob can only write to home/temperature
  ('home/humidity', (SELECT id FROM account WHERE username = 'bob'), 3),      -- Bob can read and write to home/humidity
  ('office/temperature', (SELECT id FROM account WHERE username = 'charlie'), 1), -- Charlie can only read office/temperature
  ('office/humidity', (SELECT id FROM account WHERE username = 'charlie'), 2); -- Charlie can only write office/humidity

