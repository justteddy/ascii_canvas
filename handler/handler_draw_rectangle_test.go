package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"canvas/handler/mocks"
	store "canvas/storage"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleDrawRectangle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ttl := time.Hour

	t.Run("canvas not found, should create new canvas", func(t *testing.T) {
		params := RectangleParams{
			X:       0,
			Y:       0,
			Width:   5,
			Height:  5,
			Fill:    "*",
			Outline: "-",
		}

		reqData, err := json.Marshal(params)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/canvas/canvas#1/drawRectangle", bytes.NewBuffer(reqData))
		req = wrapCtxWithParam(req, "canvasID", "canvas#1")

		mockStorage := mocks.NewMockstorage(ctrl)
		mockStorage.EXPECT().
			Get("canvas#1").
			Times(1).
			Return(nil, store.ErrNotFound)

		mockStorage.EXPECT().
			Set("canvas#1", gomock.Any(), ttl).
			Times(1).
			Return(nil)

		rr := httptest.NewRecorder()
		New(mockStorage, time.Hour).HandleDrawRectangle(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "*")
		assert.Contains(t, rr.Body.String(), "-")
	})

	t.Run("canvas found, should use previously created canvas", func(t *testing.T) {
		params := RectangleParams{
			X:       0,
			Y:       0,
			Width:   5,
			Height:  5,
			Fill:    "*",
			Outline: "-",
		}

		reqData, err := json.Marshal(params)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/canvas/canvas#1/drawRectangle", bytes.NewBuffer(reqData))
		req = wrapCtxWithParam(req, "canvasID", "canvas#1")

		mockStorage := mocks.NewMockstorage(ctrl)
		mockStorage.EXPECT().
			Get("canvas#1").
			Times(1).
			Return([]byte(`[["", ""], ["", ""]]`), nil) // no matter what size of canvas return here

		mockStorage.EXPECT().
			Set("canvas#1", gomock.Any(), ttl).
			Times(1).
			Return(nil)

		rr := httptest.NewRecorder()
		New(mockStorage, time.Hour).HandleDrawRectangle(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "*")
		assert.Contains(t, rr.Body.String(), "-")
	})

	t.Run("should respond with 500, fail on get canvas", func(t *testing.T) {
		params := RectangleParams{
			X:       0,
			Y:       0,
			Width:   5,
			Height:  5,
			Fill:    "*",
			Outline: "-",
		}

		reqData, err := json.Marshal(params)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/canvas/canvas#1/drawRectangle", bytes.NewBuffer(reqData))
		req = wrapCtxWithParam(req, "canvasID", "canvas#1")

		mockStorage := mocks.NewMockstorage(ctrl)
		mockStorage.EXPECT().
			Get("canvas#1").
			Times(1).
			Return(nil, errors.New("some error"))

		rr := httptest.NewRecorder()
		New(mockStorage, time.Hour).HandleDrawRectangle(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, `{"data": [],"error": "get canvas operation failed: some error"}`, rr.Body.String())
	})

	t.Run("should respond with 500, fail on save canvas", func(t *testing.T) {
		params := RectangleParams{
			X:       0,
			Y:       0,
			Width:   5,
			Height:  5,
			Fill:    "*",
			Outline: "-",
		}

		reqData, err := json.Marshal(params)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/canvas/canvas#1/drawRectangle", bytes.NewBuffer(reqData))
		req = wrapCtxWithParam(req, "canvasID", "canvas#1")

		mockStorage := mocks.NewMockstorage(ctrl)
		mockStorage.EXPECT().
			Get("canvas#1").
			Times(1).
			Return([]byte(`[["", ""], ["", ""]]`), nil) // no matter what size of canvas return here

		mockStorage.EXPECT().
			Set("canvas#1", gomock.Any(), ttl).
			Times(1).
			Return(errors.New("some error"))

		rr := httptest.NewRecorder()
		New(mockStorage, time.Hour).HandleDrawRectangle(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, `{"data": [],"error": "set canvas operation failed: some error"}`, rr.Body.String())
	})
}
