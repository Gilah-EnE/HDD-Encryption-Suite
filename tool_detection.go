package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/BurntSushi/rure-go"
	"github.com/Gilah-EnE/dec24-go/test_suite"
)

type SignatureData struct {
	regex  string
	sector int64
}

type AdvancedSignatureMap struct {
	regex  *rure.Regex
	sector int64
}

func foundSignaturesTotalToReadable(foundSignaturesTotal map[string]int) string {
	var readable string
	for key, value := range foundSignaturesTotal {
		readable = readable + fmt.Sprintf("%s - %d, ", key, value)
	}
	return readable
}

func toolDetection(fileName string, blockSize int, hailMaryMode bool) map[string]int {
	signatures := make(map[string]AdvancedSignatureMap)

	patterns := map[string]SignatureData{
		"FreeBSD GELI": {"(?i)(47454f4d3a3a454c49)", -1},
		"BitLocker":    {"(?i)(eb58902d4656452d46532d0002080000)", 1},
		"LUKSv1":       {"(?i)4c554b53babe0001", 1},
		"LUKSv2":       {"(?i)4c554b53babe0002", 1},
		"FileVault v2": {"(?i)41505342.{456}0800000000000000", 0},
	}

	foundSignaturesTotal := make(map[string]int)
	for name, pattern := range patterns {
		regex, err := rure.Compile(pattern.regex)
		if err != nil {
			log.Fatalf("failed to compile pattern for %s: %v", name, err)
		}
		signatures[name] = AdvancedSignatureMap{regex: regex, sector: pattern.sector}
		foundSignaturesTotal[name] = 0
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		fileCloseErr := file.Close()
		if fileCloseErr != nil {
			log.Fatal(fileCloseErr)
		}
	}(file)

	buffer := make([]byte, blockSize)

	if hailMaryMode {
		buffer := make([]byte, blockSize)
		n := 0
		for {
			bytesRead, err := file.Read(buffer)
			if bytesRead == 0 || err != nil {
				break
			}

			n += bytesRead
			fmt.Printf("%.1f ", float32(n)/1048576)

			// Convert bytes to hex string
			hexData := hex.EncodeToString(buffer[:bytesRead])
			for sigType := range signatures {
				if entry, ok := signatures[sigType]; ok {
					foundSignaturesTotal[sigType] += test_suite.FindBytesPattern(hexData, entry.regex)
				}
			}
			fmt.Print("\r")
		}
	} else {

		for sigType := range signatures {
			if entry, ok := signatures[sigType]; ok {
				skip := entry.sector

				var seekErr error

				if skip == 0 {
					buffer := make([]byte, blockSize)
					n := 0
					for {
						bytesRead, err := file.Read(buffer)
						if bytesRead == 0 || err != nil {
							break
						}

						n += bytesRead
						fmt.Printf("%.1f ", float32(n)/1048576)

						// Convert bytes to hex string
						hexData := hex.EncodeToString(buffer[:bytesRead])
						foundSignaturesTotal[sigType] += test_suite.FindBytesPattern(hexData, entry.regex)
						fmt.Print("\r")
					}
				} else if skip != 0 {
					if skip < 0 {
						_, seekErr = file.Seek(int64(blockSize*int(math.Abs(float64(skip))-2)), 2)
					} else if skip > 0 {
						skip = skip - 1
						_, seekErr = file.Seek(int64(blockSize-1)*skip, 0)
					}
					if seekErr != nil {
						log.Fatalln("Seek error: ", seekErr)
					}
					bytesRead, fileReadErr := file.Read(buffer)
					if bytesRead == 0 || fileReadErr != nil {
						break
					}
					hexData := hex.EncodeToString(buffer[:bytesRead])
					foundSignaturesTotal[sigType] += test_suite.FindBytesPattern(hexData, entry.regex)
					_, returnSeekErr := file.Seek(0, 0)
					if returnSeekErr != nil {
						log.Fatalln("Return seek error: ", returnSeekErr)
					}
				}
			}
		}
	}
	fmt.Print("\r")
	fmt.Println(foundSignaturesTotalToReadable(foundSignaturesTotal))
	return foundSignaturesTotal
}
