package messages

import (
	"Project/elevalgo"
	"Project/elevio"
	"reflect"
	"testing"
)

func TestMessages(t *testing.T) {
	// Testcase 1 - M_DoRequest (ButtonEvent data)
	test1 := M_DoRequest{2, elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want1 := test1

	bytes := MessageToBytes(test1)
	got1 := BytesToMessage(bytes)

	if want1 != got1 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test1, got1, want1)
	}

	// Testcase 2 - M_HallLights (Hall lights data)
	test2 := M_HallLights{[elevio.NumFloors][elevio.NumButtons - 1]bool{}}
	test2.Data[3][elevio.BT_HallUp] = true

	want2 := test2
	bytes = MessageToBytes(want2)
	got2 := BytesToMessage(bytes)

	if want2 != got2 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test2, got2, want2)
	}

	// Testcase 3 - M_ElevatorAlive (Elevator data)
	data3 := elevalgo.Elevator{Floor: 3, Dirn: elevio.MD_Up, Behaviour: elevalgo.EB_Moving, Obstruction: false}
	test3 := M_ElevatorAlive{data3}

	want3 := test3

	bytes = MessageToBytes(test3)
	got3 := BytesToMessage(bytes)

	if want3 != got3 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test3, got3, want3)
	}

	// Testcase 4 - getID (dont' care = -1)
	test4 := M_BackupAlive{}
	want4 := -1
	bytes = MessageToBytes(test4)
	got4 := GetID(bytes)
	if got4 != want4 {
		t.Errorf("getID(%v) = %v; want %v", bytes, got4, want4)
	}

	// Testcase 5 - getID (broadcast = 255)
	test5 := M_HallLights{[elevio.NumFloors][elevio.NumButtons - 1]bool{}}
	want5 := 255
	bytes = MessageToBytes(test5)
	got5 := GetID(bytes)
	if got5 != want5 {
		t.Errorf("getID(%v) = %v; want %v", bytes, got5, want5)
	}

	// Testcase 6 - M_PrimaryAlive
	data6 := map[int]elevalgo.Elevator{1: {Floor: 3, Dirn: elevio.MD_Up, Behaviour: elevalgo.EB_Moving, Obstruction: false}, 2: {Floor: 3, Dirn: elevio.MD_Stop, Behaviour: elevalgo.EB_Moving, Obstruction: false}}
	test6 := M_PrimaryAlive{Data: data6}

	want6 := test6

	bytes = MessageToBytes(test6)
	got6 := BytesToMessage(bytes)
	if !reflect.DeepEqual(want6, got6) {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test6, got6, want6)
	}

	// Testcase 7 - M_BackupRequest
	test7 := M_BackupHallRequest{Data: elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want7 := test7

	bytes = MessageToBytes(test7)
	got7 := BytesToMessage(bytes)
	if want7 != got7 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test7, got7, want7)
	}

	// Testcase 8 - M_AckBackupRequest
	test8 := M_AckBackupHallRequest{Data: elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want8 := test8

	bytes = MessageToBytes(test8)
	got8 := BytesToMessage(bytes)
	if want8 != got8 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test8, got8, want8)
	}

	// Testcase 9 - CompletedRequest
	test9 := M_CompletedRequest{elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want9 := test9

	bytes = MessageToBytes(test9)
	got9 := BytesToMessage(bytes)
	if want9 != got9 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test9, got9, want9)
	}

	// Testcase 10 - M_SpawnBackup
	test10 := M_SpawnBackup{Id: 2}
	want10 := test10

	bytes = MessageToBytes(test10)
	got10 := BytesToMessage(bytes)
	if want10 != got10 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test10, got10, want10)
	}

	// Testcase 11 - M_DeleteRequest
	test11 := M_DeleteHallRequest{Data: elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want11 := test11

	bytes = MessageToBytes(test11)
	got11 := BytesToMessage(bytes)
	if want11 != got11 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test11, got11, want11)
	}

	// Testcase 12 - id, if not defined message type
	test12 := "Not a message type"
	want12 := -1
	bytes = MessageToBytes(test12)
	got12 := GetID(bytes)
	if got12 != want12 {
		t.Errorf("getID(%v) = %v; want %v", bytes, got12, want12)
	}

	// Testcase 13 - backupCabRequest
	test13 := M_BackupCabRequest{2, elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want13 := test13

	bytes = MessageToBytes(test13)
	got13 := BytesToMessage(bytes)
	if want13 != got13 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test13, got13, want13)
	}

	// Testcase 14 - M_DeleteCabRequest
	test14 := M_DeleteCabRequest{2, elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want14 := test14

	bytes = MessageToBytes(test14)
	got14 := BytesToMessage(bytes)
	if want14 != got14 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test14, got14, want14)
	}

	// Testcase 15 - AckBackupCabRequest
	test15 := M_AckBackupCabRequest{2, elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want15 := test15

	bytes = MessageToBytes(test15)
	got15 := BytesToMessage(bytes)
	if want15 != got15 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test15, got15, want15)
	}

	// Testcase 16 - M_NewRequest
	test16 := M_NewRequest{elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want16 := test16

	bytes = MessageToBytes(test16)
	got16 := BytesToMessage(bytes)
	if want16 != got16 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test16, got16, want16)
	}

	// Testcase 17 - M_NewRequest (id -1)
	test17 := M_NewRequest{elevio.ButtonEvent{Floor: 3, Button: elevio.BT_Cab}}
	want17 := -1

	bytes = MessageToBytes(test17)
	got17 := GetID(bytes)

	if got17 != want17 {
		t.Errorf("getID(%v) = %v; want %v", bytes, got17, want17)
	}

	// Testcase 18 - M_KILL
	test18 := M_KILL{2}
	want18 := test18

	bytes = MessageToBytes(test18)
	got18 := BytesToMessage(bytes)
	if want18 != got18 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test18, got18, want18)
	}

	// Testcase 19 - M_Connected
	test19 := M_Connected{}
	want19 := test19

	bytes = MessageToBytes(test19)
	got19 := BytesToMessage(bytes)
	if want19 != got19 {
		t.Errorf("Network_messageToBytes(%v) = %v; want %v", test19, got19, want19)
	}

}
