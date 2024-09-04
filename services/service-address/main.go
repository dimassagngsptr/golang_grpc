package main

import (
	"chapter-c30/common/config"
	"chapter-c30/common/model"
	"context"
	"log"
	"net"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)


var localStorage *model.AddressListByUser

func init() {
	localStorage = new(model.AddressListByUser)
	localStorage.ListAddress = make(map[string]*model.AddressList)
}

type AddressesServer struct {
	model.UnimplementedAddressesServer
}

func (AddressesServer) ADD_ADDRESS(_ context.Context, param *model.AddressAndUserId) (*empty.Empty, error) {
	userId := strconv.Itoa(int(param.UserId))
	address := param.Address

	if _, ok := localStorage.ListAddress[userId]; !ok {
		localStorage.ListAddress[userId] = new(model.AddressList)
		localStorage.ListAddress[userId].List = make([]*model.Address, 0)
	}
	localStorage.ListAddress[userId].List = append(localStorage.ListAddress[userId].List, address)

	log.Println("Adding address", address.String(), "for user", userId)
	return new(empty.Empty), nil
}

func (AddressesServer) LIST_ADDRESS(_ context.Context, param *model.AddressUserId) (*model.AddressList, error) {
	userId := strconv.Itoa(int(param.UserId))
	return localStorage.ListAddress[userId], nil
}

func main(){
	srv:= grpc.NewServer()
	var addressSrv AddressesServer
	model.RegisterAddressesServer(srv, addressSrv)

	log.Println("Starting RPC server at", config.ServiceAddressPort)

	l, err := net.Listen("tcp", config.ServiceAddressPort)
	if err != nil {
		log.Fatal("could not listen to %s: %v", config.ServiceAddressPort, err)
	}

	log.Fatal(srv.Serve(l))
}