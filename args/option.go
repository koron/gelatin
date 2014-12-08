package args

// Option represents an option.
type Option struct {
	Name     string
	LongName string
	Desc     string
	ArgReq   bool
	Values   []string
}

type options struct {
	options       []*Option
	nameIndex     map[string]int
	longNameIndex map[string]int
}

func newOptions() *options {
	return &options{
		nameIndex:     make(map[string]int),
		longNameIndex: make(map[string]int),
	}
}

func (o *options) add(v *Option) error {
	if v.Name == "" && v.LongName == "" {
		return ErrorUnnamedOption
	}
	if v.Name != "" {
		if _, ok := o.nameIndex[v.Name]; ok {
			return ErrorDuplicatedOption
		}
	}
	if v.LongName != "" {
		if _, ok := o.longNameIndex[v.Name]; ok {
			return ErrorDuplicatedOption
		}
	}

	n := len(o.options)
	o.options = append(o.options, v)
	if v.Name != "" {
		o.nameIndex[v.Name] = n
	}
	if v.LongName != "" {
		o.longNameIndex[v.LongName] = n
	}
	return nil
}

func (o *options) find(s string) *Option {
	if n, ok := o.nameIndex[s]; ok {
		return o.options[n]
	}
	return nil
}

func (o *options) findLong(s string) *Option {
	if n, ok := o.longNameIndex[s]; ok {
		return o.options[n]
	}
	return nil
}

// Label returns "-X/--foobar" style name.
func (o *Option) Label() string {
	if o.Name != "" {
		if o.LongName != "" {
			return "-" + o.Name + "/--" + o.LongName
		}
		return "-" + o.Name
	}
	return "--" + o.LongName
}

func (o *Option) parse0(name string) error {
	if o.ArgReq {
		return ErrorOptionArgRequire{o}
	}
	o.Values = append(o.Values, "")
	return nil
}

func (o *Option) parse1(name, value string) error {
	if !o.ArgReq{
		return ErrorOptionArgNotRequire{o}
	}
	o.Values = append(o.Values, "")
	return nil
}
