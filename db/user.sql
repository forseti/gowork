
CREATE TABLE user(
	username VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	PRIMARY KEY (username)
) ENGINE = InnoDB;

INSERT INTO user(username, password) VALUES ('admin', 'admin')
