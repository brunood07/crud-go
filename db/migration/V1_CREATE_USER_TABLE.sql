DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id         INT AUTO_INCREMENT NOT NULL,
  first_name      VARCHAR(128) NOT NULL,
  last_name     VARCHAR(255) NOT NULL,
  age      INT NOT NULL,
  email VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);