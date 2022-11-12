package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

// process_md_block leverages the simple renderer from the gomarkdown package to allow
// you to sumbit blocks of markdown text as the stored version, but have the head block
// path be rendered as a proper HTML page
func process_md_block(_md_block_body string) (string, error) {
	opts := html.RendererOptions{
		Flags: html.FlagsNone,
	}

	renderer := html.NewRenderer(opts)

	byte_html := markdown.ToHTML([]byte(_md_block_body), nil, renderer)

	return string(byte_html), nil

}
