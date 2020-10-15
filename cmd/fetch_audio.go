package cmd

import (
	"fmt"
	"path/filepath"
	"sync"
	"unicode/utf8"

	"github.com/Harry-027/go-audio/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// split text content into chunks ...
func SplitStr(longString string, maxLen int) []string {
	var splits []string
	var l, r int
	for l, r = 0, maxLen; r < len(longString); l, r = r, r+maxLen {
		for !utf8.RuneStart(longString[r]) {
			r--
		}
		splits = append(splits, longString[l:r])
	}
	splits = append(splits, longString[l:])
	return splits
}

// Generate audio output for given text content ...
func genOutput(inPath, outPath, voiceType string) {
	fmt.Println("Processing ...")
	var wg sync.WaitGroup
	content, err := utils.ReadPdf(inPath)
	fmt.Println("Content from Pdf ::", content)
	utils.FatalErr(err)
	const maxLen = 10000
	splits := SplitStr(content, maxLen)
	for i, v := range splits {
		wg.Add(1)
		go utils.GenAudio(v, voiceType, outPath, i, &wg)
	}
	wg.Wait()
}

// Cobra audio command ...
var genCmd = &cobra.Command{
	Use:   "aud",
	Short: "Command to convert pdf into audio file. Bigger files would be processed into chunks.",
	Run: func(cmd *cobra.Command, args []string) {
		inPath, err := cmd.Flags().GetString(utils.FLAG_INPUT)
		utils.LogErr(err)
		outPath, err := cmd.Flags().GetString(utils.FLAG_OUTPUT)
		utils.LogErr(err)
		voiceType, err := cmd.Flags().GetString(utils.FLAG_VOICE)
		utils.LogErr(err)
		genOutput(inPath, outPath, voiceType)
		fmt.Println("Completed!! Output files available at ::", outPath)
	},
}

// Cobra add generate command ...
func init() {
	home, err := homedir.Dir() // Fetch the current user home dir.
	utils.PanicErr(err)        // Panic in case user dir not available
	defaultOutputDir := filepath.Join(home, utils.DEFAULT_OUTPUT_PATH)
	genCmd.Flags().String(utils.FLAG_INPUT, utils.DEFAULT_PDF_PATH, utils.FLAG_INPUT_DESC)
	genCmd.Flags().String(utils.FLAG_OUTPUT, defaultOutputDir, utils.FLAG_OUTPUT_DESC)
	genCmd.Flags().String(utils.FLAG_VOICE, utils.FEMALE_VOICE, utils.FLAG_VOICE_DESC)
	RootCmd.AddCommand(genCmd)
}
