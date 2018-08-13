package main

type errorBehavior interface {
	Retryable() bool
}


func IsRetryable(err error) bool {
	eb, ok := err.(errorBehavior)
	return ok && eb.Retryable()
}

func main() {

}

