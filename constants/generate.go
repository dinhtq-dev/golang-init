package constants

// AppConfig stores application-level configuration constants
type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
	AppLang string
	ErrorsMessages map[string]string 
}

// DBConfig stores database configuration constants
type DBConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

// Config contains both AppConfig and DBConfig
var Config = struct {
	AppConfig
	DBConfig
}{
	AppConfig: AppConfig{
		AppName: "MyApp",
		AppEnv:  "development",
		AppPort: "3000",
		AppLang: "en",
		ErrorsMessages: map[string]string{
			"required":    "The field %s is required",
			"email":       "The field %s must be a valid email address",
			"email_custom": "The email %s is not allowed",
		},
	},
	DBConfig: DBConfig{
		DBHost: "127.0.0.1",
		DBPort: "3306",
		DBUser: "root",
		DBPass: "",
		DBName: "my_golang_db",
	},
}
