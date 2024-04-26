package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/api/openapi"
	"github.com/openshift-online/rh-trex/pkg/api/presenters"
	"github.com/openshift-online/rh-trex/pkg/errors"
	"github.com/openshift-online/rh-trex/pkg/services"
)

var _ RestHandler = kindnameHandler{}

type kindnameHandler struct {
	kindname services.KindNameService
	generic  services.GenericService
}

func NewKindNameHandler(kindname services.KindNameService, generic services.GenericService) *kindnameHandler {
	return &kindnameHandler{
		kindname: kindname,
		generic:  generic,
	}
}

func (h kindnameHandler) Create(w http.ResponseWriter, r *http.Request) {
	var kindname openapi.KindName
	cfg := &handlerConfig{
		&kindname,
		[]validate{
			validateEmpty(&kindname, "Id", "id"),
		},
		func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			dino := presenters.ConvertKindName(kindname)
			dino, err := h.kindname.Create(ctx, dino)
			if err != nil {
				return nil, err
			}
			return presenters.PresentKindName(dino), nil
		},
		handleError,
	}

	handle(w, r, cfg, http.StatusCreated)
}

func (h kindnameHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var patch openapi.KindNamePatchRequest

	cfg := &handlerConfig{
		&patch,
		[]validate{},
		func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()
			id := mux.Vars(r)["id"]
			found, err := h.kindname.Get(ctx, id)
			if err != nil {
				return nil, err
			}

            //patch a field

			dino, err := h.kindname.Replace(ctx, found)
			if err != nil {
				return nil, err
			}
			return presenters.PresentKindName(dino), nil
		},
		handleError,
	}

	handle(w, r, cfg, http.StatusOK)
}

func (h kindnameHandler) List(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			ctx := r.Context()

			listArgs := services.NewListArguments(r.URL.Query())
			var kindnames = []api.KindName{}
			paging, err := h.generic.List(ctx, "username", listArgs, &kindnames)
			if err != nil {
				return nil, err
			}
			dinoList := openapi.KindNameList{
				Kind:  "KindNameList",
				Page:  int32(paging.Page),
				Size:  int32(paging.Size),
				Total: int32(paging.Total),
				Items: []openapi.KindName{},
			}

			for _, dino := range kindnames {
				converted := presenters.PresentKindName(&dino)
				dinoList.Items = append(dinoList.Items, converted)
			}
			if listArgs.Fields != nil {
				filteredItems, err := presenters.SliceFilter(listArgs.Fields, dinoList.Items)
				if err != nil {
					return nil, err
				}
				return filteredItems, nil
			}
			return dinoList, nil
		},
	}

	handleList(w, r, cfg)
}

func (h kindnameHandler) Get(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			id := mux.Vars(r)["id"]
			ctx := r.Context()
			kindname, err := h.kindname.Get(ctx, id)
			if err != nil {
				return nil, err
			}

			return presenters.PresentKindName(kindname), nil
		},
	}

	handleGet(w, r, cfg)
}

func (h kindnameHandler) Delete(w http.ResponseWriter, r *http.Request) {
	cfg := &handlerConfig{
		Action: func() (interface{}, *errors.ServiceError) {
			return nil, errors.NotImplemented("delete")
		},
	}
	handleDelete(w, r, cfg, http.StatusNoContent)
}
