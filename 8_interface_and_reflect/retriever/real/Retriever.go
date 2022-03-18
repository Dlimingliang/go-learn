package real

import (
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	Context string
}

func (r Retriever) Get(str string) string {
	resp, err := http.Get(str)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}
