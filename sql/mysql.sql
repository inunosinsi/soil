CREATE TABLE Administrator (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	login_id VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL
)ENGINE=InnoDB;