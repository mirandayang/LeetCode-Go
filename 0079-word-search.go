package main

import (
	"fmt"
)

func wordSearch(board [][]byte, word string) bool {

	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			k := 0
			fmt.Println("position: \n", i, j)
			match := matchNext(board, visited, word, i, j, k)
			if match == true {
				return true
			}
		}
	}
	return false
}

func matchNext(board [][]byte, visited [][]bool, word string, i int, j int, k int) bool {

	match := false

	// Check current char match or not
	if board[i][j] == word[k] {

		fmt.Printf("search position: %d %d for char: %#v match:%t \n", i, j, word[k], true)

		// Mapping to the end of word
		if k == len(word)-1 {
			return true
		}

		visited[i][j] = true

		// Continue comparing on the left
		if i > 0 && visited[i-1][j] == false {
			match = matchNext(board, visited, word, i-1, j, k+1)
		}

		// Continue comparing on the right
		if match == false && i < len(board)-1 && visited[i+1][j] == false {
			match = matchNext(board, visited, word, i+1, j, k+1)
		}

		// Continue comparing on the up
		if match == false && j > 0 && visited[i][j-1] == false {
			match = matchNext(board, visited, word, i, j-1, k+1)
		}

		// Continue comparing on the down
		if match == false && j < len(board[i])-1 && visited[i][j+1] == false {
			match = matchNext(board, visited, word, i, j+1, k+1)
		}

	} else {
		fmt.Printf("search position: %d %d for char: %#v match:%t \n", i, j, word[k], false)
	}

	if match == false {
		visited[i][j] = false
	}

	return match
}

// func main() {
// 	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
// 	fmt.Println(wordSearch(board, "ABCESEEEFS"))
// }
