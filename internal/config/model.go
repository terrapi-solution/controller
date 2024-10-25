package config

// Config represents the overall configuration for the application.
type Config struct {
	Servers   Servers   // Servers contains the configuration for various servers.
	Datastore Datastore // Datastore contains the configuration for the database.
}

// GrpcServer represents the configuration for a gRPC server.
type GrpcServer struct {
	Host          string      // Host is the hostname or IP address of the gRPC server.
	Port          int         // Port is the port number on which the gRPC server listens.
	Certificate   Certificate // Certificate contains the TLS certificate configuration.
	StrictCurves  bool        // StrictCurves indicates whether to enforce strict elliptic curves.
	StrictCiphers bool        // StrictCiphers indicates whether to enforce strict ciphers.
}

// Datastore represents the configuration for the database connection.
type Datastore struct {
	Host     string // Host is the hostname or IP address of the database server.
	Port     int    // Port is the port number on which the database server listens.
	Database string // Database is the name of the database.
	Username string // Username is the username for the database connection.
	Password string // Password is the password for the database connection.
}

// Certificate represents the TLS certificate configuration.
type Certificate struct {
	Status   bool   // Status indicates whether TLS is enabled.
	CertFile string // CertFile is the path to the certificate file.
	KeyFile  string // KeyFile is the path to the key file.
	CaFile   string // CaFile is the path to the CA file.
}

// Servers represents the configuration for various servers.
type Servers struct {
	Grpc   GrpcServer   // Grpc contains the configuration for the gRPC server.
	Metric MetricServer // Metric contains the configuration for the metric server.
	Rest   RestServer   // Rest contains the configuration for the REST server.
}

// RestServer represents the configuration for a REST server.
type RestServer struct {
	Host          string      // Host is the hostname or IP address of the REST server.
	Port          int         // Port is the port number on which the REST server listens.
	Certificate   Certificate // Certificate contains the TLS certificate configuration.
	StrictCurves  bool        // StrictCurves indicates whether to enforce strict elliptic curves.
	StrictCiphers bool        // StrictCiphers indicates whether to enforce strict ciphers.
}

// MetricServer represents the configuration for a metric server.
type MetricServer struct {
	Host        string      // Host is the hostname or IP address of the metric server.
	Port        int         // Port is the port number on which the metric server listens.
	Token       string      // Token is the authentication token for the metric server.
	Certificate Certificate // Certificate contains the TLS certificate configuration.
}
