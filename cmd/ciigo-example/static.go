// Code generated by github.com/shuLhan/share/lib/memfs DO NOT EDIT.

package main

import (
	"github.com/shuLhan/share/lib/memfs"
)

func generate_() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "./_example",
		Path:            "/",
		ContentType:     "",
		ContentEncoding: "",
	}
	node.SetMode(2147484141)
	node.SetName("/")
	node.SetSize(0)
	node.AddChild(_getNode("/LICENSE", generate__LICENSE))
	node.AddChild(_getNode("/custom.css", generate__custom_css))
	node.AddChild(_getNode("/favicon.ico", generate__favicon_ico))
	node.AddChild(_getNode("/html.tmpl", generate__html_tmpl))
	node.AddChild(_getNode("/index.css", generate__index_css))
	node.AddChild(_getNode("/index.html", generate__index_html))
	node.AddChild(_getNode("/sub", generate__sub))
	node.AddChild(_getNode("_example", generate__example))
	return node
}

func generate__LICENSE() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/LICENSE",
		Path:            "/LICENSE",
		ContentType:     "text/plain; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x94\x95\x5F\x6F\xDB\x36\x17\xC6\xAF\x5F\x7E\x8A\x07\xBD\x6A\x01\xC5\xFD\xF3\x02\x03\x96\xDE\x8C\x96\xE8\x98\x98\x2C\x6A\x24\x15\xD7\x97\x8A\x45\xD7\xC4\x2C\xD1\xA0\xE8\x14\xFD\xF6\x03\x29\xA7\x71\x9D\x6C\xC3\x08\x04\x39\xB6\xCE\xEF\x9C\xE7\x3C\x3C\x86\x72\x77\xFC\xEE\xED\xD7\x7D\xC0\xA7\x0F\x1F\x7F\xB9\xF9\xF4\xE1\xE3\xAF\x19\x56\x33\xA8\xFD\xE9\xB0\x6F\x07\xBC\xED\xC7\xDF\xFE\xB4\x87\xF6\xC1\x86\x99\x1D\x76\xEE\xDD\x8C\xD0\xC3\x01\x09\x19\xE1\xCD\x68\xFC\xA3\xE9\x66\x84\x48\xD3\xD9\x31\x78\xFB\x70\x0A\xD6\x0D\x68\x87\x0E\xA7\xD1\xC0\x0E\x18\xDD\xC9\x6F\x4D\xFA\xE6\xC1\x0E\xAD\xFF\x8E\x9D\xF3\xFD\x98\xE1\x9B\x0D\x7B\x38\x9F\xFE\xBB\x53\x20\xBD\xEB\xEC\xCE\x6E\xDB\x58\x20\x43\xEB\x0D\x8E\xC6\xF7\x36\x04\xD3\xE1\xE8\xDD\xA3\xED\x4C\x87\xB0\x6F\x03\xC2\xDE\x60\xE7\x0E\x07\xF7\xCD\x0E\x5F\xB1\x75\x43\x67\x23\x34\x26\xA8\x37\xE1\x96\x90\x8F\x33\xFC\x2C\x69\x84\xDB\x3D\x69\xD9\xBA\xCE\xA0\x3F\x8D\x01\xDE\x84\xD6\x0E\xA9\x60\xFB\xE0\x1E\xE3\xA3\x27\x43\x06\x17\xEC\xD6\x64\x08\x7B\x3B\x12\x00\x07\x3B\x86\x58\xE3\xB2\xDD\xD0\x5D\x69\xE9\xEC\xB8\x3D\xB4\xB6\x37\x7E\x46\xC8\xA7\x97\x1A\xEC\x70\x69\xC2\x93\x86\xA3\x77\xDD\x69\x6B\xFE\x49\x46\x54\x10\x95\xFC\x57\x19\x38\x4F\xD7\xB9\xED\xA9\x37\x43\x48\xEE\xC6\x62\xED\xD0\xBD\x77\x1E\x2E\xEC\x8D\x47\xDF\x06\xE3\x6D\x7B\x18\x9F\x8D\x4E\xB7\x93\xC8\x8B\x01\x66\x84\xFC\x7F\x86\xCA\xD8\x44\xC5\xA7\x43\xDB\x9B\x49\xCE\x93\xE0\xBD\x3B\x74\xC6\x63\x70\xCF\x09\xC9\x7B\x1B\xC6\x28\x7A\xAA\xE5\xFC\x88\xBE\xFD\x1E\x85\x3C\x98\xB8\x29\x1D\x82\x83\x19\x3A\xE7\x47\x13\x97\xE2\xE8\x5D\xEF\x82\xC1\x64\x4D\x18\xD1\x19\x6F\x1F\x4D\x87\x9D\x77\xFD\xE4\xC4\xE8\x76\xE1\x5B\xBC\xF1\xA7\x05\x02\x30\x1E\xCD\x36\x2E\x11\x8E\xDE\xC6\xD5\xF2\x71\x7D\x86\x69\x91\xC6\x71\x9A\x40\x2F\xB9\x82\x12\x0B\xBD\xA6\x92\x81\x2B\xD4\x52\xDC\xF3\x82\x15\x98\x6F\xA0\x97\x0C\xB9\xA8\x37\x92\xDF\x2D\x35\x96\xA2\x2C\x98\x54\xA0\x55\x81\x5C\x54\x5A\xF2\x79\xA3\x85\x54\x78\x43\x15\xB8\x7A\x43\xE2\x03\x5A\x6D\xC0\xBE\xD4\x92\x29\x05\x21\xC1\x57\x75\xC9\x59\x81\x35\x95\x92\x56\x9A\x33\x95\x81\x57\x79\xD9\x14\xBC\xBA\xCB\x30\x6F\x34\x2A\xA1\x51\xF2\x15\xD7\xAC\x80\x16\x59\x6C\x4A\x5E\x62\x10\x0B\xAC\x98\xCC\x97\xB4\xD2\x74\xCE\x4B\xAE\x37\x49\xC8\x82\xEB\x2A\xF6\x5A\x08\x09\x8A\x9A\x4A\xCD\xF3\xA6\xA4\x12\x75\x23\x6B\xA1\x18\xA8\x64\xA4\xE0\x2A\x2F\x29\x5F\xB1\x62\x06\x5E\xA1\x12\x60\xF7\xAC\xD2\x50\x4B\x5A\x96\x7F\x33\xA5\x90\x3F\x0F\x39\x67\x28\x39\x9D\x97\x8C\xA4\x56\xD5\x06\x05\x97\x2C\xD7\x71\x9E\xE7\x28\xE7\x05\xAB\x34\x2D\x33\xA8\x9A\xE5\x3C\x06\xEC\x0B\x5B\xD5\x25\x95\x9B\xEC\x5C\x53\xB1\x3F\x1A\x56\x69\x4E\x4B\x52\xD0\x15\xBD\x63\x0A\x6F\xFF\xC5\x93\x5A\x8A\xBC\x91\x6C\x15\x45\x8B\x05\x54\x33\x57\x9A\xEB\x46\x33\xDC\x09\x51\x44\xB1\x44\x31\x79\xCF\x73\xA6\x3E\xA3\x14\x2A\xD9\xD5\x28\x96\xA1\xA0\x9A\xA6\xC6\xB5\x14\x0B\xAE\xD5\xE7\x18\xCF\x1B\xC5\x93\x6B\xBC\xD2\x4C\xCA\xA6\xD6\x5C\x54\xEF\xB0\x14\x6B\x76\xCF\x24\xC9\x69\xA3\x58\x91\xEC\x15\x55\x1A\x55\x2F\x99\x90\x9B\x58\x34\x7A\x90\xDC\xCF\xB0\x5E\x32\xBD\x64\x32\x3A\x9A\x9C\xA2\xD1\x02\xA5\x25\xCF\xF5\x45\x1A\x11\x12\x5A\x48\x7D\x31\x23\x2A\x76\x57\xF2\x3B\x56\xE5\x2C\xAA\x11\xB1\xCA\x9A\x2B\xF6\x0E\x54\x72\x15\x13\xF8\xD4\x76\x4D\x37\x10\x4D\x1A\x39\x5E\x52\xA3\x18\x49\xE1\xC5\xCA\x66\xE9\x2A\xC1\x17\xA0\xC5\x3D\x8F\xB2\xCF\xC9\xB5\x50\x8A\x9F\x17\x25\x59\x96\x2F\x31\xD9\x3D\x23\xF1\xC7\x91\xCE\xCD\xCD\xCD\x73\xF0\xE3\xC3\x6B\xD1\x39\x81\x90\xFF\x69\x0D\xC4\x3F\x70\x0E\xCC\xE7\xF8\x71\x28\xA5\x34\x05\x65\x3A\x53\xC2\xEF\xE7\x93\xB8\x84\x5D\x73\x11\x4A\x60\x44\xF0\xC4\x21\xA2\x11\x9A\x98\x2B\x68\x6A\xF4\x83\x7A\x01\xBD\xDA\x89\x9E\xCF\x95\xC2\x67\xE8\xB5\xB1\x9E\x3B\xBD\x2E\x0F\x13\x35\x41\xD3\xB9\x92\x77\xDD\x89\xAC\xCD\xC3\x68\x83\xB9\xC5\x3E\x84\xE3\xED\xFB\xF7\x97\xAF\x54\x92\xBB\x21\xB4\xDB\x70\x8B\xAB\x77\x2D\xF9\x2B\x00\x00\xFF\xFF\x6A\xEC\xBB\x57\x9D\x07\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("LICENSE")
	node.SetSize(957)
	return node
}

func generate__custom_css() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/custom.css",
		Path:            "/custom.css",
		ContentType:     "text/css; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xCA\x30\x54\xA8\xE6\xE2\x4C\xCE\xCF\xC9\x2F\xB2\x52\x28\xCE\x4C\xCD\xCB\x4B\xB4\xE6\xAA\xE5\x02\x04\x00\x00\xFF\xFF\x1C\x35\x0D\x2C\x17\x00\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("custom.css")
	node.SetSize(47)
	return node
}

func generate__favicon_ico() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/favicon.ico",
		Path:            "/favicon.ico",
		ContentType:     "image/x-icon",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xEC\x98\x0D\x4C\x55\x65\x18\xC7\x1F\xE0\x1A\x04\x85\x18\x9A\x5D\x4B\xEF\x69\xA6\x32\xE7\xDC\x2C\x95\x73\x67\x26\xAE\x9A\xB9\xE4\x5C\xF5\x9E\x13\x73\xE8\x6E\xE9\x0C\xD4\xCC\x12\xB8\x88\x9A\x62\x69\xCD\xCD\x4F\x30\x44\xCA\x89\xB9\x34\x5D\x4D\x14\x9C\x6B\x6A\xDD\x0B\x9B\x29\x61\x1A\xF3\x13\xF9\xF4\x8A\x22\x2E\x94\xE0\xF2\xE1\x85\xF7\xDF\xDE\x73\x3F\x52\xEE\x65\x92\x0B\x5B\xDB\xFD\x6F\xCF\x38\xCF\x39\xCF\xFB\xFC\xDE\x8F\xE7\x39\x67\x5C\xA2\x00\x0A\x24\x41\xE0\x7F\x05\xFA\x2E\x82\x68\x38\x11\x45\x44\x70\x3F\x84\x12\xFB\x10\xFD\x16\x41\x14\x45\x44\x02\x11\xC5\x90\x33\xEE\x61\x92\x25\x11\xDC\x58\x8D\xB9\x57\xCD\xCD\xF1\xF3\xFD\x7C\x3F\xDF\xCF\xF7\xF3\xFF\x9F\xFC\x32\xEB\x7C\x9C\xCC\x9B\xAD\x5A\xD5\xC9\xC4\xC7\xCE\xBF\x51\x73\x06\x8E\x0E\xA8\x76\xAD\xEC\xF8\x63\xE3\xB7\x95\x27\xE1\x46\xC9\x22\x34\xDE\xAD\xF7\xF0\x6F\xD5\x5E\xC2\xF5\xE2\x85\xEA\xB3\xDE\xE4\x3B\x2A\x53\xB0\x3F\xE7\x5D\x6C\xCB\x5C\x89\x6D\xBB\x8E\x22\x67\xAF\x15\x5F\xED\x2B\x44\xE6\xD7\x79\xC8\xCA\x5C\x8E\xBC\x5D\xA6\x5E\xE5\xB7\x95\xA7\xC1\xBC\xD1\x82\xE4\x0D\x45\x3E\x6D\xE5\xE6\xA3\xFF\x2E\xFF\xDA\x0A\xB0\x86\x43\x40\xEB\x15\xA0\xDD\x06\xD8\xCF\x22\x77\xFF\x11\x95\x65\x9C\xBB\x0E\xAF\xE8\xA7\x62\xDC\xAB\xD3\x30\x6B\xE1\x16\xF5\x5E\x5E\xC1\x61\xE0\xCF\x22\xA0\xE9\x34\x58\xC3\x41\x30\xDB\xEA\x47\xE7\x5F\x5F\x0B\xDC\xAB\x43\x57\x31\x00\x25\xC5\x27\xA0\x13\x86\x7A\x6C\x44\xD4\x48\x54\x54\x5E\xF2\x8A\x45\x67\x33\x58\x5D\xF6\x23\xF1\xD1\x56\xE9\x9D\xEF\x3E\xE5\x6E\x37\x63\xDA\x94\xB1\x98\xF1\xF6\x38\xE4\xFF\x90\xD3\x7D\x60\x67\x2B\xD8\xF5\x75\x3D\xE1\x37\xF1\xFB\x6D\xE5\xC9\x60\x75\x5F\x7A\xA5\x69\x69\x69\x81\xE9\xBD\x79\x88\x1A\x39\x1A\x2B\x3E\x59\x0D\xE6\xB8\x03\x56\x93\x0A\x76\x73\xAB\x57\xEC\xF9\xF3\x17\x10\x33\xF9\x0D\x8C\x79\x79\x3C\xF2\x0B\x8E\x00\x8D\x56\x57\xFD\x24\xBB\xF9\x4D\x3E\xF8\x65\xFC\xD9\xED\x73\x4B\xC0\x1A\xF2\xBD\x72\x6E\xCF\xCE\x41\xBF\xC8\x81\x1E\xB3\x16\x16\x81\xD5\x7E\xEE\xAC\x8F\x2E\x9A\x3E\x43\xF6\xC4\xE9\x5E\x1C\x86\x0E\x7B\x85\xCA\xAF\x3F\xB7\x44\xE5\x1B\x25\xF1\x4A\x57\xBE\x62\x88\x3E\xC8\x9F\x15\x1F\x99\x07\x76\xA7\xC0\x2B\x67\x56\x56\xF6\x03\x7C\x8B\xC5\x0A\x56\xFB\x85\x4F\xBE\x61\xBA\xD1\x13\x37\x44\x78\x09\x8E\xE6\x72\x95\x7F\x3A\x7F\xAE\x7B\xFD\x87\x1E\x5C\xBB\x5E\x89\x9B\xA9\x6F\xDE\xB2\x46\xC2\xB5\xD3\x1F\x80\xD5\x6D\xF7\xCA\x69\xB7\xDB\x11\x3F\xDB\x84\x61\xC3\x47\xC2\x9C\x9A\x06\xE6\x68\x74\xED\x7F\xA6\x57\xEC\xEF\xA5\xA5\x98\x30\x31\x06\xA3\x46\x8F\x41\x5E\x1E\xEF\x89\x42\x95\x7F\xD9\x92\x88\x65\x8B\xDF\xE4\xFC\x36\xD9\xA0\x4F\x50\xD9\xD3\xF4\xAF\xC7\x4D\xD7\xB3\xD2\x63\xEF\xE3\x8F\xD2\x8F\xD0\x51\x65\x46\x67\x35\xAF\xBF\x6A\x1F\xC5\xF4\xB7\xF8\x19\x79\x6A\x95\xF7\x67\x77\xEA\x6C\x73\x9E\x93\x2B\xB6\xF5\x6A\x12\xBE\xD9\x2A\xAB\xFB\xA0\x18\xC4\x38\x45\x12\x8F\x7D\x6A\x9E\x8A\xA6\xCB\x49\xD8\xB9\xD1\x88\x9D\x9B\x66\xA2\xA5\x2C\x09\x9D\x36\xDE\x7F\xF5\xBE\x73\x36\xFF\xEA\x5C\xBB\xBB\xAF\x6C\x6B\x7C\xCF\xB7\xB3\x05\xEC\x56\x0E\x58\xB5\x19\x17\x7F\x4A\x40\xF6\xFA\x19\x38\xB4\x6B\x16\xEC\x57\x92\x9C\x7C\x49\x3C\x27\x1B\xC4\xEA\x45\x73\x63\x70\xAF\x32\x19\xED\x15\x29\x9E\x73\xCA\xF8\x4C\x82\xFD\x6A\x9A\xB3\x16\x5A\x2B\x80\xF6\x5A\xC0\x5E\x0A\x56\x9F\xDB\xCD\xBB\x2D\x15\xEC\xF6\x1E\xA0\xE9\x17\xA0\xB9\xC4\xB9\x3F\xB6\x74\x95\xB5\x7B\xAB\x8C\x5B\x67\x3F\x54\xD7\xCE\xF7\xF8\x6A\xD1\x02\x77\x1D\xD8\x15\x29\xDA\xCA\xAF\x57\x2D\x9D\x82\xC3\xB9\xF1\xD8\xB0\x2A\xD6\xD3\xA3\x89\xA6\x49\xB0\x7C\x6F\x42\x3B\xEF\xC9\x7F\xF8\x9D\xE7\xBD\xC6\xC7\x26\x98\x26\xA9\xB9\xE6\x28\x13\xB0\x27\x43\xC1\x81\x1D\x71\x98\x3F\xE7\x35\x67\x1F\xC4\x8A\x17\xF9\xF9\xCB\x92\xD8\xE1\x66\x7A\xCC\x20\xDA\xDC\xD7\xA6\xB8\x09\xC8\x58\x6B\x40\xC1\xEE\x78\x9C\x3F\x91\x80\x9B\x67\x16\xA3\xF1\xE2\xC7\xEA\x9E\x71\xE3\xD7\xFC\x1E\x7F\xC6\x63\x78\x2C\x1F\xE3\x1E\xAF\x48\xFA\x72\xAF\xFC\x92\xC8\xF8\xF9\xAB\x35\x68\x18\x3F\x51\x36\xE8\x7F\x94\x25\xB1\x52\x91\xC4\x62\x45\x12\x17\x10\x51\xA0\x31\x56\x3F\xCB\x28\x89\x55\x3E\xC6\xF6\xC8\x8C\x92\x58\x6E\x8C\x15\xDF\xE1\xFF\x02\xCB\x92\x18\x2F\x4B\x62\x91\x2B\x9F\x45\x8E\x15\xA7\x76\x7D\x07\xF8\x92\xA2\x28\x41\x33\xA5\xE8\xC9\xB2\x41\xCC\x90\x0D\xE2\x71\x3E\x47\x59\x12\x1B\x5C\xE6\x66\x35\xBA\x7C\x9B\x33\x26\x7A\x33\x5F\x53\x7A\x3A\x05\xF6\x84\xF1\x5F\xCA\x42\x14\xD4\x5B\x96\x4E\x14\xE0\x8B\xC9\xEF\x83\x48\x6D\xD1\x28\x22\x8A\x70\xFD\x56\x11\x40\x21\xDD\xCE\xB3\x5F\xE4\x40\xD2\x09\x43\x69\xF4\xA8\x11\x14\x27\x8D\xE1\xDF\x0B\x4A\x99\x1F\x4D\xCB\x12\xC7\xD2\x8A\x24\x89\x0E\x64\xBC\x45\xDF\x66\xC8\x94\xBC\xA1\x88\x4E\xED\x93\x68\xD3\x8E\xBD\xB4\x77\xCF\x7A\x3A\x95\xBF\x94\x0A\x0F\x2F\xA7\xB3\xC7\xD2\xA8\xF8\xE7\xCD\x54\x62\xCD\xA2\xA6\x0B\xC9\x84\x8A\x54\x42\x15\xB7\x65\x84\x9A\x14\x42\x75\x1A\xB1\x1A\x33\xDD\xAC\x29\x25\x7B\x8B\x83\x5A\xEF\x81\x1C\x1D\x20\x07\x03\x9F\x68\x6F\x1E\x81\x5F\x7E\xF9\xF5\x10\x09\x1A\xED\x7D\xD2\x08\xBD\xE4\x47\xF4\xD5\x0E\xE8\xDF\xDF\xED\x3F\x17\x3A\x78\x70\xA8\xD6\xED\x0F\xD0\x6A\x9F\x0E\x0D\x7D\x4A\x3B\xF0\x59\xA7\x1F\xA9\xD3\x85\x07\x04\x84\xEB\x74\x91\x4E\x7F\xD0\x10\x0A\xD4\x6A\x07\x0D\xA6\x40\xAD\x46\xD0\x3C\x11\x36\x48\xA7\xFA\x3A\x0A\x0C\xEB\xA3\xD1\x84\x85\x84\x3F\xFF\x42\xB8\x56\x1B\xAE\xD3\x85\x87\x84\x69\x04\x4D\x70\x48\x58\x18\x47\x3E\xD3\xEF\xC9\x60\x8D\xC0\x15\x1C\x1C\xA4\xD1\x04\x05\x07\x0B\x3E\x94\x1E\x40\xD4\x53\x73\x8B\x5F\xDF\xED\xCB\x5F\xB4\x44\x7F\x05\x00\x00\xFF\xFF\xF7\x87\x31\xC7\x36\x16\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("favicon.ico")
	node.SetSize(1316)
	return node
}

func generate__html_tmpl() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/html.tmpl",
		Path:            "/html.tmpl",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x8C\x54\x4F\x6F\xDC\x2E\x10\x3D\x6F\x3E\x05\xE1\x77\xFD\x61\x14\x55\x55\xAB\x16\xAF\x94\x7F\xB7\x4A\x8D\x94\x5C\x7A\xAA\x30\x9E\x5D\x50\x31\x38\x30\xDE\x64\x85\xFC\xDD\x2B\x8C\xBD\x75\xD2\x48\xCD\xC9\xC0\x9B\xF7\x78\x6F\x3C\xB6\x38\xBF\xF9\x7E\xFD\xF0\xE3\xEE\x96\x68\xEC\xEC\xF6\x4C\x94\xC7\x46\x68\x90\xED\xF6\x6C\xB3\x11\x1D\xA0\x24\x1A\xB1\x67\xF0\x38\x98\x43\x4D\xAF\xBD\x43\x70\xC8\x1E\x8E\x3D\x50\xA2\xCA\xAE\xA6\x08\xCF\xC8\x33\xFB\x2B\x51\x5A\x86\x08\x58\x0F\xB8\x63\x9F\x29\xE1\x7F\x74\x9C\xEC\xA0\xA6\x07\x03\x4F\xBD\x0F\xB8\x62\x3F\x99\x16\x75\xDD\xC2\xC1\x28\x60\xD3\xE6\x7F\x62\x9C\x41\x23\x2D\x8B\x4A\x5A\xA8\x2F\xDE\x10\x42\x0D\x1D\x30\xE5\xAD\x0F\x2B\xAD\xFF\x3E\x7C\xFA\x78\x7B\x79\x35\xD5\x67\x02\x1A\xB4\xB0\x4D\xA9\x7A\xC8\x8B\x71\x14\xBC\x9C\x64\xCC\x1A\xF7\x8B\x04\xB0\x35\x8D\x78\xB4\x10\x35\x00\x52\xA2\x03\xEC\x6A\xCA\x8D\x6B\xE1\xB9\x52\x31\xCE\x57\xA7\xC4\x48\x90\x6E\x0F\xA4\xBA\x9F\xAA\xC7\xF1\x5F\x1A\x29\x55\xE3\xB8\xA2\x83\x6B\x33\x49\xF0\xB9\xBF\xA2\xF1\xED\x71\x72\xD2\x9A\x03\x51\x56\xC6\x58\x53\xF4\x7D\x23\x03\xCD\xC7\x2F\xCE\x73\x42\x69\x1C\xCC\xD0\x6B\x0E\xCB\x9A\xC6\xED\x67\x74\x23\xE4\x12\x84\x6E\x95\x31\x7B\x2F\xB8\x9C\x89\xBC\x35\x87\xBF\x35\x3A\x70\xC3\x89\xBC\xF3\xA1\x5B\x00\x83\xD0\x51\x22\x15\x1A\xEF\x6A\xCA\x7F\x1A\x87\x10\x9C\xB4\x3C\x82\x0C\x4A\x2F\x9C\x8D\x30\xAE\x1F\x90\xE0\xB1\x87\x32\x10\x74\x7E\x51\x8F\x94\xF4\x56\x2A\xD0\xDE\xB6\x10\x6A\x7A\x5F\x78\xA5\x2F\x93\xA1\x7C\xDD\xFB\xCD\x9D\x92\xC5\xA1\xA1\xDB\xFB\xA1\x79\x23\xDA\x69\x35\x2F\x5E\x35\xB9\x97\x7B\x78\x47\x8B\xF5\xC5\x8B\xC9\xD1\x17\x6F\x39\x43\xB9\x38\x4B\xA9\xBA\x1C\x50\xFB\x30\x8D\x46\x2E\x6C\xC2\x29\x65\x4A\xD5\x8D\x44\x98\xA1\x55\xCE\x94\xAA\x2B\xDF\x1E\x0B\xB0\x4A\x70\xCE\x18\xA9\x4E\x96\x08\x63\xAB\x38\x0B\x9C\x73\x4C\xC8\xAB\x7C\x3B\xEF\x71\x89\x71\xE7\x9F\x20\x40\x4B\x9A\x23\x11\x72\xBA\xB1\xB4\x2F\x7F\xD6\xF1\x0B\xE7\x7B\x83\x7A\x68\x2A\xE5\x3B\x1E\xF5\xF0\x4D\x4B\xC7\xA7\x81\xA1\xB9\xB6\x58\x9C\xF6\xC5\x9E\x5C\xBB\x10\xBC\x4C\xB0\xE0\xE5\xCF\xF1\x3B\x00\x00\xFF\xFF\x81\xA3\x1E\xB5\x51\x04\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("html.tmpl")
	node.SetSize(520)
	return node
}

func generate__index_css() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/index.css",
		Path:            "/index.css",
		ContentType:     "text/css; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x94\x55\x5D\x8F\xAB\x36\x14\x7C\x8E\x7F\xC5\x91\x56\x7D\x89\x80\x05\xB2\x24\x2D\x48\x2B\x55\x7D\xEE\x8F\x30\xF8\x10\xAC\x35\x36\x32\xA6\xC9\xF6\x2A\xFF\xFD\xCA\x60\xBE\x02\xB9\xD2\x15\x0F\x89\xEC\xB1\xCF\x78\xE6\x78\x9C\x2B\xF6\x0D\x3F\xC8\xA1\xA6\xFA\xCA\x65\x0A\x61\x46\x0E\xA5\x92\xC6\x2F\x69\xCD\xC5\x77\x0A\x7F\x6B\x4E\x85\x07\x2D\x95\xAD\xDF\xA2\xE6\x65\x46\x0E\x39\x2D\xBE\xAE\x5A\x75\x92\xF9\x85\x12\x4A\xA7\xF0\x56\x96\x76\x42\x70\x89\x7E\x85\xFC\x5A\x99\x14\xA2\xE0\x94\x91\x83\xC1\xBB\xF1\xA9\xE0\x57\x99\x42\x81\xD2\xA0\xCE\xC8\x61\x5C\x15\xC7\x71\x46\x1E\xA4\xD1\x68\x39\xAC\xEA\xFE\x8B\x52\x28\x0F\x6A\x25\x55\xDB\xD0\x02\x47\x5E\x2D\xFF\x1F\x53\x08\x83\x3F\x2F\x89\xC6\x7A\x5B\xF3\x23\x23\x07\xF5\x1F\xEA\x52\xA8\x9B\x7F\x4F\x81\x76\x46\xAD\x28\xA7\xF0\x86\xA5\xFD\x32\x72\x68\x28\x63\x5C\x5E\xED\x7E\xE7\xD8\xED\x97\x2B\xCD\x50\xFB\x9A\x32\xDE\xB5\x76\xE6\x14\xB9\xA9\x07\xA1\x96\xE6\x48\xFE\x74\x49\x90\xE6\xE3\x11\x19\x16\x4A\x53\xC3\x95\x4C\x41\x2A\x89\x3D\x3C\xAD\x2C\x15\xBB\x68\x83\xE9\x24\x43\x6D\xB9\x5B\x20\x69\x3C\x22\xF8\xE0\xC3\xDD\xBF\x71\x66\xAA\x14\x92\x70\x20\x74\x53\x9A\xF9\x37\x4D\x9B\x14\x72\x8D\xF4\xCB\xB7\x03\xBD\x6C\x9E\x55\xCE\x23\x9D\xF0\x88\x12\x4B\x17\xA3\x60\xA2\x4C\xAA\xC8\x23\x55\xEC\x91\xEA\xE4\x91\xEA\x63\x07\x05\xE1\x8C\x5F\x28\x92\xED\x9C\xB4\x77\xE0\xE6\xB4\xCE\x95\xE8\x69\x54\xD1\x64\xDE\x60\x4E\x14\xEC\x7B\xE3\xC0\x41\xAF\x45\xDD\x19\x64\x4B\x39\x2F\x97\x4B\x0F\x88\xFB\x41\x81\x54\xA7\xA0\xED\xCA\xEC\x69\xF3\xC9\xA8\x95\xA5\x21\xE6\x65\xB2\xB6\x74\xBF\x41\xE2\xE4\xF9\x20\x52\xE9\x9A\x8A\x45\xDF\x4C\x5D\x33\x35\xD2\x9E\xFC\x55\x0C\x74\x3A\xF9\x46\x94\xD3\x46\x94\xF8\x17\x7C\x7E\xB3\xF2\x6B\x2B\x97\xDD\x3A\x40\x96\x1C\xA6\x9E\x88\xE1\x13\xDA\x86\x4A\xDB\x16\xEE\x6F\x8F\x15\x8A\x9A\x59\xF6\x29\x14\x20\x4E\x9A\x3B\x84\x73\x3A\x6C\xA4\x1B\x5D\x4C\xE2\xCB\x5F\x45\x6F\x24\x61\xAF\x5A\x92\xB1\x55\xE2\xF4\xDF\x72\x5A\x78\x0E\xB2\x7B\xE1\x1F\x84\xBC\x1F\x8F\x04\x8E\xF0\x4F\xD7\x1A\x55\x43\x21\x68\xDB\x62\x0B\xA5\xD2\xD0\xD0\x2B\xB6\x04\x8E\xEF\x84\x04\x46\x35\x39\xED\x2F\xDF\x7E\xA7\x8C\x16\x7C\x0C\xB6\xCC\x16\x54\x9C\x31\x94\x7D\xA5\x71\x13\xFB\xEB\x57\x48\x6D\x6B\x79\xF3\x68\x8D\xB2\xB3\x05\xA6\xA6\x8B\x82\x53\x74\xEA\x8D\x78\xD5\xB5\xBB\xF2\xBD\xA8\xB4\xB0\x44\x60\x69\x2C\x6E\x17\xB6\x8A\xA5\x3E\x53\x5F\x67\xD2\x96\xFC\xDA\xF4\xC7\x13\x80\xAE\xBC\x72\x31\x09\x61\x10\x6D\x13\x63\xCE\x50\xC7\xE5\x56\x71\x83\xCF\x37\x75\x0A\x93\x21\x68\xED\xBA\xD0\xED\xDA\x2A\xC1\xD9\x06\x31\x45\x71\xD2\xDC\xB7\xFC\x4A\xA5\xEB\x99\xA2\x6F\x75\x5A\x52\x79\x90\xC0\xF6\x84\x45\xB8\x58\x8D\xC2\xF0\x8F\x41\x88\x7E\xE2\x13\x82\x42\x49\x43\xB9\x44\x3D\x3B\xBB\x1C\x9D\xE2\xDB\xBD\x62\x83\x15\xEB\x82\xEE\xD6\xBA\x31\x3D\xB8\xEB\x06\x67\x81\x56\xB9\xBC\x28\x10\xD4\x68\xE6\x28\x69\xCD\xB7\xC0\x14\xB8\xA1\x82\x17\xD9\x7E\xA8\x07\xA5\x52\x66\x4B\x6D\xF3\xC0\x9E\xCF\xE7\x97\x0F\xE7\xB8\x6F\x1C\x24\xAE\x63\x1F\xE4\x67\x00\x00\x00\xFF\xFF\x2F\x17\x7F\x31\x0F\x08\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.css")
	node.SetSize(739)
	return node
}

func generate__index_html() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/index.html",
		Path:            "/index.html",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xBC\x5A\x6D\x6F\x1B\x37\xF2\x7F\xED\xFD\x14\xD3\xFD\x07\x41\x82\xBF\xA4\x95\x9D\x2B\x9A\xBA\x92\x0E\x8D\xE3\x26\xC1\xB9\x4D\x11\x3B\xB8\x2B\x8A\x83\x41\x91\xA3\x5D\xC2\x5C\x72\x43\x72\xA5\xA8\x45\xBF\xFB\x61\xC8\x5D\x69\x25\x4B\x89\x1F\x74\x97\xBC\xF0\x8A\x8F\x33\xBF\x99\xF9\x71\x38\xBB\xA3\x6F\x5E\xBF\x3F\xBB\xFA\xED\xD7\x73\x28\x7C\xA9\x26\xC9\x28\xFE\x39\x1A\x15\xC8\xC4\x24\x39\x3A\x1A\x95\xE8\x19\x14\xDE\x57\x7D\xFC\x54\xCB\xF9\x38\x3D\x33\xDA\xA3\xF6\xFD\xAB\x65\x85\x29\xF0\xF8\x6B\x9C\x7A\xFC\xEC\x33\x9A\xFD\x03\xF0\x82\x59\x87\x7E\x5C\xFB\x59\xFF\x65\x0A\xD9\x7A\x1D\xCD\x4A\x1C\xA7\x73\x89\x8B\xCA\x58\xDF\x99\xBD\x90\xC2\x17\x63\x81\x73\xC9\xB1\x1F\x7E\xF4\x40\x6A\xE9\x25\x53\x7D\xC7\x99\xC2\xF1\xF1\x8E\x85\x7C\x81\x25\xF6\xB9\x51\xC6\x76\xD6\xFA\xBF\x17\xDF\x7D\x7B\xFE\xE3\xAB\x30\x9E\x26\x78\xE9\x15\x4E\xFE\x89\x8A\x9B\x12\xC1\x1B\xE0\x52\xE6\x66\x94\xC5\x76\x1A\xA1\xA4\xBE\x01\x8B\x6A\x9C\x3A\xBF\x54\xE8\x0A\x44\x9F\x42\x61\x71\x36\x4E\x33\xA9\x05\x7E\x1E\x70\xE7\xA2\x00\xA3\xAC\x81\x66\x34\x35\x62\x19\xA6\x0B\x39\x07\xAE\x98\x73\xE3\xD4\x9B\x6A\xCA\x6C\x4A\xCD\x1B\xED\x24\x1C\x93\x1A\x9B\xAE\xED\x39\x7D\x5A\x53\xEA\xBC\xE9\x3D\x1A\xB1\x76\xF7\x74\xD2\x48\xCB\x9A\x89\x99\x90\xF3\xDB\x6B\x94\xA8\xEB\xD5\xE4\x99\xB1\x65\xDB\x21\x3D\x96\x29\x30\xEE\xA5\xD1\xE3\x34\xBB\x96\xDA\xA3\xD5\x4C\x65\x0E\x99\xE5\x45\x3B\xE7\x68\x24\x75\x55\x7B\xF0\xCB\x0A\xA3\x2D\xD3\x06\xE3\x4F\x29\x54\x8A\x71\x2C\x8C\x12\x68\xC7\xE9\x65\x9C\x17\x8D\x11\x04\xA2\xED\xEE\x2E\xDC\x4A\x33\x57\x4F\xD3\xC9\x65\x3D\xDD\xA1\xDA\xEA\xA9\x79\xD8\x02\xB9\x62\x39\xDE\x01\xE2\xE2\x78\x87\xD1\x8B\xE3\x5D\xF2\x79\xD6\xCA\x77\x59\xD4\xAA\x60\xBA\x11\x76\x6A\x57\x7A\x9E\x7C\x0B\x97\x58\x79\x2C\xA7\x68\xE1\x64\x78\xFC\xFD\x2D\x8D\x77\x85\x12\x28\xA6\xF3\x71\x8A\x3A\xA5\x86\xE0\x38\xD1\x81\xDB\x18\x49\x3F\x5E\xFD\xD4\x7F\x99\xB6\xCD\xDD\x40\xFB\x57\xFF\xE3\x8F\xFD\x33\x53\x56\xCC\xCB\xA9\xEA\xC6\xDA\xBB\xF3\x31\x8A\x80\xC1\xE3\xC3\x6A\x30\xDC\x5A\x26\x47\x8D\x96\xF9\x8D\x90\x0A\xE0\xD1\xB8\xAF\xC4\xD2\x28\xC4\xCF\x24\x49\x46\x59\xF3\xD4\x86\x4B\x88\x96\x16\x70\x66\xBD\xE4\x2A\xC8\x4F\x66\x90\x62\x9C\xD2\xA8\x60\xBA\xFD\x66\xEB\x9A\x4C\xA0\x67\x52\x39\x1A\x1F\x2D\xD0\xFE\x69\xD7\x6B\x44\xEF\x6E\x51\x59\x64\xE5\x74\xBD\x6D\xB3\x94\xC3\x10\x1C\x24\xDF\x56\x4F\xC5\x2C\xCB\x2D\xAB\x28\x4A\x46\xD5\x64\xE4\x2A\xA6\x57\x71\x55\x06\x27\x1C\xB1\xCD\x86\xC6\xB9\xFF\xAC\xAD\xEA\xE7\x46\x18\xFE\x57\x3A\x19\xC9\x32\x07\x67\xF9\x38\x25\xEB\xBA\xD3\x2C\x93\x65\x3E\x70\x85\x44\x25\xDC\x40\x9A\x6C\xCA\x44\x8E\x59\x6E\x06\x02\xE7\x7D\x8B\x33\xB4\xA8\x39\xF6\x87\xC3\xEF\xC4\xF7\xFC\xEF\xCA\xE4\x66\x9C\x9B\xA7\xF4\xF7\x8C\xA8\x6E\xBC\x28\xA4\xC7\xA7\x01\xE0\xF1\x4C\x31\xDF\x77\x9F\x6A\x66\x31\x05\xA6\xFC\x38\x7D\x63\x5E\x1B\x9E\x4E\x28\xAC\x46\x19\x89\x4C\x66\xB9\xAF\xE4\x9C\x59\xB1\x53\xF4\xDC\x58\x24\x27\xA3\x01\x03\x6E\xCA\x56\x78\xE9\x8B\x7A\x1A\x1A\x5C\x51\x5F\x14\x4C\x67\xD1\x67\x5A\x99\xE0\x43\x98\x06\x67\xCC\x8A\x0D\xE1\x46\x59\xB5\x69\xBD\x3D\xE0\x73\x23\xB0\xE5\xC2\xF0\x0C\xD2\x01\x03\x25\xA7\x96\xD9\x25\x30\x2D\x80\x41\x65\x4D\x6E\x59\x49\x8E\xB3\xB0\xD2\x23\x38\xCF\xBC\xE4\xB0\xC0\x29\x38\xB4\x73\xB4\xB0\x90\xBE\x00\x8A\x62\x21\x50\x24\x33\xA9\xD0\x41\xED\xA4\xCE\x93\x15\x35\xB5\xBA\x32\xC7\xA5\x14\x86\x7B\x63\x07\xC6\xE6\x99\x30\xDC\x65\x8B\x82\xF9\xBE\x74\xFD\xB6\x33\x4B\x27\xED\x63\x60\xB2\x92\xD9\x9B\xBA\x02\x22\x45\xE6\x07\x1B\xDA\xDD\xF6\xD4\x8E\x0B\x1E\x07\xDF\x3F\x09\x9E\x7A\x1D\xF4\xBC\x66\xEE\xBA\x51\xAF\x39\x05\x80\xB9\x56\xE1\x51\x56\x9C\x3C\xD0\x8F\x7F\x32\x16\x98\x86\xBA\x22\x98\x04\xF3\x08\xC2\xF0\xBA\x44\x4D\x58\x19\x0D\x66\x06\x85\x59\x50\x67\xED\x10\x7C\x81\x2B\x90\x1D\xC6\xDF\xB7\xA1\xAA\x6E\xF2\x41\x74\xE0\xFD\xAE\x30\x79\x63\xB6\x76\x22\x22\x27\xD0\x1E\x0D\x13\x57\xB2\x03\xD1\xD9\xC5\xBB\x47\xC0\xD3\x5D\x06\x38\xC5\x8D\xD1\x73\xB4\xBE\x07\x0D\x33\x62\x8F\xBC\x2D\x33\x36\xBA\x14\x30\x10\xD2\x22\x79\xC9\x12\x7C\xC1\x3C\x34\x27\x91\x5B\xF9\x02\x39\x59\x8F\x96\x7C\x7B\xF5\xF3\x45\xFC\x39\xD8\xE7\xF7\x24\xE7\x49\x50\xF2\x45\x54\xB2\x76\x21\x5E\x3F\xBA\x00\x56\xF1\x62\x73\xB8\x92\xCE\x4B\x9D\x4F\x95\xE1\x37\x5B\x7A\x75\x28\xB0\xB2\x38\x79\x12\xB9\x14\x7E\xEF\x7B\x2C\x2B\x45\x76\x7F\xAA\xFC\x0F\x24\xCD\xD3\xDC\xFF\xF0\xEF\x56\xCF\xD0\x2A\xA4\xA5\xC6\x51\x46\x33\xBF\x60\x92\x2D\xEC\x2E\x09\x2F\xF2\x90\x54\x48\x9B\x82\x45\x5E\x5B\x27\xE7\xA8\x96\xE4\x4E\x33\xA9\x05\x74\x31\x81\x67\x03\x26\x0C\x7F\x1E\xA2\xB7\xDD\x9E\x92\x3A\x90\xDA\x9B\xA4\x83\x56\x72\x55\x20\xAC\xE4\x4E\xA9\x31\x25\x02\x30\x15\x19\x95\xA9\x1E\x08\x9C\xB1\x5A\x79\xDA\xA7\x8D\xEE\x08\x77\x3B\x6B\x2F\xE2\x07\x84\xF0\xF7\xBE\xA9\xFD\x56\x5B\xEB\x35\x8F\xC1\xF5\xAC\xC1\x86\x29\xB5\x09\xA0\xD4\x4E\x0A\xEC\x38\xE0\x6D\xE0\x09\x5B\x5F\xA0\x4E\x02\x2C\x6B\x78\x21\x1D\x10\x3F\x3B\x53\x5B\x8E\x61\xB9\x88\xB2\xA9\x3D\x65\x82\xD4\xB0\x17\xE1\x98\x10\x5C\x47\x76\x0D\xCB\x48\x0D\xBC\xB6\x16\xB5\x4F\x56\xC2\xFC\x6F\x10\x67\x42\x58\x74\x2E\xB4\xCB\xEA\x94\x4E\x99\xD8\x15\xA3\xF3\x31\xDE\x1C\xC3\x5B\xA9\x2F\x83\x1D\x4E\x0F\x78\x7B\x75\xF5\x6B\x73\xC6\xF4\x60\xC1\x3C\x2F\x12\x5E\x30\x9D\xA3\x03\xA3\x37\x8D\x76\xCB\xDB\xBD\xE9\x30\x03\xB0\xDA\x9B\x92\x80\x65\x4A\x2D\x07\xC9\xBB\x59\x88\xA8\x56\x4D\xE9\x40\x1B\x0F\x0E\x7D\x0F\xA4\x77\x1B\x56\x39\x7D\x39\x7C\x39\x4C\xBF\xC4\xA6\x77\x20\x55\xFC\xCC\xCA\x8A\xB2\xA4\xF3\xF8\xF0\x08\x2A\xBD\x2A\xA4\x83\x66\x34\x08\x74\xDC\xCA\x29\x9D\xC9\x58\xC1\x74\x19\xFF\x4A\xED\xBC\xAD\xC3\x88\x00\x54\x73\xF2\x4C\x6B\xA9\x44\x04\xCA\x22\xF3\x98\xD0\x51\xE1\x42\x0F\x02\x25\xBA\x28\xE8\x88\x05\x65\x38\x53\x20\x70\x8E\xCA\x54\x74\xB0\x34\xD6\xB8\x9D\x2E\xEC\x75\xC6\xED\xD3\x51\x5A\xE7\x7B\xC0\x95\xD1\xF1\xB0\xDB\x91\x79\x50\x0A\xE4\x64\x70\xF1\xE4\x02\x3D\x38\xB6\x74\x91\xFA\x17\x08\x05\x9B\x63\x9C\x2E\xBE\x3E\x3F\x12\x5D\x1C\xF2\xE4\xED\xFB\x9F\xCF\xB3\xDC\x64\xCE\xF2\xBD\xA7\xE8\x3D\xB5\x39\x0B\xE8\x81\xC6\x05\xBC\x31\xDD\x7D\x09\xBC\x00\x32\xA1\xC5\x28\x41\x72\xD2\xE3\x20\xA1\xE4\xA0\xF1\x00\xBA\x2A\x74\x9C\x7D\x87\x90\x16\x4B\xE3\x71\xE0\x95\xC8\x6A\x87\x36\x2B\x97\xB4\x48\x2B\x62\xF2\x01\xC3\xCD\x11\xD2\xDD\xE3\xD2\x98\x8B\x2D\x4D\x6D\xA1\xB2\x72\x4E\x82\x1A\x0B\x55\x3D\x55\x92\x77\x31\x3E\x04\x8D\x94\x37\x42\x5A\xE8\x57\x70\x07\x05\x92\x27\xC0\xC5\x5D\x06\xDE\x97\x50\xDE\xC5\xAB\x97\xFC\x23\x7A\xD6\x1B\x03\xA5\x11\xB5\xC2\xDE\x21\x34\xCC\xC3\x6A\xE1\x7A\x07\x87\x11\xB7\x71\x9E\xD6\x05\x24\xBA\xE0\x36\x8E\x9E\x75\x0E\x64\xB7\x46\x8A\x36\xFD\x66\xFC\x86\xE5\x08\x53\xA9\xD9\xA1\xED\xC6\x4B\xB1\x36\xCF\xAA\xF5\xBA\x99\xE0\xEE\x7F\xA2\x56\xCB\x60\x84\xC6\xD9\x29\xEB\x5D\x17\x7F\x82\x3E\x1B\xD9\x03\xCC\xAC\x29\xBF\x1C\xCB\x07\xB1\x22\xDF\xF2\xD0\x7D\x3C\x90\xB5\x3C\xBD\xAE\x4F\xC1\x20\x5B\xC1\x91\x25\x0F\x59\xAA\xF0\xA5\x1A\xF8\xB2\x52\x1B\x4B\x3D\xD0\x6D\x18\xF9\x77\x93\x5D\x10\x54\xC4\x26\x04\xB8\x35\xC6\x77\x99\xC8\x9B\x75\x8A\xB4\x9D\xDF\x24\xCD\x91\x1B\x52\xCA\x95\x44\x69\x87\x95\x42\x26\x13\x2C\x45\x26\x13\x75\x59\x85\x45\x4C\x38\x34\xE5\xCA\x41\x5D\x24\xDA\xB4\x93\xB1\xC4\x8C\xE7\x02\xBD\x0B\x85\x0F\x01\xD2\x37\xF6\x6D\xC5\x19\xAC\xAD\x1C\x98\x8A\xA4\x98\x19\xA5\xCC\x82\xDC\xBF\x59\xF9\x00\x56\x4F\xDA\xB8\x29\x99\xD4\x49\x22\xCB\x70\x53\x7E\x96\x40\xF3\x2F\xDD\x7B\xA5\x4A\x9E\x27\xC9\xAC\xD6\x3C\xCC\x7C\xF6\x1C\xFE\x5C\xCD\x09\xFD\x83\x37\x8D\x26\xCF\xD2\x8E\x41\xD3\x1E\xA4\xEB\x68\xCA\xD6\x90\xF4\x3A\x18\xAF\x7D\x21\x7D\x9E\xFC\xF5\x40\x17\x20\xC0\x48\x32\xF2\x84\xC6\x05\x82\x39\x9B\x30\x5A\xC9\xD0\xA0\x7C\x00\x28\xFF\x6B\x48\x86\x5C\xF0\x16\x8C\x31\xE1\x3A\x38\x6E\x2C\x1C\xD8\x9D\x40\x68\x10\x8B\xA1\xCE\x42\xA9\xA1\xA9\x81\x7C\x25\x3E\x06\xC9\x39\xE3\xC5\xFA\x77\x8F\x4E\x58\x57\x4F\xBB\x2D\xAE\x30\xB5\x12\x31\x6F\xD9\xB7\x4F\x4C\xBC\xD8\x54\x21\x78\x93\x30\xCE\xD1\x39\x14\x94\xC5\x4D\xAD\x59\x38\xB4\x07\xB0\xDD\x18\xE0\x0A\x9D\x4F\x4E\xD7\x3C\x7C\x0A\x1D\x1E\x5A\x11\x5D\x92\xBC\x45\xA5\x4C\x0F\x16\xC6\x2A\xF1\xCD\x7D\x41\xFE\x50\xEB\x36\xD0\xD7\xD4\xD3\x51\x94\x77\x2E\x5B\x31\x23\x8F\x35\xA3\xCF\x1E\xB5\xA3\x4C\x36\x4E\xEE\xE2\x93\x6C\xD2\x50\xBC\x69\x49\x1F\xD9\xA9\x19\x9E\xED\x8A\xB9\x66\xFE\x81\xCE\xFF\xB5\x32\xF7\x43\xE4\x17\xB3\x00\x5B\xEB\x4E\xBA\xDA\x95\x76\x83\x03\x63\xF7\xEB\xF3\x57\x1F\xDF\xB4\xED\xA8\xE7\xD2\x1A\x1D\xB2\xEF\x39\xB3\x32\x78\x49\x48\xFC\x0F\x41\x8E\x4F\x20\x6C\x36\x3E\x26\x05\x49\xC8\x4D\xD1\xEE\xA7\xE8\x8F\x7A\x09\xDA\x68\xF8\x03\xAD\x81\x39\x53\x35\xC2\xCA\x9E\x7B\x75\x72\x32\xD7\x4C\xC5\xB3\xAB\xD6\x9A\xB8\xBF\x5B\x67\xDC\xB8\xEA\x49\x0D\x69\x70\x8C\x74\xF3\xCE\xB8\xFB\xF0\x22\x5F\x89\xB7\x54\x5A\xBC\xB5\x9E\x88\x85\x8F\x38\x4C\xED\x4F\xA3\xB6\x74\x7B\x5F\x61\xB4\xE0\x02\xA7\x6D\x54\x02\x6B\xCF\xB4\x70\x51\x2A\x8C\xF3\x81\xB2\x3A\xDE\x4E\x97\xA9\xCD\xED\x83\x1F\x0F\x92\xDF\x4C\xDD\x52\x83\x43\x84\x74\x23\xE6\x52\x60\x6E\x4D\xEF\x74\x39\xBB\xAB\x98\x57\x05\xA3\xD3\xD8\x7F\x73\xBF\x1B\x8C\xB1\x50\x57\xA1\x58\xC9\xF4\x72\x85\xB1\xB1\x90\xF6\x4A\x71\x17\xB0\x7B\x24\x6E\xD2\x1A\x6E\x21\x95\xDA\xBC\x68\x77\xF4\x27\xBD\x56\xF7\xF1\x1E\x4C\x6B\x4F\x97\x14\x70\x9E\x26\x69\xA4\x11\x26\xB1\x38\xB3\xE8\x8A\x5B\x88\x7B\x03\xCA\xB0\xB8\x08\xB1\xF8\x7A\xD9\x90\x70\xDC\xB9\xEC\x27\xB0\x52\x66\x59\x86\x48\x78\xBD\x7A\xBE\x5D\x00\xDC\x7D\x7D\x5D\x60\x2B\x28\x94\xEC\x06\xC1\xD5\x16\xE3\x0D\x75\x5F\x09\xA9\x0B\x1A\xB3\x98\x34\x2C\x18\xD7\x58\x31\x9B\xC5\x55\xAE\x46\x0A\x36\xE5\xF5\xE6\x78\x3F\xD4\x35\xE6\xA1\x34\x76\x89\xDC\x68\xD1\x6B\x8A\x07\xAD\x61\x9A\xBA\x7F\x50\x3E\x06\x5B\x23\x75\x27\x35\x6C\x82\x64\x9B\x96\x0F\xA5\x50\x14\xE8\xE1\xB4\x75\x55\x48\x2B\x7A\xE0\xD1\xF9\x6D\xB5\xA6\xCB\x15\x29\x51\x4F\xEB\xDF\x64\x2B\x53\xA1\x5E\xD7\x42\x76\x12\x40\x62\x74\xD7\x79\x0F\xA2\xEF\xE0\x81\x4A\xFE\x24\x35\xC5\x61\x0F\xA2\xE7\x6F\xA8\xE3\x4D\x2C\x13\x44\xA5\xEF\x4A\x35\x23\xE7\xAD\xD1\xF9\xE4\x97\xF7\x57\xE7\xA7\xA3\xAC\xF9\x05\xAF\x96\x6D\xD1\xAC\xB7\x7E\x29\xA4\x14\x90\x92\xA8\xE9\x40\x68\x6B\x6D\x11\xB9\xE1\x20\xFC\x6F\x39\x93\x79\x08\x59\x65\xEC\xEC\x80\x19\x4A\x75\x44\x14\x6D\xE4\xD5\x8E\xD8\xCA\xF8\x02\x6D\x98\xD2\x0B\xBD\xE1\x95\x42\x38\x2E\x28\x43\x58\x11\x74\x27\x3B\x20\x4E\x1D\xEC\x29\xF5\xDC\xBB\x9C\xA7\x64\x29\xE3\xEB\x16\x77\xCD\xB4\xB8\xBE\xD1\x66\xA1\xAF\xA7\x75\xEE\xD2\xC9\xC5\xBA\x2F\x78\xCC\x3F\xA8\x0F\x5E\xD5\xB9\x7B\x44\xD1\x6F\xC7\x1D\x39\xC0\xAB\x8D\x87\x82\x69\xA1\x70\x4D\xBC\xC0\xD1\x7A\x39\x93\x9C\x08\xE5\x19\x0E\xF2\x41\x53\xBF\xA3\x6B\xD9\xB9\xE6\x76\x59\xF9\xE7\xA1\xCE\x99\xC4\xD7\x55\x4C\x94\x52\x4B\xE7\xC3\x0B\xEA\x58\x2C\x2C\x70\x63\x15\x66\x11\x72\x46\x98\xA3\xA0\xF3\x61\xC5\xBF\x77\xF5\x9A\x8F\x41\x00\xB7\x2C\xC3\x57\x21\x66\xF3\x40\xDF\xBA\xC1\x44\xDF\x6F\xD5\xEC\xDC\x4A\x9B\x22\x6D\x5D\x91\xDD\x51\xC0\x12\x1F\xFC\x2E\xD0\x62\xBC\x44\xBB\x74\xF2\xA1\x7D\x7C\x54\x4D\x16\x37\x6E\xE5\x33\x43\xDC\x48\xE9\x9A\x99\xF9\x05\xA1\x47\x0E\xBA\xAE\xB2\x32\x7F\xFB\x6D\xDF\xFE\xB7\xBD\xCD\x9E\x53\x66\x31\x9D\x7C\x6D\x78\x78\x5F\x5A\x6B\x81\x16\x78\xED\xBC\x29\x61\xFD\x75\xC8\xC5\xBB\xB3\xF3\x5F\x2E\xCF\xD3\xC9\xAB\xCB\xD7\xA0\x24\x47\xED\xBE\xFA\xAA\x70\xEB\x13\x80\x99\x31\x3E\x7E\x52\xB0\xD9\xD2\x0F\x9F\xB7\x4C\x92\x0B\xE6\x7C\x93\x54\x08\x38\x19\x9E\x0C\xFB\xC7\xC7\xFD\xE1\xDF\x60\x38\x3C\x3D\x79\x71\xFA\xED\xF7\xF0\xFF\xC3\xEF\x86\xC3\x5B\x9B\xC4\xCF\x7E\x46\x59\xF3\xA5\x54\xF7\x1B\x90\xA4\xF3\x63\xFB\xA3\x95\x95\x34\x47\x47\x47\xBF\x9A\x45\x70\xCF\xE9\x12\x46\x2C\x7C\x3C\x72\x47\x7C\x69\x6C\xFC\xDA\x24\xFC\x8E\xBB\xB3\xCE\x97\x32\x47\xDB\xF2\xFD\x27\x00\x00\xFF\xFF\xE6\xD5\x74\x58\xE1\x25\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.html")
	node.SetSize(2808)
	return node
}

func generate__sub() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/sub",
		Path:            "/sub",
		ContentType:     "",
		ContentEncoding: "",
	}
	node.SetMode(2147484141)
	node.SetName("sub")
	node.SetSize(0)
	node.AddChild(_getNode("/sub/index.html", generate__sub_index_html))
	return node
}

func generate__sub_index_html() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/sub/index.html",
		Path:            "/sub/index.html",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xA4\x55\xDD\x6E\xDB\x36\x14\xBE\x96\x9F\x82\x63\x81\xDD\x6C\x0A\xED\xA6\x5D\xBB\x94\x32\xD0\x65\x19\x30\xA0\xC0\x0A\x34\x05\xB6\xAB\xE1\x48\x3A\x36\x89\x51\x24\x4B\x1E\x39\xF1\xDB\x0F\xA4\x24\x5B\x51\xBC\xA2\xC0\xAE\x44\x9E\xDF\xEF\xFC\x7C\x94\xFC\xEE\xD7\x3F\x6E\xEF\xFF\xFA\x78\xC7\x14\x75\x66\xBB\x92\xC3\xA7\x90\x0A\xA1\xDD\xAE\x8A\x42\x76\x48\xC0\x14\x91\x2F\xF1\x4B\xAF\x0F\x15\xBF\x75\x96\xD0\x52\x79\x7F\xF4\xC8\x59\x33\xDC\x2A\x4E\xF8\x48\x22\x79\xBF\x63\x8D\x82\x10\x91\xAA\x9E\x76\xE5\x5B\xCE\xC4\x39\x8E\x85\x0E\x2B\x7E\xD0\xF8\xE0\x5D\xA0\x99\xF7\x83\x6E\x49\x55\x2D\x1E\x74\x83\x65\xBE\xFC\xC8\xB4\xD5\xA4\xC1\x94\xB1\x01\x83\xD5\xE6\x42\x20\x52\xD8\x61\xD9\x38\xE3\xC2\x2C\xD6\x8B\xEB\x37\xAF\xEF\xDE\xFF\x92\xED\x93\x03\x69\x32\xB8\xFD\xD4\xD7\xAC\xD5\x01\x1B\x72\xE1\x28\xC5\x20\x4C\x6A\xA3\xED\x3F\x2C\xA0\xA9\x78\xA4\xA3\xC1\xA8\x10\x89\x33\x15\x70\x57\x71\xA1\x6D\x8B\x8F\x57\x4D\x8C\x53\xF6\xAF\x5A\x37\x7D\x24\xD7\x9D\xCD\xA5\x18\xDB\x28\x6B\xD7\x1E\xB3\x7F\xAB\x0F\xAC\x31\x10\x63\xC5\xC9\xF9\x1A\x02\x4F\xE2\x27\xF2\x54\x08\x68\x8B\xA3\x6A\xE9\x53\xA6\x98\xDA\xEE\x47\x6D\x21\x61\x4A\xCF\xB7\x8D\xD6\x7B\x27\x05\x8C\x8E\xA2\xD5\x87\xE7\x31\x3A\xB4\xFD\xC9\x79\xE7\x42\x37\x29\x34\x61\xC7\x19\x34\xA4\x9D\xAD\xB8\xF8\x5B\x5B\xC2\x60\xC1\x88\x88\x10\x1A\x35\xF9\x14\x52\x5B\xDF\x13\xA3\xA3\xC7\x61\xEE\x7C\x9C\xC7\x17\xCE\xBC\x81\x06\x95\x33\x2D\x86\x8A\x7F\x1A\xFC\x86\xD6\x65\x40\x29\xDD\xB7\x83\x3B\x55\x16\xFB\x9A\xA7\x09\x5E\x28\xED\x74\x1A\x0F\x8B\x26\x7B\xD8\xE3\x37\xB4\x58\x6D\x96\x0B\xA2\x36\x97\xC0\x11\x4C\xE0\x46\x88\x75\x38\x55\xF7\x72\xFD\x72\x5D\xAE\x7F\x2A\xD7\xAF\xD9\xE6\xFA\x66\xF3\xEA\xE6\xFA\x67\xF6\xFD\x8B\x57\xD7\xEF\xD6\x6F\xD6\xEB\x67\x35\x5F\x22\x1E\x33\x60\xF7\x15\x47\xCB\x93\x20\xAF\xCE\xB0\xEE\x13\xA3\xF8\xE7\xFB\xDF\xCA\xB7\x7C\x12\xCF\x69\xF9\x67\xF9\xF9\x7D\x79\xEB\x3A\x0F\xA4\x6B\x33\x67\xE6\xEF\x77\x15\xB6\xB9\x0B\xFF\x9F\x84\x57\xEB\x45\x98\x3D\x5A\x0C\x40\x4F\x08\x98\xB7\x30\xD9\x7D\x8D\x79\x32\xF3\x67\xBB\x5A\x49\x31\x9E\x26\xB6\x64\xB2\x4C\x2D\x87\x40\xBA\x31\x19\x7C\x1A\x84\x6E\x2B\x9E\xAC\xF2\xE4\xFE\x63\x6A\xF3\x89\xB5\x48\xA0\x4D\x4C\xC6\x43\xEF\xA7\xCF\x14\x6C\x04\x3D\x8F\xEF\x03\x42\x57\x9F\x73\x8E\xA1\x22\x66\x62\x24\x70\x0B\x8D\x87\x00\xFB\x00\x3E\x31\x44\xFA\xED\xBD\xD2\x91\xE9\xC8\xC0\x32\x7C\x84\xCE\x1B\x64\x6E\x37\x75\x87\x69\xCB\xE2\x1C\x33\xEB\xA3\xB6\x7B\x36\x3C\x1D\xEC\xFC\xA8\x5C\x49\xE1\x9F\xA1\xBE\x5C\xC3\xCE\x39\x1A\x1A\xF2\x54\x52\x66\x6E\x6E\x57\x1F\x20\x12\xEB\x7D\x0B\x84\x2D\xBB\xB4\xA5\x3F\xE4\x0D\x5D\x26\x19\xDE\x2C\x29\xC6\x5F\xC2\x7C\x7D\x57\xB3\xCB\x92\x71\x27\x34\x45\x51\x7C\x74\x0F\x18\xB0\x65\xF5\x91\x49\xC8\x7B\x3F\x10\x3A\x2D\x6E\xBC\x11\x62\xAF\x49\xF5\xF5\x55\xE3\x3A\x11\x55\xFF\x41\x81\x15\xC3\xF2\x24\xDB\x81\x28\xF9\x3E\x64\x87\x19\xCD\x8B\x25\xBE\x7F\x03\x00\x00\xFF\xFF\xA1\x1D\x5B\x4A\xCA\x06\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.html")
	node.SetSize(716)
	return node
}

func generate__example() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example",
		Path:            "_example",
		ContentType:     "",
		ContentEncoding: "",
	}
	node.SetMode(2147484141)
	node.SetName("_example")
	node.SetSize(0)
	node.AddChild(_getNode("_example/html.tmpl", generate__example_html_tmpl))
	return node
}

func generate__example_html_tmpl() *memfs.Node {
	node := &memfs.Node{
		SysPath:         "_example/html.tmpl",
		Path:            "_example/html.tmpl",
		ContentType:     "text/html; charset=utf-8",
		ContentEncoding: "gzip",
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x8C\x54\x4F\x6F\xDC\x2E\x10\x3D\x6F\x3E\x05\xE1\x77\xFD\x61\x14\x55\x55\xAB\x16\xAF\x94\x7F\xB7\x4A\x8D\x94\x5C\x7A\xAA\x30\x9E\x5D\x50\x31\x38\x30\xDE\x64\x85\xFC\xDD\x2B\x8C\xBD\x75\xD2\x48\xCD\xC9\xC0\x9B\xF7\x78\x6F\x3C\xB6\x38\xBF\xF9\x7E\xFD\xF0\xE3\xEE\x96\x68\xEC\xEC\xF6\x4C\x94\xC7\x46\x68\x90\xED\xF6\x6C\xB3\x11\x1D\xA0\x24\x1A\xB1\x67\xF0\x38\x98\x43\x4D\xAF\xBD\x43\x70\xC8\x1E\x8E\x3D\x50\xA2\xCA\xAE\xA6\x08\xCF\xC8\x33\xFB\x2B\x51\x5A\x86\x08\x58\x0F\xB8\x63\x9F\x29\xE1\x7F\x74\x9C\xEC\xA0\xA6\x07\x03\x4F\xBD\x0F\xB8\x62\x3F\x99\x16\x75\xDD\xC2\xC1\x28\x60\xD3\xE6\x7F\x62\x9C\x41\x23\x2D\x8B\x4A\x5A\xA8\x2F\xDE\x10\x42\x0D\x1D\x30\xE5\xAD\x0F\x2B\xAD\xFF\x3E\x7C\xFA\x78\x7B\x79\x35\xD5\x67\x02\x1A\xB4\xB0\x4D\xA9\x7A\xC8\x8B\x71\x14\xBC\x9C\x64\xCC\x1A\xF7\x8B\x04\xB0\x35\x8D\x78\xB4\x10\x35\x00\x52\xA2\x03\xEC\x6A\xCA\x8D\x6B\xE1\xB9\x52\x31\xCE\x57\xA7\xC4\x48\x90\x6E\x0F\xA4\xBA\x9F\xAA\xC7\xF1\x5F\x1A\x29\x55\xE3\xB8\xA2\x83\x6B\x33\x49\xF0\xB9\xBF\xA2\xF1\xED\x71\x72\xD2\x9A\x03\x51\x56\xC6\x58\x53\xF4\x7D\x23\x03\xCD\xC7\x2F\xCE\x73\x42\x69\x1C\xCC\xD0\x6B\x0E\xCB\x9A\xC6\xED\x67\x74\x23\xE4\x12\x84\x6E\x95\x31\x7B\x2F\xB8\x9C\x89\xBC\x35\x87\xBF\x35\x3A\x70\xC3\x89\xBC\xF3\xA1\x5B\x00\x83\xD0\x51\x22\x15\x1A\xEF\x6A\xCA\x7F\x1A\x87\x10\x9C\xB4\x3C\x82\x0C\x4A\x2F\x9C\x8D\x30\xAE\x1F\x90\xE0\xB1\x87\x32\x10\x74\x7E\x51\x8F\x94\xF4\x56\x2A\xD0\xDE\xB6\x10\x6A\x7A\x5F\x78\xA5\x2F\x93\xA1\x7C\xDD\xFB\xCD\x9D\x92\xC5\xA1\xA1\xDB\xFB\xA1\x79\x23\xDA\x69\x35\x2F\x5E\x35\xB9\x97\x7B\x78\x47\x8B\xF5\xC5\x8B\xC9\xD1\x17\x6F\x39\x43\xB9\x38\x4B\xA9\xBA\x1C\x50\xFB\x30\x8D\x46\x2E\x6C\xC2\x29\x65\x4A\xD5\x8D\x44\x98\xA1\x55\xCE\x94\xAA\x2B\xDF\x1E\x0B\xB0\x4A\x70\xCE\x18\xA9\x4E\x96\x08\x63\xAB\x38\x0B\x9C\x73\x4C\xC8\xAB\x7C\x3B\xEF\x71\x89\x71\xE7\x9F\x20\x40\x4B\x9A\x23\x11\x72\xBA\xB1\xB4\x2F\x7F\xD6\xF1\x0B\xE7\x7B\x83\x7A\x68\x2A\xE5\x3B\x1E\xF5\xF0\x4D\x4B\xC7\xA7\x81\xA1\xB9\xB6\x58\x9C\xF6\xC5\x9E\x5C\xBB\x10\xBC\x4C\xB0\xE0\xE5\xCF\xF1\x3B\x00\x00\xFF\xFF\x81\xA3\x1E\xB5\x51\x04\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("html.tmpl")
	node.SetSize(520)
	return node
}

//
// _getNode is internal function to minimize duplicate node created on
// Node.AddChild() and on GeneratedPathNode.Set().
//
func _getNode(path string, fn func() *memfs.Node) *memfs.Node {
	node := memfs.GeneratedPathNode.Get(path)
	if node != nil {
		return node
	}
	return fn()
}

func init() {
	memfs.GeneratedPathNode = memfs.NewPathNode()
	memfs.GeneratedPathNode.Set("/",
		_getNode("/", generate_))
	memfs.GeneratedPathNode.Set("/LICENSE",
		_getNode("/LICENSE", generate__LICENSE))
	memfs.GeneratedPathNode.Set("/custom.css",
		_getNode("/custom.css", generate__custom_css))
	memfs.GeneratedPathNode.Set("/favicon.ico",
		_getNode("/favicon.ico", generate__favicon_ico))
	memfs.GeneratedPathNode.Set("/html.tmpl",
		_getNode("/html.tmpl", generate__html_tmpl))
	memfs.GeneratedPathNode.Set("/index.css",
		_getNode("/index.css", generate__index_css))
	memfs.GeneratedPathNode.Set("/index.html",
		_getNode("/index.html", generate__index_html))
	memfs.GeneratedPathNode.Set("/sub",
		_getNode("/sub", generate__sub))
	memfs.GeneratedPathNode.Set("/sub/index.html",
		_getNode("/sub/index.html", generate__sub_index_html))
	memfs.GeneratedPathNode.Set("_example",
		_getNode("_example", generate__example))
	memfs.GeneratedPathNode.Set("_example/html.tmpl",
		_getNode("_example/html.tmpl", generate__example_html_tmpl))
}
