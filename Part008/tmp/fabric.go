package main

func NewOperator(id string, name string, cca bool, ctm bool, cba bool) Operator {
	a := Operator{User: User{id, name}, CanCreateAccounts: cca, CanTransferMoney: ctm, CanBlockAccounts: cba}
	return a
}

func NewAccountService() IAccountService {
	return &Account{}
}

func NewPaymentService() IPaymentService {
	return &Payment{}
}

func NewLogger() ILogger {
	return &Logger{}
}
