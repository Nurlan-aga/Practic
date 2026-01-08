package main

func (a *Account) CreateNewAccount(amount Money) (*Account, error) {
	panic("implement me")
}

func (a *Account) UpdateAccountBalance(account *Account, amount Money) error {
	panic("implement me")
}

func (a *Account) TransferMoney(from *Account, to *Account, amount Money) error {
	panic("implement me")
}

func (a *Account) BlockAccount(account *Account, reason string) error {
	panic("implement me")
}

func (a *Account) SetPaymentService(paymentService IPaymentService) {
	panic("implement me")
}

func (a *Account) SetOperator(operator Operator) {
	panic("implement me")
}

func (a *Account) SetLogger(logger ILogger) {
	panic("implement me")
}

func (p *Payment) CreateNewPayment(accountOut Account, amount Money, paymentType PaymentType, comment string) (*Payment, error) {
	panic("implement me")
}

func (p *Payment) UpdatePaymentStatus(payment *Payment, status PaymentStatus) error {
	panic("implement me")
}

func (p *Payment) SetLogger(logger ILogger) {
	panic("implement me")
}

func (p *Payment) SetOperator(operator Operator) {
	panic("implement me")
}

func (l *Logger) Info(msg string) {
	l.infoLogs = append(l.infoLogs, msg)
}

func (l *Logger) Fatal(msg string) {
	l.fatalLogs = append(l.fatalLogs, msg)
}
