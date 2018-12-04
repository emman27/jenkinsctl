package queue

// Item is an Item in the queue
type Item struct {
	Blocked    bool        `json:"blocked"`
	Buildable  bool        `json:"buildable"`
	ID         int         `json:"id"`
	URL        string      `json:"url"`
	Executable *Executable `json:"executable"`
	Task       Task        `json:"task"`
}

// Executable is an execution of the queue item
type Executable struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
}

// Task is the scheduled task for the queue item
type Task struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Executing returns a boolean indicating whether a queue Item is currently executing
func (i *Item) Executing() bool {
	return i.Executable != nil
}
