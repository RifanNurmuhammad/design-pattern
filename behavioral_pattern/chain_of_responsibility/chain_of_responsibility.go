package chain

import "fmt"

// Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers.
// Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

type user struct {
	name                   string
	basicProfileComplete   bool
	workExperienceComplete bool
	educationComplete      bool
	registrationComplete   bool
}

type HandlerRegistration interface {
	setNext(h HandlerRegistration)
	execute(data *user)
}

type UserProfileHandler struct {
	next HandlerRegistration
}

func (u *UserProfileHandler) execute(data *user) {
	if data.basicProfileComplete {
		fmt.Println("User registration already done")
		u.next.execute(data)
		return
	}
	fmt.Println("User profile step")
	data.basicProfileComplete = true
	u.next.execute(data)
}

func (u *UserProfileHandler) setNext(h HandlerRegistration) {
	u.next = h
}

type WorkExperienceHandler struct {
	next HandlerRegistration
}

func (w *WorkExperienceHandler) execute(data *user) {
	if data.workExperienceComplete {
		fmt.Println("Work Experience already done")
		w.next.execute(data)
		return
	}

	fmt.Println("Work Experience step")
	data.workExperienceComplete = true
	w.next.execute(data)
}

func (w *WorkExperienceHandler) setNext(h HandlerRegistration) {
	w.next = h
}

type EducationHandler struct {
	next HandlerRegistration
}

func (e *EducationHandler) execute(data *user) {
	if data.educationComplete {
		fmt.Println("Education already done")
		return
	}

	fmt.Println("Education step")
}

func (e *EducationHandler) setNext(h HandlerRegistration) {
	e.next = h
}

func Init() {
	ed := &EducationHandler{}

	we := &WorkExperienceHandler{}
	we.setNext(ed)

	up := &UserProfileHandler{}
	up.setNext(we)

	user := &user{name: "rinum"}

	up.execute(user)
}
