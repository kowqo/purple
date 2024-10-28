package links

import (
	"net/http"
	"purple/pkg/request"
	"purple/pkg/resp"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := h.LinkRepository.GetByHash(hash)
		if err != nil {
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)

	}
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)

		for {
			existedLink, _ := h.LinkRepository.GetByHash(link.Hash)

			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		resp.JSON(w, createdLink, http.StatusOK)
	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		resp.JSON(w, id, http.StatusOK)
	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
