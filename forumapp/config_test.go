package forumapp

import (
	"testing"
)

func TestForumInit(t *testing.T) {
	if _, err := ForumInit(); err != nil {
		t.Errorf("expected %v got %v", nil, err)
	}
}
