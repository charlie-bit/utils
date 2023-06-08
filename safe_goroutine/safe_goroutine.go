package safe_goroutine

import "log"

func SafeGo(f func()) {
	go func() {
		defer RecoverPanic()
		f()
	}()
}

func RecoverPanic() {
	if e := recover(); e != nil {
		log.Println("panic at")
	}
}
