package main

import (
	"io"
	"log"
	"os"
)

/*
https://www.digitalocean.com/community/tutorials/understanding-defer-in-go

One of the primary uses of a defer statement is for cleaning up resources, such as open files, network connections, and database handles.
When your program is finished with these resources, it’s important to close them to avoid exhausting the program’s limits and to allow other programs access to those resources.
defer makes our code cleaner and less error prone by keeping the calls to close the file/resource in proximity to the open call.

A defer statement adds the function call following the defer keyword onto a stack. All of the calls on that stack are called when the function in which they were added returns. Because the calls are placed on a stack, they are called in last-in-first-out order.
*/

func simpleDeffer() {
	log.Println("Hi 1")

	defer log.Println("Bye") // defer statement is executed, and places fmt.Println("Bye") on a list to be executed prior to the function returning
	log.Println("Hi 2")

	// fmt.Println*("Bye") is now invoked, as we are at the end of the function scope
	/*
	Although this code illustrates the order in which defer would be run, it’s not a typical way it would be used when writing a Go program. 
	It’s more likely we are using defer to clean up a resource, such as a file handle. 
	*/	
}

func write_to_file(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close() // We have now ensured that even if we add more code and create another branch that exits the function in the future, we will always clean up and close the file.
	// However, we have introduced yet another bug by adding the defer. We are no longer checking the potential error that can be returned from the Close method. This is because when we use defer, there is no way to communicate any return value back to our function.
	
	_, err = io.WriteString(file, text)
	
	if err != nil {
		// If the call to io.WriteString fails, the function will return without closing the file and releasing the resource back to the system.
		// We could fix the problem by adding another file.Close() statement, which is how you would likely solve this in a language without defer
		// file.Close()
		// Now even if the call to io.WriteString fails, we will still close the file. While this was a relatively easy bug to spot and fix, with a more complicated function, it may have been missed.
		// Instead of adding the second call to file.Close(), we can use a defer statement to ensure that regardless of which branches are taken during execution, we always call Close().
		return err
	}

	// file.Close()
	
	// In Go, it is considered a safe and accepted practice to call Close() more than once without affecting the behavior of your program. If Close() is going to return an error, it will do so the first time it is called. This allows us to call it explicitly in the successful path of execution in our function.
	return file.Close()
}

func fileCopy(source string, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close() // closing of the source file we just opened.

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	log.Printf("Copied %d bytes from %s to %s\n", n, source, destination)

	if err := src.Close(); err != nil {
		return err
	}

	return dst.Close()
}

func main() {
	simpleDeffer() // Output: Hi 1, Hi 2, Bye

	err := write_to_file("readme.txt", "This is a readme file")
	if err != nil {
		log.Fatal("failed to write file:", err)
	}

	fileCopy("readme.txt", "readme2.txt")
}