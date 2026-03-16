package cli

import (
	"encoding/json"
	"fmt"
)

func printJSON(payload any) {
	enc := json.NewEncoder(logOutputWriter())
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(payload); err != nil {
		logError("failed to render JSON output", "error", err)
		fmt.Fprintln(logOutputWriter(), "{}")
	}
}
