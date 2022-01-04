package cats

import (
	"rest_api/ent"
	"rest_api/swagger/models"
	"time"
)

func catsTransform(value []*ent.Cat) (cats []*models.Cat) {
	cats = make([]*models.Cat, len(value))
	for index, elem := range value {
		cats[index] = catTransform(elem)
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
