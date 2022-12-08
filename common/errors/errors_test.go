package errors

import "testing"

func TestNewError(t *testing.T) {
	err := NewError("this is the message")

	if err == nil {
		t.Fatal("err should not be nil")
	}
	if err.Error() != "this is the message" {
		t.Errorf("wrong err.Error(): %s", err.Error())
	}

}
