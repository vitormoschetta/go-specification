package shared

type Specification interface {
	IsSatisfiedBy(any) bool
}
