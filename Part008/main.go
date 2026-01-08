package main

import (
	"errors"
	"fmt"
	"time"
)

const (
	STATUS_NEW     PaymentStatus = 1
	STATUS_PROCESS PaymentStatus = 2
	STATUS_SUCCESS PaymentStatus = 3
	STATUS_FAILED  PaymentStatus = 4

	PAYMENT_TYPE_IN  PaymentType = 1
	PAYMENT_TYPE_OUT PaymentType = 2
)

var (
	ErrPermissionDenied        = errors.New("permission denied")         // ошибка если недостаточно прав
	ErrInsufficientFunds       = errors.New("insufficient amount")       // ошибка если недостаточно средств
	ErrAccountBlocked          = errors.New("account is blocked")        // ошибка если аккаунт блокирован
	ErrInvalidStatusTransition = errors.New("invalid status transition") // ошибка статуса платежа
)

// Money - специальный тип деньги - в центах
type Money int

// PaymentStatus статус платежа
type PaymentStatus int

// PaymentType - тип платежа
type PaymentType int

// Account - entity аккаунта с балансом
type Account struct {
	Id          string    // id uuid
	Title       string    // название аккаунта
	Balance     Money     // баланс
	Blocked     bool      // статус блокировки
	BlockReason string    // причина блокировки
	CreatedAt   time.Time // дата создания
	UpdatedAt   time.Time // дата обновления
}

type Payment struct {
	Id        string        // id uuid (можно просто любой)
	Amount    Money         // сумма платежа
	Account   Account       // аккаунт на который или с которого зачиляются средства
	Type      PaymentType   // тип операции PAYMENT_TYPE_IN PAYMENT_TYPE_OUT
	Status    PaymentStatus // статус платежа STATUS_NEW | STATUS_PROCESS | STATUS_SUCCESS | STATUS_FAILED
	CreatedAt time.Time     // дата создания
	UpdatedAt time.Time     // дата обновления платежа
}

type User struct {
	Id   string
	Name string
}

type Operator struct {
	User
	CanCreateAccounts bool
	CanTransferMoney  bool
	CanBlockAccounts  bool
}

type Logger struct {
	logtype   int
	infoLogs  []string
	fatalLogs []string
}

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
	GetInfo()
}

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
	operator.CanBlockAccounts = true
	panic("implement me")
}

func (a *Account) SetLogger(logger ILogger) {
	logger.Info("start logging account")
}

func (p *Payment) CreateNewPayment(accountOut Account, amount Money, paymentType PaymentType, comment string) (*Payment, error) {
	panic("implement me")
}

func (p *Payment) UpdatePaymentStatus(payment *Payment, status PaymentStatus) error {
	panic("implement me")
}

func (p *Payment) SetLogger(logger ILogger) {
	logger.Info("start logging payment")
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

func (l *Logger) GetInfo() {
	fmt.Println(l.infoLogs, l.fatalLogs)
}

func BusinessLogic(operator Operator, accountService IAccountService, paymentService IPaymentService, log ILogger) {
	// TODO сделай так, чтобы это заработало с твоими реализацими
	// хранение аккаунтов логов и платежей сделать in memory

	// даем сервису платежей свои зависимости
	paymentService.SetLogger(log)
	paymentService.SetOperator(operator)

	// даем сервису аккаунтов свои зависимости
	accountService.SetPaymentService(paymentService)
	accountService.SetOperator(operator)
	accountService.SetLogger(log)

	// созлаем 2 аккаунта ---------

	// создаем первый аккаунт, проверяем права на кадлое действие у оператора
	// внутри метода CreateNewAccount - проводим операцию пополнения через сервис платежей и смену статуса через автомат
	// статус платежа может устанавливаться только в определенной последовательности new -> process -> (success | failed)
	// все должно логироваться и сопровождаться соответсвующими ошибками
	accountA, err := accountService.CreateNewAccount(Money(100000))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Info("Account A created")

	// создаем второй аккаунт соответственно, сопровождая логами и ошибками
	accountB, err := accountService.CreateNewAccount(Money(0))
	if err != nil {
		e := accountService.BlockAccount(accountA, "второй аккаунт не смог создаться, по этому первый заблокировали")
		log.Fatal(err.Error())
		if e != nil {
			log.Fatal(e.Error())
		}
		return
	}
	log.Info("account B created")

	// делаем перевод между аккаунтами
	log.Info("creating transfer from account A to account B")
	err = accountService.TransferMoney(accountA, accountB, Money(1000))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// ! примечание ! все сервисы должны поддерживать ошибки
	//если у оператора нет CanCreateAccounts — CreateNewAccount возвращает ошибку;
	//если нет CanTransferMoney — ошибка в TransferMoney;
	//если нет CanBlockAccounts — ошибка в BlockAccount;
	//если баланс меньше суммы перевода — ошибка ErrInsufficientFunds;
	//если аккаунт заблокирован — ошибки в операциях по нему;
	//если нарушен автомат статусов (например, NEW -> SUCCESS напрямую) — ErrInvalidStatusTransition.

	// меняйте код с разными вводными чтобы проверить случаи ошибок

}

func main() {
	// TODO сделай все реализации и запусти с ними бизнес логику
	operator := NewOperator("1", "User1", true, true, true)
	accountService := NewAccountService()
	paymentService := NewPaymentService()
	log := NewLogger()

	_ = operator
	_ = accountService
	_ = paymentService
	_ = log

	//BusinessLogic(operator, accountService, paymentService, log)
	accountService.SetLogger(log)
	paymentService.SetLogger(log)

	log.GetInfo()

}
