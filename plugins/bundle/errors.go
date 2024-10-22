package bundle

import (
	"fmt"
	"strings"
)

// Errors represents a map of errors that occurred during a bundle load indexed by bundle name.
type Errors map[string]error

func (e Errors) Error() string {

	if len(e) == 0 {
		return "no error(s)"
	}

	var s []string
	for name, err := range e {
		s = append(s, fmt.Sprintf("'%s': %v", name, err.Error()))
	}

	return fmt.Sprintf("%d error(s) occurred:\n%s", len(e), strings.Join(s, "\n"))
}
