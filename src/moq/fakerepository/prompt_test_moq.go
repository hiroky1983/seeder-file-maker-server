// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fakerepository

import (
	"github.com/sashabaranov/go-openai"
	"seeder-app/domain/repository"
	"sync"
)

// Ensure, that PromptMock does implement repository.Prompt.
// If this is not the case, regenerate this file with moq.
var _ repository.Prompt = &PromptMock{}

// PromptMock is a mock implementation of repository.Prompt.
//
//	func TestSomethingThatUsesPrompt(t *testing.T) {
//
//		// make and configure a mocked repository.Prompt
//		mockedPrompt := &PromptMock{
//			PromptFunc: func(client *openai.Client, s string) (string, error) {
//				panic("mock out the Prompt method")
//			},
//		}
//
//		// use mockedPrompt in code that requires repository.Prompt
//		// and then make assertions.
//
//	}
type PromptMock struct {
	// PromptFunc mocks the Prompt method.
	PromptFunc func(client *openai.Client, s string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Prompt holds details about calls to the Prompt method.
		Prompt []struct {
			// Client is the client argument value.
			Client *openai.Client
			// S is the s argument value.
			S string
		}
	}
	lockPrompt sync.RWMutex
}

// Prompt calls PromptFunc.
func (mock *PromptMock) Prompt(client *openai.Client, s string) (string, error) {
	if mock.PromptFunc == nil {
		panic("PromptMock.PromptFunc: method is nil but Prompt.Prompt was just called")
	}
	callInfo := struct {
		Client *openai.Client
		S      string
	}{
		Client: client,
		S:      s,
	}
	mock.lockPrompt.Lock()
	mock.calls.Prompt = append(mock.calls.Prompt, callInfo)
	mock.lockPrompt.Unlock()
	return mock.PromptFunc(client, s)
}

// PromptCalls gets all the calls that were made to Prompt.
// Check the length with:
//
//	len(mockedPrompt.PromptCalls())
func (mock *PromptMock) PromptCalls() []struct {
	Client *openai.Client
	S      string
} {
	var calls []struct {
		Client *openai.Client
		S      string
	}
	mock.lockPrompt.RLock()
	calls = mock.calls.Prompt
	mock.lockPrompt.RUnlock()
	return calls
}