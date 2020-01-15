// Given two words (beginWord and endWord), and a dictionary's word list, find the length 
// of shortest transformation sequence from beginWord to endWord, such that:
// * Only one letter can be changed at a time.
// * Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
// Note:
// * Return 0 if there is no such transformation sequence.
// * All words have the same length.
// * All words contain only lowercase alphabetic characters.
// * You may assume no duplicates in the word list.
// * You may assume beginWord and endWord are non-empty and are not the same.
package main

import(
	"fmt"
	"container/list"
)

type WordLevel struct {
	word string
	level int
}

// The trick here is to use "generic words" with a placeholder * character. So
// e.g. "hot" becomes "*ot", "h*t" and "ho*". We put the generic words as key
// into a map, with the list of words (having the same generic word) as value.
// E.g. map[*ot] = hot, dot, bot etc. - those words are then our "adjancey nodes"
// which we visit in a BFS manner.
//
// We can limit the search space with Bidirectional Breadth First Search, that is
// we start searching from both the start- and the end word.
// -> The search would stop if we visit a node which has already been visited
//    "by the other BFS". The "word level" is then the sum of both BFS levels
func ladderLength(beginWord string, endWord string, wordList []string) int {

    // Since all words are of same length.
    L := len(beginWord)

    // Dictionary to hold combination of words that can be formed,
    // from any given word. By changing one letter at a time.
    allComboDict := make(map[string][]string)

	for _, word := range wordList {
		for i := 0; i < L; i++ {
			// Key is the generic word
			// Value is a list of words which have the same intermediate generic word
			newWord := fmt.Sprintf("%s*%s", word[0:i], word[i+1:])
			transformations, ok := allComboDict[newWord]
			if !ok {
				transformations = make([]string, 0)
			}
			transformations = append(transformations, word)
			allComboDict[newWord] = transformations
		}
	}

	// Queue for BFS
	queue := list.New()
	queue.PushFront(WordLevel{beginWord, 1})

	// Visited to make sure we don't repeat processing same word.
	visited := make(map[string]bool)
	visited[beginWord] = true
    
	// Breadth First Search (using a FIFO queue)
    for queue.Len() > 0 {

		node := queue.Back().Value.(WordLevel)
		queue.Remove(queue.Back())

		word := node.word
		level := node.level
		for i := 0; i < L; i++ {

			// Intermediate words for current word
			newWord := fmt.Sprintf("%s*%s", word[0:i], word[i+1:])
			// Next states are all the words which share the same intermediate state
			adjacencyWordList, ok := allComboDict[newWord]
			if ok {
				for _, adjacentWord := range adjacencyWordList {
					// If at any point if we find what we are looking for
					// i.e. the end word - we can return with the answer
					if (adjacentWord == endWord) {
						return level + 1
					}
					// Otherwise, add it to the BFS Queue. Also mark it visited
					if !visited[adjacentWord] {
						visited[adjacentWord] = true
						queue.PushFront(WordLevel{adjacentWord, level + 1})
					}
				}
			}
		}

	}
	// end node with endWord not found
    return 0;
  }

  func main() {
	beginWord := "hit"
    endWord := "cog"
	wordList := []string{"hot","dot","dog","lot","log","cog"}

	ladderLength := ladderLength(beginWord, endWord, wordList)
	fmt.Printf("Ladder length for begin %s and end %s and list %v: %d\n", beginWord, endWord, wordList, ladderLength)
  }