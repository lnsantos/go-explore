package domain

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func GetCommandOutput(
	command string,
	arg ...string,
) string {

	var out bytes.Buffer
	var stder bytes.Buffer

	cmd := exec.Command(command, arg...)
	cmd.Stdout = &out
	cmd.Stderr = &stder

	if err := cmd.Start(); err != nil {
		log.Fatal(fmt.Sprint(err, "::", stder.String()))
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(fmt.Sprint(err, "::", stder.String()))
	}

	return out.String()
}
