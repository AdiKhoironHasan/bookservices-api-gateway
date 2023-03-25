package contract

// ProtectedMethods is a function to hold grpc service methods
// false value indicates that the method is not protected (no authorization needed)
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/user.UserService/Ping":   true,
		"/user.UserService/List":   true,
		"/user.UserService/Store":  true,
		"/user.UserService/Detail": true,
		"/user.UserService/Update": true,
		"/user.UserService/Delete": true,

		"/book.BookService/Ping":   true,
		"/book.BookService/List":   true,
		"/book.BookService/Store":  true,
		"/book.BookService/Detail": true,
		"/book.BookService/Update": true,
		"/book.BookService/Delete": true,
	}
}
