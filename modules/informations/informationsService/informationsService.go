package informationsService

import (
	"github.com/gin-api/modules/informations"
	"github.com/gin-api/modules/informations/informationsRepository"
	"github.com/gin-api/pkg/utils"
)

type InformationsService interface {
	GetInformationsService() ([]informations.InformationsRes, error)
}

type informationsService struct {
	informationsRepo informationsRepository.InformationsRepository
}

func NewInformationsService(infoRepo informationsRepository.InformationsRepository) informationsService {
	return informationsService{informationsRepo: infoRepo}
}

func (s *informationsService) GetInformationsService() ([]informations.InformationsRes, error) {
	info, err := s.informationsRepo.GetInfo()
	if err != nil {
		return nil, utils.NewInternalServerError()
	}
	return info, nil
}
