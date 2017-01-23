package main

import (
	"io/ioutil"
	"log"

	"github.com/adamwg/proto-example/model"
	"github.com/golang/protobuf/proto"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/sha3"
)

func main() {
	me := &model.User{
		Id:       uuid.New(),
		Name:     "Adam Wolfe Gordon",
		Password: sha3.New512().Sum([]byte("password")),
	}

	photo, err := createPhoto("rabbit.txt")
	if err != nil {
		panic("file read error")
	}
	post := &model.Post{
		Id:     uint64(1),
		UserId: me.Id,
		Content: &model.PostContent{
			Text: "Look at my rabbit!",
			Attachments: []*model.Attachment{
				&model.Attachment{
					Attachment: &model.Attachment_File{photo},
				},
			},
		},
	}

	if err := writeProto(me, "saved-user"); err != nil {
		log.Panicf("failed to write user: %s", err)
	}
	log.Println("Wrote user to saved-user")
	if err := writeProto(post, "saved-post"); err != nil {
		log.Panicf("failed to write post: %s", err)
	}
	log.Println("Wrote post to saved-post")
}

func createPhoto(fname string) (*model.File, error) {
	ret := &model.File{
		MimeType: "text/plain",
	}

	var err error
	ret.Content, err = ioutil.ReadFile(fname)

	return ret, err
}

func writeProto(m proto.Message, fname string) error {
	bs, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fname, bs, 0644)
	if err != nil {
		return err
	}

	return nil
}
