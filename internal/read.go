package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
)

func Read(filename string, vbscript domain.VBScript) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		CountControlFlow(vbscript, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("err: %v", err)
	}
}