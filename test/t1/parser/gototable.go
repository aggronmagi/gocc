
/*
*/
package parser

const numNTSymbols = 3
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // A
		2, // B
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	
}
