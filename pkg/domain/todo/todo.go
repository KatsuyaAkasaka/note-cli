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

func (t *Todo) ToLine(fileType FileType) string {
	switch fileType {
	case FileTypeMarkdown:
		return FormatMD.ToLine(t)
	}
	return t.Content
}

func (f FileType) String() string {
	return string(f)
}

func (ts Todos) FilterBy(filter func(t *Todo) bool) Todos {
	dst := make(Todos, 0, len(ts))
	for i := range ts {
		if filter(ts[i]) {
			dst = append(dst, ts[i])
		}
	}
	return dst
}
