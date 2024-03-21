package main

import (
	"testing"
)

func TestSearchAddedContent(t *testing.T) {
	// Test when there is added content at the beginning of the updated text
	t.Run("Added content at the beginning", func(t *testing.T) {
		oldText := "world"
		updatedText := "hello world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})

	
	// Test when there is added content at the end of the updated text
	t.Run("Added content at the end", func(t *testing.T) {
		oldText := "hello"
		updatedText := "hello world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})
	
	// Test when there is added content in the middle of the updated text
	t.Run("Added content in the middle", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello there world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})
	// Test when there is no added content
	t.Run("No added content", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})
}
func TestSearchDeletedContent(t *testing.T) {
	// Test when there is deleted content at the beginning of the old text
	t.Run("Deleted content at the beginning", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})

	// Test when there is deleted content at the end of the old text
	t.Run("Deleted content at the end", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})

	// Test when there is deleted content in the middle of the old text
	t.Run("Deleted content in the middle", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello there"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})

	// Test when there is no deleted content
	t.Run("No deleted content", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})
}
func TestSearchModifiedContent(t *testing.T) {
	//Test when there is modified content at the beginning of the old text
	t.Run("Modified content at the beginning", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "jello world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})

	// Test when there is modified content at the end of the old text
	t.Run("Modified content at the end", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello worlx"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})

	// Test when there is modified content in the middle of the old text
	t.Run("Modified content in the middle", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello xorld"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})

	// Test when there is no modified content
	t.Run("No modified content", func(t *testing.T) {
		oldText := "hello world"
		updatedText := "hello world"
		windowSize := 2
		delta := checkString(oldText, updatedText, windowSize,1)		
		expectedAddedContent := replaceDelta(oldText, delta)
		if updatedText != expectedAddedContent {
			t.Errorf("Test failed. Expected: %s Got: %s", updatedText, expectedAddedContent)
		}
	})
}
