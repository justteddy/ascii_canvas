package handler

import (
	"context"
	"net/http"
	"strings"

	"canvas/canvas"
	store "canvas/storage"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

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

func writeErrorResponse(ctx context.Context, w http.ResponseWriter, code int, err error) {
	log.WithContext(ctx).WithError(err).Error("request failed")
	w.WriteHeader(code)
	if _, err = w.Write([]byte(`{"error":"` + strings.TrimSpace(err.Error()) + `"}`)); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to write erroneous response")
	}
}
