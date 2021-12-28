package service

import (
	"bank/dto"
	"bank/repo"
)

func NewLogService(historyRepo *repo.LogRepo) LogService {
	return &LogServiceImpl{
		LogRepo: *historyRepo,
	}
}

type LogServiceImpl struct {
	LogRepo repo.LogRepo
}

func (service *LogServiceImpl) Create(log dto.Log) error {
	err := service.LogRepo.Create(log)
	return err
}
