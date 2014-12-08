package args

import "errors"

// ErrorDuplicatedMode raised when same sub mode is defined alreadly.
var ErrorDuplicatedMode = errors.New("same mode is defined")

// ErrorUnnamedOption raised when option defined without any names.
var ErrorUnnamedOption = errors.New("option isnot given any names")

// ErrorDuplicatedOption raised when same name/long name option have been
// defined already.
var ErrorDuplicatedOption = errors.New("same option is defined")

// ErrorUnknownOption raised when detect unknown option.
type ErrorUnknownOption struct {
	Mode   *Mode
	Option string
}

func (e ErrorUnknownOption) Error() string {
	return "unknown option:" + e.Option + " in mode:" + e.Mode.Name
}

// ErrorOptionArgRequire need argument for the option.
type ErrorOptionArgRequire struct {
	option *Option
}

func (e ErrorOptionArgRequire) Error() string {
	return "need argument for " + e.option.Label()
}

// ErrorOptionArgNotRequire don't need argument for the option.
type ErrorOptionArgNotRequire struct {
	option *Option
}

func (e ErrorOptionArgNotRequire) Error() string {
	return "not need argument for " + e.option.Label()
}
