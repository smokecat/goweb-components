package fw

type DatabaseConf struct {
	Dsn           string `json:"dsn"`
	MigrationsDir string `json:"migrationsDir"`
}
