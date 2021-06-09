package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	mechanism := plain.Mechanism{
		Username: "producer",
		Password: "prod-sec@Gd~CTrH]-sV[g]h",
	}
	tlsConfig, err := NewTLSConfigP("kafka/certificate.pem", "kafka/key.pem", "kafka/CARoot.pem")
	if err != nil {
		log.Fatal("Certificate generation failed", err)
	}
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		TLS:           tlsConfig,
		SASLMechanism: mechanism,
	}
	// This can be used on test server if domain does not match cert:
	tlsConfig.InsecureSkipVerify = true

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"192.168.20.31:9092"},
		Topic:    "iot_new_terminal",
		Balancer: &kafka.Hash{},
		Dialer:   dialer,
	})

	var msg kafka.Message
	var s Msg
	s.DeviceId = "KIBO339053"
	s.Ip = "192.168.3.10"
	s.Mac = "2e:3d:fb:1a:64:7d"
	s.Os = "Linux 4.4"
	s.Manufacturer = "Intel Corporate"
	s.TerminalType = "其他"
	s.Status = 2
	s.OccurredTime = time.Now()
	jsonByte, err := json.MarshalIndent(s, "", " ") //看上去更加格式化
	msg.Value = jsonByte
	err = w.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatal("Failed to send message", err)
		return
	}
	err = w.Close()
	if err != nil {
		return
	}
}

type Msg struct {
	DeviceId     string      `json:"deviceId"`
	Ip           string      `json:"ip"`
	Mac          string      `json:"mac"`
	Os           string      `json:"os"`
	Manufacturer string      `json:"manufacturer"`
	TerminalType string      `json:"terminalType"`
	Status       int         `json:"status"`
	OccurredTime interface{} `json:"occurredTime"`
}

// NewTLSConfigP NewTLSConfig generates a TLS configuration used to authenticate on server with
// certificates.
// Parameters are the three pem files path we need to authenticate: client cert, client key and CA cert.
func NewTLSConfigP(clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return &tlsConfig, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = caCertPool
	tlsConfig.BuildNameToCertificate()
	return &tlsConfig, err
}
