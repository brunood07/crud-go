-- +goose Up
CREATE TABLE users (
  id INT NOT NULL,
  first_name VARCHAR(128) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  age INT NOT NULL,
  email VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE user;