package pipeline

import (
	"fmt"
	"net/http"
	"strings"
)

type DefferedResponseWriter struct {
	http.ResponseWriter
	strings.Builder
	statusCode int
}

func (dw *DefferedResponseWriter) Write(data []byte) (int, error) {
	return dw.Builder.Write(data)
}

func (dw *DefferedResponseWriter) FlushData() {
	if dw.statusCode == 0 {
		dw.statusCode = http.StatusOK
	}
	fmt.Println(dw.ResponseWriter.Header())
	dw.ResponseWriter.WriteHeader(dw.statusCode)
	dw.ResponseWriter.Write([]byte(dw.Builder.String()))
}

func (dw *DefferedResponseWriter) WriteHeader(statusCode int) {
	dw.statusCode = statusCode
}
