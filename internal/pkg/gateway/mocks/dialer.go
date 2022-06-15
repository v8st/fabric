// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"google.golang.org/grpc"
)

type Dialer struct {
	Stub        func(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []grpc.DialOption
	}
	returns struct {
		result1 *grpc.ClientConn
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 *grpc.ClientConn
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Dialer) Spy(arg1 context.Context, arg2 string, arg3 ...grpc.DialOption) (*grpc.ClientConn, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []grpc.DialOption
	}{arg1, arg2, arg3})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("dialer", []interface{}{arg1, arg2, arg3})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return returns.result1, returns.result2
}

func (fake *Dialer) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *Dialer) Calls(stub func(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error)) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *Dialer) ArgsForCall(i int) (context.Context, string, []grpc.DialOption) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2, fake.argsForCall[i].arg3
}

func (fake *Dialer) Returns(result1 *grpc.ClientConn, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 *grpc.ClientConn
		result2 error
	}{result1, result2}
}

func (fake *Dialer) ReturnsOnCall(i int, result1 *grpc.ClientConn, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 *grpc.ClientConn
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 *grpc.ClientConn
		result2 error
	}{result1, result2}
}

func (fake *Dialer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Dialer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
