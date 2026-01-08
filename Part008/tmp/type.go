package main

import (
	"errors"
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
	infoLogs  []string
	fatalLogs []string
}
