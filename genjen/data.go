package main

var keywords = []string{"break", "default", "func", "select", "chan", "else", "const", "fallthrough", "type", "continue", "var", "goto", "defer", "go", "range"}

var identifiers = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "true", "false", "iota", "nil", "err"}

var groups = []struct {
	name        string   // name of the function / method
	comment     string   // comment appended to name
	variadic    bool     // is the parameter variadic?
	opening     string   // opening token
	closing     string   // closing token
	separator   string   // separator token
	parameters  []string // parameter names
	preventFunc bool     // prevent the fooFunc function/method
}{
	{
		name:       "Parens",
		comment:    "renders a single item in parenthesis. Use for type conversion or to specify evaluation order.",
		variadic:   false,
		opening:    "(",
		closing:    ")",
		separator:  "",
		parameters: []string{"item"},
	},
	{
		name:       "List",
		comment:    "renders a comma separated list. Use for multiple return functions.",
		variadic:   true,
		opening:    "",
		closing:    "",
		separator:  ",",
		parameters: []string{"items"},
	},
	{
		name:       "Values",
		comment:    "renders a comma separated list enclosed by curly braces. Use for slice literals.",
		variadic:   true,
		opening:    "{",
		closing:    "}",
		separator:  ",",
		parameters: []string{"values"},
	},
	{
		name:       "Index",
		comment:    "renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.",
		variadic:   true,
		opening:    "[",
		closing:    "]",
		separator:  ":",
		parameters: []string{"items"},
	},
	{
		name:       "Block",
		comment:    "renders a statement list enclosed by curly braces. Use for code blocks.",
		variadic:   true,
		opening:    "{",
		closing:    "}",
		separator:  "\n",
		parameters: []string{"statements"},
	},
	{
		name:       "Defs",
		comment:    "renders a statement list enclosed in parenthesis. Use for definition lists.",
		variadic:   true,
		opening:    "(",
		closing:    ")",
		separator:  "\n",
		parameters: []string{"definitions"},
	},
	{
		name:       "Call",
		comment:    "renders a comma separated list enclosed by parenthesis. Use for function calls.",
		variadic:   true,
		opening:    "(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"params"},
	},
	{
		name:       "Params",
		comment:    "renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.",
		variadic:   true,
		opening:    "(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"params"},
	},
	{
		name:       "Assert",
		comment:    "renders a period followed by a single item enclosed by parenthesis. Use for type assertions.",
		variadic:   false,
		opening:    ".(",
		closing:    ")",
		separator:  "",
		parameters: []string{"typ"},
	},
	{
		name:       "Map",
		comment:    "renders the keyword followed by a single item enclosed by square brackets. Use for map definitions.",
		variadic:   false,
		opening:    "map[",
		closing:    "]",
		separator:  "",
		parameters: []string{"typ"},
	},
	{
		name:       "If",
		comment:    "renders the keyword followed by a semicolon separated list.",
		variadic:   true,
		opening:    "if ",
		closing:    "",
		separator:  ";",
		parameters: []string{"conditions"},
	},
	{
		name:       "Return",
		comment:    "renders the keyword followed by a comma separated list.",
		variadic:   true,
		opening:    "return ",
		closing:    "",
		separator:  ",",
		parameters: []string{"results"},
	},
	{
		name:       "For",
		comment:    "renders the keyword followed by a semicolon separated list.",
		variadic:   true,
		opening:    "for ",
		closing:    "",
		separator:  ";",
		parameters: []string{"conditions"},
	},
	{
		name:       "Switch",
		comment:    "renders the keyword followed by a semicolon separated list.",
		variadic:   true,
		opening:    "switch ",
		closing:    "",
		separator:  ";",
		parameters: []string{"conditions"},
	},
	{
		name:       "Interface",
		comment:    "renders the keyword followed by a method list enclosed by curly braces.",
		variadic:   true,
		opening:    "interface{",
		closing:    "}",
		separator:  "\n",
		parameters: []string{"methods"},
	},
	{
		name:       "Struct",
		comment:    "renders the keyword followed by a field list enclosed by curly braces.",
		variadic:   true,
		opening:    "struct{",
		closing:    "}",
		separator:  "\n",
		parameters: []string{"fields"},
	},
	{
		name:       "Case",
		comment:    "renders the keyword followed by a comma separated list.",
		variadic:   true,
		opening:    "case ",
		closing:    "",
		separator:  ",",
		parameters: []string{"cases"},
	},
	{
		name:       "Sel",
		comment:    "renders a period separated list. Use for a chain of selectors.",
		variadic:   true,
		opening:    "",
		closing:    "",
		separator:  ".",
		parameters: []string{"selectors"},
	},
	{
		name:       "Append",
		comment:    "renders the append built-in function.",
		variadic:   true,
		opening:    "append(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"args"},
	},
	{
		name:       "Cap",
		comment:    "renders the cap built-in function.",
		variadic:   false,
		opening:    "cap(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"v"},
	},
	{
		name:       "Close",
		comment:    "renders the close built-in function.",
		variadic:   false,
		opening:    "close(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"c"},
	},
	{
		name:       "Complex",
		comment:    "renders the complex built-in function.",
		variadic:   false,
		opening:    "complex(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"r", "i"},
	},
	{
		name:       "Copy",
		comment:    "renders the copy built-in function.",
		variadic:   false,
		opening:    "copy(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"dst", "src"},
	},
	{
		name:       "Delete",
		comment:    "renders the delete built-in function.",
		variadic:   false,
		opening:    "delete(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"m", "key"},
	},
	{
		name:       "Imag",
		comment:    "renders the imag built-in function.",
		variadic:   false,
		opening:    "imag(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"c"},
	},
	{
		name:       "Len",
		comment:    "renders the len built-in function.",
		variadic:   false,
		opening:    "len(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"v"},
	},
	{
		name:        "Make",
		comment:     "renders the make built-in function. The final parameter of the make function is optional, so it is represented by a variadic parameter list.",
		variadic:    true,
		opening:     "make(",
		closing:     ")",
		separator:   ",",
		parameters:  []string{"args"},
		preventFunc: true, // the underlying function is not variadic, so we prevent the MakeFunc variation
	},
	{
		name:       "New",
		comment:    "renders the new built-in function.",
		variadic:   false,
		opening:    "new(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"typ"},
	},
	{
		name:       "Panic",
		comment:    "renders the panic built-in function.",
		variadic:   false,
		opening:    "panic(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"v"},
	},
	{
		name:       "Print",
		comment:    "renders the print built-in function.",
		variadic:   true,
		opening:    "print(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"args"},
	},
	{
		name:       "Println",
		comment:    "renders the println built-in function.",
		variadic:   true,
		opening:    "println(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"args"},
	},
	{
		name:       "Real",
		comment:    "renders the real built-in function.",
		variadic:   false,
		opening:    "real(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"c"},
	},
	{
		name:       "Recover",
		comment:    "renders the recover built-in function.",
		variadic:   false,
		opening:    "recover(",
		closing:    ")",
		separator:  ",",
		parameters: []string{},
	},
}
