package marshaler

import (
	"github.com/KatsuyaAkasaka/nt/pkg/domain/todo"
)

const todoDelimiter = " : "

func TodoToOutput(t *todo.Todo, withID bool) string {
	content := t.Content
	if withID {
		content = t.ID + todoDelimiter + content
	}
	if t.Done {
		return "✅" + todoDelimiter + content
	}

	return "✍️ " + todoDelimiter + content
}

type OutputTodosParams struct {
	WithID bool
}

func OutputTodos(ts todo.Todos, params *OutputTodosParams) []string {
	dst := make([]string, len(ts))
	for i := range ts {
		dst[i] = TodoToOutput(ts[i], params.WithID)
	}

	return dst
}
