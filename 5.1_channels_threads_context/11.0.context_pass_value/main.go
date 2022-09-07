package main

import (
	"context"
	"fmt"
)
/*
! Note: Contexts can be a powerful tool with all the values they can hold, but a balance needs to be struck between data being stored in a context and data being passed to a function as parameters.
! It may seem tempting to put all of your data in a context and use that data in your functions instead of parameters, but that can lead to code that is hard to read and maintain.
! A good rule of thumb is that any data required for a function to run should be passed as parameters. Sometimes, for example,
! it can be useful to keep values such as usernames in context values for use when logging information for later.
! However, if the username is used to determine if a function should display some specific information, you’d want to include it as a function parameter even if it’s already available from the context.
! This way when you, or someone else, looks at the function in the future, it’s easier to see which data is actually being used.
*/

// The variable’s name is ctx, which is commonly used for context values. It’s also recommended to put the context.Context parameter as the first parameter in a function, and you’ll see it there in the Go standard library. 
func doSomething(ctx context.Context) {
	fmt.Println("Doing something!")

	anotherCtx := context.WithValue(ctx, "keyA", "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("doSomething: keyA value is %s\n", ctx.Value("keyA"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: keyA's value is %s\n", ctx.Value("keyA"))
}

func main() {
	// ctx := context.TODO() // Empty context
	// The context.Background function creates an empty context like context.TODO does, but it’s designed to be used where you intend to start a known context. Fundamentally the two functions do the same thing: they return an empty context that can be used as a context.Context. The biggest difference is how you signal your intent to other developers. If you’re unsure which one to use, context.Background is a good default option.
	ctx := context.Background()

	// In a larger program running on a server, this value could be something like the time the program started running, or the server the program is running on.
	// When using contexts, it’s important to know that the values stored in a specific context.Context are immutable, meaning they can’t be changed. When you called the context.WithValue, you passed in the parent context and you also received a context back. You received a context back because the context.WithValue function didn’t modify the context you provided. Instead, it wrapped your parent context inside another one with the new value.
	ctx = context.WithValue(ctx, "keyA", "valueA")
	
	doSomething(ctx)
}