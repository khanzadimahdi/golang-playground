package cleanCodeTips

// 1. Our variables should be named from more to less specific the deeper we go into nested scopes.
//    as a variable's scope becomes smaller, it becomes increasingly clear for the reader what that
//	  variable represents, thereby eliminating the need for specific naming.

// 2. You shouldn’t name your variables after their types for the same reason you wouldn’t name your
//    pets 'dog' or 'cat'. – Dave Cheney

// var beerList []Beer
// for _, brand := range beerBrands {
// 	for _, beer := range brand {
// 		beerList = append(beerList, beer)
// 	}
// }

// 3. we can also improve readability by declaring variables as close to their usage as possible.

// 4. When passing pointers as an input parameter of a function,
//	  we are expanding the scope of the variable whose data is being pointed to.

// 5. Consider the case where our application uses multiple threads.
//	  In this scenario, passing pointers to the same memory location can also potentially
//	  result in a race condition.
