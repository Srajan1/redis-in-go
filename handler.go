package main

import "sync"

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET":  set,
	"GET":  get,
}

var redisMap = map[string]string{}
var redisMutex = sync.RWMutex{}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].bulk}
}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}
	key := args[0].bulk
	value := args[1].bulk

	redisMutex.Lock()
	redisMap[key] = value
	redisMutex.Unlock()

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	redisMutex.RLock()
	val, ok := redisMap[key]
	redisMutex.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: val}
}
