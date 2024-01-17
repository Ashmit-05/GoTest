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
