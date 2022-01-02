package cats

import (
	"log"
	"rest_api/ent"
	"rest_api/swagger/restapi/operations"
)

type Module struct {
	service *service
	// other external services if needed
}

func (this *Module) Init(api *operations.CatAPIAPI, dbClient *ent.Client) {
	this.service = &service{client: dbClient.Cat}

	api.CatsGetHandler = &getCats{catsService: this.service}
	api.CatsGetIDHandler = &getCatID{catsService: this.service}
	api.CatsPostHandler = &postCat{catsService: this.service}
	api.CatsPutIDHandler = &putCatID{catsService: this.service}
	api.CatsDeleteIDHandler = &deleteCatID{catsService: this.service}
	log.Println("Cats: module configured")
}
