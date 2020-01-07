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

func stringInText(s1 string, s2 string) {
	bool := strings.Contains(s2, s1)

	if bool == true {
		fmt.Println(s1)
	}
	
	waitingGroup.Done()
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

func main(){
	fmt.Println("DEBUG START")

	bd := openText("bd_test.txt")
	string_tab := splitText("prez.txt")
	for s:=0; s<len(string_tab); s++ {
		waitingGroup.Add(1)
		go stringInText(string_tab[s], bd)
	}

	fmt.Println("DEBUG WAITING")
	waitingGroup.Wait()

	fmt.Println("DEBUG END")
}