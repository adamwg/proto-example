package server

import (
	"bytes"
	"sync"

	"github.com/adamwg/proto-example/model"
	"github.com/golang/protobuf/proto"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	mu sync.Mutex

	// users maps user IDs to users.
	users map[string]*model.User
	// posts is the posts in the system. A post's index in the array is its ID.
	posts []*model.Post
	// sessions maps session tokens to user IDs.
	sessions map[string]string
}

var _ model.PostsServer = &server{}

var (
	errUserIdSpecified = grpc.Errorf(codes.InvalidArgument, "user ID specified in registration request")
	errInvalidUser     = grpc.Errorf(codes.Unauthenticated, "invalid user ID or password")
	errUserNotFound    = grpc.Errorf(codes.NotFound, "user not found")
	errInvalidToken    = grpc.Errorf(codes.PermissionDenied, "invlaid authentication token")
	errInvalidPostID   = grpc.Errorf(codes.InvalidArgument, "invlaid post ID")
)

func (s *server) RegisterUser(ctx context.Context, req *model.RegisterUserRequest) (*model.RegisterUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.User.Id != "" {
		return nil, errUserIdSpecified
	}

	req.User.Id = uuid.New()

	s.users[req.User.Id] = &model.User{}
	proto.Merge(s.users[req.User.Id], req.User)

	ret := &model.RegisterUserResponse{
		Id: req.User.Id,
	}
	return ret, nil
}

func (s *server) Authenticate(ctx context.Context, req *model.AuthenticateRequest) (*model.AuthenticateResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[req.UserId]
	if !ok {
		return nil, errInvalidUser
	}

	if !bytes.Equal(req.Password, u.Password) {
		return nil, errInvalidUser
	}

	token := uuid.New()
	s.sessions[token] = u.Id

	ret := &model.AuthenticateResponse{
		Token: token,
	}
	return ret, nil
}

func (s *server) GetUserName(ctx context.Context, req *model.GetUserNameRequest) (*model.GetUserNameResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[req.UserId]
	if !ok {
		return nil, errUserNotFound
	}

	ret := &model.GetUserNameResponse{
		UserName: u.Name,
	}
	return ret, nil
}

func (s *server) Post(ctx context.Context, req *model.PostRequest) (*model.PostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	uid, ok := s.sessions[req.Token]
	if !ok {
		return nil, errInvalidToken
	}

	if req.Post.UserId != uid {
		return nil, errInvalidToken
	}

	post := &model.Post{}
	proto.Merge(post, req.Post)
	post.Id = uint64(len(s.posts))
	s.posts = append(s.posts, post)

	ret := &model.PostResponse{
		Id: post.Id,
	}
	return ret, nil
}

func (s *server) Read(req *model.ReadRequest, srv model.Posts_ReadServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(req.PostIds) != 0 {
		for _, id := range req.PostIds {
			if id >= uint64(len(s.posts)) {
				return errInvalidPostID
			}
			if err := srv.Send(&model.ReadResponse{Post: s.posts[id]}); err != nil {
				return err
			}
		}
		return nil
	}

	for _, post := range s.posts {
		if err := srv.Send(&model.ReadResponse{Post: post}); err != nil {
			return err
		}
	}

	return nil
}

func New() model.PostsServer {
	return &server{
		users:    make(map[string]*model.User),
		sessions: make(map[string]string),
	}
}
