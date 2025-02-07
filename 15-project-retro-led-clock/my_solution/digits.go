package main

type digit []string

var zero digit = []string{

	"█████████████",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█████████████",
}

var one digit = []string{

	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
}

var two digit = []string{

	"█████████████",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"█████████████",
	"█            ",
	"█            ",
	"█            ",
	"█            ",
	"█████████████",
}
var three digit = []string{

	"█████████████",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"█████████████",
	"           ██",
	"           ██",
	"           ██",
	"           ██",
	"█████████████",
}

var four digit = []string{

	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█████████████",
	"            █",
	"            █",
	"            █",
	"            █",
}

var five digit = []string{

	"█████████████",
	"█            ",
	"█            ",
	"█            ",
	"█            ",
	"█████████████",
	"            █",
	"            █",
	"            █",
	"            █",
	"█████████████",
}

var six digit = []string{

	"█████████████",
	"█            ",
	"█            ",
	"█            ",
	"█            ",
	"█████████████",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█████████████",
}
var seven digit = []string{

	"█████████████",
	"            █",
	"            █",
	"            █",
	"            █",
	"            █",
	"            █",
	"            █",
	"            █",
	"            █",
	"            █",
}
var eight digit = []string{

	"█████████████",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█████████████",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█████████████",
}

var nine digit = []string{

	"█████████████",
	"█           █",
	"█           █",
	"█           █",
	"█           █",
	"█████████████",
	"            █",
	"            █",
	"            █",
	"            █",
	"█████████████",
}

var digits = [...]*digit{&zero, &one, &two, &three, &four, &five, &six, &seven, &eight, &nine}

var seperator digit = []string{

	"             ",
	"     █ █     ",
	"      █      ",
	"     █ █     ",
	"             ",
	"             ",
	"     █ █     ",
	"      █      ",
	"     █ █     ",
	"             ",
	"             ",
}
