package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeTaskState(t *testing.T) {
	task := NewTask("title", "description")
	assert.Equal(t, "TODO", task.CurrentState())

	task = task.Assign("John")
	assert.Equal(t, "ASSIGNED", task.CurrentState())
	assert.Equal(t, "John", task.Assignee)

	task = task.UnAssign()
	assert.Equal(t, "TODO", task.CurrentState())
	assert.Equal(t, "", task.Assignee)

	task = task.Close()
	assert.Equal(t, "CLOSED", task.CurrentState())
}
