package todo

import (
	"strings"
)

const (
	donePrefix  = "- [x] "
	doingPrefix = "- [ ] "
	Delimiter   = " : "
)

type markdown struct{}

func (m *markdown) ToLine(t *Todo) string {
	prefix := doingPrefix
	if t.Done {
		prefix = donePrefix
	}

	return prefix + t.ID + Delimiter + t.Content
}

func (m *markdown) ToLineAll(ts Todos) []string {
	dst := make([]string, len(ts))
	for i := range ts {
		dst[i] = m.ToLine(ts[i])
	}

	return dst
}

func (m *markdown) Parse(content string) *Todo {
	contentWithID := ""
	done := false
	if strings.HasPrefix(content, donePrefix) {
		contentWithID = strings.Replace(content, donePrefix, "", 1)
		done = true
	}
	if strings.HasPrefix(content, doingPrefix) {
		contentWithID = strings.Replace(content, doingPrefix, "", 1)
		done = false
	}
	s := strings.Split(contentWithID, Delimiter)

	return &Todo{
		ID:      s[0],
		Content: strings.Join(s[1:], Delimiter),
		Done:    done,
	}
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
