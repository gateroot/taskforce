package task

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeTaskState(t *testing.T) {
	task := NewTask("title", "description")
	assert.Equal(t, "TODO", task.CurrentState())

	task = task.Start()
	assert.Equal(t, "DOING", task.CurrentState())

	task = task.Pause()
	assert.Equal(t, "PAUSED", task.CurrentState())

	task = task.Complete()
	assert.Equal(t, "COMPLETED", task.CurrentState())

	task = task.Stop()
	assert.Equal(t, "TODO", task.CurrentState())

	task = task.Close()
	assert.Equal(t, "CLOSED", task.CurrentState())
}
