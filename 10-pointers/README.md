
# Pointers

1. Creating pointers
    - Pointer types use an asterisk(*) as a prefix to type pointed to
        - *int-a pointer to an integer
    - Use the address operator (&) to get address of variable
2. Derefrencing pointers
    - Derefrence a pointer by preceding with an asterisk(*)
    - Complex types (e.g. structs) are automatically derefrenced
    - Create pointers to objects
        - Can use the address operator (&) if value type already exists
            - ms := myStruct{foo: 42}
            - p := &ms
        - Use of address operator (&) before initializer
            - &myStruct{foo: 42}
        - Use the **new** keyword
            - Can't initialize fields at the same time
3. Types with internal pointers
    - All assignment operations in Go are copy operations
    - **Slices** and **maps** **contain internal pointers**, so **copies point to same underlying data**