package todo

type Format interface {
	ToLine(t *Todo) string
	ToLineAll(ts Todos) []string
	Parse(content string) *Todo
	ParseAll(contents []string) Todos
}

var FormatMD = NewMarkdownFormat()
