package main

// named type for holding the digits
type placeholder = [5]string

// Define  the  digits from  0  through  9 as  strings  of  size 5  using
// the  character █.  The  placeholder  can be  used  to represent  the
// complete digit  set using a multidimensional  array of size 10  X 5 as
// [10][5]string

var zero = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var digits = [...]placeholder{
	zero,
	one,
	two,
	three,
	four,
	five,
	six,
	seven,
	eight,
	nine,
}

// const sep = "░"
var sep = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

// alarm array
var alarm = [...]placeholder{
	{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	},
	{
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	},
	{
		"█  ",
		"█  ",
		"█  ",
		"█  ",
		"███",
	},
	{
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	},
	{
		"███",
		"█ █",
		"██ ",
		"█ █",
		"█ █",
	},
	{
		"█ █",
		"███",
		"█ █",
		"█ █",
		"█ █",
	},
	{
		" █ ",
		" █ ",
		" █ ",
		"   ",
		" █ ",
	},
	{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	},
}
