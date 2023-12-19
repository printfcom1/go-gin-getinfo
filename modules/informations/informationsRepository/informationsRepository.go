package informationsRepository

import (
	"fmt"

	"github.com/gin-api/modules/informations"
	"github.com/jmoiron/sqlx"
)

type InformationsRepository interface {
	GetInfo() ([]informations.InformationsRes, error)
}

type InformationsRepositoryDB struct {
	db *sqlx.DB
}

func NewInformationsRepositoryDB(db *sqlx.DB) InformationsRepositoryDB {
	return InformationsRepositoryDB{db: db}
}

func (r *InformationsRepositoryDB) GetInfo() ([]informations.InformationsRes, error) {
	var info []informations.Informations
	query := "SELECT * FROM informations where parent_id = 0 ORDER BY id"

	err := r.db.Select(&info, query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var infoRes []informations.InformationsRes

	for _, v := range info {
		var infoChid []informations.InformationsChild
		query := "SELECT * FROM informations WHERE parent_id = ? ORDER BY id"

		err := r.db.Select(&infoChid, query, v.ID)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		infoData := informations.InformationsRes{
			ID:         v.ID,
			ParentID:   v.ParentID,
			Title:      v.Title,
			Name:       v.Name,
			Route:      v.Route,
			Icon:       v.Icon,
			IsChildren: v.IsChildren,
			Children:   infoChid,
		}
		infoRes = append(infoRes, infoData)
	}

	return infoRes, nil
}
