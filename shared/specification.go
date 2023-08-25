package shared

type Specification interface {
	IsSatisfiedBy(any) bool
}

type AndSpecification struct {
	One, Other Specification
}

func (s AndSpecification) IsSatisfiedBy(candidate any) bool {
	return s.One.IsSatisfiedBy(candidate) && s.Other.IsSatisfiedBy(candidate)
}

type OrSpecification struct {
	One, Other Specification
}

func (s OrSpecification) IsSatisfiedBy(candidate any) bool {
	return s.One.IsSatisfiedBy(candidate) || s.Other.IsSatisfiedBy(candidate)
}

type NotSpecification struct {
	One Specification
}

func (s NotSpecification) IsSatisfiedBy(candidate any) bool {
	return !s.One.IsSatisfiedBy(candidate)
}
