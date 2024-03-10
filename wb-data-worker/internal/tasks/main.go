package tasks

import "wb-data-service-golang/wb-data-worker/internal/domain"

type _WbTasks struct {
	Logger      domain.Logger
	HttpSession domain.HttpSession
}

func NewWbTasks(
	logger domain.Logger,
	session domain.HttpSession,
) WbTasks {
	return &_WbTasks{
		Logger:      logger,
		HttpSession: session,
	}
}
