package presenters

import (
	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/api/openapi"
	"github.com/openshift-online/rh-trex/pkg/util"
)

func ConvertKindName(kindname openapi.KindName) *api.KindName {
	c := &api.KindName{
		Meta: api.Meta{
			ID: util.NilToEmptyString(kindname.Id),
		},
	}

	if kindname.CreatedAt != nil {
		c.CreatedAt = *kindname.CreatedAt
		c.UpdatedAt = *kindname.UpdatedAt
	}

	return c
}

func PresentKindName(kindname *api.KindName) openapi.KindName {
	reference := PresentReference(kindname.ID, kindname)

	reference.CreatedAt = PresentTime(kindname.CreatedAt)
	reference.UpdatedAt = PresentTime(kindname.UpdatedAt)

	return openapi.KindName{
		Id:         reference.Id,
		CreatedAt:  PresentTime(*reference.CreatedAt),
		UpdatedAt:  PresentTime(*reference.UpdatedAt),
		Kind:       reference.Kind,
		Href:       reference.Href,
	}
}
