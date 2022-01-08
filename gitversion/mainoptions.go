package gitversion

// genopts --opt_type=MainOption --prefix=Main --outfile=gitversion/mainoptions.go 'tag:string'

type MainOption func(*mainOptionImpl)

type MainOptions interface {
	Tag() string
}

func MainTag(tag string) MainOption {
	return func(opts *mainOptionImpl) {
		opts.tag = tag
	}
}

type mainOptionImpl struct {
	tag string
}

func (m *mainOptionImpl) Tag() string { return m.tag }

func makeMainOptionImpl(opts ...MainOption) *mainOptionImpl {
	res := &mainOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeMainOptions(opts ...MainOption) MainOptions {
	return makeMainOptionImpl(opts...)
}
