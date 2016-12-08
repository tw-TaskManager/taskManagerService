package model

import (
	"os/exec"
	"log"
	"fmt"
)

type Tasks struct {
	Id   string
	Task string
}

func GenerateUUID() (error, string) {
	out, err := exec.Command("uuidgen").Output();
	if (err != nil) {
		log.Fatal(err.Error())
		return err, "";
	}
	id := fmt.Sprintf("%s", out)
	return nil, id;
}
