package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/services"
)

type FeedHandler interface {
	GetFeed(w http.ResponseWriter, r *http.Request)
}

type feedHandler struct {
	rc pkgs.RequestContext
	fs services.FeedService
}

func NewFeedHandler(
	requestContext pkgs.RequestContext,
	feedService services.FeedService) FeedHandler {
	return &feedHandler{
		rc: requestContext,
		fs: feedService,
	}
}

func (f *feedHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "feed"),
		slog.String("method", "GetFeed"),
	)

	userID, ok := f.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("userID not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10
	offset := 0

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}
	if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
		offset = o
	}

	feed, err := f.fs.GetFeed(r.Context(), userID, limit, offset)
	if err != nil {
		logger.Error("error getting feed", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, feed)
}
