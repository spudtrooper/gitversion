package gen

// genopts --opt_type=MainOption --prefix=Main --outfile=gen/mainoptions.go 'tag:string' 'incTag:bool'

type MainOption func(*mainOptionImpl)

type MainOptions interface {
	Tag() string
	IncTag() bool
}

func MainTag(tag string) MainOption {
	return func(opts *mainOptionImpl) {
		opts.tag = tag
	}
}

func MainIncTag(incTag bool) MainOption {
	return func(opts *mainOptionImpl) {
		opts.incTag = incTag
	}
}

type mainOptionImpl struct {
	tag    string
	incTag bool
}

func (m *mainOptionImpl) Tag() string  { return m.tag }
func (m *mainOptionImpl) IncTag() bool { return m.incTag }

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
