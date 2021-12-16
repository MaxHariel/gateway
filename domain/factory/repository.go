package factory

import "github.com/maxhariel/gateway/domain/repository"

type RespositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
