package cats

import (
	"context"
	"rest_api/ent"
	"rest_api/ent/cat"
	"rest_api/swagger/models"
	"time"
)

type service struct {
	client *ent.CatClient
}

func (this *service) getCats(since int64, limit int32) ([]*models.Cat, error) {
	result, err := this.client.Query().
		Offset(int(since)).
		Limit(int(limit)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return catsTransform(result), nil
}

func (this *service) getCatByID(id int64) (*models.Cat, error) {
	result, err := this.client.Get(context.Background(), int(id))
	if err != nil {
		return nil, err
	}
	return catTransform(result), nil
}

func (this *service) postCat(body *models.Cat) (*models.Cat, error) {
	result, err := this.client.Create().
		SetAge(int(body.Age)).
		SetColor(cat.Color(*body.Color)).
		SetCreatedAt(time.Now()).
		SetFeral(body.Feral).
		SetName(*body.Name).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return catTransform(result), nil
}

func (this *service) putCatByID(id int64, body *models.Cat) (*models.Cat, error) {
	result, err := this.client.UpdateOneID(int(id)).
		SetAge(int(body.Age)).
		SetColor(cat.Color(*body.Color)).
		SetCreatedAt(time.Now()).
		SetFeral(body.Feral).
		SetName(*body.Name).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return catTransform(result), nil
}

func (this *service) deleteCatByID(id int64) error {
	return this.client.DeleteOneID(int(id)).Exec(context.Background())
}
