package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
	"flag"
	"io/ioutil"
)
/*
	Want to make it where they can make a github repo with it, but it needs to make sure that its not nested!
		pass in flag for something like this
	want to make sure that default we pass in
	
	.gitignore (contains .flow.yml)
	.flow.yml (get from other program)
	main.go (with simple output from gofiler saying your file is complete)
	README.md
*/

type newFile struct{
	Name string
	Content string
}

func goFiler()(error){
	dirName := flag.String("dirName","","name of the new directory made by gofiler")
//	github := flag.Bool("makeGitRepo",false,"make a github repository")
//	files := flag.String("files","","add in extra files")
	flag.Parse()

	if *dirName ==""{
		getDirName(dirName)
	}
	fmt.Println(*dirName)
	return err

	err := os.Mkdir(*dirName,os.ModePerm)

	if err != nil{
		return err
	}

	err = os.Chdir(*dirName)

	if err != nil{
		return err
	}
	// gets file data from customizer into array
	newFiles, err := getFileData()

	if err != nil{
		return err
	}

	for _ , file := range newFiles{
		err = ioutil.WriteFile(file.Name,[]byte(file.Content),0644)
		if err != nil{
			return err
		}
	}

	return err
}

func getDirName(dirName *string){
	fmt.Print("Please enter new directory name:")
	for *dirName ==""{
		reader := bufio.NewReader(os.Stdin)
		text , _ := reader.ReadString('\n')
		*dirName = string(text)
		*dirName = strings.ReplaceAll(" ","",*dirName)
	}
	return
}

/*
This function reads through customizer and assigns name and content to new files 
*/

func getFileData()([]newFile,error){
	var array []newFile
	curDir, err := os.Getwd()

	if err != nil{
		return array, err
	}

	file , err := os.Open(curDir)

	if err != nil{
		return array, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	started := false
	var thisFile newFile
	for scanner.Scan(){
		if !started{
			scanner.Scan()
			thisFile.Name = scanner.Text()
			started = true
			continue
		}else if scanner.Text() == "..."{
			array = append(array, thisFile)
			scanner.Scan()
			thisFile.Name = scanner.Text()
			thisFile.Content = ""
		}else{
			thisFile.Content += scanner.Text()
		}
	}
	return array, nil
}


