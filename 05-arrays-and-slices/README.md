# Arrays and Slices

1. Arrays
    - Collection of items with same type
    - Fixed size
    - Declaration styles
        - a := [3]int{1, 2, 3}
        - a := [...]int{1, 2, 3}
        - var a [3]int
    - Access via zero-based index
        - a := [3]int{1, 3, 5} // a[1] == 3
    - **len** function returns size of array
    - Copies refer to different underlying data
1. Slices
    - Backed by array
    - Creation styles
        - Slice existing array or slice
        - Literal style
        - Via make function
            1. a:= make([]int, 10)
            2. a:= make([]int, 10, 100)
    - **len** function returns length of slice
    - **cap** function returns length ou underlying array
    - **append** function to add elements to slice
        - May cause expensive copy operation if underlying array is too small
    - Copies refer to same underlying array
