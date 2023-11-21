package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLang string 
	TargetLang string
	SourceText string // text to be translated
}

// The translate api url
const translateUrl = "https://translate.googleapis.com/translate_a/single"

// RequestTranslate creates a request to the google translate api
func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {

	client := &http.Client{}

	//this step does not make a request, it just initializes the request
	//client.Do actually makes the request
	req, err := http.NewRequest("GET", translateUrl, nil)

	query := req.URL.Query()
	query.Add("client", "gtx") //gtx=google translate
	//received the body object in this function, it has been sent from main func
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)
	//need to encode to json
	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatalf("1 There was a problem: %s", err)
	}

	//this is the place where the request is actually made
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("2 There was a problem: %s", err)
	}
	//close the response at the end using "defer"
	defer res.Body.Close()

	//you may get blocked if there are too many requests to the API
	if res.StatusCode == http.StatusTooManyRequests {
		// <- this syntax is how you publish to the channel
		str <- "You have been rate limited, Try again later."
		wg.Done()
		return
	}

	//you want to parse the json using gabs package
	parsedJson, err := gabs.ParseJSONBuffer(res.Body)
	if err != nil {
		log.Fatalf("3 There was a problem - %s", err)
	}

	//get the nested elements at 0th root of parsedJson variable
	nestOne, err := parsedJson.ArrayElement(0)
	if err != nil {
		log.Fatalf("4 There was a problem - %s", err)
	}

	//get one level deeper nested element
	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatalf("5 There was a problem - %s", err)
	}

	translatedStr, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatalf("6 There was a problem - %s", err)
	}
	//input into channel
	str <- translatedStr.Data().(string)
	wg.Done()
}
