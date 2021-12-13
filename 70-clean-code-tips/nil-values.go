package cleanCodeTips

// tips:
// 1. nil is essentially an uninitialised pointer.
// 2. Things break when we try to access methods or properties of a nil value
// 3. it's recommended to avoid returning a nil value when possible.
//	  This way, the users of your code are less likely to accidentally access nil values.
