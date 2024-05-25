package encoder

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type MultiEncoder struct {
	w http.ResponseWriter
}

func NewMultiEncoder(w http.ResponseWriter) *MultiEncoder {
	return &MultiEncoder{w: w}
}

func (e *MultiEncoder) Encode(r *http.Request, data any) error {
	acceptHeader := r.Header.Get("Accept")
	switch {
	case acceptHeader == "application/xml" || acceptHeader == "text/xml":
		e.w.Header().Set("Content-Type", "application/xml")
		if err := xml.NewEncoder(e.w).Encode(data); err != nil {
			return fmt.Errorf("encode: %w", err)
		}
	default:
		e.w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(e.w).Encode(data); err != nil {
			return fmt.Errorf("encode: %w", err)
		}
	}
	return nil
}
