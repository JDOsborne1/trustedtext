package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func process_md_block(_md_block_body string) (string, error) { 
	opts := html.RendererOptions{
		Flags: html.FlagsNone,
	}
	
	renderer := html.NewRenderer(opts)

		byte_html := markdown.ToHTML([]byte(_md_block_body), nil, renderer)

	return string(byte_html), nil

}