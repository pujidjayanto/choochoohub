package bootstrap

type App struct {
	Env *Env
}

func InitApp() (*App, error) {
	if err := LoadEnv(); err != nil {
		return nil, err
	}

	envs := Env{
		AppEnv:           GetEnv("APP_ENV", "development"),
		DatabaseHost:     GetEnv("DB_HOST", "localhost"),
		DatabasePort:     GetEnv("DB_PORT", "3000"),
		DatabaseUser:     GetEnv("DB_USER", "admin"),
		DatabasePassword: GetEnv("DB_PASS", "admin"),
		DatabaseName:     GetEnv("DB_NAME", "user-api"),
	}

	return &App{
		Env: &envs,
	}, nil
}
