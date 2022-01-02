package cats

import (
	"rest_api/ent"
	"rest_api/swagger/models"
	"time"
)

func catsTransform(value []*ent.Cat) (cats []*models.Cat) {
	cats = make([]*models.Cat, len(value))
	for index, elem := range value {
		color := elem.Color.String()
		name := elem.Name
		cats[index] = &models.Cat{
			Age:       int32(elem.Age),
			Color:     &color,
			CreatedAt: elem.CreatedAt.Format(time.RFC3339),
			Feral:     elem.Feral,
			ID:        int64(elem.ID),
			Name:      &name,
		}
	}
	return
}

func catTransform(value *ent.Cat) (cat *models.Cat) {
	color := value.Color.String()
	name := value.Name
	cat = &models.Cat{
		Age:       int32(value.Age),
		Color:     &color,
		CreatedAt: value.CreatedAt.Format(time.RFC3339),
		Feral:     value.Feral,
		ID:        int64(value.ID),
		Name:      &name,
	}
	return
}
