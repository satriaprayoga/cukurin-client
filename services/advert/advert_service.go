package serviceadvert

import "cukurin-client/models"

type Repository interface {
	GetDataBy(ID int) (result *models.Advert, err error)
}
