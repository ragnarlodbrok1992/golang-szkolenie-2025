package main

// Create a hangman (wisielec) game in go
// 1. Program prepares a random word for user to guess.
// 2. It asks for input - a single character.
// 3. If user guess correctly program shows this character or characters to user.
// 4. If not - it draws another piece of hangman.
// 5. Game ends if user guesses the word or if hangman is drawn fully.
// Hangman stages are provided for convienance.

var stages = []string{
    `
  +---+
  |   |
      |
      |
      |
      |
=========
    `,
    `
  +---+
  |   |
  O   |
      |
      |
      |
=========
    `,
    `
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
    `,
    `
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
    `,
    `
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
    `,
    `
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
    `,
    `
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
    `,
}

func main() {
	// Put code here.
}
