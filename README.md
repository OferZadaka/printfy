# printfy
a package to bring python like print formatting to golang


## Examples:
```
func main() {
	age := 12
	name := "ofer"
	height := 184.5
  printfy.Printfy("age = {0}, name = {1}, height = {2}", age, name, height)
  ```
**Output:**
```
age = 12, name = ofer, height = 184.5
```

--------------------------------------------------------------
```
func main() {
  p1 := Person{
      Name: "Warren",
      Age:  31,
      Addr: Address{
        City:    "Denver",
        State:   "CO",
        Country: "U.S.A.",
      },
    }
  printfy.Printfy("\nPerson: {0}", p1)
  ```
**Output:**
```
Person: {Warren 31 {Denver CO U.S.A.}}
```
--------------------------------------------------------------
```
func main() {
	name := flag.String("name", "", "name of person")
	age := flag.Int("age", -1, "age of person")
	address := flag.String("address", "", "address of person")
	flag.Parse()

	if *age == -1 {
		log.Print(printfy.Stringfy("{0} not a valid age!", *age))
	}

	p1 := Person{
		Name: *name,
		Age:  *age,
		Addr: *address,
	}
	if p1.Age <= 12 {
		log.Print(printfy.Stringfy("age must be 13 or above! {0} age is: {1}", p1.Name, p1.Age))
	}

}
```
**Output**:
```
2022/03/23 15:00:13 -1 not a valid age!
2022/03/23 15:00:13 age must be 13 or above!  age is: -1
```
