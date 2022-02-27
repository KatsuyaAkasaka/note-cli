package todo

type Todo struct {
	ID      string
	Content string
	Done    bool
}

func (t *Todo) ToCheckMarkdown() string {
	prefix := ""
	if t.Done {
		prefix = "- [ ] "
	} else {
		prefix = "- [x] "
	}
	return prefix + t.Content
}
