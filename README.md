# HDD-Encryption-Suite
## Installation
### Linux
Compile the Rust Regex (rure) library from source:
```bash
git clone https://github.com/rust-lang/regex
cargo build --release --manifest-path ./regex/regex-capi/Cargo.toml
cd ./regex/target/release
sudo cp librure.so /usr/local/lib
sudo ldconfig
```
Install Go and Qt5 libraries:
```bash
sudo apt install golang-go qtbase5-dev libqscintilla2-qt5-dev libqt5svg5-dev libqt5webchannel5-dev libqt5webkit5-dev qtbase5-private-dev qtmultimedia5-dev qtpdf5-dev qtwebengine5-dev qtwebengine5-private-dev build-essential
```
Build the app itself:
```bash
go build
```
### macOS
Install the Rust Regex (rure) library from Homebrew:
```bash
brew install rure
```
Install Go and Qt5 libraries:
```bash
xcode-select --install
brew install golang
brew install pkg-config
brew install qt@5
```
Build the app itself:
```bash
go build
```

## Usage
- Open the application;
- Select the image to be tested;
- Select appropriate sector size according to the media type;
- Use Hail Mary mode to enforce search of all patterns in all sectors **(optionally)**;
- Press the Start button and wait for the result!

## Pattern definition format
Default pattern collection is:
```go
patterns := map[string]SignatureData{
    "FreeBSD GELI": {"(?i)(47454f4d3a3a454c49)", -1},
    "BitLocker":    {"(?i)(eb58902d4656452d46532d0002080000)", 1},
    "LUKSv1":       {"(?i)4c554b53babe0001", 1},
    "LUKSv2":       {"(?i)4c554b53babe0002", 1},
    "FileVault v2": {"(?i)41505342.{456}0800000000000000", 0},
}
```
The pattern definition structure is `Name: {Regular expression, Sector number}`. Sector number should be:
- positive for forward search mode (for example 1 corresponds to the sector 1);
- negative for reverse search mode (for example, -1 corresponds to the last sector, -2 to the second last etc);
- `0` for all sector search mode.