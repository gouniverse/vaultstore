package vaultstore

import (
	"context"
	"testing"
)

func Test_Store_RecordCount(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatalf("Test_Store_RecordCount: Expected [err] to be nil received [%v]", err.Error())
	}

	record := NewRecord().SetToken("test_token1").SetValue("test_value")

	ctx := context.Background()
	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatalf("Test_Store_RecordCount Failure: [%v]", err.Error())
	}

	record = NewRecord().SetToken("test_token2").SetValue("test_value")

	err = store.RecordCreate(context.Background(), *record)
	if err != nil {
		t.Fatalf("Test_Store_RecordCount Failure: [%v]", err.Error())
	}

	record = NewRecord().SetToken("test_token3").SetValue("test_value")

	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatalf("Test_Store_RecordCount Failure: [%v]", err.Error())
	}

	count, err := store.RecordCount(ctx, RecordQueryOptions{
		Token: "test_token2",
	})

	if err != nil {
		t.Fatal("Test_Store_RecordCount: Expected [err] to be nil received " + err.Error())
	}

	if count != 1 {
		t.Fatalf("Test_Store_RecordCount: Expected [count] to be 1 received [%v]", count)
	}

}

func Test_Store_RecordCreate(t *testing.T) {
	store, err := initStore(":memory:")

	if err != nil {
		t.Fatal("Test_Store_RecordCreate: Expected [err] to be nil received: ", err.Error())
	}

	record := NewRecord().SetToken("test_token").SetValue("test_value")

	ctx := context.Background()
	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatal("Test_Store_RecordCreate: Expected [err] to be nil received " + err.Error())
	}

	exists, err := store.TokenExists(ctx, "test_token")

	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("Test_Store_RecordCreate: token should exist")
	}
}

func Test_Store_RecordFindByID(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByID: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()

	// Test with empty ID
	record, err := store.RecordFindByID(ctx, "")
	if err == nil {
		t.Fatal("Test_Store_RecordFindByID: Expected error for empty ID but got nil")
	}
	if record != nil {
		t.Fatal("Test_Store_RecordFindByID: Expected nil record for empty ID")
	}

	// Create a record
	newRecord := NewRecord().SetToken("test_token_find_by_id").SetValue("test_value_find_by_id")
	recordID := newRecord.ID()

	err = store.RecordCreate(ctx, *newRecord)
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByID: Failed to create record: [%v]", err.Error())
	}

	// Find by ID
	foundRecord, err := store.RecordFindByID(ctx, recordID)
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByID: Expected [err] to be nil received [%v]", err.Error())
	}
	if foundRecord == nil {
		t.Fatal("Test_Store_RecordFindByID: Expected to find record but got nil")
	}
	if foundRecord.ID() != recordID {
		t.Fatalf("Test_Store_RecordFindByID: Expected ID [%s] but got [%s]", recordID, foundRecord.ID())
	}
	if foundRecord.Token() != "test_token_find_by_id" {
		t.Fatalf("Test_Store_RecordFindByID: Expected Token [test_token_find_by_id] but got [%s]", foundRecord.Token())
	}
	if foundRecord.Value() != "test_value_find_by_id" {
		t.Fatalf("Test_Store_RecordFindByID: Expected Value [test_value_find_by_id] but got [%s]", foundRecord.Value())
	}

	// Test with non-existent ID
	nonExistentRecord, err := store.RecordFindByID(ctx, "non_existent_id")
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByID: Expected [err] to be nil for non-existent ID but got [%v]", err.Error())
	}
	if nonExistentRecord != nil {
		t.Fatal("Test_Store_RecordFindByID: Expected nil record for non-existent ID")
	}
}

func Test_Store_RecordFindByToken(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByToken: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()

	// Test with empty token
	record, err := store.RecordFindByToken(ctx, "")
	if err == nil {
		t.Fatal("Test_Store_RecordFindByToken: Expected error for empty token but got nil")
	}
	if record != nil {
		t.Fatal("Test_Store_RecordFindByToken: Expected nil record for empty token")
	}

	// Create a record
	token := "test_token_find"
	newRecord := NewRecord().SetToken(token).SetValue("test_value_find")

	err = store.RecordCreate(ctx, *newRecord)
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByToken: Failed to create record: [%v]", err.Error())
	}

	// Find by token
	foundRecord, err := store.RecordFindByToken(ctx, token)
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByToken: Expected [err] to be nil received [%v]", err.Error())
	}
	if foundRecord == nil {
		t.Fatal("Test_Store_RecordFindByToken: Expected to find record but got nil")
	}
	if foundRecord.Token() != token {
		t.Fatalf("Test_Store_RecordFindByToken: Expected Token [%s] but got [%s]", token, foundRecord.Token())
	}
	if foundRecord.Value() != "test_value_find" {
		t.Fatalf("Test_Store_RecordFindByToken: Expected Value [test_value_find] but got [%s]", foundRecord.Value())
	}

	// Test with non-existent token
	nonExistentRecord, err := store.RecordFindByToken(ctx, "non_existent_token")
	if err != nil {
		t.Fatalf("Test_Store_RecordFindByToken: Expected [err] to be nil for non-existent token but got [%v]", err.Error())
	}
	if nonExistentRecord != nil {
		t.Fatal("Test_Store_RecordFindByToken: Expected nil record for non-existent token")
	}
}

func Test_Store_RecordList(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatalf("Test_Store_RecordList: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()

	// Create multiple records
	for i := 1; i <= 5; i++ {
		token := "test_token_list_" + string(rune('A'+i-1))
		record := NewRecord().SetToken(token).SetValue("test_value_" + string(rune('A'+i-1)))

		err = store.RecordCreate(ctx, *record)
		if err != nil {
			t.Fatalf("Test_Store_RecordList: Failed to create record: [%v]", err.Error())
		}
	}

	// Test listing all records
	records, err := store.RecordList(ctx, RecordQueryOptions{})
	if err != nil {
		t.Fatalf("Test_Store_RecordList: Expected [err] to be nil received [%v]", err.Error())
	}
	if len(records) != 5 {
		t.Fatalf("Test_Store_RecordList: Expected 5 records but got %d", len(records))
	}

	// Test with token filter
	filteredRecords, err := store.RecordList(ctx, RecordQueryOptions{
		Token: "test_token_list_C",
	})
	if err != nil {
		t.Fatalf("Test_Store_RecordList: Expected [err] to be nil received [%v]", err.Error())
	}
	if len(filteredRecords) != 1 {
		t.Fatalf("Test_Store_RecordList: Expected 1 record with token filter but got %d", len(filteredRecords))
	}
	if filteredRecords[0].Token() != "test_token_list_C" {
		t.Fatalf("Test_Store_RecordList: Expected Token [test_token_list_C] but got [%s]", filteredRecords[0].Token())
	}

	// Test with limit
	limitedRecords, err := store.RecordList(ctx, RecordQueryOptions{
		Limit: 2,
	})
	if err != nil {
		t.Fatalf("Test_Store_RecordList: Expected [err] to be nil received [%v]", err.Error())
	}
	if len(limitedRecords) != 2 {
		t.Fatalf("Test_Store_RecordList: Expected 2 records with limit but got %d", len(limitedRecords))
	}

	// Test with token list filter
	tokenListRecords, err := store.RecordList(ctx, RecordQueryOptions{
		TokenIn: []string{"test_token_list_A", "test_token_list_E"},
	})
	if err != nil {
		t.Fatalf("Test_Store_RecordList: Expected [err] to be nil received [%v]", err.Error())
	}
	if len(tokenListRecords) != 2 {
		t.Fatalf("Test_Store_RecordList: Expected 2 records with token list filter but got %d", len(tokenListRecords))
	}
}

func Test_Store_RecordUpdate(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatalf("Test_Store_RecordUpdate: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()

	// Create a record
	record := NewRecord().SetToken("test_token_update").SetValue("test_value_original")
	recordID := record.ID()

	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatalf("Test_Store_RecordUpdate: Failed to create record: [%v]", err.Error())
	}

	// Retrieve the record
	foundRecord, err := store.RecordFindByID(ctx, recordID)
	if err != nil {
		t.Fatalf("Test_Store_RecordUpdate: Expected [err] to be nil received [%v]", err.Error())
	}
	if foundRecord == nil {
		t.Fatal("Test_Store_RecordUpdate: Expected to find record but got nil")
	}

	// Update the record
	foundRecord.SetValue("test_value_updated")
	err = store.RecordUpdate(ctx, *foundRecord)
	if err != nil {
		t.Fatalf("Test_Store_RecordUpdate: Failed to update record: [%v]", err.Error())
	}

	// Verify the update
	updatedRecord, err := store.RecordFindByID(ctx, recordID)
	if err != nil {
		t.Fatalf("Test_Store_RecordUpdate: Expected [err] to be nil received [%v]", err.Error())
	}
	if updatedRecord == nil {
		t.Fatal("Test_Store_RecordUpdate: Expected to find updated record but got nil")
	}
	if updatedRecord.Value() != "test_value_updated" {
		t.Fatalf("Test_Store_RecordUpdate: Expected updated Value [test_value_updated] but got [%s]", updatedRecord.Value())
	}

	// Test update with no changes
	err = store.RecordUpdate(ctx, *updatedRecord)
	if err != nil {
		t.Fatalf("Test_Store_RecordUpdate: Expected [err] to be nil for no changes but got [%v]", err.Error())
	}
}

func Test_Store_RecordDeleteByID(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByID: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()

	// Test with empty ID
	err = store.RecordDeleteByID(ctx, "")
	if err == nil {
		t.Fatal("Test_Store_RecordDeleteByID: Expected error for empty ID but got nil")
	}

	// Create a record
	record := NewRecord().SetToken("test_token_delete_id").SetValue("test_value_delete")
	recordID := record.ID()

	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByID: Failed to create record: [%v]", err.Error())
	}

	// Verify record exists
	foundRecord, err := store.RecordFindByID(ctx, recordID)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByID: Expected [err] to be nil received [%v]", err.Error())
	}
	if foundRecord == nil {
		t.Fatal("Test_Store_RecordDeleteByID: Expected to find record but got nil")
	}

	// Delete the record
	err = store.RecordDeleteByID(ctx, recordID)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByID: Failed to delete record: [%v]", err.Error())
	}

	// Verify record is deleted
	deletedRecord, err := store.RecordFindByID(ctx, recordID)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByID: Expected [err] to be nil after deletion but got [%v]", err.Error())
	}
	if deletedRecord != nil {
		t.Fatal("Test_Store_RecordDeleteByID: Expected nil record after deletion but got a record")
	}

	// Test deleting non-existent record
	err = store.RecordDeleteByID(ctx, "non_existent_id")
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByID: Expected [err] to be nil for non-existent ID but got [%v]", err.Error())
	}
}

func Test_Store_RecordDeleteByToken(t *testing.T) {
	store, err := initStore(":memory:")
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByToken: Expected [err] to be nil received [%v]", err.Error())
	}

	ctx := context.Background()

	// Test with empty token
	err = store.RecordDeleteByToken(ctx, "")
	if err == nil {
		t.Fatal("Test_Store_RecordDeleteByToken: Expected error for empty token but got nil")
	}

	// Create a record
	token := "test_token_delete"
	record := NewRecord().SetToken(token).SetValue("test_value_delete")

	err = store.RecordCreate(ctx, *record)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByToken: Failed to create record: [%v]", err.Error())
	}

	// Verify record exists
	foundRecord, err := store.RecordFindByToken(ctx, token)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByToken: Expected [err] to be nil received [%v]", err.Error())
	}
	if foundRecord == nil {
		t.Fatal("Test_Store_RecordDeleteByToken: Expected to find record but got nil")
	}

	// Delete the record
	err = store.RecordDeleteByToken(ctx, token)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByToken: Failed to delete record: [%v]", err.Error())
	}

	// Verify record is deleted
	deletedRecord, err := store.RecordFindByToken(ctx, token)
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByToken: Expected [err] to be nil after deletion but got [%v]", err.Error())
	}
	if deletedRecord != nil {
		t.Fatal("Test_Store_RecordDeleteByToken: Expected nil record after deletion but got a record")
	}

	// Test deleting non-existent record
	err = store.RecordDeleteByToken(ctx, "non_existent_token")
	if err != nil {
		t.Fatalf("Test_Store_RecordDeleteByToken: Expected [err] to be nil for non-existent token but got [%v]", err.Error())
	}
}
