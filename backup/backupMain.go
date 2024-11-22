package backup

import (
	"Project/network/messages"
	"Project/network/tcpnetwork"
	"Project/primary"
	"fmt"
	"time"
)

func BackupMain(primary_udp_broadcast chan string) {
	fmt.Println("Backup started")

	timeInitialized := time.Now()
	
	// Connect to primary
	ip_primary := <-primary_udp_broadcast

	primaryTCP_socket := tcpnetwork.NewBackupToPrimaryTCPClient(ip_primary)
	primaryTCP_socket.Run()

	for{
		if primaryTCP_socket.IsActive(){
			break
		}
		if time.Since(timeInitialized) > _connectTimeout{
			fmt.Println("Backup: Could not connect to primary, exiting...")
			return
		}
	}

	// Send Connected
	primaryTCP_socket.Out <- messages.MessageToBytes(messages.M_Connected{})

	// Channels
	backup_hallRequest := make(chan messages.M_BackupHallRequest)
	backup_cabRequest := make(chan messages.M_BackupCabRequest)
	delete_hallRequest := make(chan messages.M_DeleteHallRequest)
	delete_cabRequest := make(chan messages.M_DeleteCabRequest)
	primary_alive := make(chan messages.M_PrimaryAlive)
	send_alive := make(chan bool)
	primary_dead := make(chan bool)

	// Start go routines
	go pollPrimaryAlive(primaryTCP_socket, primary_dead)
	go periodicallySendAlive(primaryTCP_socket, send_alive)

	// Main loop
	for {
		select {
		case a := <-primaryTCP_socket.In:
			// decode message and send to correct channel
			go decodeMessage(a, backup_hallRequest, backup_cabRequest,
				delete_hallRequest, delete_cabRequest, primary_alive)
		case a := <-backup_hallRequest:
			// save request and ack to primary
			go event_backupHallRequest(a, primaryTCP_socket)
		case a := <-backup_cabRequest:
			// save request and ack to primary
			go event_backupCabRequest(a, primaryTCP_socket)
		case a := <-delete_hallRequest:
			// delete request
			event_deleteHallRequest(a)
		case a := <-delete_cabRequest:
			// delete request
			event_deleteCabRequest(a)
		case a := <-primary_alive:
			// update elevators map
			event_primaryAlive(a)				
		case <-send_alive:
			// send alive message to primary
			go event_sendAlive(primaryTCP_socket)
		case <-primary_dead:
			// become primary and return
			fmt.Println("Backup: Primary dead, 'calling go primary.PrimaryMain()'")
			primaryTCP_socket.Stop()
			go primary.PrimaryMain(elevators, confirmedHallRequests, confirmedCabRequests)
			return
		}
	}
}
