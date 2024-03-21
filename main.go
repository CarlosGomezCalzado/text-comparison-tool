/*
This Go code performs text comparison between two strings, identifying added, deleted, and modified content. It utilizes a rolling hash algorithm for efficient text search.

Package main implements a text comparison tool.

The TextSearch struct represents a search context for sliding window hashing.

The CustomError struct defines a custom error type for handling errors.

The following functions are implemented:

1. GetWindowString:
   - Parameters: None
   - Results: Returns the current window of text.
   - Description: Returns the current window of text being analyzed.

2. Slide:
   - Parameters: None
   - Results: Returns a custom error, the updated hash, and the current window of text.
   - Description: Slides the window to calculate the hash of the next text segment.

3. GetHash:
   - Parameters: None
   - Results: Returns the current hash value.
   - Description: Retrieves the current hash value of the text.

4. CreateBuffer:
   - Parameters: input (string), windowSize (int)
   - Results: None
   - Description: Initializes the text buffer with a specific window size.

5. SetStart:
   - Parameters: index (int), window (int)
   - Results: None
   - Description: Sets the starting point of the window for hashing.

6. SearchFirstDif:
   - Parameters: text1 (string), text2 (string), windowSize (int)
   - Results: Equal text until first difference, index of first difference, boolean indicating completion, error
   - Description: Searches for the first difference between two texts.

7. searchAddedContent:
   - Parameters: text1 (string), text2 (string), windowSize (int)
   - Results: Added content, old index, new index, boolean indicating completion
   - Description: Searches for added content in the updated text.

8. searchDeletedContent:
   - Parameters: text1 (string), text2 (string), windowSize (int)
   - Results: Deleted content, old index, new index, boolean indicating completion
   - Description: Searches for deleted content in the updated text.

9. searchModifiedContent:
   - Parameters: text1 (string), text2 (string), windowSize (int)
   - Results: Previous content, new content, old index, new index, boolean indicating completion
   - Description: Searches for modified content in the updated text.

10. checkString:
    - Parameters: old (string), updated (string), windowSize (int), oldGeneralIndex (int)
    - Results: Comparison result
    - Description: Recursively checks for differences between two texts.

11. readLine:
    - Parameters: None
    - Results: User input string
    - Description: Reads a line of input from standard input.

12. getInput:
    - Parameters: None
    - Results: Old text, updated text, window size
    - Description: Gets user input for text comparison.

13. displayResult:
    - Parameters: old (string), updated (string), result (string)
    - Results: None
    - Description: Displays the old text, updated text, and comparison result.

14. main:
    - Parameters: None
    - Results: None
    - Description: Orchestrates the text comparison process, obtaining input, performing comparison, and displaying results.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"bufio"
	"os"
	"strings"
)

type TextSearch struct {
	buffer     string
	hash       int
	index      int
	length     int
	prime      int
	windowSize int
	lastError  error
}

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

// Obtain current window string
func (ts *TextSearch) GetWindowString() string {
	return ts.buffer[ts.index:]
}

// Slide the window to calculate the hash of the next text segment.
func (ts *TextSearch) Slide() (*CustomError, int, string) {
	if ts.index+ts.windowSize >= ts.length {
		ts.lastError = &CustomError{message: "EOF"}
		return ts.lastError.(*CustomError), ts.hash, ts.GetWindowString()
	}
	// Remove the contribution of the oldest character.
	ts.hash = (ts.hash - int(ts.buffer[ts.index])*int(math.Pow(256, float64(ts.windowSize-1)))) % ts.prime
	if ts.hash < 0 {
		ts.hash += ts.prime // Ensure that the result is positive
	}

	// Add the contribution of the new character
	ts.hash = (ts.hash*256 + int(ts.buffer[ts.index+ts.windowSize])) % ts.prime
	ts.index++
	return nil, ts.hash, ts.GetWindowString()
}

// Get the current hash of the text
func (ts *TextSearch) GetHash() int {
	return ts.hash
}

// Create a text buffer with a specific window size
func (ts *TextSearch) CreateBuffer(input string, windowSize int) {
	ts.buffer = input
	ts.hash = 0
	ts.prime = 5381
	ts.length = len(input)
	ts.windowSize = windowSize
	ts.lastError = nil
}

// Set the starting point of the window
func (ts *TextSearch) SetStart(index, window int) {
	ts.index = index
	ts.windowSize = window
	ts.hash = 0
	ts.lastError = nil
	for i := index; i < index + window; i++ {
		ts.hash = (ts.hash*256 + int(ts.buffer[i])) % ts.prime
	}
}



func SearchFirstDif(text1, text2 string, windowSize int) (string, int, bool, error) {
	// We create two instances of TextSearch for the two texts
	var text1Search, text2Search TextSearch
	text1Search.CreateBuffer(text1, windowSize)
	text1Search.SetStart(0, windowSize)
	text2Search.CreateBuffer(text2, windowSize)
	text2Search.SetStart(0, windowSize)

	// Variables to track the index of the first difference
	index := 0
	text1Search.SetStart(index, 1)
	text2Search.SetStart(index, 1)

	// Get the new hashes
	newHash1 := text1Search.GetHash()
	newHash2 := text2Search.GetHash()

	boolRes := false
	// If the hashes are different and the window size is 1, we find the exact index of the first different character
	if newHash1 != newHash2 {
		return "", index, boolRes, nil	
	} 
	
	// Iterate until we reach the end of one of the texts
	for {
		// Get the hashes of the two texts
		hash1 := text1Search.GetHash()
		hash2 := text2Search.GetHash()

		// If the hashes are different, we find the first difference
		if hash1 != hash2 {
			// Reduce the window size until finding the exact index of the first different character
			for i:= 1; i < windowSize; i++{
				// Volver a calcular el hash desde el punto donde se detectó la diferencia
				text1Search.SetStart(index, i)
				text2Search.SetStart(index, i)

				// Get the new hashes
				newHash1 := text1Search.GetHash()
				newHash2 := text2Search.GetHash()

				// If the hashes are different and the window size is 1, we have found the exact index of the first different character
				if newHash1 != newHash2 {
					break
				} else {
					index ++
				}
				
			}
			break
		} else {
			// Advance the windows
			text1Search.Slide()
			text2Search.Slide()

			// We increment the index
			index++
		}
		

		// Check if we have reached the end of either of the texts
		if err := text1Search.lastError; err != nil || text2Search.lastError != nil {
			boolRes = true
			break
		}
	}

	// Build the text string that is the same in both strings up to the first difference
	equalText := text1[:index]

	return equalText, index, boolRes, nil
}

func searchAddedContent(text1, text2 string, windowSize int) (string, int, int, bool){
	if len(text1) <= windowSize || len(text2) <= windowSize{
		windowSize = 1
	}
	var text1Search, text2Search TextSearch
	text1Search.CreateBuffer(text1, windowSize)
	text1Search.SetStart(0, windowSize)
	text2Search.CreateBuffer(text2, windowSize)
	text2Search.SetStart(0, windowSize)
	indexNew := 0
	indexOld := 0
	addedContent := ""
	boolRes := true
	for {
		if err := text1Search.lastError; err != nil || text2Search.lastError != nil {
			boolRes = false
			indexNew++
			break
		} 
		// Get the hashes of the two texts
		hash2 := text2Search.GetHash()
		hash1 := text1Search.GetHash()
		// If the hashes are different, we have found the first difference
		if hash1 != hash2 {
			addedContent = addedContent + string(text2[indexNew])
			text2Search.Slide()
			if err := text1Search.lastError; err != nil || text2Search.lastError != nil {
				indexNew++
			}
		} else {
			break
		}
	}
	return addedContent, indexOld, indexNew, boolRes

}

func searchDeletedContent(text1, text2 string, windowSize int) (string, int, int, bool){
	if len(text1) <= windowSize || len(text2) <= windowSize{
		windowSize = 1
	}
	var text1Search, text2Search TextSearch
	text1Search.CreateBuffer(text1, windowSize)
	text1Search.SetStart(0, windowSize)
	text2Search.CreateBuffer(text2, windowSize)
	text2Search.SetStart(0, windowSize)
	indexOld := 0
	indesUpd := 0
	deletedContent := ""
	boolRes := true
	for {
		if  text1Search.lastError != nil || text2Search.lastError != nil  {
			boolRes = false
			if text1Search.lastError != nil {
				indexOld++
			}
			break
		}
		// Get the hashes of the two texts
		hash2 := text2Search.GetHash()
		hash1 := text1Search.GetHash()
		// If the hashes are different, we have found the first difference
		if hash1 != hash2 {
			deletedContent = deletedContent + string(text1[indexOld])
			text1Search.Slide()
			if text1Search.lastError == nil {
				indexOld++
			}
		} else {
			break
		}
	}
	return deletedContent, indexOld, indesUpd, boolRes

}

func searchModifiedContent(text1, text2 string, windowSize int) (string, string, int, int, bool){
	if len(text1) <= windowSize || len(text2) <= windowSize{
		windowSize = 1
	}
	var text1Search, text2Search TextSearch
	text1Search.CreateBuffer(text1, windowSize)
	text1Search.SetStart(0, windowSize)
	text2Search.CreateBuffer(text2, windowSize)
	text2Search.SetStart(0, windowSize)
	
	indexOld := 0
	indexNew := 0
	previousContent := ""
	newContent := ""
	boolRes := false
	hash2 := text2Search.GetHash()
	hash1 := text1Search.GetHash()
	
	for {
		// Check if we have reached the end of one of the texts
		if  text1Search.lastError != nil || text2Search.lastError != nil  {
			boolRes = true
			if (text1Search.lastError != nil || text2Search.lastError != nil )&& windowSize >1{
				text1Search.SetStart(indexOld, 1)
				text2Search.SetStart(indexNew, 1)
			} else if len(string(text1[indexOld])) > 1 && len(string(text2[indexNew])) > 1 && windowSize == 1{
				indexOld++
				indexNew++
			} else {
				break
			}
		} 
		// Get the hashes of the two texts
		hash2 = text2Search.GetHash()
		hash1 = text1Search.GetHash()
		// If the hashes are different, we have found the first difference
		if hash1 != hash2 {
			previousContent = previousContent + string(text1[indexOld])
			newContent = newContent + string(text2[indexNew])
			text1Search.Slide()
			text2Search.Slide()
			if len(string(text1[indexOld:])) >= 1{ 
				indexOld++
			}
			if len(string(text2[indexNew:])) >= 1  {
				indexNew++
			}
			if len(string(text1[indexOld:])) <= 1 || len(string(text2[indexNew:])) <= 1 {
				boolRes = true
				break
			}			
		} else {
			boolRes = true
			break
		}
	}
	return previousContent, newContent, indexOld, indexNew, boolRes

}

func checkString(old, updated string, windowSize int, oldGeneralIndex int) string{
	if len(old) < windowSize || len(updated) < windowSize{
		windowSize = 1
	}
	// Search for the first difference between the two texts
	_, firstDiffIndex, isEnd, err := SearchFirstDif(old, updated, windowSize)
	if err != nil {
		return ""
	}
	oldGeneralIndex = oldGeneralIndex + firstDiffIndex
	extra := ""
	addedContent := ""
	deletedContent := ""
	previousContent := ""
	newContent := ""
	oldAddIndex := 0
	newAddIndex := 0
	oldDelIndex := 0
	newPatternIndex := 0
	isAdded := false
	isDel := false
	oldModifiedIndex := 0
	newModifiedIndex := 0
	isModified := false
	old = old[firstDiffIndex:]
	updated = updated[firstDiffIndex:]
	if !isEnd {
		// If we have differences in the following parts
		addedContent, oldAddIndex, newAddIndex, isAdded = searchAddedContent(old, updated,windowSize)
		deletedContent, oldDelIndex, newPatternIndex,isDel = searchDeletedContent(old, updated,1)
		previousContent, newContent, oldModifiedIndex, newModifiedIndex, isModified = searchModifiedContent(old, updated,windowSize)

		if isModified {// If it is a modification
			old = old[oldModifiedIndex:]
			updated = updated[newModifiedIndex:]
			extra = "Start character: "+strconv.Itoa(oldGeneralIndex)+" [--- "+ previousContent+"][+++ "+newContent+"]\n"
			oldGeneralIndex += oldModifiedIndex
		} else if isAdded {// If it is an added content
			old = old[oldAddIndex:]
			updated = updated[newAddIndex:]
			addedContent = "Start character: "+strconv.Itoa(oldGeneralIndex)+" [+++ "+addedContent+"]\n"
			extra = addedContent
			oldGeneralIndex += oldAddIndex
		} else if isDel {// If it is a deleted
			old = old[oldDelIndex:]
			updated = updated[newPatternIndex:]
			deletedContent = "Start character: "+strconv.Itoa(oldGeneralIndex)+" [--- "+deletedContent+"]\n"
			extra = deletedContent
			oldGeneralIndex += oldDelIndex
		} else { // end case
			old = old[oldModifiedIndex:]
			updated = updated[newModifiedIndex:]
			extra = "Start character: "+strconv.Itoa(oldGeneralIndex)+" [--- "+ previousContent+"][+++ "+newContent+"]"
			oldGeneralIndex += oldModifiedIndex
		}	
	} 
	
	recursiveResult := ""
	if len(old) > 1 && len(updated) > 1{
		recursiveResult = checkString(old, updated, windowSize, oldGeneralIndex ) // Recursive call for check the rest of the content
	}else if len(old) == 1 || len(updated) == 1 { // Last characters checkings
		recursiveResult = checkString(old, updated, 1, oldGeneralIndex )
	} else if len(old) == 0 && len(updated) > 0 {
		recursiveResult = "Start character: "+strconv.Itoa(oldGeneralIndex)+" [+++ "+updated+"]"
	} else if len(old) > 0 && len(updated) == 0 {
		recursiveResult = "Start character: "+strconv.Itoa(oldGeneralIndex)+" [--- "+old+"]"
	} 

	return  extra + recursiveResult
}
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func getInput() (string, string, int) {
	// This function gets user input for the two texts and the window size
	var old, updated string
	var windowSize int

	// Prompt the user to enter the old text
	fmt.Println("Enter the old text:")
	old = readLine()

	// Prompt the user to enter the updated text
	fmt.Println("Enter the updated text:")
	updated = readLine()

	// Prompt the user to enter the window size
	fmt.Println("Enter the window size for comparison:")
	fmt.Scanln(&windowSize)
	fmt.Println("_______________________________________")

	return old, updated, windowSize
}

func displayResult(old, updated, result string) {
	// This function displays the old text, updated text, and comparison result
	fmt.Println("Old text:", old)
	fmt.Println("Updated text:", updated)
	fmt.Println("Comparison result:")
	fmt.Println(result)
}

func replaceDelta(old, delta string) string {
	// Find the start index "Start character: X"
	if delta == ""{
		return old
	}
	lines := strings.Split(delta, "\n")
	result := old
	for _, value := range lines{
		if len(value)>0{
			startIndexStrSlide := strings.Split(value, "Start character: ")
			if len(startIndexStrSlide) <= 1{
				return "fallo "+value
			}
			startIndexStr :=startIndexStrSlide[1]
			startIndex, err := strconv.Atoi(strings.Split(startIndexStr, " ")[0])
			if err != nil {
				// If index conversion fails, return the original string
				return old
			}
			startIndex--
			startIndexMark := strings.Split(value, "[--- ")
			if len(startIndexMark) > 1{
				startIndexMark2 := strings.Split(startIndexMark[1], "]")
				numCharDel := len(startIndexMark2[0])
				result = fmt.Sprintf("%s%s", result[:startIndex], old[startIndex+numCharDel:])
				if  startIndex+numCharDel < len(old){
					result = fmt.Sprintf("%s%s", result[:startIndex], old[startIndex+numCharDel:])
				} else {
					result = fmt.Sprintf("%s", result[:startIndex])
				}
			}
			startIndexMark = strings.Split(value, "[+++ ")
			if len(startIndexMark) > 1{
				startIndexMark2 := strings.Split(startIndexMark[1], "]")
				numCharAdd := len(startIndexMark2[0])
				if  startIndex+numCharAdd < len(old){
					result = fmt.Sprintf("%s%s%s", result[:startIndex], startIndexMark2[0], old[startIndex+numCharAdd:])
				} else {
					result = fmt.Sprintf("%s%s", result[:startIndex], startIndexMark2[0])
				}
				
			}
			old = result
		}
	}

	return result
}

func main() {
	// Separate input/output operations from calculations
	old, updated, windowSize := getInput()
	result := checkString(old, updated, windowSize, 1)
	displayResult(old, updated, result)
	result = replaceDelta(old, result)
	fmt.Println(result)
}
