/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
	    	return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}
	
	fmt.Printf("Running Init method")
	err := stub.PutState("hello_world", []byte(args[0]))
	
	if err != nil {
        return nil, err
    }
	return nil, nil
}

// Invoke is our entry point to invoke a chaincode 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("invoke is running ")
     function, args := stub.GetFunctionAndParameters()
	 fmt.Println("args === " + args[0])
	 errst := stub.PutState("hello_world", []byte(args[0]))
	if errst != nil {
        return nil, errst
    }
	
	// Handle different s
	if   function == "init" {						
	//initialize the chaincode state, used as reset
		return t.Init(stub)
	}else if function == "Read" {
		// Assign ownership
		return t.read(stub, args)
	fmt.Println("invoke did not find func: ")					//error

	return nil, errors.New("Received unknown  invocation: ")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
    function, args := stub.GetFunctionAndParameters()
	fmt.Println("args === " + args[0])
	// Handle different 
	if function == "dummy_query" {	
         //read a variable
        return t.read(stub, args)
    }
	
	fmt.Println("query did not find func: ")						

	return nil, errors.New("Received unknown  query: ")
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var name, jsonResp string
    var err error
	//function, args := stub.GetFunctionAndParameters()
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    }

    name = args[0]
    valAsbytes, err := stub.GetState(name)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
        return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil
}
