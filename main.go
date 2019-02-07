package main

import (
	"encoding/base64"
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
	optionQR_Mode = QRModeUTF
)

type QR_Mode int

const (
	QRModeUTF QR_Mode = iota
	QRModeiTerm
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
	for _, env := range os.Environ() {
		s := strings.SplitN(env, "=", 2)
		if len(s) == 2 {
			if s[0] == "TERM_PROGRAM" && s[1] == "iTerm.app" {
				optionQR_Mode = QRModeiTerm
			}
		}
	}

	//var flag string
	//var i int
	//for i, flag = range os.Args[1:] {
	//	fmt.Println(i)
	//	fmt.Println(flag)
	//}

}
func main() {
	args()
	in := strings.Join(os.Args[1:], " ")
	if len(os.Args) == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		in = strings.TrimRight(string(b), "\n")
	}
	if optionQR_Mode == QRModeUTF {
		QR_UTF(in)
	} else if optionQR_Mode == QRModeiTerm {
		QR_iTerm(in)
	} else {
		log.Fatal("oh no...")
	}

}
func QR_UTF(in string) {
	defer color.Unset()
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
func QR_iTerm(in string) {
	code, err := qr.Encode(in, qr.Q)
	if err != nil {
		log.Fatal(err)
	}
	img := code.PNG()

	//img, err := ioutil.ReadFile("go.ico")
	//fatal(err)
	data := base64.StdEncoding.EncodeToString(img)
	fmt.Printf("\033]1337;File=;inline=1:")
	fmt.Printf("%s", data)
	fmt.Printf("\a\033\\")
}
