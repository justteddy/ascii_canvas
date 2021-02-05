package handler

import (
	"context"
	"fmt"
	"net/http"

	"canvas/canvas"
	store "canvas/storage"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const responseTemplate = `{"data": %s,"error":"%s"}`

func (h *Handler) storeCanvas(canvasID string, c *canvas.Canvas) ([]byte, error) {
	data, err := c.Marshal()
	if err != nil {
		return nil, err
	}

	if err := h.storage.Set(canvasID, data, h.ttl); err != nil {
		return nil, err
	}

	return data, nil
}

func (h *Handler) restoreCanvas(ctx context.Context, canvasID string) (*canvas.Canvas, error) {
	var canv *canvas.Canvas
	data, err := h.storage.Get(canvasID)
	if err != nil {
		if err == store.ErrNotFound {
			canv = canvas.New()
			log.WithContext(ctx).Infof("canvas %s not found, new one will be created", canvasID)
		} else {
			return nil, errors.Wrap(err, "get canvas operation failed")
		}
	}

	if data != nil {
		canv, err = canvas.Unmarshal(data)
		if err != nil {
			return nil, err
		}
	}

	return canv, nil
}

func writeSuccessResponse(ctx context.Context, w http.ResponseWriter, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(fmt.Sprintf(responseTemplate, data, ""))); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to write erroneous response")
	}
}

func writeErrorResponse(ctx context.Context, w http.ResponseWriter, code int, handlerErr error) {
	log.WithContext(ctx).WithError(handlerErr).Error("request failed")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if _, err := w.Write([]byte(fmt.Sprintf(responseTemplate, "[]", handlerErr))); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to write erroneous response")
	}
}
