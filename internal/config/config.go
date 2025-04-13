package config

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

// App holds the application configuration with tags for environment variables, YAML, and validation
type App struct {
	SERVERHOST                     string `env:"SERVER_HOST" yaml:"server_host" validate:"required"`
	SERVERSSL                      bool   `env:"SERVER_SSL" yaml:"server_ssl"`
	SERVERCERT                     string `env:"SERVER_CERT" yaml:"server_cert"`
	SERVERKEY                      string `env:"SERVER_KEY" yaml:"server_key"`
	SERVER_LOCALE_PATH             string `env:"SERVER_LOCALE_PATH" yaml:"server_locale_path"`
	GRPCPORT                       string `env:"GRPC_PORT" yaml:"grpc_port" validate:"required"`
	HTTPPORT                       string `env:"HTTP_PORT" yaml:"http_port" validate:"required"`
	MSSQLDB                        string `env:"MSSQL_DB" yaml:"mssql_db" validate:"required"`
	MSSQLUSER                      string `env:"DB_SQLSERVER_USER_APPS" yaml:"mssql_user" validate:"required"`
	MSSQLPASSWORD                  string `env:"DB_SQLSERVER_PASS_APPS" yaml:"mssql_password" validate:"required"`
	MSSQLHOST                      string `env:"DB_SQLSERVER_HOST_APPS" yaml:"mssql_host" validate:"required"`
	MSSQLPORT                      int    `env:"MSSQL_PORT" yaml:"mssql_port" validate:"required"`
	LOGLEVEL                       int    `env:"LOG_LEVEL" yaml:"log_level"`
	LOGTIMEFORMAT                  string `env:"LOG_TIME_FORMAT" yaml:"log_time_format"`
	MIGRATIONSFOLDER               string `env:"MIGRATIONS_FOLDER" yaml:"migrations_folder"`
	MEMCACHE_HOST                  string `env:"MEMCACHE_HOST" yaml:"memcache_host"`
	MEMCACHE_PORT                  string `env:"MEMCACHE_PORT" yaml:"memcache_port"`
	EMAIL_SERVICE_CONFIG           string `env:"NOTIF_SERVICE_CONFIG" yaml:"email_service_config"`
	X_API_KEY                      string `env:"X_API_KEY" yaml:"x_api_key"`
	AES_SECRET_KEY                 string `env:"AES_SECRET_KEY" yaml:"aes_secret_key"`
	X_LINK_ENDPOINT                string `env:"X_LINK_ENDPOINT" yaml:"x_link_endpoint"`
	X_LINK_INQUIRY_ENDPOINT        string `env:"X_LINK_INQUIRY_ENDPOINT" yaml:"x_link_inquiry_endpoint"`
	HOST_FUND_TRANSACTION          string `env:"HOST_FUND_TRANSACTION" yaml:"host_fund_transaction"`
	HOST_SWITCHING                 string `env:"HOST_SWITCHING" yaml:"host_switching"`
	X_API_KEY_SWITCHING            string `env:"X_API_KEY_SWITCHING" yaml:"x_api_key_switching"`
	APP_MODE                       string `env:"APP_MODE" yaml:"app_mode"`
	ELASTIC_APM_CAPTURE_BODY       string `env:"ELASTIC_APM_CAPTURE_BODY" yaml:"elastic_apm_capture_body"`
	ELASTIC_APM_CENTRAL_CONFIG     string `env:"ELASTIC_APM_CENTRAL_CONFIG" yaml:"elastic_apm_central_config"`
	ELASTIC_APM_CAPTURE_HEADERS    string `env:"ELASTIC_APM_CAPTURE_HEADERS" yaml:"elastic_apm_capture_headers"`
	ELASTIC_APM_LOG_FILE           string `env:"ELASTIC_APM_LOG_FILE" yaml:"elastic_apm_log_file"`
	ELASTIC_APM_VERIFY_SERVER_CERT string `env:"ELASTIC_APM_VERIFY_SERVER_CERT" yaml:"elastic_apm_verify_server_cert"`
	ELASTIC_APM_SERVICE_NAME       string `env:"ELASTIC_APM_SERVICE_NAME" yaml:"elastic_apm_service_name"`
	ELASTIC_APM_ENVIRONMENT        string `env:"ELASTIC_APM_ENVIRONMENT" yaml:"elastic_apm_environment"`
	ELASTIC_APM_LOG_LEVEL          string `env:"ELASTIC_APM_LOG_LEVEL" yaml:"elastic_apm_log_level"`
	ELASTIC_APM_SERVER_URL         string `env:"ELASTIC_APM_SERVER_URL" yaml:"elastic_apm_server_url"`
	KAFKA_SERVER                   string `env:"KAFKA_SERVER" yaml:"kafka_server"`
	KAFKA_USERNAME                 string `env:"KAFKA_USERNAME" yaml:"kafka_username"`
	KAFKA_PASSWORD                 string `env:"KAFKA_PASSWORD" yaml:"kafka_password"`
	KAFKA_JKS_FILE                 string `env:"KAFKA_JKS_FILE" yaml:"kafka_jks_file"`
	KAFKA_JKS_PASSWORD             string `env:"KAFKA_JKS_PASSWORD" yaml:"kafka_jks_password"`
	KAFKA_TOPIC_NOTIFICATION       string `env:"KAFKA_TOPIC_NOTIFICATION" yaml:"kafka_topic_notification"`
	KAFKA_TOPIC_NOTIFICATION_IF    string `env:"KAFKA_TOPIC_NOTIFICATION_IF" yaml:"kafka_topic_notification_if"`
	KAFKA_TOPIC_TRANSACTION_LOG    string `env:"KAFKA_TOPIC_TRANSACTION_LOG" yaml:"kafka_topic_transaction_log"`
	LMD_API                        string `env:"LMD_API" yaml:"lmd_api"`
	LMD_API_KEY                    string `env:"LMD_API_KEY" yaml:"lmd_api_key"`
	TMP_FILE                       string `env:"TMP_FILE" yaml:"tmp_file"`
	TMP_PDF_FILE                   string `env:"TMP_PDF_FILE" yaml:"tmp_pdf_file"`
	HIDE_FITUR                     string `env:"HIDE_FITUR" yaml:"hide_fitur"`
	ENABLE_BPJS                    string `env:"ENABLE_BPJS" yaml:"enable_bpjs"`
	MS_FACILITY                    string `env:"MS_FACILITY" yaml:"ms_facility"`
	X_API_KEY_MS_FACILITY          string `env:"X_API_KEY_MS_FACILITY" yaml:"x_api_key_ms_facility"`
	MS_GENERIC_REST_HOST           string `env:"MS_GENERIC_REST_HOST" yaml:"ms_generic_rest_host"`
	MASTER_KEY                     string `env:"MASTER_KEY" yaml:"master_key"`
	MS_GENERIC_REST_HOST_IN        string `env:"MS_GENERIC_REST_HOST_IN" yaml:"ms_generic_rest_host_in"`
	DEV_MODE                       bool   `env:"DEV_MODE" yaml:"dev_mode"`
}

var ConfigApp *App

// InitConfig initializes the application configuration
func InitConfig() (config *App, err error) {
	config = &App{
		SERVERHOST:                     viper.GetString("SERVER_HOST"),
		SERVERSSL:                      viper.GetBool("SERVER_SSL"),
		SERVERCERT:                     viper.GetString("SERVER_CERT"),
		SERVERKEY:                      viper.GetString("SERVER_KEY"),
		SERVER_LOCALE_PATH:             viper.GetString("SERVER_LOCALE_PATH"),
		GRPCPORT:                       viper.GetString("GRPC_PORT"),
		HTTPPORT:                       viper.GetString("HTTP_PORT"),
		MSSQLDB:                        viper.GetString("MSSQL_DB"),
		MSSQLUSER:                      viper.GetString("DB_SQLSERVER_USER_APPS"),
		MSSQLPASSWORD:                  viper.GetString("DB_SQLSERVER_PASS_APPS"),
		MSSQLHOST:                      viper.GetString("DB_SQLSERVER_HOST_APPS"),
		MSSQLPORT:                      viper.GetInt("MSSQL_PORT"),
		LOGLEVEL:                       viper.GetInt("LOG_LEVEL"),
		LOGTIMEFORMAT:                  viper.GetString("LOG_TIME_FORMAT"),
		MIGRATIONSFOLDER:               viper.GetString("MIGRATIONS_FOLDER"),
		MEMCACHE_HOST:                  viper.GetString("MEMCACHE_HOST"),
		MEMCACHE_PORT:                  viper.GetString("MEMCACHE_PORT"),
		EMAIL_SERVICE_CONFIG:           viper.GetString("NOTIF_SERVICE_CONFIG"),
		X_API_KEY:                      viper.GetString("X_API_KEY"),
		AES_SECRET_KEY:                 viper.GetString("AES_SECRET_KEY"),
		X_LINK_ENDPOINT:                viper.GetString("X_LINK_ENDPOINT"),
		X_LINK_INQUIRY_ENDPOINT:        viper.GetString("X_LINK_INQUIRY_ENDPOINT"),
		HOST_FUND_TRANSACTION:          viper.GetString("HOST_FUND_TRANSACTION"),
		HOST_SWITCHING:                 viper.GetString("HOST_SWITCHING"),
		X_API_KEY_SWITCHING:            viper.GetString("X_API_KEY_SWITCHING"),
		APP_MODE:                       viper.GetString("APP_MODE"),
		ELASTIC_APM_CAPTURE_BODY:       viper.GetString("ELASTIC_APM_CAPTURE_BODY"),
		ELASTIC_APM_CENTRAL_CONFIG:     viper.GetString("ELASTIC_APM_CENTRAL_CONFIG"),
		ELASTIC_APM_CAPTURE_HEADERS:    viper.GetString("ELASTIC_APM_CAPTURE_HEADERS"),
		ELASTIC_APM_LOG_FILE:           viper.GetString("ELASTIC_APM_LOG_FILE"),
		ELASTIC_APM_VERIFY_SERVER_CERT: viper.GetString("ELASTIC_APM_VERIFY_SERVER_CERT"),
		ELASTIC_APM_SERVICE_NAME:       viper.GetString("ELASTIC_APM_SERVICE_NAME"),
		ELASTIC_APM_ENVIRONMENT:        viper.GetString("ELASTIC_APM_ENVIRONMENT"),
		ELASTIC_APM_LOG_LEVEL:          viper.GetString("ELASTIC_APM_LOG_LEVEL"),
		ELASTIC_APM_SERVER_URL:         viper.GetString("ELASTIC_APM_SERVER_URL"),
		KAFKA_SERVER:                   viper.GetString("KAFKA_SERVER"),
		KAFKA_USERNAME:                 viper.GetString("KAFKA_USERNAME"),
		KAFKA_PASSWORD:                 viper.GetString("KAFKA_PASSWORD"),
		KAFKA_JKS_FILE:                 viper.GetString("KAFKA_JKS_FILE"),
		KAFKA_JKS_PASSWORD:             viper.GetString("KAFKA_JKS_PASSWORD"),
		KAFKA_TOPIC_NOTIFICATION:       viper.GetString("KAFKA_TOPIC_NOTIFICATION"),
		KAFKA_TOPIC_NOTIFICATION_IF:    viper.GetString("KAFKA_TOPIC_NOTIFICATION_IF"),
		KAFKA_TOPIC_TRANSACTION_LOG:    viper.GetString("KAFKA_TOPIC_TRANSACTION_LOG"),
		LMD_API:                        viper.GetString("LMD_API"),
		LMD_API_KEY:                    viper.GetString("LMD_API_KEY"),
		TMP_FILE:                       viper.GetString("TMP_FILE"),
		TMP_PDF_FILE:                   viper.GetString("TMP_PDF_FILE"),
		HIDE_FITUR:                     viper.GetString("HIDE_FITUR"),
		ENABLE_BPJS:                    viper.GetString("ENABLE_BPJS"),
		MS_FACILITY:                    viper.GetString("MS_FACILITY"),
		X_API_KEY_MS_FACILITY:          viper.GetString("X_API_KEY_MS_FACILITY"),
		MS_GENERIC_REST_HOST:           viper.GetString("MS_GENERIC_REST_HOST"),
		MASTER_KEY:                     viper.GetString("MASTER_KEY"),
		MS_GENERIC_REST_HOST_IN:        viper.GetString("MS_GENERIC_REST_HOST_IN"),
		DEV_MODE:                       viper.GetBool("DEV_MODE"),
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, err
	}
	ConfigApp = config
	return config, nil
}

func Init() {
	viper.SetConfigFile(".env") // Load configuration from .env file
	viper.AutomaticEnv()        // Automatically map environment variables

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

// Show displays all loaded configuration settings
func Show() {
	log.Println("Loaded Configuration:")
	log.Println(viper.AllSettings())
}
