# google-translate-tool

**Tools used:**
* `Go Routines` for concurrency of task execution
* `Channels` for communication of Go Routines
* `Wait Groups` sync mechanism that blocks code to allow for other functions to complete execution, similar to `await` keyword


One thing I think is important to understand for this project is that it utilizes concurrency through Go Routines, aka tasks are always being actively worked on when there is downtime during API calls and things alike, however, there is no paralellism, such as when you have dynamically scaling worker servers/containers. 

We use a channel to communicate between goroutines (concurrent threads of execution) and synchronize their operations.

``` Go
strChan := make(chan string)
```

We also create a wait group, shown below, which we pass into our go routine defined in "cli.go":

``` Go
var wg sync.WaitGroup
```

Here, a channel of strings named strChan is created using the make function. Channels are used for communication between goroutines. In this case, the channel is specifically designed to transmit strings.

``` Go
go cli.RequestTranslate(reqBody, strChan, &wg)
```

This line invokes the RequestTranslate function from the cli package as a goroutine. The strChan channel is passed as an argument to this function. The goroutine runs concurrently with the main program.

``` Go
processedStr := strings.ReplaceAll(<-strChan, " + ", " ")
```

The <-strChan syntax is used to receive data from the channel (strChan). This operation blocks until there is data available on the channel. The received string is then processed using strings.ReplaceAll to replace any occurrences of " + " with an empty string

In go.cli, we connect to the google api

``` Go
const translateUrl = "https://translate.googleapis.com/translate_a/single"
```

Then we make the GET request:

``` Go
req, err := http.NewRequest("GET", translateUrl, nil)
```

And we add input parameters then actually run the request:

``` Go
res, err := client.Do(req)
	if err != nil {
		log.Fatalf("2 There was a problem: %s", err)
	}
```

We then parse through the request result and output the data we're interested in (the translated text).

Here is the output with no flags:

<img width="606" alt="Screenshot 2023-11-21 at 5 13 00 PM" src="https://github.com/mfkimbell/google-translate-tool/assets/107063397/eb1c4960-0783-4e28-9d06-f1b09d820268">

Here is the output with correct flags:

<img width="399" alt="Screenshot 2023-11-21 at 8 31 59 PM" src="https://github.com/mfkimbell/google-translate-tool/assets/107063397/7b7d6d46-c6a4-4f36-9acd-6d8b8e12c21c">
