-- +goose Up
CREATE TABLE notification (
  id INT NOT NULL,
  title VARCHAR(255) NOT NULL,
  content VARCHAR(255) NOT NULL,
  readAt TimeStamp,
  recipientId INT NOT NULL,
  PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE notification;