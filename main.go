// Package main. This program converts a markdown file to a pdf file by applying a custom stylesheet.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"github.com/gomarkdown/markdown"
	renderer "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	// inputFilename is the file path to the input markdown file
	inputFilename string
	// cssFilename is the file path to the custom stylesheet
	cssFilename string
)

// Run me: go run main.go --help
// You need to install go first.
func main() {
	parseArgs()
	fileContents := readFileOrPanic(inputFilename)
	htmlContents := markdownToHtml(fileContents)
	htmlFilename := inputFilename + ".html"
	writeFileOrPanic(htmlContents, htmlFilename)
	htmlToPdf(htmlFilename, inputFilename+".pdf")
}

// htmlToPdf converts htmlFilename to a pdf file by starting a chromium in headless mode and let it print the html to pdf
func htmlToPdf(htmlFilename, pdfFilename string) {
	cmd := exec.Command(
		"chromium-browser",
		// "--no-sandbox",
		"--disable-gpu",
		"--headless", // Do not open a chromium window
		"--print-to-pdf-no-header", // Disables printing of file name and date
		"--run-all-compositor-stages-before-draw", // Wait until document fully printed
		"--print-to-pdf="+pdfFilename,
		htmlFilename,
	)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Could not execute chromium. Maybe chromium-browser is not installed?")
		panic(err)
	}
}

// markdownToHtml converts markdownContent to html
func markdownToHtml(markdownContent []byte) []byte {
	parser := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs | parser.Footnotes | parser.SuperSubscript)
	renderer := renderer.NewRenderer(renderer.RendererOptions{
		Flags:                      renderer.CommonFlags | renderer.CompletePage | renderer.FootnoteReturnLinks,
		Title:                      inputFilename,
		CSS:                        cssFilename,
		FootnoteReturnLinkContents: "↩",
		// Import prism.js fpr code highlighting
		Head:                       []byte(`<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.28.0/themes/prism.min.css" integrity="sha512-tN7Ec6zAFaVSG3TpNAKtk4DOHNpSwKHxxrsiw4GHKESGPs5njn/0sMCUMl2svV4wo4BK/rCP7juYz+zx+l6oeQ==" crossorigin="anonymous" referrerpolicy="no-referrer" /><script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.28.0/components/prism-core.min.js" integrity="sha512-9khQRAUBYEJDCDVP2yw3LRUQvjJ0Pjx0EShmaQjcHa6AXiOv6qHQu9lCAIR8O+/D8FtaCoJ2c0Tf9Xo7hYH01Q==" crossorigin="anonymous" referrerpolicy="no-referrer"></script><script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.28.0/plugins/autoloader/prism-autoloader.min.js" integrity="sha512-fTl/qcO1VgvKtOMApX2PdZzkziyr2stM65GYPLGuYMnuMm1z2JLJG6XVU7C/mR+E7xBUqCivykuhlzfqxXBXbg==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>`),
	})
	return markdown.ToHTML(markdownContent, parser, renderer)
}

// readFileOrPanic returns the contents of the file specified through the global variable inputFilename
func readFileOrPanic(filename string) []byte {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Can not read file " + filename + ": " + err.Error())
	}
	return fileContents
}

// writeFileOrPanic writes data to filename or panics if an error occurs.
func writeFileOrPanic(data []byte, filename string) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic("Can not write file " + filename + ": " + err.Error())
	}
}

// parseArgs parses command line arguments and saves them in the global variables
func parseArgs() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "LiTeX\nConvert markdown files to PDF\n\n")
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&inputFilename, "input", "", "Markdown file to convert to PDF")
	flag.StringVar(&cssFilename, "css", "", "CSS file to apply to markdown")
	flag.Parse()
	if len(inputFilename) == 0 || len(cssFilename) == 0 {
		panic("Please fill command line arguments. See --help for help")
	}
}
