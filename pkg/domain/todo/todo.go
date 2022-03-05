package todo

type (
	Todo struct {
		ID      string
		Content string
		Done    bool
	}
	Todos    []*Todo
	FileType string
)

const (
	FileTypeMarkdown FileType = "md"
)

func (t *Todo) ToContent(fileType FileType) string {
	prefix := ""
	switch fileType {
	case FileTypeMarkdown:
		if t.Done {
			prefix = "- [x] "
		} else {
			prefix = "- [ ] "
		}

	}
	return prefix + t.Content
}

func (f FileType) String() string {
	return string(f)
}
