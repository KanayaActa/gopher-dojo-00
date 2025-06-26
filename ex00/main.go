package main

import (
	"fmt"
	"os"
	"convert-ex00/convert"
)

func main() {
	cnfig, err := convert.ParseArgs(os.Args[1:])
	if (err == nil)
	{

	}
	err := convert.Convert(cnfig)
	if (err == nil)
	{
		
	}
	fmt.Println(args)
}
