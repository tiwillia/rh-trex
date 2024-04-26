package integration

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"gopkg.in/resty.v1"

	"github.com/openshift-online/rh-trex/pkg/api/openapi"
	"github.com/openshift-online/rh-trex/test"
)

func TestKindNameGet(t *testing.T) {
	h, client := test.RegisterIntegration(t)

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)

	// 401 using no JWT token
	_, _, err := client.DefaultApi.ApiRhTrexV1KindNamesIdGet(context.Background(), "foo").Execute()
	Expect(err).To(HaveOccurred(), "Expected 401 but got nil error")

	// GET responses per openapi spec: 200 and 404,
	_, resp, err := client.DefaultApi.ApiRhTrexV1KindNamesIdGet(ctx, "foo").Execute()
	Expect(err).To(HaveOccurred(), "Expected 404")
	Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

	dino, err := h.Factories.NewKindName(h.NewID())
    Expect(err).NotTo(HaveOccurred())

	kindname, resp, err := client.DefaultApi.ApiRhTrexV1KindNamesIdGet(ctx, dino.ID).Execute()
	Expect(err).NotTo(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(http.StatusOK))

	Expect(*kindname.Id).To(Equal(dino.ID), "found object does not match test object")
	Expect(*kindname.Kind).To(Equal("KindName"))
	Expect(*kindname.Href).To(Equal(fmt.Sprintf("/api/rh-trex/v1/kindnames/%s", dino.ID)))
	Expect(*kindname.CreatedAt).To(BeTemporally("~", dino.CreatedAt))
	Expect(*kindname.UpdatedAt).To(BeTemporally("~", dino.UpdatedAt))
}

func TestKindNamePost(t *testing.T) {
	h, client := test.RegisterIntegration(t)

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)

	// POST responses per openapi spec: 201, 409, 500
	dino := openapi.KindName{

	}

	// 201 Created
	kindname, resp, err := client.DefaultApi.ApiRhTrexV1KindNamesPost(ctx).KindName(dino).Execute()
	Expect(err).NotTo(HaveOccurred(), "Error posting object:  %v", err)
	Expect(resp.StatusCode).To(Equal(http.StatusCreated))
	Expect(*kindname.Id).NotTo(BeEmpty(), "Expected ID assigned on creation")
	Expect(*kindname.Kind).To(Equal("KindName"))
	Expect(*kindname.Href).To(Equal(fmt.Sprintf("/api/rh-trex/v1/kindnames/%s", *kindname.Id)))

	// 400 bad request. posting junk json is one way to trigger 400.
	jwtToken := ctx.Value(openapi.ContextAccessToken)
	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", jwtToken)).
		SetBody(`{ this is invalid }`).
		Post(h.RestURL("/kindnames"))

	Expect(restyResp.StatusCode()).To(Equal(http.StatusBadRequest))
}

func TestKindNamePatch(t *testing.T) {
	h, client := test.RegisterIntegration(t)

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)

	// POST responses per openapi spec: 201, 409, 500

	dino, err := h.Factories.NewKindName(h.NewID())
    Expect(err).NotTo(HaveOccurred())

	// 200 OK
	kindname, resp, err := client.DefaultApi.ApiRhTrexV1KindNamesIdPatch(ctx, dino.ID).KindNamePatchRequest(openapi.KindNamePatchRequest{}).Execute()
	Expect(err).NotTo(HaveOccurred(), "Error posting object:  %v", err)
	Expect(resp.StatusCode).To(Equal(http.StatusOK))
	Expect(*kindname.Id).To(Equal(dino.ID))
	Expect(*kindname.CreatedAt).To(BeTemporally("~", dino.CreatedAt))
	Expect(*kindname.Kind).To(Equal("KindName"))
	Expect(*kindname.Href).To(Equal(fmt.Sprintf("/api/rh-trex/v1/kindnames/%s", *kindname.Id)))

	jwtToken := ctx.Value(openapi.ContextAccessToken)
	// 500 server error. posting junk json is one way to trigger 500.
	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", jwtToken)).
		SetBody(`{ this is invalid }`).
		Patch(h.RestURL("/kindnames/foo"))

	Expect(restyResp.StatusCode()).To(Equal(http.StatusBadRequest))
}

func TestKindNamePaging(t *testing.T) {
	h, client := test.RegisterIntegration(t)

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)

	// Paging
	_, err := h.Factories.NewKindNameList("Bronto", 20)
    Expect(err).NotTo(HaveOccurred())

	list, _, err := client.DefaultApi.ApiRhTrexV1KindNamesGet(ctx).Execute()
	Expect(err).NotTo(HaveOccurred(), "Error getting kindname list: %v", err)
	Expect(len(list.Items)).To(Equal(20))
	Expect(list.Size).To(Equal(int32(20)))
	Expect(list.Total).To(Equal(int32(20)))
	Expect(list.Page).To(Equal(int32(1)))

	list, _, err = client.DefaultApi.ApiRhTrexV1KindNamesGet(ctx).Page(2).Size(5).Execute()
	Expect(err).NotTo(HaveOccurred(), "Error getting kindname list: %v", err)
	Expect(len(list.Items)).To(Equal(5))
	Expect(list.Size).To(Equal(int32(5)))
	Expect(list.Total).To(Equal(int32(20)))
	Expect(list.Page).To(Equal(int32(2)))
}

func TestKindNameListSearch(t *testing.T) {
	h, client := test.RegisterIntegration(t)

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)

	kindnames, _ := h.Factories.NewKindNameList("bronto", 20)

	search := fmt.Sprintf("id in ('%s')", kindnames[0].ID)
	list, _, err := client.DefaultApi.ApiRhTrexV1KindNamesGet(ctx).Search(search).Execute()
	Expect(err).NotTo(HaveOccurred(), "Error getting kindname list: %v", err)
	Expect(len(list.Items)).To(Equal(1))
	Expect(list.Total).To(Equal(int32(20)))
	Expect(*list.Items[0].Id).To(Equal(kindnames[0].ID))
}
