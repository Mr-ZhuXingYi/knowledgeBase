package models

type Mode interface {
	String() string
}

type ModelAttrFunc func(m Mode)
type UserAttsFuncs []ModelAttrFunc

func (this UserAttsFuncs) Apply(m Mode) {
	for _, f := range this {
		f(m)
	}
}
