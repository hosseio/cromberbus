// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cromberbus

import (
	"sync"
)

var (
	lockCommandBusMockDispatch sync.RWMutex
)

// Ensure, that CommandBusMock does implement CommandBus.
// If this is not the case, regenerate this file with moq.
var _ CommandBus = &CommandBusMock{}

// CommandBusMock is a mock implementation of CommandBus.
//
//     func TestSomethingThatUsesCommandBus(t *testing.T) {
//
//         // make and configure a mocked CommandBus
//         mockedCommandBus := &CommandBusMock{
//             DispatchFunc: func(command Command) error {
// 	               panic("mock out the Dispatch method")
//             },
//         }
//
//         // use mockedCommandBus in code that requires CommandBus
//         // and then make assertions.
//
//     }
type CommandBusMock struct {
	// DispatchFunc mocks the Dispatch method.
	DispatchFunc func(command Command) error

	// calls tracks calls to the methods.
	calls struct {
		// Dispatch holds details about calls to the Dispatch method.
		Dispatch []struct {
			// Command is the command argument value.
			Command Command
		}
	}
}

// Dispatch calls DispatchFunc.
func (mock *CommandBusMock) Dispatch(command Command) error {
	if mock.DispatchFunc == nil {
		panic("CommandBusMock.DispatchFunc: method is nil but CommandBus.Dispatch was just called")
	}
	callInfo := struct {
		Command Command
	}{
		Command: command,
	}
	lockCommandBusMockDispatch.Lock()
	mock.calls.Dispatch = append(mock.calls.Dispatch, callInfo)
	lockCommandBusMockDispatch.Unlock()
	return mock.DispatchFunc(command)
}

// DispatchCalls gets all the calls that were made to Dispatch.
// Check the length with:
//     len(mockedCommandBus.DispatchCalls())
func (mock *CommandBusMock) DispatchCalls() []struct {
	Command Command
} {
	var calls []struct {
		Command Command
	}
	lockCommandBusMockDispatch.RLock()
	calls = mock.calls.Dispatch
	lockCommandBusMockDispatch.RUnlock()
	return calls
}
