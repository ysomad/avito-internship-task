package v1

import (
	"context"

	"github.com/google/uuid"

	"github.com/ysomad/avito-internship-task/internal"
	"github.com/ysomad/avito-internship-task/internal/domain"
	"github.com/ysomad/avito-internship-task/internal/service/dto"
)

var _ ServerInterface = &handler{}

type atomicWrapper interface {
	Wrap(context.Context, func(context.Context) error) error
}

type accountService interface {
	DepositFunds(ctx context.Context, userID uuid.UUID, amount string) (domain.Account, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (domain.Account, error)
	ReserveFunds(context.Context, dto.ReserveFundsArgs) (dto.AccountWithReservation, error)
	DeclareRevenue(context.Context, dto.DeclareRevenueArgs) (*domain.Reservation, error)
	GetTransactionList(ctx context.Context, args dto.GetTransactionListArgs) (domain.TransactionList, error)
	TransferFunds(ctx context.Context, args dto.TransferFundsArgs) (domain.Account, error)
}

type handler struct {
	log     internal.Logger
	atomic  atomicWrapper
	account accountService
}

func NewHandler(l internal.Logger, ar atomicWrapper, as accountService) *handler {
	return &handler{
		log:     l,
		atomic:  ar,
		account: as,
	}
}
