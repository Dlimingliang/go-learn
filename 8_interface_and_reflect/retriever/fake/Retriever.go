package fake

type Retriever struct {
	Context string
}

func (r *Retriever) String() string {
	return "I am stringer"
}

func (r *Retriever) Post(url string, param map[string]string) {
	r.Context = param["context"]
}

func (r *Retriever) Get(str string) string {
	return r.Context
}
