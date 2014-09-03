package system

// Application represents an application as a collection of dependent tasks.
// The tasks are assumed to form a directed acyclic graph with a single root.
type Application struct {
	Tasks []Task
}

// Task represents a task of an application. A task can have a number of
// children, which are tasks that depend on the current one (they can only
// proceed when this task is done). Each task is also given a type (Type),
// which is used for looking up the execution time and power consumption of the
// task when it is being executed of a core (see the definition of Core).
type Task struct {
	ID       uint32
	Type     uint32
	Children []*Task
}
