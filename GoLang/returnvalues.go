// Named Return Values
/*
Return values may be given names, and if they are, then they are treated the 
same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose 
of the returned value

A return statement without arguments returns the nameed return values. This
is known as a "naked" return. Naked return statements should be used only in 
short functions. They can harm readability in longer functions.
*/

func getCoords() (x, y int){
	//x and y are initialized with zero values

	return //automatically returns x and y
}

//is the same as

func getCoords() (int , int){
	var x int
	var y int

	return(x,y)
}