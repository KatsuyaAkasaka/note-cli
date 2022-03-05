package marshaler

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
)

func TodoToOutput(t *todo.Todo) string {
	if t.Done {
		return "✅ " + t.Content
	} else {
		return "✍️  " + t.Content
	}
}

func TodosToOutput(ts todo.Todos) []string {
	dst := make([]string, len(ts))
	for i := range ts {
		dst[i] = TodoToOutput(ts[i])
	}
	return dst
}
