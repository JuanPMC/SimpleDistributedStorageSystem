package main

import (
	"net"
	"strings"
	"log"
	"time"
	"sync"
	"strconv"
)

// Declare a global
var (
	// map to store nodes
	nodeMap = make(map[string]*Node)
	// mu to control acces to said variable
	mu      sync.Mutex
)

// regisrter a node
func registerNode(node *Node) {
	// lock the much and defer its unlock
	mu.Lock()
	defer mu.Unlock()

	// Check if the node is already registered
	if _, exists := nodeMap[node.ID]; exists {
		log.Printf("Node %s is already registered\n", node.ID)
		return
	}

	// If not registered, add it to the map
	nodeMap[node.ID] = node
	log.Printf("Node %s has been registered\n", node.ID)
}

// for unregistering nodes
func unregisterNode(node *Node) {
	// save the node ide in a variable
	nodeID := node.ID
	// lock access to the nodeMap and defer the unlock
	mu.Lock()
	defer mu.Unlock()

	// Check if the node is registered
	if _, exists := nodeMap[nodeID]; exists {
		// remove from map
		delete(nodeMap, nodeID)
		log.Printf("Node %s has been unregistered\n", nodeID)
	} else {
		log.Printf("Node %s is not registered\n", nodeID)
	}
}

// this function processes the commands form the client
func process_comand(bytesRead int,buffer []byte) ([]byte, string){
	// reeds the command and stores it in orden
	orden := strings.Fields(string(buffer[:bytesRead]))
	// Defines the default output and extra logs
	output := []byte("orden dosent exist\n")
	extra_logs := "Command: "+ string(buffer[:bytesRead])

	// command is register
	if orden[0] == "register"{
		// check the ammount of arguments are correct
		if len(orden) >= 3{
			// register the node callign register node a nd creating a new node with NewNode
			registerNode(NewNode(string(orden[1]),string(orden[2]) ))
			output = []byte("Node: " +  string(orden[1]) +":" + string(orden[2]) + " registered")
		}
	}
	// command is unregister
	if orden[0] == "unregister"{
		// check the ammount of arguments are correct
		if len(orden) >= 2{
			// unregistering the node by calling unregisterNode
			unregisterNode(nodeMap[string(orden[1])])
			output = []byte("Node: "+ string(orden[1]) +" unregistered")
		}
	}
	// command is add
	if orden[0] == "add"{
		// check the ammount of arguments are correct
		if len(orden) >= 4{
			// get the size by converting the argument ot int
			realsize, _ := strconv.Atoi(orden[3])
			// adding the file by calling addFile
			nodeMap[string(orden[1])].addFile(orden[2],realsize)
			output = []byte("Node "+string(orden[1])+" Added file " + string(orden[2]))
		}
	}
	// command is rm
	if orden[0] == "rm"{
		// check the ammount of arguments are correct
		if len(orden) >= 3{
			// delete file by calling deleteFile
			nodeMap[string(orden[1])].deleteFile(orden[2])
			output = []byte("removed file")
		}
	}
	// command is query
	if orden[0] == "query"{
		// declare output list
		var nodeList string
		// check the ammount of arguments are correct
		if len(orden) >= 2{
			// loop all the nodes
			for _, node := range nodeMap{
				// if a node contains the file requested
				if _, exists := node.fileMap[orden[1]]; exists {
					// add the node to the output list
					nodeList += node.ID +"\n"
				}
			}
		}
		// retunr the list
		output = []byte(nodeList)
	}
	// command is state ( this to get the state of all nodes and the files )
	if orden[0] == "state"{
		// declare output lost
		var filelist string
		// loop all nodes
		for _, node := range nodeMap {
			// add the output of printinfo to filelist
			filelist += node.printInfo()+ "\n"
		}
		output = []byte(filelist)
	}

	// return the output and the extralogs
	return output, extra_logs
}

// No loop to make it rest-full
func handleConnection(conn net.Conn) {
	// close the connecton at the end of the function
	defer conn.Close()
	var extra_logs string

	// Log the client connection
	clientAddr := conn.RemoteAddr()
	log.Printf("Client connected: %s\n", clientAddr)

	buffer := make([]byte, 1024)

	// Read data from the connection
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err)
		return
	}

	output, extra_logs := process_comand(bytesRead,buffer)

	// Return the output of the processed request
	startTime := time.Now()
	_, err = conn.Write(output)
	if err != nil {
		log.Println("Error writing:", err)
		return
	}
	// log the ammount of bytes sent
	log.Printf("Bytes: %d Client: %s Transfer_Time: %s %s\n", len(output),clientAddr, time.Since(startTime),extra_logs)

}