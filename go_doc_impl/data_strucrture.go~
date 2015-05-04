package main

//packages to import
import (
	"fmt"
	"io/ioutil"
)

//struct defined as per requirement 
//since dealing with website, we need
//body to struccture the web page
type Page struct {
    Title string
    Body  []byte
}

//to STORE page data into memory
//this method saves page's body to text
//return is of type error
//0600 --> read&write permissions for the current user
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

//to LOAD pages
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

//function MAIN
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}

