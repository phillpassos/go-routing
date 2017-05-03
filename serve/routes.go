package serve

import (
	"regexp"

	"github.com/phillpassos/go-routing/serve/hi"
)

func (h *RegexpHandler) MapRoutes() {
	h.HandleFunc(regexp.MustCompile("/sayhi"), hi.SayHi)
}
