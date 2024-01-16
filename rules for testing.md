[See this for more info](https://quii.gitbook.io/learn-go-with-tests/)


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