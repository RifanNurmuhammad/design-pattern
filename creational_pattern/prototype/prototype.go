package prototype

// Prototype is a creational design pattern that lets you copy existing objects
// without making your code dependent on their classes.
import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Template struct {
	ID       string
	Title    string
	htmlBody string
}

func NewTemplate(p Prototype) Prototype {
	return p.Clone()
}

func (t Template) Clone() Prototype {
	return &Template{
		ID:       t.ID,
		Title:    t.Title,
		htmlBody: t.htmlBody,
	}
}

func Init() {
	t := Template{
		ID:    "T001",
		Title: "template 1",
	}

	t2 := t.Clone()
	template2, ok := t2.(*Template)
	if ok {
		template2.Title = "Template 2"
	}

	fmt.Println(t)
	fmt.Println(template2)
}
