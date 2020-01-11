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

//Type phrase. Contient un string qui correspond au contenu de la phrase et un booléen qui indique si la phrase est une phrase plagiée ou non
type sentence struct{
	text string
	boolean bool
}

var wg sync.WaitGroup
var wg2 sync.WaitGroup

//Fonction qui prend en paramètre un fichier .txt, l'ouvre et nous renvoie ce texte sous forme de String 
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

//Fonction qui prend en paramètre un fichier .txt, l'ouvre et nous renvoie un tableau de sentence
func splitText(text string) []sentence{

	//Ouvre le texte et stocke le String lui correspondant
	string := openText(text)

	//On prend la chaine de caractere qu'on split en phrase avec le point comme separateur
	string_tab := strings.Split(string, ".")

	var sentence_tab []sentence
	var sentence_in_text sentence

	for s:=0; s<len(string_tab); s++{
		sentence_in_text = sentence{text: string_tab[s], boolean : false}
		sentence_tab = append(sentence_tab, sentence_in_text)
		//fmt.Println(string(sentence_tab[s].boolean))
	}

  	fmt.Println("DEBUG Fin du split")
  	
  	return sentence_tab
}

//Fonction qui renvoie un tableau contenant le nom de tous les fichiers .txt du répertoire courant
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

//Fonction qui prend en paramètre la phrase à comparer, le texte dans lequel il faut regarder et son nom et 
//qui affiche dans la console si cette phrase appartient au texte
func stringInText(s1 *sentence, s2 string, s2_name string){
	
	//On teste si la phrase est dans le texte
	string_in_text := strings.Contains(s2, s1.text)

	//Si oui, on affiche
	if string_in_text {
		fmt.Println("From " + s2_name + " : " + s1.text)
		s1.boolean = true
	}

	wg2.Done()
}

//Fonction qui prend en paramètre le texte de la base de données et le texte reçu pour les comparer
func compareText(database_text_name string, received_text *[]sentence){
	
	//On ouvre le fichier de la base de données
	database_text := openText(database_text_name)

	//Pour chaque phrase du texte reçu on regarde si elle est dans le texte de la base de données
	for s:=0; s<len(*received_text); s++ {
		wg2.Add(1)
		go stringInText(&(*received_text)[s], database_text, database_text_name)
	}

	wg2.Wait()

	wg.Done()
}

func main() {
	fmt.Println("DEBUG START")
	received_text := splitText("prez.txt")
	text_files := textFilesInDirectory()
	for t:=0; t<len(text_files); t++ {
		if text_files[t] != "prez.txt"{
			wg.Add(1)
			go compareText(text_files[t], &received_text)
		}
	}

	fmt.Println("DEBUG WAITING")
	wg.Wait()

	var counter int
	var total int

	for k:=0; k<len(received_text); k++{
		if received_text[k].boolean == true {
			counter = counter + len(received_text[k].text)
		}
		total = total + len(received_text[k].text)
	}

	ratio := (float32(counter)/float32(total))*100

	fmt.Printf("Plagiarism score : %f%", ratio)
	fmt.Println()
	fmt.Println("DEBUG END")
}