package uuid

type Repository interface {
	Gen() string
}
