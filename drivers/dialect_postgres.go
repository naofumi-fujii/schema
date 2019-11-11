package drivers

// No setup needed for test db: default database and schema are empty on Postgres.

var PostgresDialect = TestParams{
	CreateDDL: []string{
		`CREATE TABLE web_resource (
					id				INTEGER NOT NULL,
					url				TEXT NOT NULL UNIQUE,
					content			BYTEA,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	TEXT NOT NULL,
					etag			TEXT NOT NULL,
					last_modified	TEXT NOT NULL,
					created_at		TIMESTAMP NOT NULL,
					modified_at		TIMESTAMP,
					PRIMARY KEY (id)
		)`,
		"CREATE INDEX idx_web_resource_url ON web_resource(url)",
		"CREATE INDEX idx_web_resource_created_at ON web_resource(created_at)",
		"CREATE INDEX idx_web_resource_modified_at ON web_resource(modified_at)",
		"CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource",
		// Tests for correct identifer escaping.
		`CREATE TABLE "blanks in name" (id INTEGER, PRIMARY KEY (id))`,
		`CREATE TABLE "[brackets] in name" (id INTEGER, PRIMARY KEY (id))`,
		`CREATE TABLE """d.quotes"" in name" (id INTEGER, PRIMARY KEY (id))`,
		`CREATE TABLE "'s.quotes' in name" (id INTEGER, PRIMARY KEY (id))`,
		`CREATE TABLE "{braces} in name" (id INTEGER, PRIMARY KEY (id))`,
		"CREATE TABLE \"`backticks` in name\" (id INTEGER, PRIMARY KEY (id))",
		`CREATE TABLE "backslashes\in\name" (id INTEGER, PRIMARY KEY (id))`,
	},
	DropDDL: []string{
		`DROP TABLE "backslashes\in\name"`,
		"DROP TABLE \"`backticks` in name\"",
		`DROP TABLE "{braces} in name"`,
		`DROP TABLE "'s.quotes' in name"`,
		`DROP TABLE """d.quotes"" in name"`,
		`DROP TABLE "[brackets] in name"`,
		`DROP TABLE "blanks in name"`,
		"DROP VIEW web_resource_view",
		"DROP INDEX idx_web_resource_modified_at",
		"DROP INDEX idx_web_resource_created_at",
		"DROP INDEX idx_web_resource_url",
		"DROP TABLE web_resource",
	},

	TableNamesExpRes: []string{
		"web_resource",
		// Tests for correct identifer escaping.
		"blanks in name",
		"[brackets] in name",
		`"d.quotes" in name`,
		"'s.quotes' in name",
		"{braces} in name",
		"`backticks` in name",
		`backslashes\in\name`,
	},
	ViewNamesExpRes: []string{
		"web_resource_view",
	},
}