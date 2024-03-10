package worker

import (
	"wb-data-service-golang/wb-data-service/internal/domain"
)

type Dependency struct {
	Logger   domain.Logger
	Database domain.DatabaseManager
	WbWorker domain.WbWorker
}

func NewWbWorkerModule(dependency Dependency) {

}
