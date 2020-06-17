# Summary

1. we can declare variables in 3 different ways
   - var foo int
   - var foo int = 42
   - foo := 42
2. we **can't redeclare** variable, but **can shadow** them
3. All variables must be used
4. Visibility
   - **Lower case first letter** for **package scope**
   - **Upper case first letter to export**
   - No private scope (we have block scopes)
5. Naming conventions
   - Pascal or CamelCase
     - Capitalize acronyms (HTTP, URL)
   - As short as reasonable
     - longer names for longer lives

6. Type conversions
   - type conversions must be explicit (example: destinationType(variable))
   - use **strconv** package to working with strings

