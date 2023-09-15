package main

import "fmt"

// ServerConfig 包含服务器的配置选项
type ServerConfig struct {
	Host     string
	Port     int
	Protocol string
	Timeout  int
}

// ServerOption 是一个函数类型，用于配置 ServerConfig
type ServerOption func(*ServerConfig)

// WithHost 设置服务器的主机名
func WithHost(host string) ServerOption {
	return func(c *ServerConfig) {
		c.Host = host
	}
}

// WithPort 设置服务器的端口号
func WithPort(port int) ServerOption {
	return func(c *ServerConfig) {
		c.Port = port
	}
}

// WithProtocol 设置服务器的协议
func WithProtocol(protocol string) ServerOption {
	return func(c *ServerConfig) {
		c.Protocol = protocol
	}
}

// WithTimeout 设置服务器的超时时间
func WithTimeout(timeout int) ServerOption {
	return func(c *ServerConfig) {
		c.Timeout = timeout
	}
}

// NewServer 创建一个新的服务器配置
func NewServer(options ...ServerOption) *ServerConfig {
	serverConfig := &ServerConfig{
		Host:     "localhost",
		Port:     8080,
		Protocol: "HTTP",
		Timeout:  30,
	}

	for _, option := range options {
		// 调用每一个option方法，从而修改serverConfig的值
		option(serverConfig)
	}

	return serverConfig
}

func main() {
	// 创建服务器配置，可以选择性地配置主机、端口、协议和超时时间
	server1 := NewServer(
		WithHost("example.com"),
		WithPort(8081),
	)

	server2 := NewServer(
		WithProtocol("HTTPS"),
		WithTimeout(60),
	)

	fmt.Println("Server 1:", server1)
	fmt.Println("Server 2:", server2)
}
