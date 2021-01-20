package models

type Mode interface {
	String() string
}

type ModelAttrFunc func(m Mode)
type UserAtts []ModelAttrFunc

func (this UserAtts) Apply(m Mode) {
	for _, f := range this {
		f(m)
	}
}
