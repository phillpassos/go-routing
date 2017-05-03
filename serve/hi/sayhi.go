package hi

import (
	"io"
	"net/http"
)

func SayHi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi!")
}
