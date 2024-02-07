/*
Interfaces are collections of method signatures.
A type “implements” an interface if it has all of the methods of the given interface defined on it.
*/

/*
In the following example, a "shape" must be able to return its area and perimeter. For instance,
Both rect and circle fulfill the interface.
*/

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.pi * 2 * c.radius
}