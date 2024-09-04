package main

import (
	"chapter-c30/common/config"
	"chapter-c30/common/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.ServiceGaragePort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func serviceAddress() model.AddressesClient {
	port := config.ServiceAddressPort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewAddressesClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "1",
		Name:     "Dimas",
		Password: "123",
		Gender:   model.UserGender_MALE,
	}
	garage1 := model.Garage{
		Id:   "1",
		Name: "Bandung",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.234232434,
			Longitude: 89.0129831281,
		},
	}
	address1 := model.Address{
		Id:      1,
		Address: "Jalan jalannn kemana saja",
	}
	user := serviceUser()
	fmt.Printf("\n %s user test\n", strings.Repeat("=", 10))

	user.Register(context.Background(), &user1)

	res1, err := user.List(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatal("error listing user %d", err.Error())
	}
	resString1, _ := json.Marshal(res1.List)
	log.Println(string(resString1))

	garage := serviceGarage()

	fmt.Printf("\n %s> garage test A\n", strings.Repeat("=", 10))

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})
	resGarag1, err := garage.List(context.Background(), &model.GarageUserID{
		UserId: user1.Id,
	})
	if err != nil {
		log.Fatal("error listing garage %d", err.Error())
	}

	resStringGarag1, _ := json.Marshal(resGarag1.List)
	log.Println(string(resStringGarag1))

	address := serviceAddress()

	fmt.Printf("\n %s> address test A\n", strings.Repeat("=", 10))

	userId, _ := strconv.ParseUint(user1.Id, 10, 32)
	address.ADD_ADDRESS(context.Background(), &model.AddressAndUserId{
		UserId:  uint32(userId),
		Address: &address1,
	})

	resAddress1, err := address.LIST_ADDRESS(context.Background(), &model.AddressUserId{
		UserId: uint32(userId),
	})
	if err != nil {
		log.Fatal("error listing address %d", err.Error())
	}
	resStringAddress1, _ := json.Marshal(resAddress1.List)
	log.Println(string(resStringAddress1))

}
