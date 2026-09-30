package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Asset struct {
	ID        string `json:"id"`
	Owner     string `json:"owner"`
	Value     int    `json:"value"`
	CreatedAt string `json:"createdAt"`
}

func (s *SmartContract) CreateAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
	owner string,
	value int,
) error {

	txTimestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("failed to get transaction timestamp: %v", err)
	}

	createdAt := time.Unix(
		txTimestamp.Seconds,
		int64(txTimestamp.Nanos),
	).UTC().Format(time.RFC3339Nano)

	asset := Asset{
		ID:        id,
		Owner:     owner,
		Value:     value,
		CreatedAt: createdAt,
	}

	bytes, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, bytes)
}

func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface,
	id string, newOwner string) error {

	bytes, err := ctx.GetStub().GetState(id)
	if err != nil || bytes == nil {
		return fmt.Errorf("asset not found")
	}

	var asset Asset
	json.Unmarshal(bytes, &asset)
	asset.Owner = newOwner

	updated, _ := json.Marshal(asset)
	return ctx.GetStub().PutState(id, updated)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface,
	id string) (*Asset, error) {

	bytes, err := ctx.GetStub().GetState(id)
	if err != nil || bytes == nil {
		return nil, fmt.Errorf("asset not found")
	}

	var asset Asset
	json.Unmarshal(bytes, &asset)
	return &asset, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		log.Panic(err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panic(err)
	}
}
