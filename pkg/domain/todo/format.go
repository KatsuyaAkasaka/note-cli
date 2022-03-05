package todo

type Format interface {
	Content(t *Todo) string
	ContentAll(ts Todos) []string
	Parse(content string) *Todo
	ParseAll(contents []string) Todos
}

var (
	FormatMD = NewMarkdownFormat()
)
