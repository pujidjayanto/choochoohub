package config

type ExternalClientConfig struct {
	UserApiConfig UserApi
}

type UserApi struct {
	Host string
	Port string
}

func NewExternalClientConfig(userApiConfig UserApi) ExternalClientConfig {
	return ExternalClientConfig{
		UserApiConfig: userApiConfig,
	}
}
