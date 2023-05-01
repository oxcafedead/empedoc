package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type MarkdownFiles struct {
	Content map[string]ast.Node
	Order   []string
}

var (
	indexMd      string
	inputDir     string
	outputDir    string
	htmlTemplate string
)

func main() {
	// Define default input and output directories
	defaultInputDir := "./docs"
	defaultOutputDir := "./docs_gen"

	// Define command line flags
	flag.StringVar(&indexMd, "i", "README.md", "A root Markdown file name")
	flag.StringVar(&inputDir, "dir", defaultInputDir, "Source directory with Markdown documentation")
	flag.StringVar(&inputDir, "d", defaultInputDir, "Source directory with Markdown documentation")
	flag.StringVar(&outputDir, "out", defaultOutputDir, "Output directory with HTML documentation")
	flag.StringVar(&outputDir, "o", defaultOutputDir, "Output directory with HTML documentation")
	flag.StringVar(&htmlTemplate, "tpl", "", "HTML template as a string")
	flag.StringVar(&htmlTemplate, "t", "", "HTML template as a string")

	// Parse command line flags
	flag.Parse()

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Read all the files in the input directory
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}
	readmePath := filepath.Join(".", indexMd)
	readmeInfo, err := os.Stat(readmePath)
	if err != nil {
		log.Fatal(err)
	}

	// pre-process MD files
	var mdFiles MarkdownFiles
	mdFiles.Content = make(map[string]ast.Node)
	mdFiles.Order = make([]string, 0)
	preProcessMd(&mdFiles, readmeInfo.Name())
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			preProcessMd(&mdFiles, file.Name())
		}
	}

	// Loop through all the files and directories and copy them to the output directory
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			copyFile(file.Name())
		} else {
			convertFile(&mdFiles, file.Name())
		}
	}

	// Convert the README.md file to HTML
	convertFile(&mdFiles, readmeInfo.Name())
}

func preProcessMd(mdFiles *MarkdownFiles, file string) {
	in := inputDir
	if file == indexMd {
		in = "."
	}
	inputBytes, err := os.ReadFile(in + "/" + file)
	if err != nil {
		log.Fatal(err)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	mdFiles.Content[file] = p.Parse(inputBytes)
	mdFiles.Order = append(mdFiles.Order, file)

	log.Printf("Preprocessed %s\n", file)
}

const HtmlTemplate = `<html><head><style>
h1, h2, h3, h4 {
	margin-top: -50px;
	padding-top: 70px;
}	
nav a {
	color: #777;
}
nav a:hover {
	color: #333;
}
nav ul {
	padding-left: 5px;
	padding-top: 5px;
	margin-left: 5px;
}
li a {
	text-decoration: none;
}
body {
	font-family: 'Helvetica', 'Arial';
	margin: 0 auto;
	color: rgb(29, 31, 32);
}
header {
	height: 50px;
	margin-left: 5px;
	padding-left: 5px;
}
nav {
	width: 200px;
	float: left;
	height: 100vh;
	line-height: 150%;
	padding-left: 0px;
	border-right: 1px solid #e8e8e8;
}
header + nav {
	height: calc(100vh - 70px);
}
nav li {
	list-style-type: none;
}
nav li.h1 {
	padding-left: 0px;
}
nav li.h2 {
	padding-left: 5px;
}
nav li.h3 {
	padding-left: 10px;
}
nav li.h4 {
	padding-left: 15px;
}
nav li.h5,li.h6,li.h7,li.h8,li.h9,li.h10 {
	padding-left: 20px;
}
main {
	margin-left: 250px;
	height: 100vh;
	overflow-y: auto;
	padding-right: 50px;
	line-height: 1.5;
}
header + nav + main {
	height: calc(100vh - 70px);
}
table {
	border-collapse: collapse;
}
th, td {
	padding: 5px;
	text-align: left;
}
tr {
	border: 1px solid #d0d0d0;
}
@media (max-width: 768px) {
	nav {
		display: none;
	}
	main {
		margin: 5px;
	}
}
code {
	background: rgba(29, 31, 32, .05);
	display: inline-block;
	padding: 5px;
	border-radius: 5px;
	padding: 2px;
	margin: -2px;
}
blockquote {
    background: #e5e6fc;
    padding: 5px 10px;
    margin: 0px;
    border-left: 5px solid #d3d4f5;
}
blockquote p {
    padding: 0px;
    margin: 0px;
}
header ul {
	display: flex;
	list-style: none;
	padding: 0;
	padding-top: 10px;
	align-items: center;
}
header li {
	margin-right: 30px;
	font-size: 120%;
}
header li:last-child {
	margin-right: 0;
}
header li a {
	display: block;
	padding: 3px 6px;
	text-decoration: none;
	background-color: #dedbdb;
	color: #2b2b2b;
	border-radius: 5px;
}
header li a:hover {
	background-color: #cdc9c9;
}
.para-link {
	text-decoration: none;
	color: #d2d2d2;
	font-size: smaller;
	position: relative;
	left: 10px;
	background: none;
	border: none;
}	
main a {
	color: #1d1f20;
	background-color: #3333331a;
	text-decoration: underline #b5b5b594;
}
</style>
</head>
<body>
	{{header}}
	<nav>
		{{nav}}
	</nav>
	<main>
		{{main}}
	</main>
</body>`

func convertFile(mdFiles *MarkdownFiles, file string) {
	in := inputDir
	if file == indexMd {
		in = "."
	}

	fileOutputDir := outputDir + "/" + in
	if err := os.MkdirAll(fileOutputDir, 0755); err != nil {
		log.Fatal(err)
	}
	outputPath := filepath.Join(fileOutputDir, file+".html")
	if file == indexMd {
		outputPath = filepath.Join(fileOutputDir, "index.html")
	}
	singleReadme := len(mdFiles.Content) == 1

	outputBytes := []byte(``)
	headers := ""
	if !singleReadme {
		headers = "<header>" + getHeaderWithLinks(file, mdFiles) + "</header>"
	}
	htmlWithParts := htmlTemplate
	if len(htmlWithParts) == 0 {
		htmlWithParts = HtmlTemplate
	}
	htmlWithParts = strings.Replace(htmlWithParts, "{{header}}", headers, 1)
	htmlWithParts = strings.Replace(htmlWithParts, "{{nav}}", getTableOfContents(mdFiles.Content[file]), 1)
	htmlWithParts = strings.Replace(htmlWithParts, "{{main}}", string(mdToHTML(mdFiles.Content[file])), 1)
	outputBytes = append(outputBytes, []byte(htmlWithParts)...)

	if err := os.WriteFile(outputPath, outputBytes, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("Converted %s to %s\n", file, outputPath)
}

func getHeaderWithLinks(currentFile string, mdFiles *MarkdownFiles) string {
	var buf bytes.Buffer

	// Start the list of links
	buf.WriteString("<ul>")

	// Create a map of filenames to their corresponding HTML links

	// Add a link to each file in the links map
	for _, filename := range mdFiles.Order {
		fileMdContent := mdFiles.Content[filename]
		var heading string
		fileMdContent.GetChildren()
		if filename == indexMd {
			heading = "Home"
		} else {
			heading = findHeadingTitle(fileMdContent, heading)
		}
		if filename != currentFile {
			linkPath := filename
			if filename != indexMd && currentFile == indexMd {
				linkPath = "docs/" + linkPath
			}
			if filename == indexMd {
				linkPath = "../index"
			}
			buf.WriteString("<li><a href=\"" + linkPath + ".html\">" + heading + "</a></li>")
		} else {
			buf.WriteString("<li current>" + heading + "</li>")
		}
	}

	buf.WriteString("</ul>")

	return buf.String()
}

func findHeadingTitle(fileMdContent ast.Node, heading string) string {
	for _, child := range fileMdContent.GetChildren() {
		foundHeading := false
		if h, ok := child.(*ast.Heading); ok {
			heading, foundHeading = findFirstChildText(h)
		}
		if foundHeading {
			break
		}
	}
	return heading
}

func findFirstChildText(h *ast.Heading) (string, bool) {
	var heading string
	var foundHeading bool
	for _, innerHeaderElement := range h.Container.GetChildren() {
		if t, ok := innerHeaderElement.(*ast.Text); ok {
			heading = string(t.Literal)
			if len(heading) > 0 {
				foundHeading = true
				break
			}
		}
	}
	return heading, foundHeading
}

type TableOfContentsNodeVisitorFunc func(node ast.Node, entering bool) ast.WalkStatus

func (f TableOfContentsNodeVisitorFunc) Visit(node ast.Node, entering bool) ast.WalkStatus {
	return f(node, entering)
}

type TOCHeader struct {
	Title string
	Id    string
	Level int
}

func getTableOfContents(mdFile ast.Node) string {
	// Find all the headings in the markdown file
	var headings []TOCHeader
	visitor := TableOfContentsNodeVisitorFunc(func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}
		if heading, ok := node.(*ast.Heading); ok {
			var h TOCHeader
			title, _ := findFirstChildText(heading)
			h.Title = title
			h.Id = heading.HeadingID
			h.Level = heading.Level
			headings = append(headings, h)
		}
		return ast.GoToNext
	})
	ast.Walk(mdFile, visitor)

	// Create an HTML string representing the table of contents
	var tocBuffer bytes.Buffer
	if len(headings) > 0 {
		tocBuffer.WriteString("<ul>")
		for _, heading := range headings {
			tocBuffer.WriteString("<li class=\"h" + strconv.Itoa(heading.Level) + "\">")
			tocBuffer.WriteString("<a href=\"#")
			tocBuffer.WriteString(heading.Id)
			tocBuffer.WriteString("\">")
			tocBuffer.WriteString(heading.Title)
			tocBuffer.WriteString("</a>")
			tocBuffer.WriteString("</li>")
		}
		tocBuffer.WriteString("</ul>")
	}

	return tocBuffer.String()
}

func copyFile(file string) {
	inputPath := filepath.Join(inputDir, file)
	// Create the output directory if it doesn't exist
	fileOutputDir := outputDir + "/" + inputDir
	if err := os.MkdirAll(fileOutputDir, 0755); err != nil {
		log.Fatal(err)
	}
	outputPath := filepath.Join(fileOutputDir, file)

	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(outputPath, inputBytes, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("Copied %s to %s\n", inputPath, outputPath)
}

func mdToHTML(doc ast.Node) []byte {
	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags, RenderNodeHook: func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
		switch n := node.(type) {
		case *ast.Link:
			if entering && bytes.Contains(n.Destination, []byte(indexMd)) {
				n.Destination = bytes.Replace(n.Destination, []byte(indexMd), []byte("index.html"), 1)
			} else if entering && bytes.Contains(n.Destination, []byte(".md")) && !bytes.Contains(n.Destination, []byte(".md.html")) {
				n.Destination = bytes.ReplaceAll(n.Destination, []byte(".md"), []byte(".md.html"))
			}
		case *ast.Heading:
			if !entering {
				w.Write([]byte("<a class=\"para-link\" href=\"#" + n.HeadingID + "\">#</a>"))
			}
		}
		return ast.GoToNext, false
	}}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
