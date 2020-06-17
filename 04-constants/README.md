# Constants

1. Immutable, but can be shadowed
2. Replaced by the compiler at compile time
    - Value must be calculable at compile time
3. Named like variables
    - PascalCase for exported constants
    - CamelCase for internalConstants
4. Types constants work like immutable variables
    - Can interoperate only with same type
5. Untyped constants work like literals
    - Can interoperate with similar types
6. Enumerated constants
    - Special symbol **iota** allows related constants to be created easily
    - **Iota** starts at 0 in each const block and increments by one
    - Watch out of constant values that match zero values for variables
7. Enumerated expressions
    - Operations that can be determined at compile time are allowed
        - Arithemetic
        - Bitwise operations
        - Bitshifting
