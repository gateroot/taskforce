package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeTaskState(t *testing.T) {
	sf := NewStateFactory()
	tf := NewTaskFactory(sf)

	task := tf.New(1, "title", "description", sf.Todo())
	assert.IsType(t, sf.Todo(), task.GetState())
	assert.Equal(t, "TODO", task.GetState().String())

	task = tf.Start(task)
	assert.IsType(t, sf.Doing(), task.GetState())
	assert.Equal(t, "DOING", task.GetState().String())

	task = tf.Pause(task)
	assert.IsType(t, sf.Paused(), task.GetState())
	assert.Equal(t, "PAUSED", task.GetState().String())

	task = tf.Complete(task)
	assert.IsType(t, sf.Completed(), task.GetState())
	assert.Equal(t, "COMPLETED", task.GetState().String())

	task = tf.Stop(task)
	assert.IsType(t, sf.Todo(), task.GetState())
	assert.Equal(t, "TODO", task.GetState().String())

	task = tf.Close(task)
	assert.IsType(t, sf.Closed(), task.GetState())
	assert.Equal(t, "CLOSED", task.GetState().String())
}
