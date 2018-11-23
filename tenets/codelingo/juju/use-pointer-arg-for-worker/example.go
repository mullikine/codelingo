// Package mypack This package serves as example code for the tenet `use-pointer-arg-for-worker`
package mypack

// good
func NewGoodWorkerGood() (*GoodWorker) {
}

// bad
func NewBadWorkerBad() (worker.Worker) {
}
