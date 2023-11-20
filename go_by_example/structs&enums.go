package example

// import "fmt"

// type DayOfWeek int

// const (
// 	Sunday DayOfWeek = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )

// type Person struct {
// 	Name    string
// 	Country string
// 	Age     int
// }

// func StructsEnums() {
// 	p := Person{Name: "Alice", Age: 25, Country: "USA"}
// 	fmt.Println("Country:", p.Country)
// 	fmt.Println("Age:", p.Age)
// 	fmt.Println("Name:", p.Name)
// 	today := Sunday
// 	fmt.Println("Today is", today)
// }

// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string) *person {

// 	p := person{name: name}
// 	p.age = 42
// 	return &p
// }

// func Structy() {

// 	fmt.Println(person{"Bob", 20})

// 	fmt.Println(person{name: "Alice", age: 30})

// 	fmt.Println(person{name: "Fred"})

// 	fmt.Println(&person{name: "Ann", age: 40})

// 	fmt.Println(newPerson("Jon"))

// 	s := person{name: "Sean", age: 50}
// 	fmt.Println(s.name)

// 	sp := &s
// 	fmt.Println(sp.age)

// 	sp.age = 51
// 	fmt.Println(sp.age)

// 	dog := struct {
// 		name   string
// 		isGood bool
// 	}{
// 		"Rex",
// 		true,
// 	}
// 	fmt.Println(dog)
// }

// type rect struct {
// 	width, height int
// }

// func (r *rect) area() int {
// 	return r.width * r.height
// }

// func (r rect) perim() int {
// 	return 2*r.width + 2*r.height
// }

// func CalcArea() {
// 	r := rect{width: 10, height: 5}

// 	fmt.Println("area: ", r.area())
// 	fmt.Println("perim:", r.perim())

// 	rp := &r
// 	fmt.Println("area: ", rp.area())
// 	fmt.Println("perim:", rp.perim())
// }
