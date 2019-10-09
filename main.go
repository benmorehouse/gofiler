package main

import(
	"log"
)

func main(){
	/*
	first thing it needs to do is get working directory so that we know where to put new directory
	we also need to read in which files they want from the file flag thing that they have
	then we make the directory itself
	THEN we go through and concurrently make it make these files within this directory (while making sure to change to directory)
	*/
	err := goFiler();
	if err != nil{
		log.Fatal("goFiler returned an error in main:",err)
	}

}



