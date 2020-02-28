package todo

// Todo describes how a todo task should look like
type Todo struct {
	Title       string
	Description string
}

// New creates new instances of Todo struct
func New(title, description string) Todo {
	//if add(title, description) {
	return Todo{title, description}
	// }
	//return Todo{}
}

//////// Another approach//////
// // Title field getter
// func (t Todo) Title() string {
// 	return t.title
// }

// // SetTitle title field setter
// func (t *Todo) SetTitle(title string) {
// 	t.title = title
// }

/*
func add(title, description string) bool {
	if title == "" || description == "" {
		return false
	}
	return true
}
*/
