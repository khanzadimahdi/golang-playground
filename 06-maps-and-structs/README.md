# Maps and Structs

1. Maps 
    - Collection of value types that are accessed via keys
    - Created via literals or via **make** function
    - Members accessed via **[key]** syntax
        - myMap["kay"] = "value"
    - Check for presence with "value, ok" form of result
    - Multiple asignments refer to same underlying data
2. Structs
    - Collections of disparate data types that describe a single concept
    - Keyed by named fields
    - Normally created as types, but anonymous structs are allowed
    - Structs are value types
    - No inheritance, but can use composition via embedding
    - Tags can be added to struct fields to describe field
    