# Functions

1. Basic syntax
    - func foo() {
        ...
    }

2. Parameters
    - Commad delimited list of variables and types
        - func foo(bar string, baz int)
    - Parameters of same type list type once
        - func foo(bar, baz int)
    - When pointers are passed in, the function change the value in the caller
        - This is always true for data of slices and maps
    - User variadic parameters to send list of same types in
        - Must be last parameter
        - Received as a slice
        - func foo(bar string, baz ...int)
3. Return values
    - Single return values just list type
        - func foo() int
    - Multiple return value list types surrounded by parentheses
        - func foo() (int, error)
        - The (result type, error) paradigm is a very common idiom
    - Can use named return values
        - Initialize returned variable
        - Return using return keyword on its own
    - Can return addresses of local variables
        - Automatically promoted from local memory (stack) to shared memory (heap)
4. Anonymous functions
    - Functions don't have names if they are:
        - Immediately invoked
            - func() {
                ...
            }()
        - Assigned to a variable or passed as an argument to a function
            - a := func() {
                ...
            }
            a()
5. Functions as types
    - Can assign functions to variables or use as arguments and return values in functions
    - Type signature is like function signature, with no parameter names
        - var f func(string, string, int) (int, error)
6. Methods
    - Function that executes in contect of a type
    - Format 
        - func (g greeter) greet() {
            ...
        }
    - Receiver can be a value or pointer
        - Value reciever gets copy of type
        - Pointer receiver gets pointer to type