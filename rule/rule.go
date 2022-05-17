package rule

type Rule interface {
	Check()
	Name() string
	Desc() string
}
