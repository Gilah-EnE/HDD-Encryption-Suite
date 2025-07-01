package main

import (
	qt "github.com/mappu/miqt/qt"
	"log"
	"os"
)

func main() {
	qt.NewQApplication(os.Args)
	window := qt.NewQMainWindow(nil)
	window.SetMinimumSize(qt.NewQSize2(600, 192))
	window.Show()
	widget := qt.NewQWidget2()

	// File picker
	filePickerLayout := qt.NewQHBoxLayout2()
	fileNameField := qt.NewQLineEdit(widget)
	filePickerButton := qt.NewQPushButton(widget)
	filePickerButton.SetText("Select file")
	filePickerButton.OnClicked(func() {
		wDir, wDirErr := os.Getwd()
		if wDirErr != nil {
			log.Fatal(wDirErr)
		}
		filePickerDialog := qt.NewQFileDialog6(widget, "Select image file", wDir, "ALl files (*)")

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
	resultsLabel := qt.NewQLabel5("Encryption suite", widget)
	resultsField := qt.NewQLineEdit(widget)
	resultsField.SetReadOnly(true)
	resultsLayout.AddWidget2(resultsLabel.QWidget, 0, 0)
	resultsLayout.AddWidget2(resultsField.QWidget, 0, 1)

	// Sector size selection
	sectorSizeSwitchLayout := qt.NewQHBoxLayout2()
	standardSectorSizeRadio := qt.NewQRadioButton4("512 bytes (default)", widget)
	standardSectorSizeRadio.SetChecked(true)
	opticalSectorSizeRadio := qt.NewQRadioButton4("2048 bytes (optical media)", widget)
	advancedFormatSectorSizeRadio := qt.NewQRadioButton4("4096 bytes (AF disks)", widget)
	sectorSizeSwitchLayout.AddWidget(standardSectorSizeRadio.QWidget)
	sectorSizeSwitchLayout.AddWidget(opticalSectorSizeRadio.QWidget)
	sectorSizeSwitchLayout.AddWidget(advancedFormatSectorSizeRadio.QWidget)

	// Hail Mary mode
	hailMarySwitch := qt.NewQCheckBox4("Hail Mary mode (all patterns checked in all sectors - slow, horrible and painful)", widget)

	// Start button
	startButton := qt.NewQPushButton5("Run", widget)
	startButton.OnClicked(func() {
		var sectorSize int
		if standardSectorSizeRadio.IsChecked() {
			sectorSize = 512
		} else if opticalSectorSizeRadio.IsChecked() {
			sectorSize = 2048
		} else if advancedFormatSectorSizeRadio.IsChecked() {
			sectorSize = 4096
		}
		resultsField.SetText("")
		toolDetectionResult := toolDetection(fileNameField.Text(), sectorSize, hailMarySwitch.IsChecked())
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
	globalLayout.AddLayout(sectorSizeSwitchLayout.QLayout)
	globalLayout.AddWidget(hailMarySwitch.QWidget)
	globalLayout.AddWidget(startButton.QWidget)

	window.SetCentralWidget(widget)
	window.Show()
	qt.QApplication_Exec()
}
