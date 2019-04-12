package main

import "wiretutorial/wiredi"
import "wiretutorial/simpledi"
import "flag"

var isSimple *bool

func init() {
	isSimple = flag.Bool("simple", false, "")
	flag.Parse()
}

func main() {
	if *isSimple {
		simpledi.RunSimpleDI()
	} else {
		wiredi.RunWireDi()
	}
}
