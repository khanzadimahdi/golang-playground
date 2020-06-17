# Primitives

1. Boolean type
   - Values are true or false
   - Not an alias for other types (e.g. int)
   - Zero value is false
2. Numeric types
   - Integers
     - Signed integers
       - int type has varying size, but min 32 bits
       - 8 bit (int8) through 64, bit (int64)
     - Unsigned integers
       - 8 bit (byte and uint8) through 32 bit(uint32)
     - Arithmetic operations
       - addition(+)
       - subtraction(-)
       - multiplication(*)
       - division(/)
       - remainder(%)
     - Bitwise operations
       - and (&)
       - or (|)
       - xor (^)
       - and not (&^)
     - Zero value is 0 (initial value)
     - Cant mix types in same family! (uint16 + uint32 = error!!!)
   - Floating point numbers
     - Follow IEEE-754 standard
     - Zero value is 0
     - 32 and 64 bit versions
     - Literal styles
       - Decimal(3.14)
       - Exponential(13e18 or 2E10)
       - mixed(13.7e12)
     - Arithmetic operations
       - addition(+)
       - subtraction(-)
       - multiplication
       - division
   - Complex numbers
     - Zero value is (0+0i)
     - 64 and 128 bit versions
     - Built-in functions
       - complex: make complex number from tow floats
       - real: get real part as float
       - imag: get imaginary part as float
     - Arithmetic operations are the same as float numbers
3. Text types
   1. Strings
      1. UTF-8
      2. Immutable
      3. Can be concatenated with plus(+) operator
      4. can be converted to []byte
   2. Rune
      1. UTF-32
      2. Alias for int32
      3. Special methods normally required to process
         1. e.g. strings.Reader#ReadRune