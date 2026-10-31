package handler

import (
	"amarfin/src/constant"
	"amarfin/src/customerror"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func Error(err error, r *http.Request) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), constant.ERROR_CONTEXT_KEY, err))
}

func GetAuth(r *http.Request) (string, error) {
	val := r.Context().Value(constant.AUTH_KEY)
	uid, ok := val.(string)
	if !ok {
		return "", customerror.NewError(
			"invalid credentials",
			errors.New("given user id is not a string"),
			customerror.Unauthenticate,
		)
	}
	_, err := uuid.Parse(uid)
	if err != nil {
		return "", customerror.NewError(
			"invalid credentials",
			err,
			customerror.Unauthenticate,
		)
	}
	return uid, nil
}

func WriteJSON(res any, statusCode int, w http.ResponseWriter) error {
	b, err := json.Marshal(res)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add(constant.STATUS_CODE, strconv.Itoa(statusCode))
	w.WriteHeader(statusCode)
	_, err = w.Write(b)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}

	return nil
}

func ReadJSON(req any, r *http.Request) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	if err := json.Unmarshal(b, req); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}
	return err
}
