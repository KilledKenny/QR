package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"rsc.io/qr"
	"strings"
)

var (
	blocks = map[[2]bool]string{
		[2]bool{false, false}: "\xE2\x96\x88",
		[2]bool{true, true}:   " ",
		[2]bool{true, false}:  "\xE2\x96\x84",
		[2]bool{false, true}:  "\xE2\x96\x80",
	}
)

func FprintlnErr(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}
func help() {
	FprintlnErr("test")
}
func args() {
	if len(os.Args) == 1 {
		//Ceck if std in is terminal
	}
	var flag string
	var i int
	for i, flag = range os.Args[1:] {
		fmt.Println(i)
		fmt.Println(flag)
	}
}
func main() {
	//args()
	defer color.Unset()
	var err error
	in := strings.Join(os.Args[1:], " ")
	if len(os.Args) == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		in = string(b)
	}
	c, err := qr.Encode(in, qr.Q)
	if err != nil {
		log.Fatal(err)
	}
	for x := -2; x < c.Size+2; x = x + 2 {
		color.Set(color.FgWhite, color.BgBlack)
		for y := -2; y < c.Size+2; y++ {
			fmt.Printf(blocks[[2]bool{
				c.Black(x, y),
				c.Black(x+1, y),
			}])
			//fmt.Printf(q)
		}
		color.Unset()
		fmt.Println("")
	}
}
