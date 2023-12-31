package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/mfkimbell/google-translate-project/cli"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	//flag can have bool, string etc. in this case we have taken stringVar so store value in sourceLand
	flag.StringVar(&sourceLang, "s", "en", "Source language [en]")
	//t is the paramter for target language, default is id
	flag.StringVar(&targetLang, "t", "fr", "Target language [fr]")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

func main() {
	flag.Parse()

	//NFlag just returns the number of flags that have been set
	//aka seeing if it's zero

	if flag.NFlag() == 0 {
		//if zero flags have been set, we will show the usage options
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	strChan := make(chan string) //this is the channel

	//wg add, adds a counter, done reduces by 1 and wait waits for it to hit 0

	wg.Add(1) //ADD go cli.RequestTranslate

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	go cli.RequestTranslate(reqBody, strChan, &wg) //go keyword makes it a go routine

	processedStr := strings.ReplaceAll(<-strChan, " + ", " ") //replace the + in the channel

	fmt.Printf("%s\n", processedStr)

	close(strChan)
	wg.Wait() //WAIT ON go cli.RequestTranslate
}
