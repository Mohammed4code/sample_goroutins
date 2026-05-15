package main

import (
	"fmt"
	"io"
	"net/http"
	"sync" 
)

// تعريف متغير الـ Mutex للحماية 
var m sync.Mutex

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		// نرسل من cلي msg
		msg := <-c

		//  بداية القسم الحرج 
		m.Lock() 
		
		fmt.Println(msg)
		fmt.Println("-----------------------------------")
		
		m.Unlock() 
		//  نهاية القسم الحرج
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintf("[-] %s -> Error: %v", link, err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c <- fmt.Sprintf("[-] %s -> Error reading body: %v", link, err)
		return
	}

	bodyString := string(bodyBytes)
	//اعرص اول 300 حرف
	if len(bodyString) > 300 {
		c <- fmt.Sprintf("[+] %s (Status: %s)\n%s\n... [Truncated] ...", 
			link, resp.Status, bodyString[:300])
	} else {
		c <- fmt.Sprintf("[+] %s (Status: %s)\n%s", link, resp.Status, bodyString)
	}
}

func sep(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
