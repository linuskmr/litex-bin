# LiTeX

LiTeX is an alternative to LaTeX. LiTeX creates a PDF from a markdown file (which may contain regular HTML as well) according to a CSS stylesheet.

Alternativly, simply use **pandoc**:

```bash
$ pandoc document.md -s -c style.css -o document.pdf --pdf-engine=weasyprint
# or, with auto-reload
$ ls document.md | entr -c pandoc document.md -s -c style.css -o document.pdf --pdf-engine=weasyprint
```

(likely requires `sudo apt install weasyprint`)

## Installation

Install go.

Install chromium-browser, for example with `sudo apt-get install chromium-browser`. It is used for rendering the PDF.

Install litex `go install github.com/linuskmr/litex@latest`

## Usage

Invoke the litex executable with a markdown file (`-input`) and a CSS stylesheet (`-css`). A HTML and PDF file will be generated.

```
LiTeX
Convert markdown files to PDF

Arguments:
  -css string
    	CSS file to apply to markdown
  -input string
    	Markdown file to convert to PDF
```

> Installation successfull, but command `litex` can not be found? If the executable is in `~/go/bin`, your GOPATH might not be configured.

## Example

An example letter. The result will be written to `example_letter.md.pdf`.

```
litex -input example_letter.md -css letter.css
```

An example document. The result will be written to `example_document.md.pdf`.

```
litex -input example_document.md -css document.css
```
