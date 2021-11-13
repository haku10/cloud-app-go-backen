
-- +migrate Up
CREATE TABLE users (
    id int(10) NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    age int(3),
    gender int(1),
    created_at datetime,
    updated_at datetime,
    PRIMARY KEY(id)
);
-- +migrate Down
DROP TABLE users;
