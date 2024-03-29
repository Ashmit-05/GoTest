[See this for more info](https://quii.gitbook.io/learn-go-with-tests/)

[Another really cool resource](https://gobyexample.com)


Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier.


## Some new concepts related to function return type
- In our function signature we have made a named return value (prefix string).
- This will create a variable called prefix in your function.
- It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
- You can return whatever it's set to by just calling return rather than return prefix.
- This will display in the Go Doc for your function so it can make the intent of your code clearer.

*It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.*

## A note on Methods & Interfaces
1. Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".
2. Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.

- Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
- Circle has a method called Area that returns a float64 so it satisfies the Shape interface
r string does not have such a method, so it doesn't satisfy the interface
etc.
In Go interface resolution is implicit. If the type you pass in matches what the interface is ksking for, it will compile.

### More on Interfaces
We have enjoyed the type-safety that Go has offered us in terms of functions that work with known types, such as string, int and our own types like BankAccount.
This means that we get some documentation for free and the compiler will complain if you try and pass the wrong type to a function.
You may come across scenarios though where you want to write a function where you don't know the type at compile time.
Go lets us get around this with the type interface{} which you can think of as just any type (in fact, in Go any is an alias for interface{}).

## Maps
An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
This may make them feel like a "reference type", but as Dave Cheney describes they are not.
> A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.
A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. You can read more about maps here.
Therefore, you should never initialize an empty map variable:
> var m map[string]string

Instead, you can initialize an empty map like we were doing above, or use the make keyword to create a map for you:
```
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```
Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime panic.

## Concurrency
[This is a really cool resource](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency)
Now we can finally talk about concurrency which, for the purposes of the following, means "having more than one thing in progress." This is something that we do naturally everyday.
For instance, this morning I made a cup of tea. I put the kettle on and then, while I was waiting for it to boil, I got the milk out of the fridge, got the tea out of the cupboard, found my favourite mug, put the teabag into the cup and then, when the kettle had boiled, I put the water in the cup.
What I didn't do was put the kettle on and then stand there blankly staring at the kettle until it boiled, then do everything else once the kettle had boiled.
If you can understand why it's faster to make tea the first way, then you can understand how we will make CheckWebsites faster. Instead of waiting for a website to respond before sending a request to the next website, we will tell our computer to make the next request while it is waiting.
Normally in Go when we call a function `doSomething()` we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is blocking - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a goroutine. Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts, it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page.
To tell Go to start a new goroutine we turn a function call into a go statement by putting the keyword `go` in front of it: `go doSomething()`.
