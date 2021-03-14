package controllers

import (
	"fmt"
	"net/http"

	"github.com/twinemarron/bookstore_items-api/domain/items"
	"github.com/twinemarron/bookstore_items-api/services"
	"github.com/twinemarron/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error json to the user.
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemService.Create(item)
	if err != nil {
		// TODO: Return error json to the user.
	}

	fmt.Println(result)
	// TODO: Return created item with HTTP status 201 - Created.
}
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}
