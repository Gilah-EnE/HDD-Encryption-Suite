package main

import (
	"github.com/mappu/miqt/qt"
	"log"
	"os"
)

func main() {
	qt.NewQApplication(os.Args)
	window := qt.NewQMainWindow(nil)
	window.SetMinimumSize(qt.NewQSize2(600, 144))
	window.Show()
	widget := qt.NewQWidget2()

	// File picker
	filePickerLayout := qt.NewQHBoxLayout2()
	fileNameField := qt.NewQLineEdit(widget)
	filePickerButton := qt.NewQPushButton(widget)
	filePickerButton.SetText("Выбор файла")
	filePickerButton.OnClicked(func() {
		wDir, wDirErr := os.Getwd()
		if wDirErr != nil {
			log.Fatal(wDirErr)
		}
		filePickerDialog := qt.NewQFileDialog6(widget, "Выберите файл для анализа", wDir, "Все файлы (*)")

		if filePickerDialog.Exec() == int(qt.QDialog__Accepted) {
			selectedFile := filePickerDialog.SelectedFiles()
			if len(selectedFile) > 0 {
				fileName := selectedFile[0]
				fileNameField.SetText(fileName)
			}
		}
	})
	filePickerLayout.AddWidget(fileNameField.QWidget)
	filePickerLayout.AddWidget(filePickerButton.QWidget)

	// Results display
	resultsLayout := qt.NewQGridLayout2()
	resultsLabel := qt.NewQLabel5("Утилита шифрования", widget)
	resultsField := qt.NewQLineEdit(widget)
	resultsLayout.AddWidget2(resultsLabel.QWidget, 0, 0)
	resultsLayout.AddWidget2(resultsField.QWidget, 0, 1)

	// Start button
	startButton := qt.NewQPushButton5("Запуск", widget)
	startButton.OnClicked(func() {
		toolDetectionResult := toolDetection(fileNameField.Text(), 512)
		for suite, value := range toolDetectionResult {
			if value > 0 {
				resultsField.SetText(suite)
			}
		}
	})

	// Global layout
	globalLayout := qt.NewQVBoxLayout(widget)
	globalLayout.AddLayout(filePickerLayout.QLayout)
	globalLayout.AddLayout(resultsLayout.QLayout)
	globalLayout.AddWidget(startButton.QWidget)

	window.SetCentralWidget(widget)
	window.Show()
	qt.QApplication_Exec()
}
