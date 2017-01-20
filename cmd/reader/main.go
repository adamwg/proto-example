package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/adamwg/proto-example/model"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	user := &model.User{}
	post := &model.Post{}

	if err := readProto("saved-user", user); err != nil {
		log.Panicf("failed to read user: %s", err)
	}
	log.Println("read user from saved-user")
	if err := readProto("saved-post", post); err != nil {
		log.Panicf("failed to read post: %s", err)
	}
	log.Println("read post from saved-post")

	marshaler := &jsonpb.Marshaler{Indent: "    "}

	userString, err := marshaler.MarshalToString(user)
	if err != nil {
		log.Panicf("failed to format user: %s", err)
	}
	fmt.Printf("user:\n%s\n", userString)

	postString, err := marshaler.MarshalToString(post)
	if err != nil {
		log.Panicf("failed to format post: %s", err)
	}
	fmt.Printf("post:\n%s\n", postString)
}

func readProto(fname string, m proto.Message) error {
	bs, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(bs, m)
	if err != nil {
		return err
	}

	return nil
}
