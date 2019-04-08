package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleChainCode struct {
}

//-------------------
// Assets definition
//-------------------
// Aircraft
type Aircraft struct {
	Id           string        `json:"id"`
	Registration string        `json:"registration"`
//	Owner        OwnerRelation `json:"owner"`
}

// Owners
//type Owner struct {
//	Id           string `json:"id"`
//	Registration string `json:"registration"`
//}

//type OwnerRelation struct {
//	Id           string `json:"id"`
//}


// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    fmt.Println("We are airborne")
    return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
    fn, args := stub.GetFunctionAndParameters()
    fmt.Println("invoking [" + fn + "]")

    if fn == "addAircraft" {
        return addAircraft(stub, args)
    } else if( fn == "getAircrafts") {
        return getAircrafts( stub)
    }
    // error out
    msg := "Never heard of [" + fn + "] function, please try something else"
    fmt.Println( msg)
    return shim.Error( msg)
}

// Adds a new aircraft to the ledger
// Input : array if string 
// 	0=id
//	1=registration
func addAircraft(stub shim.ChaincodeStubInterface, args []string) (peer.Response) {
	var err error
	expectedArgs := 2
	if len(args) != expectedArgs {
		return shim.Error("Incorrect number of arguments. Expecting " + strconv.Itoa(expectedArgs))
	}
	//input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}
	id := args[0]
	registration := args[1]
	//check if aircraft id already exists
	aircraft, err := getAircraft(stub, id)
	if err == nil {
		msg := "Aircraft [" + id + "] already exists"
		fmt.Println( msg)
		fmt.Println( aircraft)
		return shim.Error( msg)
	}

	// build the aircraft object
	str := `{ "id":"` + id + `","registration":"` + registration + `"}`
	err = stub.PutState(id, []byte(str))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Get returns the value of the specified aircraft by id
func getAircraft(stub shim.ChaincodeStubInterface, id string) (Aircraft, error) {
	var aircraft Aircraft
	aircraftAsBytes, err := stub.GetState(id)
	if err != nil {
		return aircraft, errors.New("Failed to get aircraft " + id)
	}
	json.Unmarshal(aircraftAsBytes, &aircraft)

	if len(aircraft.Id) == 0 {
		return aircraft, errors.New("Aircraft [" + id + "] does not exist")
	}
	
	return aircraft, nil
}

func getAircrafts(stub shim.ChaincodeStubInterface) peer.Response{
	var aircrafts []Aircraft

	// ---- Get All Marbles ---- //
	resultsIterator, err := stub.GetStateByRange("a0", "a9999999999999999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	
	for resultsIterator.HasNext() {
		aKeyValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on aircraft id - ", queryKeyAsStr)
		var aircraft Aircraft
		json.Unmarshal(queryValAsBytes, &aircraft)
		// add this aircraf to the list
		aircrafts = append(aircrafts, aircraft)
	}

	//change to array of bytes
	aircraftsAsBytes, _ := json.Marshal(aircrafts)
	return shim.Success(aircraftsAsBytes)
}

// Input Sanitation - dumb input checking, look for empty strings
func sanitize_arguments(strs []string) error{
	for i, val:= range strs {
		if len(val) <= 0 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be a non-empty string")
		}
		if len(val) > 32 {
			return errors.New("Argument " + strconv.Itoa(i) + " must be <= 32 characters")
		}
	}
	return nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
    if err := shim.Start(new(SimpleChainCode)); err != nil {
            fmt.Printf("Error starting AireLogs chaincode: %s", err)
    }
}
