package repository

import (
	"os"
	"testing"

	"github.com/maxhariel/gateway/adpater/repository/fixture"
	"github.com/maxhariel/gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)

	repository := NewTransactionRepositoyDb(db)
	err := repository.Insert("1", "1", 300, entity.APPROVED, "")

	assert.Nil(t, err)
}
