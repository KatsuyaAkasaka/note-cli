package todo

import "strings"

const (
	donePrefix  = "- [x] "
	doingPrefix = "- [ ] "
)

type markdown struct{}

func (m *markdown) Content(t *Todo) string {
	prefix := ""
	if t.Done {
		prefix = donePrefix
	} else {
		prefix = doingPrefix
	}
	return prefix + t.Content
}

func (m *markdown) ContentAll(ts Todos) []string {
	dst := make([]string, len(ts))
	for i := range ts {
		dst[i] = m.Content(ts[i])
	}
	return dst
}

func (m *markdown) Parse(content string) *Todo {
	if strings.HasPrefix(content, donePrefix) {
		return &Todo{
			ID:      "",
			Content: strings.Replace(content, donePrefix, "", 1),
			Done:    true,
		}
	}
	if strings.HasPrefix(content, doingPrefix) {
		return &Todo{
			ID:      "",
			Content: strings.Replace(content, doingPrefix, "", 1),
			Done:    false,
		}
	}
	return &Todo{}
}

func (m *markdown) ParseAll(contents []string) Todos {
	dst := make(Todos, len(contents))
	for i := range contents {
		dst[i] = m.Parse(contents[i])
	}
	return dst
}

func NewMarkdownFormat() Format {
	return &markdown{}
}
