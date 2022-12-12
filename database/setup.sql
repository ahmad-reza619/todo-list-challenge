CREATE TABLE IF NOT EXISTS activities (
	id INT NOT NULL AUTO_INCREMENT,
	email VARCHAR(255) NOT NULL,
	title VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS todos (
	id INT NOT NULL AUTO_INCREMENT,
	activity_group_id INT NOT NULL,
	title VARCHAR(255) NOT NULL,
	is_active VARCHAR(255) NOT NULL,
	priority VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL DEFAULT NULL,
	PRIMARY KEY (id),
    FOREIGN KEY (activity_group_id) REFERENCES activities(id)
)
