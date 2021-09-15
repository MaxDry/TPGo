package main

import (
 "fmt"
 "net/http"
 "time"
 "os"
 "bufio"
)

func firstHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprintf(w, time.Now().Format("15h04"))
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		entry := req.PostForm.Get("entry")
		author := req.PostForm.Get("author")
		result := author + ":" + entry
		fmt.Fprintf(w, result)

		readEntries := readEntries()
		saveEntries(readEntries, entry)
	}
}

func saveEntries(readAuthors string, description string) {
	saveFile, err := os.OpenFile("./save.data", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	defer saveFile.Close()

	w := bufio.NewWriter(saveFile)
	if err == nil {
		fmt.Fprintf(w, fmt.Sprintf("%s\n", description))
	}
	w.Flush()
}

func entriesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprintf(w, readEntries())
	}
}

func readEntries() string {
	saveData, err := os.ReadFile("./save.data")
	if err == nil {
		fmt.Println(string(saveData))
		return string(saveData)
	}
	return "Error"
}

func main() {
	http.HandleFunc("/", firstHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/entries", entriesHandler)
	http.ListenAndServe(":9002", nil)
}