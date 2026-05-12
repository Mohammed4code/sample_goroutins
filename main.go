package main

import(
	"fmt"
	"net/http"
)

func main(){
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.orn",
		"https://amazon.com",
	}
	//Create a new channel called C
	c := make(chan string)

	for _, link := range links{
	// Start goroutin
		go checkLink(link,c);
	}
	//Do not stop until the links are finished.
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string){
	_, err :=http.Get(link)

	if err != nil {
		c <- link + "might be down!!"
		return 
	}
	c <- link +"is up!!"
}