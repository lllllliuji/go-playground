package mytype

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func TestType() {
	lisa := user{"lisa", "lisa@email.com", 123, true}
	bob := user{name: "bob", email: "bob@email.com", ext: 234, privileged: false}
	// fmt.Println(lisa)
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	lisa.notify();
	bob.notify()
	fred.person.notify()
	lisa.changeEmail("myemail@email.com")
}
