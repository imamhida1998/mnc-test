package usecase

import (
	"errors"
	"github.com/google/uuid"
	"mnc-test/model"
	"mnc-test/model/request"
	"mnc-test/model/response"
	"mnc-test/service/repository"
)

type TransactionUsecase interface {
	CreateTopUp(input *request.TopUp, users *model.User) (*response.TopUp, error)
	CreatePayment(input *request.Payment, users *model.User) (*response.Payment, error)
	Transfer(input *request.Transfer, users *model.User) (*response.Transfer, error)
	TransactionReport(userId string) (*[]model.Transaction, error)
}
type transactionUsercase struct {
	user        repository.UserRepository
	topup       repository.TopUpRepository
	pay         repository.PaymentRepository
	transfer    repository.TransferRepository
	transaction repository.TransactionRepository
}

func NewTransactionUsecase(user repository.UserRepository, topup repository.TopUpRepository, pay repository.PaymentRepository, transfer repository.TransferRepository, transaction repository.TransactionRepository) TransactionUsecase {
	return &transactionUsercase{user, topup, pay, transfer, transaction}
}

func (tx *transactionUsercase) CreateTopUp(input *request.TopUp, users *model.User) (*response.TopUp, error) {
	TopUpid, _ := uuid.NewV7()
	_, err := tx.topup.InsertTopUp(&model.TopUp{
		TopUpId:     TopUpid.String(),
		UserId:      users.UserId,
		AmountTopUp: input.Amount,
	})
	if err != nil {
		return nil, err
	}
	transactionId, _ := uuid.NewV7()
	currentAmount := input.Amount + users.Balance

	resTransaction, err := tx.transaction.InsertTransaction(&model.Transaction{
		TransactionId:   transactionId.String(),
		TopUpID:         TopUpid.String(),
		Status:          "SUCCESS",
		UserID:          users.UserId,
		TransactionType: "CREDIT",
		Amount:          input.Amount,
		BalanceBefore:   users.Balance,
		BalanceAfter:    currentAmount,
	})
	if err != nil {
		return nil, err
	}
	users.Balance = currentAmount
	_, err = tx.user.UpdateUser(users)
	if err != nil {
		return nil, err
	}

	return &response.TopUp{
		TopUpID:       resTransaction.TopUpID,
		AmountTopUp:   input.Amount,
		BalanceBefore: resTransaction.BalanceBefore,
		BalanceAfter:  resTransaction.BalanceAfter,
		CreatedDate:   resTransaction.CreatedAt.Format("2006-1-2 15:04:05"),
	}, nil

}

func (tx *transactionUsercase) CreatePayment(input *request.Payment, users *model.User) (*response.Payment, error) {
	if users.Balance < input.Amount {
		return nil, errors.New("Balance is not enough")
	}

	paymentId, _ := uuid.NewV7()
	_, err := tx.pay.InsertPayment(&model.Payment{
		PaymentId: paymentId.String(),
		UserId:    users.UserId,
		Amount:    input.Amount,
		Remarks:   input.Remarks,
	})
	if err != nil {
		return nil, err
	}
	transactionId, _ := uuid.NewV7()

	currentAmount := users.Balance - input.Amount
	resTransaction, err := tx.transaction.InsertTransaction(&model.Transaction{
		TransactionId:   transactionId.String(),
		PaymentId:       paymentId.String(),
		Status:          "SUCCESS",
		UserID:          users.UserId,
		TransactionType: "DEBIT",
		Amount:          input.Amount,
		Remarks:         input.Remarks,
		BalanceBefore:   users.Balance,
		BalanceAfter:    currentAmount,
	})
	if err != nil {
		return nil, err
	}
	users.Balance = currentAmount
	_, err = tx.user.UpdateUser(users)
	if err != nil {
		return nil, err
	}

	return &response.Payment{
		TopUpID:       resTransaction.PaymentId,
		Amount:        input.Amount,
		Remarks:       resTransaction.Remarks,
		BalanceBefore: resTransaction.BalanceBefore,
		BalanceAfter:  resTransaction.BalanceAfter,
		CreatedDate:   resTransaction.CreatedAt.Format("2006-1-2 15:04:05"),
	}, nil

}

func (tx *transactionUsercase) Transfer(input *request.Transfer, users *model.User) (*response.Transfer, error) {
	if users.Balance < input.Amount {
		return nil, errors.New("Balance is not enough")
	}

	resTarget, err := tx.user.FindById(input.TargetUser)
	if err != nil {
		return nil, err
	}

	transferId, _ := uuid.NewV7()
	_, err = tx.transfer.InsertTransfer(&model.Transfer{
		TransferId: transferId.String(),
		UserId:     users.UserId,
		Amount:     input.Amount,
		Remarks:    input.Remarks,
	})
	if err != nil {
		return nil, err
	}
	transactionId, _ := uuid.NewV7()

	resTarget.Balance += input.Amount
	_, err = tx.user.UpdateUser(resTarget)
	if err != nil {
		return nil, err
	}

	currentAmount := users.Balance - input.Amount

	resTransaction, err := tx.transaction.InsertTransaction(&model.Transaction{
		TransactionId:   transactionId.String(),
		TransferId:      transferId.String(),
		Status:          "SUCCESS",
		UserID:          users.UserId,
		TransactionType: "DEBIT",
		Amount:          input.Amount,
		Remarks:         input.Remarks,
		BalanceBefore:   users.Balance,
		BalanceAfter:    currentAmount,
	})
	if err != nil {
		return nil, err
	}
	users.Balance = currentAmount
	_, err = tx.user.UpdateUser(users)
	if err != nil {
		return nil, err
	}

	return &response.Transfer{
		TransferId:    resTransaction.TransferId,
		Amount:        input.Amount,
		Remarks:       resTransaction.Remarks,
		BalanceBefore: resTransaction.BalanceBefore,
		BalanceAfter:  resTransaction.BalanceAfter,
		CreatedDate:   resTransaction.CreatedAt.Format("2006-1-2 15:04:05"),
	}, nil

}

func (tx *transactionUsercase) TransactionReport(userId string) (*[]model.Transaction, error) {
	res, err := tx.transaction.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return res, nil

}
