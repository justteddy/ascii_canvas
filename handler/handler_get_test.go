package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"canvas/handler/mocks"
	store "canvas/storage"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandleGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("not found canvas", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/canvas/not_exist_id", nil)
		req = wrapCtxWithParam(req, "canvasID", "not_exist_id")

		mockStorage := mocks.NewMockstorage(ctrl)
		mockStorage.EXPECT().
			Get("not_exist_id").
			Times(1).
			Return(nil, store.ErrNotFound)

		rr := httptest.NewRecorder()
		New(mockStorage, time.Hour).HandleGet(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, `{"data": [],"error": "not found"}`, rr.Body.String())
	})

	t.Run("return founded canvas", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/canvas/canvas#1", nil)
		req = wrapCtxWithParam(req, "canvasID", "canvas#1")

		mockStorage := mocks.NewMockstorage(ctrl)
		mockStorage.EXPECT().
			Get("canvas#1").
			Times(1).
			Return([]byte(`[["", ""], ["", ""]]`), nil)

		rr := httptest.NewRecorder()
		New(mockStorage, time.Hour).HandleGet(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, `{"data": [["", ""], ["", ""]],"error": ""}`, rr.Body.String())
	})
}

func wrapCtxWithParam(req *http.Request, key, value string) *http.Request {
	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add(key, value)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	return req
}
