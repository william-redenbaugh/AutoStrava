package main

/*
	@brief Data structure containing everything relevant for processing athelete activity
*/
type AtheleteActivity struct {
	Name           string
	Distance       float64
	MovingTime     float64
	ElevationGain  float64
	AverageSpeed   float64
	AverageCadence float64
	AverageTemp    float64
	AverageWattage float64
}

/*
	@brief Linked List setup of Athlete activity so we can easily stack and iterate through list
*/
type ActivityLL struct {
	athleteActivity AtheleteActivity

	// Pointer to next element in the array
	ptr_next *ActivityLL
}

func (ActivityLinkedList ActivityLL) next() ActivityLL {
	return *ActivityLinkedList.ptr_next
}
