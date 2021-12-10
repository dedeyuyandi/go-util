package query

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// this is fucntion for dinamic filter for query where
// for eg return : WHERE "city" = "indonesia" AND "country": "bandung"
func Where(request map[string]interface{}) string {
	var w []string
	for k, v := range request {
		switch v.(type) {
		case uint64, uint32, int32, int64, uint8, uint16, int8, int16, int:
			w = append(w, fmt.Sprintf("%s = %d", k, v))
		case float32, float64, bool, uuid.UUID:
			w = append(w, fmt.Sprintf("%s = %v", k, v))
		default:
			w = append(w, fmt.Sprintf("%s = '%s'", k, v))
		}
	}
	if len(w) > 0 {
		return fmt.Sprintf("WHERE %s", strings.Join(w, " AND "))
	}
	return ""
}
