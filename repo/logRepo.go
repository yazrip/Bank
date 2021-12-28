package repo

import "bank/dto"

type LogRepo interface {
	Create(log dto.Log) error
}
