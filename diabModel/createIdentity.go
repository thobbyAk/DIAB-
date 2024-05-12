package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

type IdentityChaincode struct {
}

type Identity struct {
    FirstName   string `json:"firstName"`
    LastName    string `json:"lastName"`
    DateOfBirth string `json:"dateOfBirth"`
    Address     string `json:"address"`
    Nationality string `json:"nationality"`
    Height      string `json:"height"`
}

func (t *IdentityChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}

func (t *IdentityChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    function, args := stub.GetFunctionAndParameters()

    if function == "createIdentity" {
        return t.createIdentity(stub, args)
    }

    return shim.Error("Invalid function name")
}

//CREATING IDENTITY
func (t *IdentityChaincode) createIdentity(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 6 {
        return shim.Error("Incorrect number of arguments. Expecting 6")
    }

    firstName := args[0]
    lastName := args[1]
    dateOfBirth := args[2]
    address := args[3]
    nationality := args[4]
    height := args[5]

   
    identity := Identity{
        FirstName:   firstName,
        LastName:    lastName,
        DateOfBirth: dateOfBirth,
        Address:     address,
        Nationality: nationality,
        Height:      height,
    }

    // Convert Identity struct to JSON
    identityJSON, err := json.Marshal(identity)
    if err != nil {
        return shim.Error(err.Error())
    }

    // Save identity JSON to the fabric ledger
    err = stub.PutState(firstName+lastName, identityJSON)
    if err != nil {
        return shim.Error(err.Error())
    }

    return shim.Success([]byte("Identity created successfully"))
}

func main() {
    err := shim.Start(new(IdentityChaincode))
    if err != nil {
        fmt.Printf("Error starting IdentityChaincode: %s", err)
    }
}
