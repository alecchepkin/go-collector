package components

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type FormatRequest struct {
}

// formatRequest generates ascii representation of a request
func (t FormatRequest) Format(r *http.Request) string {
	var request []string
	// Add the request string
	u := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, u)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	// If this is a POST, add post data
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
		}
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	var res bytes.Buffer
	res.WriteString("\n")
	res.WriteString(strings.Join(request, "\n"))
	res.WriteString("\n")
	res.WriteString("\n")

	return res.String()
}
