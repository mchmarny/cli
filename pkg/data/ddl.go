package data

const (
	ddl string = `
		CREATE TABLE IF NOT EXISTS sample (
			id INTEGER NOT NULL,
			date TEXT NOT NULL,
			PRIMARY KEY (id, date)
		);
	`
)
