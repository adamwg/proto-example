package web

import (
	"encoding/base64"
	"html/template"
	"log"
	"net/http"

	"github.com/adamwg/proto-example/model"
)

var funcMap = template.FuncMap{
	"base64": base64.StdEncoding.EncodeToString,
}

type Server struct {
	client model.PostsClient
	tmpl   *template.Template
}

func (s *Server) ListenAndServe() error {
	srv := http.Server{
		Handler: s,
		Addr:    "localhost:3000",
	}

	return srv.ListenAndServe()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &model.ReadRequest{}

	postReader, err := s.client.Read(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ch := make(chan *model.Post)
	go func() {
		for {
			resp, err := postReader.Recv()
			// TODO: Handle errors other than EOF better
			if err != nil {
				close(ch)
				return
			}
			ch <- resp.Post
		}
	}()

	err = s.tmpl.ExecuteTemplate(w, "index.tmpl", ch)
	if err != nil {
		log.Fatal(err)
	}
}

func NewServer(cl model.PostsClient) *Server {
	return &Server{
		client: cl,
		tmpl:   template.Must(template.New("").Funcs(funcMap).ParseFiles("index.tmpl")),
	}
}
