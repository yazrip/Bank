package service

import "bank/dto"

type LogService interface {
	Create(log dto.Log) error
}
