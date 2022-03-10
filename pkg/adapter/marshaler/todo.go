package marshaler

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
)

func TodoToOutput(t *todo.Todo, withID bool) string {
	content := t.Content
	if withID {
		content = t.ID + " : " + content
	}
	if t.Done {
		return "✅ " + content
	} else {
		return "✍️  " + content
	}
}

func TodosToOutput(ts todo.Todos, withID bool) []string {
	dst := make([]string, len(ts))
	for i := range ts {
		dst[i] = TodoToOutput(ts[i], withID)
	}
	return dst
}
