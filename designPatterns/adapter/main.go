package main

// https://golangbyexample.com/adapter-design-pattern-go/

func main() {
	c := client{}
	m := &mac{}
	c.insertSquareUsbInComputer(m)

	w := &windows{}
	wa := &windowsAdapter{
		windowMachine: w,
	}

	c.insertSquareUsbInComputer(wa)
}