package gen

//go:generate genopts --opt_type=MainOption --prefix=Main --outfile=mainoptions.go "tag:string" "incTag:bool" "verbose"

type MainOption func(*mainOptionImpl)

type MainOptions interface {
	Tag() string
	IncTag() bool
	Verbose() bool
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

func MainVerbose(verbose bool) MainOption {
	return func(opts *mainOptionImpl) {
		opts.verbose = verbose
	}
}

type mainOptionImpl struct {
	tag     string
	incTag  bool
	verbose bool
}

func (m *mainOptionImpl) Tag() string   { return m.tag }
func (m *mainOptionImpl) IncTag() bool  { return m.incTag }
func (m *mainOptionImpl) Verbose() bool { return m.verbose }

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
