# go-audio
---
A solution to convert PDFs into audiobooks (offline).
This solution kit consists of an Opentts engine hosted on docker and a CLI client that parses the given PDF file for text content & connects with Opentts to generate audio files.

## Installation & setup :-
---
* Go,Docker & Make should be pre-installed.
* Clone the repository.
* Run the command `make download`.
* Run the command `make setup`.
* CLI & Opentts engine setup is now ready to use.
* Run the command `go-audio --help` to explore various operations.

## Generate the audio files
---
* Run the command `go-audio aud --input=PATH_TO_PDF --output=PATH_TO_OUTPUT --voice=male|female`.
* Default value for given flags ::
    * input=./sample_pdf/test.pdf
    * output=homeDir/audio-go/output
    * voice=female.


