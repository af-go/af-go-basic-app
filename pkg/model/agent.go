package model

// RestServiceOptions http server options
type RestServiceOptions struct {
	Port           int    `json:"port" yaml:"port"`
	CAFile         string `json:"caFile" yaml:"caFile"`
	PrivateKeyFile string `json:"privateKetFile" yaml:"privateKeyFile"`
	PublicCertFile string `json:"publicCertFile" yaml:"publicCertFile"`
}

// MetricsServiceOptions metrics server options
type MetricsServiceOptions struct {
	Port int `json:"port" yaml:"port"`
}

// GRPCServiceOptions gRPC server options
type GRPCServiceOptions struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

// HTTPError error response
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// StatusResponse status response
type StatusResponse struct {
	Message string `json:"message"`
}
