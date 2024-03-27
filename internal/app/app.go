package app

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"github.com/albakov/go-hangman/internal/config"
	"github.com/albakov/go-hangman/internal/entity"
	"math/big"
	"os"
	"strings"
)

const maxAttempts = 5

const (
	inputPrefix         = "-> "
	messageWelcome      = "[Н]овая игра или [В]ыйти\n"
	messageGuessedWord  = "Загаданное слово: %s\n"
	messageGameOver     = "Вы проиграли!\n"
	messageAttemptsLeft = "Осталось попыток: %d\n"
	messageWin          = "Вы выиграли!\n"
	messageTextWrong    = "Введен некорректный символ!\n"
)

type App struct {
	isBeginning             bool
	attempts                int
	wordsCount              int
	hiddenWord, guessedWord []string
	words                   []string
	reader                  *bufio.Reader
}

func New(credentials *config.Config) *App {
	return &App{
		isBeginning: true,
		reader:      bufio.NewReader(os.Stdin),
		words:       credentials.Words,
		wordsCount:  len(credentials.Words),
	}
}

func (a *App) Start() {
	for {
		if a.isBeginning {
			a.onStart()
			continue
		}

		if a.isLetterFound() {
			a.onGuessed()
		} else {
			a.onNotGuessed()
		}
	}
}

func (a *App) onStart() {
	a.showMessage(messageWelcome)
	a.showMessage(inputPrefix)

	text := a.textFromInput()

	// Exit game when type "в"
	if text == "в" {
		os.Exit(0)
	}

	// Start new game when type "н"
	if text != "н" {
		return
	}

	a.generateWord()
	a.showMessage(messageGuessedWord, strings.Repeat("*", len(a.hiddenWord)))
	a.isBeginning = false
	a.attempts = maxAttempts

	for i := range a.hiddenWord {
		a.guessedWord[i] = "*"
	}
}

func (a *App) isLetterFound() bool {
	a.showMessage(inputPrefix)
	text := a.textFromInput()
	found := false

	for i, letter := range a.hiddenWord {
		if strings.EqualFold(letter, text) {
			a.guessedWord[i] = letter
			found = true
		}
	}

	return found
}

func (a *App) onGuessed() {
	strUserWord := strings.Join(a.guessedWord, "")
	a.showMessage("%s\n", strUserWord)

	if !strings.Contains(strUserWord, "*") {
		a.showMessage("%s\n", messageWin)
		a.reset()
	}
}

func (a *App) onNotGuessed() {
	a.attempts--

	if a.attempts == 0 {
		a.showMessage(messageGameOver)
		a.drawHangman()
		a.showMessage(messageGuessedWord, strings.Join(a.hiddenWord, ""))
		a.isBeginning = true

		return
	}

	a.showMessage(messageAttemptsLeft, a.attempts)
	a.drawHangman()
}

func (a *App) textFromInput() string {
	text, err := a.reader.ReadString('\n')
	if err != nil {
		a.showMessage(messageTextWrong)

		return ""
	}

	text = strings.Replace(text, "\n", "", -1)

	return strings.ToLower(text)
}

func (a *App) generateWord() {
	i := a.randNumber(0, int64(a.wordsCount-1))
	a.hiddenWord = strings.Split(a.words[i], "")
	a.guessedWord = make([]string, len(a.hiddenWord))
}

func (a *App) drawHangman() {
	a.showMessage("%s\n", entity.HangmanStages[a.attempts])
}

func (a *App) showMessage(message string, args ...interface{}) {
	fmt.Printf(message, args...)
}

func (a *App) randNumber(min, max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))

	return n.Int64() + min
}

func (a *App) reset() {
	a.isBeginning = true
	a.attempts = maxAttempts
}
