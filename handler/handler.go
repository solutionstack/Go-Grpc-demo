package handler

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	logger "github.com/rs/zerolog"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pb "proto_test/proto"
	"strings"
	"time"
)

var Log = logger.New(os.Stdout)

type Handler struct {
	userSvc pb.UserClient //handler holds interface containing client methods
}

func NewHandler(svc pb.UserClient) Handler {
	return Handler{
		userSvc: svc,
	}
}
func (h *Handler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var userReq pb.PostParam
	if err := UnmarshalRequestBody(r, &userReq); err != nil || len(userReq.Name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := ErrorResponse{Message: errors.Wrap(err, "error decoding post body").Error()}
		jsonErr, _ := json.Marshal(errMsg)
		w.Write(jsonErr)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	newUser, err := h.userSvc.CreateUSer(ctx, &userReq) //call grpc method
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMsg := ErrorResponse{Message: errors.Wrap(err, "error creating user").Error()}
		jsonErr, _ := json.Marshal(errMsg)
		w.Write(jsonErr)
		return

	}

	jsonUser, _ := json.Marshal(newUser)
	w.Write(jsonUser)

}
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userList, err := h.userSvc.GetUsers(ctx, &empty.Empty{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMsg := ErrorResponse{Message: errors.Wrap(err, "error fetching users").Error()}
		jsonErr, _ := json.Marshal(errMsg)
		w.Write(jsonErr)
		return

	}

	jsonResponse, _ := json.Marshal(userList)
	w.Write(jsonResponse)

}
func (h *Handler) GetSingle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//query id
	userID := r.URL.Query().Get("user_id")
	if len(userID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		errMsg := ErrorResponse{Message: errors.New("expecting valid user_id in query").Error()}
		jsonErr, _ := json.Marshal(errMsg)
		w.Write(jsonErr)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := h.userSvc.GetUser(ctx, &pb.GetParam{Id: userID})
	var errMsg ErrorResponse
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			errMsg = ErrorResponse{Message: errors.New("user not found").Error()}
			w.WriteHeader(http.StatusNotFound)

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			errMsg = ErrorResponse{Message: errors.Wrap(err, "error fetching user").Error()}
		}

		jsonErr, _ := json.Marshal(errMsg)
		w.Write(jsonErr)
		return
	}

	jsonResponse, _ := json.Marshal(user)
	w.Write(jsonResponse)

}

func (h *Handler) ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "{\"message\": \"bad route\"}")

}

func UnmarshalRequestBody(r *http.Request, output interface{}) error {
	if r.Body == nil {
		return errors.New("invalid body in request")
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}

type ErrorResponse struct {
	Message string `json:"message"`
}
