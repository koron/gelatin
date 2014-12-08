package args

import (
	"strings"

	"github.com/koron/gelatin/omap"
)

// Mode represents a mode.
type Mode struct {
	Name string
	Desc string

	Selected     bool
	SelectedMode *Mode
	Args         []string

	options  *options
	subModes *omap.OMap
}

// Parse arguments
func (m *Mode) Parse(a ...string) error {
	skipMode := false
	i := 0
	// Parse mode options
	for ; i < len(a); i++ {
		s := a[i]
		if s == "--" {
			i++
			skipMode = true
			break
		} else if s[0] != '-' {
			break
		}
		d, err := m.parseOption(s, a[i+1:]...)
		if err != nil {
			return err
		}
		i += d
	}
	// Parse sub-mode.
	if !skipMode && m.subModes.Count() > 0 && i < len(a) {
		if sm := m.Mode(a[i]); sm != nil {
			err := sm.Parse(a[i+1:]...)
			if err == nil {
				sm.Selected = true
				m.SelectedMode = sm
			}
			return err
		}
	}
	// Parse as args.
	m.Args = append(m.Args, a[i:]...)
	return nil
}

func (m *Mode) parseOption(n string, a ...string) (skip int, err error) {
	// Split into name and value
	var nv []string
	var p string
	if len(n) >= 2 && n[1] == '-' {
		nv = splitOption(n[2:])
		p = "--"
	} else {
		nv = splitOption(n[1:])
		p = "-"
	}
	// Find an option.
	o := m.options.findLong(nv[0])
	if o == nil {
		return 0, ErrorUnknownOption{Mode: m, Option: p + nv[0]}
	}
	// Parse as option.
	if len(nv) >= 2 {
		err = o.parse1(nv[0], nv[1])
	} else if o.ArgReq {
		err = o.parse1(nv[0], a[0])
		skip = 1
	} else {
		err = o.parse0(nv[0])
	}
	return skip, err
}

// DefineMode define a new sub mode.
func (m *Mode) DefineMode(name, desc string) *Mode {
	s, err := m.defineMode(name, desc)
	if err != nil {
		panic(err)
	}
	return s
}

func (m *Mode) defineMode(name, desc string) (*Mode, error) {
	s := &Mode{
		Name: name,
		Desc: desc,
	}
	if err := m.subModes.Add(name, s); err != nil {
		return nil, ErrorDuplicatedMode
	}
	s.options = newOptions()
	s.subModes = omap.New()
	return s, nil
}

// Mode get a sub mode for that name.
func (m *Mode) Mode(name string) *Mode {
	v := m.subModes.Get(name)
	if v == nil {
		return nil
	}
	return v.(*Mode)
}

// DefineOption define a new option.
func (m *Mode) DefineOption(name, longName, desc string, argReq bool) (*Option, error) {
	o := &Option{
		Name:     name,
		LongName: longName,
		Desc:     desc,
		ArgReq:   argReq,
	}
	if err := m.options.add(o); err != nil {
		return nil, err
	}
	return o, nil
}

// Option retrieve an option for this mode, if it was defined.
func (m *Mode) Option(name string) *Option {
	if o := m.options.find(name); o != nil {
		return o
	}
	if o := m.options.findLong(name); o != nil {
		return o
	}
	return nil
}

func splitOption(s string) []string {
	nv := make([]string, 0, 2)
	if n := strings.Index(s, "="); n >= 0 {
		return append(nv, s[:n], s[n+1:])
	}
	return append(nv, s)
}
