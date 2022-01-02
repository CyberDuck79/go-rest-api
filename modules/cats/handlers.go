package cats

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"rest_api/ent"
	"rest_api/swagger/restapi/operations/cats"
)

type getCats struct {
	catsService *service
}

func (this *getCats) Handle(params cats.GetParams) middleware.Responder {
	since := int64(0)
	if params.Since != nil {
		since = *params.Since
	}
	result, err := this.catsService.getCats(since, *params.Limit)
	if err != nil {
		middleware.Error(500, fmt.Sprintf("internal error: %s", err))
	}
	return cats.NewGetOK().WithPayload(result)
}

type getCatID struct {
	catsService *service
}

func (this *getCatID) Handle(params cats.GetIDParams) middleware.Responder {
	result, err := this.catsService.getCatByID(params.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return middleware.Error(404, fmt.Sprintf("cat not found with ID: %v", params.ID))
		}
		return middleware.Error(500, fmt.Sprintf("internal error: %s", err))
	}
	return cats.NewGetIDOK().WithPayload(result)
}

type postCat struct {
	catsService *service
}

func (this *postCat) Handle(params cats.PostParams) middleware.Responder {
	result, err := this.catsService.postCat(params.Body)
	if err != nil {
		if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return middleware.Error(400, fmt.Sprintf("Validation error: %s", err))
		}
		return middleware.Error(500, fmt.Sprintf("internal error: %s", err))
	}
	return cats.NewPostCreated().WithPayload(result)
}

type putCatID struct {
	catsService *service
}

func (this *putCatID) Handle(params cats.PutIDParams) middleware.Responder {
	result, err := this.catsService.putCatByID(params.ID, params.Body)
	if err != nil {
		if ent.IsNotFound(err) {
			return middleware.Error(404, fmt.Sprintf("cat not found with ID: %v", params.ID))
		} else if ent.IsValidationError(err) || ent.IsConstraintError(err) {
			return middleware.Error(400, fmt.Sprintf("Validation error: %s", err))
		}
		return middleware.Error(500, fmt.Sprintf("internal error: %s", err))
	}
	return cats.NewPostCreated().WithPayload(result)
}

type deleteCatID struct {
	catsService *service
}

func (this *deleteCatID) Handle(params cats.DeleteIDParams) middleware.Responder {
	err := this.catsService.deleteCatByID(params.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return middleware.Error(404, fmt.Sprintf("cat not found with ID: %v", params.ID))
		}
		return middleware.Error(500, fmt.Sprintf("internal error: %s", err))
	}
	return cats.NewDeleteIDNoContent()
}
