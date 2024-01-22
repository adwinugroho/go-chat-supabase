package test

import (
	"database/sql"
	"go-chat-supabase/entity"
	"go-chat-supabase/repository/postgres"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := postgres.NewMessageRepository(db)
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	message := entity.Message{
		Content:     []string{"content1", "content2"},
		Description: "Test description",
		CreatedAt:   time.Now().Format("2006-01-02"),
		UserID:      []string{"user1", "user2"},
	}

	mock.ExpectExec(`INSERT INTO table_message \(message_id, content, user_id, description, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resultMessageID, err := repo.Insert(message)

	assert.NoError(t, mock.ExpectationsWereMet())

	assert.NoError(t, err)
	assert.NotEmpty(t, resultMessageID)
}

func TestInsertMessage_ErrorInExec(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := postgres.NewMessageRepository(db)
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	message := entity.Message{
		Content:     []string{"content1", "content2"},
		Description: "Test description",
		CreatedAt:   time.Now().Format("2006-01-02"),
		UserID:      []string{"user1", "user2"},
	}

	mock.ExpectExec(`INSERT INTO table_message \(message_id, content, user_id, description, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(sql.ErrConnDone)

	resultMessageID, err := repo.Insert(message)

	assert.NoError(t, mock.ExpectationsWereMet())

	assert.Error(t, err)
	assert.Empty(t, resultMessageID)
}

func TestListAllMessages(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := postgres.MessageImp{DB: db}
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	filters := map[string]interface{}{
		"search_text": "content",
	}

	mock.ExpectQuery(`SELECT \* FROM table_message WHERE "content" LIKE \'%content%\'`).
		WillReturnRows(sqlmock.NewRows([]string{"message_id", "content", "user_id", "description", "created_at"}).
			AddRow("1", pq.Array([]string{"content1", "content2"}), pq.Array([]string{"user1", "user2"}), "Test description", time.Now()).
			AddRow("2", pq.Array([]string{"content3", "content4"}), pq.Array([]string{"user3", "user4"}), "Another description", time.Now()))

	resultMessages, err := repo.ListAll(filters)

	assert.NoError(t, mock.ExpectationsWereMet())

	assert.NoError(t, err)
	assert.Len(t, resultMessages, 2)
}
