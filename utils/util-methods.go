package utils

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func ReadPdf(filePath string) (string, error) {
	var content bytes.Buffer
	f, r, err := pdf.Open(filePath)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}

	totalPage := r.NumPage()
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			for _, word := range row.Content {
				content.WriteString(word.S)
			}
		}
	}
	return content.String(), nil
}

func getVoice(voiceType string) string {
	if voiceType == MALE_VOICE {
		return MALE_SET_VOICE
	}
	return DEFAULT_VOICE
}

func GenAudio(content, voiceType, outPath string, i int, wg *sync.WaitGroup) {
	fmt.Println("File in Process :: ", i)
	err := os.MkdirAll(outPath, 0755) // make directory if not exists
	FatalErr(err)

	voice := getVoice(voiceType)

	params := url.Values{}
	params.Add(VOICE_KEY, voice)
	params.Add(TEXT_KEY, content)
	query := params.Encode()

	urlStr := fmt.Sprintf("%s%s", DEFAULT_SERVER_URL, query)
	resp, err := http.Get(urlStr)
	FatalErr(err)
	var op string
	op = fmt.Sprintf("%s.wav",strconv.Itoa(i))
	filePath := filepath.Join(outPath, op) // Set the filepath
	out, err := os.Create(filePath)        // Create the file
	FatalErr(err)

	if resp.StatusCode == 200 {
		_, err = io.Copy(out, resp.Body) // Write the body to file
		fmt.Println("Processed successfully file :: ", i)
	} else {
		fmt.Println("Processing failed for content file :: ", i)
	}
	defer func() {
		resp.Body.Close()
		out.Close()
		wg.Done()
	}()
}
