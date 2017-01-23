package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/adamwg/proto-example/model"
	"github.com/codegangsta/cli"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/sha3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	rpcAddrFlag = cli.StringFlag{
		Name:  "rpc.addr",
		Value: "localhost:8080",
		Usage: "host:port where the Posts server can be reached",
	}
	userNameFlag = cli.StringFlag{
		Name:  "user.name",
		Usage: "the name for a user",
	}
	userIDFlag = cli.StringFlag{
		Name:  "user.id",
		Usage: "the id of a user",
	}
	passwordFlag = cli.StringFlag{
		Name:  "password",
		Usage: "a user's plaintext password",
	}
	textFlag = cli.StringFlag{
		Name:  "post.text",
		Usage: "the text of a post",
	}
	attachmentFileFlag = cli.StringFlag{
		Name:  "post.attachment.file",
		Usage: "the filename of a file attachment for a post",
		Value: "",
	}
	attachmentURLFlag = cli.StringFlag{
		Name:  "post.attachment.url",
		Usage: "a URL attachment for a post",
		Value: "",
	}
)

var globalFlags = []cli.Flag{
	rpcAddrFlag,
}

var commands = []cli.Command{
	{
		Name:        "register",
		Usage:       "register a user",
		Description: "registers a user",
		Flags: []cli.Flag{
			userNameFlag,
			passwordFlag,
		},
		Action: registerCommand,
	},
	{
		Name:        "post",
		Usage:       "post a post",
		Description: "posts a post",
		Flags: []cli.Flag{
			userIDFlag,
			passwordFlag,
			textFlag,
			attachmentFileFlag,
			attachmentURLFlag,
		},
		Action: postCommand,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "posts-client"
	app.Author = "Adam Wolfe Gordon"
	app.Email = "awg@xvx.ca"
	app.Usage = "A commandline client for the hot new social network Posts"
	app.Version = "0.0.1"

	app.Flags = globalFlags
	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func registerCommand(ctx *cli.Context) {
	client := getPostsClient(ctx)

	userName := ctx.String(userNameFlag.Name)
	password := ctx.String(passwordFlag.Name)

	if userName == "" || password == "" {
		log.Fatal("must specify username and password to register")
	}

	req := &model.RegisterUserRequest{
		User: &model.User{
			Name:     userName,
			Password: sha3.New512().Sum([]byte(password)),
		},
	}

	resp, err := client.RegisterUser(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	printProto(resp)
}

func postCommand(ctx *cli.Context) {
	client := getPostsClient(ctx)

	userID := ctx.String(userIDFlag.Name)
	password := ctx.String(passwordFlag.Name)

	if userID == "" || password == "" {
		log.Fatal("must specify user ID and password to authenticate")
	}

	authReq := &model.AuthenticateRequest{
		UserId:   userID,
		Password: sha3.New512().Sum([]byte(password)),
	}

	authResp, err := client.Authenticate(context.Background(), authReq)
	if err != nil {
		log.Fatalf("failed to authenticate: %v", err)
	}
	printProto(authResp)

	token := authResp.Token
	text := ctx.String(textFlag.Name)
	if text == "" {
		log.Fatal("must provide text for a post")
	}
	attachmentURL := ctx.String(attachmentURLFlag.Name)
	attachmentFile := ctx.String(attachmentFileFlag.Name)

	var attachments []*model.Attachment
	if attachmentURL != "" {
		a := &model.Attachment{
			Attachment: &model.Attachment_Url{attachmentURL},
		}
		attachments = append(attachments, a)
	}
	if attachmentFile != "" {
		f := &model.File{
			MimeType: "image/jpeg",
		}
		var err error
		f.Content, err = ioutil.ReadFile(attachmentFile)
		if err != nil {
			log.Fatalf("failed to read file attachment: %v", err)
		}

		a := &model.Attachment{
			Attachment: &model.Attachment_File{f},
		}
		attachments = append(attachments, a)
	}

	post := &model.Post{
		UserId: userID,
		Content: &model.PostContent{
			Text:        text,
			Attachments: attachments,
		},
	}

	postReq := &model.PostRequest{
		Token: token,
		Post:  post,
	}

	postResp, err := client.Post(context.Background(), postReq)
	if err != nil {
		log.Fatalf("failed to post: %v", err)
	}
	printProto(postResp)
}

func getPostsClient(ctx *cli.Context) model.PostsClient {
	rpcAddr := ctx.GlobalString(rpcAddrFlag.Name)
	if rpcAddr == "" {
		log.Fatal("must provide an RPC address")
	}

	conn, err := grpc.Dial(rpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return model.NewPostsClient(conn)
}

func printProto(m proto.Message) {
	marshaler := &jsonpb.Marshaler{Indent: "    "}
	str, err := marshaler.MarshalToString(m)
	if err != nil {
		log.Fatalf("failed to marshal proto: %v", err)
	}

	fmt.Printf("%s\n", str)
}
