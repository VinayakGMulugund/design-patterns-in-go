package main

type PaymentStrategy interface {
	Pay(amount float64) string
}
