package repository

import (
	"awesomeProject/model"
)

type BackendRepository interface {
	All() []model.Backend
}

type BackendRepositoryImpl struct {
}

func (b BackendRepositoryImpl) All() []model.Backend {

	return []model.Backend{
		{
			Url: "http://127.0.0.1:8091",
		},
		{
			Url: "http://127.0.0.1:8091",
		},
	}

}
