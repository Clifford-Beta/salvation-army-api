package store

import (
	"testing"
	"time"
	"salv_prj/model"
)

func TestSqlMessageStore_CreateMessage(t *testing.T) {
	msg := model.Message{
		Id:1,
		Title:"Test",
		Content:"This is a test message",
		Attachment:"",
		From:"admin@system.com",
		To:"user@system.com",
		Status:1,
		TimeStamp:time.Now(),
		DateSent:time.Now(),
	}

}
