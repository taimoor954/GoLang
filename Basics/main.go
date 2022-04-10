// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory). 
package main


import "fmt";


//@NOTE This is an imported 3rd party package
import "rsc.io/quote/v4"

// @NOTE Implement a main function to print a message to the console. A main function executes by default when you run the main package. 
func main() {
    name := "Go Developers"
    fmt.Println("Azure for", name)

    //Imported package usage
    fmt.Println(quote.Go());
    fmt.Println(quote.Hello());
    fmt.Println(quote.Opt());
}

//When 3rd parties are downloaded two things will happen
//1) Go will add the quote module as a requirement

//2)go.sum file for use in authenticating the module
// When the go command downloads a module zip file or go.mod file into the module cache, it computes a cryptographic hash and compares it with a known value to verify the file hasn’t changed since it was first downloaded. 
// The go command reports a security error if a downloaded file does not have the correct hash.



// go run main.go To run file in cmd