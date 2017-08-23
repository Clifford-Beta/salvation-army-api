package store

import (
	"salvation-army-api/model"
	"testing"
	"time"
)

var msgStore = SqlMessageStore{Database}

func TestSqlMessageStore_CreateMessage(t *testing.T) {
	msg := model.Message{
		Id:         1,
		Title:      "Test",
		Content:    "This is a test message",
		Attachment: "",
		From:       "admin@system.com",
		To:         "user@system.com",
		Status:     1,
		TimeStamp:  time.Now(),
		DateSent:   time.Now(),
	}
	res := <-msgStore.CreateMessage(&msg)
	if res.Err != nil {
		t.Errorf("CreateMessage Failed with", res.Err)
	}

}

func TestSqlMessageStore_RetrieveMessage(t *testing.T) {
	res := <-msgStore.RetrieveMessage(1)
	if res.Err != nil {
		t.Errorf("RetrieveMessage Failed with", res.Err)
	}
}

func TestSqlMessageStore_RetrieveAll(t *testing.T) {
	res := <-msgStore.RetrieveAll()
	if res.Err != nil {
		t.Errorf("RetrieveAllMessages Failed with", res.Err)

	}
}
