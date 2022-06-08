# go-linq

go get github.com/LeoAlexandrov/go-linq


## Usage examples
```
type Person struct {
	Name string
	Age  int
}


persons := api.Collection[Person]{
	Person{Name: "Bob", Age: 22},
	Person{Name: "Alice", Age: 20},
	Person{Name: "John Doe", Age: 44},
	Person{Name: "Peter", Age: 27},
}

p1 := persons.
	Where(func(p Person) bool { return p.Age < 40 }).
	Order(func(p1, p2 Person) int { return strings.Compare(p1.Name, p2.Name) })

fmt.Println(p1) // output: [{Alice 20} {Bob 22} {Peter 27}]

a := api.Collection[string]{"00", "11", "2x", "3a", "4v", "5h", "6w", "7e", "8f", "9d"}

b := a.
	RemoveRange(3, 4).
	OrderDesc(func(s1, s2 string) int { return strings.Compare(s1, s2) })

if b.Any(func(s string) bool { return strings.HasPrefix(s, "8") }) {
	fmt.Println("There is at least one item starting with \"8\"")
}

fmt.Println(b) // output: [9d 8f 7e 2x 11 00]
```
