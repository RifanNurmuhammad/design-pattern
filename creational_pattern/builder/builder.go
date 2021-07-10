package builder

import "fmt"

// Builder is a creational design pattern that lets you construct complex objects step by step.
// The pattern allows you to produce different types and representations of an object using the same construction code.

type Workspace struct {
	numOfMonitor    int
	numOfKeyboard   int
	adjustAbleTable bool
}

func NewWorkspace(numOfMonitor int, numOfKeyboard int, adjustAbleTable bool) *Workspace {
	return &Workspace{
		numOfMonitor:    numOfMonitor,
		numOfKeyboard:   numOfKeyboard,
		adjustAbleTable: adjustAbleTable,
	}
}

type WorkspaceBuilder struct {
	workspace Workspace
}

func (w *WorkspaceBuilder) BuildMonitor(num int) *WorkspaceBuilder {
	w.workspace.numOfMonitor = num
	return w
}

func (w *WorkspaceBuilder) BuildKeyboard(num int) *WorkspaceBuilder {
	w.workspace.numOfKeyboard = num
	return w
}

func (w *WorkspaceBuilder) BuildTable() *WorkspaceBuilder {
	w.workspace.adjustAbleTable = true
	return w
}

func (w *WorkspaceBuilder) CreateWorkspace() (*WorkspaceBuilder, error) {
	if w.workspace.numOfKeyboard > 1 {
		return nil, fmt.Errorf("maximum keyboard only 1")
	}
	return w, nil
}

func Init() {
	w := WorkspaceBuilder{}
	myworkspace, err := w.BuildMonitor(1).BuildKeyboard(1).BuildTable().CreateWorkspace()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myworkspace)
}
