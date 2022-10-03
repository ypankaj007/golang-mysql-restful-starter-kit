CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(250) NOT NULL,
    email VARCHAR(50) NOT NULL,
    createdAt DATETIME,
    PRIMARY KEY ( id )
);
