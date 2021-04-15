package delivery

import (
	"encoding/json"
	"github.com/Ythosa/gostorage/internal/dto"
	"net/http"
)

// @tags file
// @summary Upload file
// @description Saves file on the server and returns href to saved file
// @accept json
// @produce json
// @param payload body dto.File true "file payload"
// @success 200 {object} dto.FileURL
// @failure 400 {object} APIError
// @failure 500 {object} APIError
// @router /api/file/upload [post].
func (h *Handler) uploadFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &dto.File{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, http.StatusBadRequest, err)

			return
		}

		if err := req.Validate(); err != nil {
			h.error(w, http.StatusBadRequest, err)

			return
		}

		href, err := h.service.File.UploadFile(*req)
		if err != nil {
			h.error(w, http.StatusInternalServerError, err)

			return
		}

		h.respond(w, http.StatusOK, href)
	}
}

// @tags file
// @summary Updates file
// @description Updates file on the server and returns href to saved file
// @accept json
// @produce json
// @param payload body dto.File true "file payload"
// @success 200 {object} dto.FileURL
// @failure 400 {object} APIError
// @failure 500 {object} APIError
// @router /api/file/update [patch].
func (h *Handler) updateFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &dto.File{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, http.StatusBadRequest, err)

			return
		}

		if err := req.Validate(); err != nil {
			h.error(w, http.StatusBadRequest, err)

			return
		}

		href, err := h.service.File.UpdateFile(*req)
		if err != nil {
			h.error(w, http.StatusInternalServerError, err)

			return
		}

		h.respond(w, http.StatusOK, href)
	}
}

// @tags file
// @summary Deletes file
// @description Deletes file on the server and returns error if error
// @accept json
// @produce json
// @param payload body dto.FileCredentials true "file credentials payload"
// @success 200
// @failure 400 {object} APIError
// @failure 500 {object} APIError
// @router /api/file/delete [delete].
func (h *Handler) deleteFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &dto.FileCredentials{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, http.StatusBadRequest, err)

			return
		}

		if err := req.Validate(); err != nil {
			h.error(w, http.StatusBadRequest, err)

			return
		}

		err := h.service.File.DeleteFile(*req)
		if err != nil {
			h.error(w, http.StatusInternalServerError, err)

			return
		}

		h.respond(w, http.StatusOK, struct{}{})
	}
}
