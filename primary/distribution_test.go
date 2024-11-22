package primary

import (
	"Project/elevalgo"
	"Project/elevio"
	"testing"
	//"reflect"
)

// unit tests
func TestTimeToHandle(t *testing.T) {
	// Define a test case

	r1 := [elevio.NumFloors][elevio.NumButtons]bool{}
	testCase1 := struct {
		elev   elevalgo.Elevator
		floor  int
		btntyp elevio.ButtonType
		want   int
	}{
		elev: elevalgo.Elevator{
			Behaviour: elevalgo.EB_Idle,
			Dirn:      elevio.MD_Stop,
			Floor:     0,
			Requests:  r1,
		},
		floor:  1,
		btntyp: elevio.BT_HallUp,
		want:   travel_cost + door_open_cost,
	}

	r2 := [elevio.NumFloors][elevio.NumButtons]bool{}
	r2[2][elevio.BT_HallUp] = true
	r2[2][elevio.BT_Cab] = true // shall add nothing to time
	testCase2 := struct {
		elev   elevalgo.Elevator
		floor  int
		btntyp elevio.ButtonType
		want   int
	}{
		elev: elevalgo.Elevator{
			Behaviour: elevalgo.EB_Idle,
			Dirn:      elevio.MD_Stop,
			Floor:     0,
			Requests:  r2,
		},
		floor:  3,
		btntyp: elevio.BT_HallDown,
		want:   3*travel_cost + 2*door_open_cost,
	}

	

	// Call the function with the test case
	got1 := timeToHandle(testCase1.elev, testCase1.floor, testCase1.btntyp)
	got2 := timeToHandle(testCase2.elev, testCase2.floor, testCase2.btntyp)

	// Check if the result matches the expected value
	if got1 != testCase1.want {
		t.Errorf("TimeToHandle(%v, %v, %v) = %v; want %v", testCase1.elev, testCase1.floor, testCase1.btntyp, got1, testCase1.want)

	}
	if got2 != testCase2.want {
		t.Errorf("TimeToHandle(%v, %v, %v) = %v; want %v", testCase2.elev, testCase2.floor, testCase2.btntyp, got2, testCase2.want)

	}

}

/* // Needs to be tested without network
func TestAvailableElevators(t *testing.T) {
    // Define a test case
    connectedElevators = []int{0, 1, 2}
    elevators = map[int]elevalgo.Elevator{
        0: {Obstruction: false},
        1: {Obstruction: true},
        2: {Obstruction: false},
    }

    // Call the function
    got := availableElevators()

    // Define the expected result
    want := []int{0, 2}

    // Check if the result matches the expected value
    if !reflect.DeepEqual(got, want) {
        t.Errorf("AvailableElevators() = %v; want %v", got, want)
    }
}

func TestElevatorToHandle(t *testing.T) {
    // Define a test case
	emptyRequests := [elevio.NumFloors][elevio.NumButtons]bool{}
	fullRequests := [elevio.NumFloors][elevio.NumButtons]bool{}
	for i := range fullRequests {
    	for j := range fullRequests[i] {
        	fullRequests[i][j] = true
    	}	
	}
    connectedElevators = []int{0, 1, 2}
    elevators = map[int]elevalgo.Elevator{
        0: {Floor: 0, Dirn: elevio.MD_Up, Behaviour: elevalgo.EB_Moving, Requests: emptyRequests},
        1: {Floor: 1, Dirn: elevio.MD_Up, Behaviour: elevalgo.EB_Moving, Requests: fullRequests},
        2: {Floor: 3, Dirn: elevio.MD_Down, Behaviour: elevalgo.EB_Moving, Requests: emptyRequests},
    }
	
	for i := 0; i < len(elevators); i++ {
		elevalgo.ElevatorPrint(elevators[i])
	}

    // Call the function with the test case
    got := elevatorToHandle(1, elevio.BT_HallUp)

    // Define the expected result
    want := 0 // elevator id with minimum time to handle

    // Check if the result matches the expected value
    if got != want {
        t.Errorf("ElevatorToHandle(1, elevio.BT_HallUp) = %v; want %v", got, want)
    }
}

*/