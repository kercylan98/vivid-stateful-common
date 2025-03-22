package application

import "errors"

var (
	// ErrorServiceRegistered 服务已经注册
	ErrorServiceRegistered = errors.New("service already registered")
	// ErrorServiceTypeEmpty 服务类型为空
	ErrorServiceTypeEmpty = errors.New("service type is empty")
	// ErrorInvalidService 无效的服务
	ErrorInvalidService = errors.New("invalid service")
)
