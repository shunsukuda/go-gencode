package gencode

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	DoGoFmt = true
)

func GenCode(name string, input string, output string, data interface{}) (inBytes int, outBytes int) {
	inF, err := os.Open(input)
	defer inF.Close()
	if err != nil {
		log.Fatal(err, "cannot open input file!")
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, inF)
	inBytes = buf.Len()
	fmt.Printf("input: %s %d bytes\n", input, inBytes)
	tmpl, err := template.New(name).Parse(buf.String())
	if err != nil {
		log.Fatal(err, "cannot new templete!")
	}
	outF, err := os.Create(output)
	defer outF.Close()
	if err != nil {
		log.Fatal(err, "cannot create output file!")
	}
	if err = tmpl.Execute(outF, data); err != nil {
		log.Fatal(err, "cannot do template execute!")
	}
	if err = exec.Command("go", "fmt", output).Start(); err != nil {
		log.Fatal(err, "cannot do go fmt!")
	}
	buf.Reset()
	io.Copy(buf, outF)
	outBytes = buf.Len()
	fmt.Printf("output: %s %d bytes\n", output, outBytes)
	return
}
