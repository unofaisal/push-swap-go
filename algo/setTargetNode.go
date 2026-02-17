package algo

// find target nodes for given nodelist1 against nodelist2

// loop over the nodelist1
    // for each node in nodelist1 loop over nodeList2
		//  compare if the curr node of nodelist1 value > curr node of nodeList2 value and curr node value of nodeList 2 is less than best_match(this will be initialised with the max value of nodeList 2, helper func to find max value in the nodelist)
			//  if true for the above, set target node for the curr node of nodeLst 1 to the value found
			

	// if the above loop ends then there was no target for the curr value that is > nodevalue  then we find the smallest value(this needs helper func) and set it to target
// this loop ends after populating all nodeList1 nodes with a target value