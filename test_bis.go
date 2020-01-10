package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "strings"
    "sync"
    //"github.com/agnivade/levenshtein"
)

var waitingGroup sync.WaitGroup
var waitingGroup2 sync.WaitGroup

func openText(text string) string{
	//Ouverture du fichier
    file, err := os.Open(text)
    if err != nil {
        log.Fatal(err)
    }

    //A la fin on fermera le fichier
    defer file.Close()

    //ReadAll nous renvoie un immense tableau de bytes
	bytes, err := ioutil.ReadAll(file)

	//On convertit le tableau de bytes en un string
	string := string(bytes)

	return string
}

func splitText(text string) []string{

	string := openText(text)

	//On prend la chaine de caractere qu'on split en phrase avec le point comme separateur
	string_tab := strings.Split(string, ".")

	//Affichage des phrases
	//for i := 0; i<len(string_tab); i++ {
  	//	fmt.Println(string_tab[i])
  	//}

  	fmt.Println("DEBUG Fin du split")
  	
  	return string_tab
}

func stringInText(s1 string, s2 string, s2_name string) {
	bool := strings.Contains(s2, s1)

	if bool == true {
		fmt.Println("From " + s2_name + " : " + s1)
	}
	
	waitingGroup2.Done()
	//distance := levenshtein.ComputeDistance(s1, s2)
	//fmt.Printf("The distance between %s and %s is %d.\n", s1, s2, distance)
	// Output:
	// The distance between kitten and sitting is 3.
}

func textFilesInDirectory() []string {
	files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }
	
	var textFiles []string
    
    for _, file := range files {
        if strings.Contains(file.Name(), ".txt") {
        	textFiles = append(textFiles, file.Name())
        }
    }

    return textFiles
}

func compareText(text string, textToCompare []string){
	bd := openText(text)
	for s:=0; s<len(textToCompare); s++ {
		waitingGroup2.Add(1)
		go stringInText(textToCompare[s], bd, text)
	}

	waitingGroup2.Wait()

	waitingGroup.Done()
}

func main() {
	fmt.Println("DEBUG START")
	receivedText := splitText("prez.txt")
	textFiles := textFilesInDirectory()
	for t:=0; t<len(textFiles); t++ {
		waitingGroup.Add(1)
		go compareText(textFiles[t], receivedText)
	}

	fmt.Println("DEBUG WAITING")
	waitingGroup.Wait()

	fmt.Println("DEBUG END")
}