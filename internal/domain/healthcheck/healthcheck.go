package healthcheck

// HealthCheck model untuk status aplikasi
type HealthCheck struct {
	DBStatus    string `json:"db_status"`
	KafkaStatus string `json:"kafka_status"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}
