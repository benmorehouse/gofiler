package main

import(
	"fmt"
	"os/exec"
	"os"
)

func gofile(){
	curDir , err := os.GetWd()

	if err != nil{
		log.Fatal("Couldnt get the current directory")
	}

	os.Mkdir(
}

