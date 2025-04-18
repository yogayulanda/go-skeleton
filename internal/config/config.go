package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

// App holds the application configuration with tags for environment variables, YAML, and validation
type App struct {
	// PENAMAAN VARIABEL HARUS SESUAI DENGAN NAMA DI FILE .ENV
	// Server Configuration
	SERVER_HOST        string `env:"SERVER_HOST" yaml:"server_host" validate:"required"`
	SERVER_SSL         bool   `env:"SERVER_SSL" yaml:"server_ssl"`
	SERVER_CERT        string `env:"SERVER_CERT" yaml:"server_cert"`
	SERVER_KEY         string `env:"SERVER_KEY" yaml:"server_key"`
	SERVER_LOCALE_PATH string `env:"SERVER_LOCALE_PATH" yaml:"server_locale_path"`

	// Ports Configuration
	GRPC_PORT string `env:"GRPC_PORT" yaml:"grpc_port" validate:"required"`
	HTTP_PORT string `env:"HTTP_PORT" yaml:"http_port" validate:"required"`

	// Database Configuration
	MSSQL_DB   string `env:"MSSQL_DB" yaml:"mssql_db" validate:"required"`
	MSSQL_USER string `env:"MSSQL_USER" yaml:"mssql_user" validate:"required"`
	MSSQL_PASS string `env:"MSSQL_PASS" yaml:"mssql_password" validate:"required"`
	MSSQL_HOST string `env:"MSSQL_HOST" yaml:"mssql_host" validate:"required"`
	MSSQL_PORT int    `env:"MSSQL_PORT" yaml:"mssql_port" validate:"required"`

	// Logging Configuration
	LOG_LEVEL       string `env:"LOG_LEVEL" yaml:"log_level"`
	LOG_TIME_FORMAT string `env:"LOG_TIME_FORMAT" yaml:"log_time_format"`

	// Misc Configuration
	MIGRATIONS_FOLDER    string `env:"MIGRATIONS_FOLDER" yaml:"migrations_folder"`
	MEMCACHE_HOST        string `env:"MEMCACHE_HOST" yaml:"memcache_host"`
	MEMCACHE_PORT        string `env:"MEMCACHE_PORT" yaml:"memcache_port"`
	EMAIL_SERVICE_CONFIG string `env:"NOTIF_SERVICE_CONFIG" yaml:"email_service_config"`

	// API & Keys Configuration
	X_API_KEY      string `env:"X_API_KEY" yaml:"x_api_key"`
	HOST_SWITCHING string `env:"HOST_SWITCHING" yaml:"host_switching"`
	APP_MODE       string `env:"APP_MODE" yaml:"app_mode"`

	// Elastic APM Configuration
	ELASTIC_APM_CAPTURE_BODY       string `env:"ELASTIC_APM_CAPTURE_BODY" yaml:"elastic_apm_capture_body"`
	ELASTIC_APM_CENTRAL_CONFIG     string `env:"ELASTIC_APM_CENTRAL_CONFIG" yaml:"elastic_apm_central_config"`
	ELASTIC_APM_CAPTURE_HEADERS    string `env:"ELASTIC_APM_CAPTURE_HEADERS" yaml:"elastic_apm_capture_headers"`
	ELASTIC_APM_LOG_FILE           string `env:"ELASTIC_APM_LOG_FILE" yaml:"elastic_apm_log_file"`
	ELASTIC_APM_VERIFY_SERVER_CERT string `env:"ELASTIC_APM_VERIFY_SERVER_CERT" yaml:"elastic_apm_verify_server_cert"`
	ELASTIC_APM_SERVICE_NAME       string `env:"ELASTIC_APM_SERVICE_NAME" yaml:"elastic_apm_service_name"`
	ELASTIC_APM_ENVIRONMENT        string `env:"ELASTIC_APM_ENVIRONMENT" yaml:"elastic_apm_environment"`
	ELASTIC_APM_LOG_LEVEL          string `env:"ELASTIC_APM_LOG_LEVEL" yaml:"elastic_apm_log_level"`
	ELASTIC_APM_SERVER_URL         string `env:"ELASTIC_APM_SERVER_URL" yaml:"elastic_apm_server_url"`

	// Kafka Configuration
	KAFKA_SERVER                string `env:"KAFKA_SERVER" yaml:"kafka_server"`
	KAFKA_USERNAME              string `env:"KAFKA_USERNAME" yaml:"kafka_username"`
	KAFKA_PASSWORD              string `env:"KAFKA_PASSWORD" yaml:"kafka_password"`
	KAFKA_JKS_FILE              string `env:"KAFKA_JKS_FILE" yaml:"kafka_jks_file"`
	KAFKA_JKS_PASSWORD          string `env:"KAFKA_JKS_PASSWORD" yaml:"kafka_jks_password"`
	KAFKA_TOPIC_NOTIFICATION    string `env:"KAFKA_TOPIC_NOTIFICATION" yaml:"kafka_topic_notification"`
	KAFKA_TOPIC_TRANSACTION_LOG string `env:"KAFKA_TOPIC_TRANSACTION_LOG" yaml:"kafka_topic_transaction_log"`

	// LMD Configuration
	LMD_API     string `env:"LMD_API" yaml:"lmd_api"`
	LMD_API_KEY string `env:"LMD_API_KEY" yaml:"lmd_api_key"`

	// File Paths Configuration
	TMP_FILE     string `env:"TMP_FILE" yaml:"tmp_file"`
	TMP_PDF_FILE string `env:"TMP_PDF_FILE" yaml:"tmp_pdf_file"`

	// Environment and TLS Configuration
	DEV_MODE      bool `env:"DEV_MODE" yaml:"dev_mode"`
	ENABLE_TLS    bool
	TLS_CERT_PATH string
	TLS_KEY_PATH  string
}

func Init() error {
	viper.SetConfigFile(".env") // Load configuration from .env file
	viper.AutomaticEnv()        // Automatically map environment variables

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}
	return nil
}

// LoadConfig initializes and validates the configuration.
func LoadConfig() (*App, error) {
	config := &App{}

	// Memetakan konfigurasi dari viper ke dalam struct
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("Error unmarshalling config: %v", err)
	}

	// Debugging setelah Unmarshal
	log.Println("After UnmarshalLoadconfig, SERVER_HOST:", config.SERVER_HOST)

	// Validasi konfigurasi
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("Validation error: %v", err)
	}
	return config, nil
}

// Show displays all loaded configuration settings.
func ShowConfig() {
	// Tampilkan semua pengaturan yang dimuat
	fmt.Println("Loaded configuration:")
	fmt.Printf("Server Host: %s\n", viper.GetString("SERVER_HOST"))
	fmt.Printf("gRPC Port: %s\n", viper.GetString("GRPC_PORT"))
	fmt.Printf("HTTP Port: %s\n", viper.GetString("HTTP_PORT"))
	// Anda bisa menambahkan lebih banyak item konfigurasi di sini
}
