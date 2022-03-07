package ui

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
	"math/rand"
	"strconv"

	clipboard "github.com/atotto/clipboard"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

//MyForm is the global form object
var MyForm *tview.Form

//OutputValue contains value of generated random string
var OutputValue string
var TheTime int64
var lowerLetterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var upperLetterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")
var symbolRunes = []rune("~=+%^*/()[]{}!@#$?|\\<>-_")
var punctuationRunes = []rune(".,:;'\"")

const (
	seedInputLabel        = "Seed"
	lengthInputLabel      = "Length"
	upperCaseInputLabel   = "Upper Case"
	numberInputLabel      = "Numbers"
	symbolInputLabel      = "Symbol"
	punctuationInputLabel = "Punctuation"
	outputValueLabel      = "Output:"
	showInputLabel        = "Show"
	generateButtonLabel   = "Generate"
	useButtonLabel        = "Copy"
	quitButtonLabel       = "Quit"
)

func initSeed(seedString string) {
	hash := sha256.New()
	io.WriteString(hash, seedString)
	//theTime := time.Now().UnixNano()
	timeBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(timeBytes, uint64(TheTime))
	var seedFromUser uint64 = binary.BigEndian.Uint64(hash.Sum(timeBytes))
	finalSeed := TheTime - int64(seedFromUser)
	rand.Seed(int64(finalSeed))
}

//newValue creates a new OutputValue based on the form's parameters
func newValue() {
	_, selectedLengthStr := (MyForm.GetFormItemByLabel(lengthInputLabel).(*tview.DropDown)).GetCurrentOption()
	selectedLength, _ := strconv.Atoi(selectedLengthStr)
	capitalLetters := (MyForm.GetFormItemByLabel(upperCaseInputLabel).(*tview.Checkbox)).IsChecked()
	numbers := (MyForm.GetFormItemByLabel(numberInputLabel).(*tview.Checkbox)).IsChecked()
	symbols := (MyForm.GetFormItemByLabel(symbolInputLabel).(*tview.Checkbox)).IsChecked()
	punctuation := (MyForm.GetFormItemByLabel(punctuationInputLabel).(*tview.Checkbox)).IsChecked()
	OutputValue = generator(selectedLength, capitalLetters, numbers, symbols, punctuation)
}

//updateFormWithValue creates or updates the form element with the label designated by outputValueLabel
func updateFormWithValue() {

	if MyForm.GetFormItemByLabel(outputValueLabel) == nil {
		MyForm.AddCheckbox(showInputLabel, false, func(show bool) {
			removeFormField()
			updatePasswordField(show, OutputValue)
		})
		updatePasswordField(false, OutputValue)
	} else {
		removeFormField()
		showObj := MyForm.GetFormItemByLabel(showInputLabel)
		show := showObj.(*tview.Checkbox).IsChecked()
		updatePasswordField(show, OutputValue)
	}

}

//removeFormField removes the previously generated output field if it exists
func removeFormField() {
	formItemNumber := MyForm.GetFormItemIndex(outputValueLabel)
	if formItemNumber != -1 {
		MyForm.RemoveFormItem(formItemNumber)
	}
}

//updatePasswordField creates a InputField or a PasswordField
func updatePasswordField(show bool, value string) {
	if show {
		MyForm.AddInputField(outputValueLabel, value, len(value), nil, nil)
	} else {
		MyForm.AddPasswordField(outputValueLabel, value, len(value), '*', nil)
	}

}

//moveToClipboard copies value from form to system clipboard
func moveToClipboard() {
	object := MyForm.GetFormItemByLabel(outputValueLabel)

	if object != nil {
		clipboard.WriteAll(object.(*tview.InputField).GetText())
		// object.(*tview.InputField).SetText()
	}
}

//generator creates random string using the characters specified
func generator(selectedLength int, capitalLetters bool, numbers bool, symbols bool, punctuation bool) string {
	allRunes := lowerLetterRunes
	//ensure one of each catagory is used
	//needs plenty of error checking

	buf := make([]rune, selectedLength)

	if capitalLetters {
		allRunes = append(allRunes, upperLetterRunes...)
	}
	if numbers {
		allRunes = append(allRunes, numberRunes...)
	}
	if symbols {
		allRunes = append(allRunes, symbolRunes...)
	}
	if punctuation {
		allRunes = append(allRunes, punctuationRunes...)
	}

	for i := 0; i < selectedLength; i++ {
		buf[i] = allRunes[rand.Intn(len(allRunes))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}

//CreateGeneratorForm returns a tview Form object that randomly generates passwords
func CreateGeneratorForm(app *tview.Application) *tview.Form {
	tview.Styles.ContrastBackgroundColor = tcell.Color236
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorCoral
	//all text color
	tview.Styles.PrimaryTextColor = tcell.ColorOldLace
	// color of prompt text
	tview.Styles.SecondaryTextColor = tcell.ColorPaleGoldenrod

	tview.Styles.BorderColor = tcell.ColorOliveDrab

	initSeed("")
	form := tview.NewForm().
		AddInputField(seedInputLabel, "", 32, nil, initSeed).
		AddDropDown(lengthInputLabel, []string{"16", "32", "64", "128", "256"}, 0, nil).
		AddCheckbox(upperCaseInputLabel, true, nil).
		AddCheckbox(numberInputLabel, true, nil).
		AddCheckbox(symbolInputLabel, true, nil).
		AddCheckbox(punctuationInputLabel, false, nil).
		AddButton(generateButtonLabel, func() {
			newValue()
			updateFormWithValue()
		}).
		AddButton(useButtonLabel, moveToClipboard).
		AddButton(quitButtonLabel, func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Password Options").SetTitleAlign(tview.AlignLeft)
	MyForm = form
	return MyForm
}
