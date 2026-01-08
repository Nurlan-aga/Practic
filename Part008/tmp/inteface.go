package main

type IAccountService interface {
	CreateNewAccount(amount Money) (*Account, error)
	UpdateAccountBalance(account *Account, amount Money) error
	TransferMoney(from *Account, to *Account, amount Money) error
	BlockAccount(account *Account, reason string) error
	SetPaymentService(paymentService IPaymentService)
	SetOperator(operator Operator)
	SetLogger(logger ILogger)
}

type IPaymentService interface {
	// CreateNewPayment - создание платежа,
	// возвращает новый платеж с указанием аккаунта и суммой в статусе new
	CreateNewPayment(accountOut Account, amount Money, paymentType PaymentType, comment string) (*Payment, error)
	// UpdatePaymentStatus - автомат статусов, переводит платеж по жизненному циклу
	UpdatePaymentStatus(payment *Payment, status PaymentStatus) error
	SetLogger(logger ILogger)
	SetOperator(operator Operator)
}

type ILogger interface {
	Info(msg string)
	Fatal(msg string)
}
