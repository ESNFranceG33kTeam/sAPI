package adherent

func CreateAdherentsTable() {
	_, err := TheDb().Exec(`
		CREATE TABLE IF NOT EXISTS adherents (
			id INT NOT NULL AUTO_INCREMENT,
			firstname VARCHAR(45) NOT NULL,
			lastname VARCHAR(45) NOT NULL,
			email VARCHAR(45) NOT NULL,
			dateofbirth VARCHAR(45) NOT NULL,
			student TINYINT NOT NULL,
			university VARCHAR(45) NULL DEFAULT NULL,
			homeland VARCHAR(45) NOT NULL,
			speakabout VARCHAR(45) NULL DEFAULT NULL,
			newsletter TINYINT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NULL DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE INDEX id_UNIQUE (id ASC)
		);
	`)
	if err != nil {
		TheLogger().LogCritical("database", "create table adherents got a problem.", err)
	} else {
		TheLogger().LogInfo("database", "adherents table successfully created.")
	}
}
