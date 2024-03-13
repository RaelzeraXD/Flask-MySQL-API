USE pydb;

CREATE TABLE pytable (
    id int PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name varchar(255) not null,
    age int not null,
    email varchar(255) not null
);
INSERT INTO pytable (name,age,email) VALUES ("teste",100,"teste@gmail.com");