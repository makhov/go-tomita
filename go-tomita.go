package tomita

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

type TomitaParser struct {
	execPath string
	confPath string
	debug    bool
}

func New(execPath, confPath string) (TomitaParser, error) {
	tp := TomitaParser{
		execPath,
		confPath,
		false,
	}
	return tp, nil
}

func (tp *TomitaParser) SetDebug(debug bool) {
	tp.debug = debug
}

func (tp *TomitaParser) Run(text string) (string, error) {
	tp.debugMsg("Run Tomita")

	cmd := exec.Command(tp.execPath, tp.confPath)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Panic(err)
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		log.Panic(err)
		return "", err
	}

	go func() {
		tp.debugMsg("Write input")
		_, err = stdin.Write([]byte(text))
		stdin.Close()
		if err != nil {
			log.Panic(err)		
		}
	}()

	out := make(chan []byte)
	go func() {
		tp.debugMsg("Read output")
		output, err := ioutil.ReadAll(stdout)
		if err != nil {
			log.Panic(err)			
		}

		tp.debugMsg(fmt.Sprintf("Output: \"%s\"", string(output)))
		out <- output
	}()
	output := <-out

	return string(output), nil
}

func (tp *TomitaParser) debugMsg(message string) {
	if tp.debug {
		log.Print(message)
	}
}
