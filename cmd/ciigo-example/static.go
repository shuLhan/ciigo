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
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\xBC\x5A\x5F\x8F\xDB\xB6\xB2\x7F\x8E\x3E\xC5\x54\x0D\x82\x04\xB0\xAD\xA4\xB9\x17\xED\xDD\xDA\xBE\x68\x36\xDB\x24\x38\xDB\xA6\x48\xB6\x0F\x45\x71\xB0\xA0\xC8\xB1\x44\x2C\x45\xAA\x24\x65\xC7\x3D\x38\xDF\xFD\x60\x48\x49\x96\xBD\xF6\xD6\xBB\xEB\x9E\xDD\x07\x4B\xFC\x33\x9C\xF9\xCD\x1F\x0E\x87\x9A\x7E\xF5\xF6\xE3\xF9\xD5\x6F\xBF\x5C\x40\xE9\x2B\x35\x4F\xA6\xF1\xE7\xC9\xB4\x44\x26\xE6\xC9\x93\x27\xD3\x0A\x3D\x83\xD2\xFB\x7A\x8C\x7F\x34\x72\x39\x4B\xCF\x8D\xF6\xA8\xFD\xF8\x6A\x5D\x63\x0A\x3C\xBE\xCD\x52\x8F\x5F\x7C\x46\xB3\xBF\x07\x5E\x32\xEB\xD0\xCF\x1A\xBF\x18\x7F\x97\x42\xB6\xA1\xA3\x59\x85\xB3\x74\x29\x71\x55\x1B\xEB\x07\xB3\x57\x52\xF8\x72\x26\x70\x29\x39\x8E\xC3\xCB\x08\xA4\x96\x5E\x32\x35\x76\x9C\x29\x9C\xBD\xDA\x43\xC8\x97\x58\xE1\x98\x1B\x65\xEC\x80\xD6\xD7\xAF\xBF\xFD\xDF\x8B\x1F\xDE\x84\xF1\x34\xC1\x4B\xAF\x70\x3E\xCD\xE2\x2F\xB5\x28\xA9\x6F\xC0\xA2\x9A\xA5\xCE\xAF\x15\xBA\x12\xD1\xA7\x50\x5A\x5C\xCC\xD2\x4C\x6A\x81\x5F\x26\xDC\xB9\xB8\xE0\x34\x6B\xA1\x98\xE6\x46\xAC\xC3\x74\x21\x97\xC0\x15\x73\x6E\x96\x7A\x53\xE7\xCC\xA6\xD4\xBC\xD5\x4E\xCC\x30\xA9\xB1\xED\xDA\x9D\x33\x26\x9A\x52\x17\x6D\xEF\x93\x29\xEB\x56\x4F\xE7\x5C\xCA\xC2\x4C\x33\xD6\x4E\xCC\x84\x5C\xDE\xA6\x51\xA1\x6E\xFA\xC9\x0B\x63\xAB\xAE\x43\x7A\xAC\x52\x60\xDC\x4B\xA3\x67\x69\x76\x2D\xB5\x47\xAB\x99\xCA\x1C\x32\xCB\xCB\x6E\xCE\x93\xA9\xD4\x75\xE3\xC1\xAF\x6B\x8C\xBA\x4B\x5B\x4C\xFF\x48\xA1\x56\x8C\x63\x69\x94\x40\x3B\x4B\x3F\xC7\x79\x11\xFC\xC0\x10\x2D\x77\x3C\x73\xBD\x64\xAE\xC9\xD3\xF9\xE7\x26\xDF\x23\x5A\xFF\xD4\x3E\xEC\x80\x5C\xB3\x02\x8F\x80\xB8\x7C\x35\x9F\x66\xE5\xAB\x7D\xFC\x78\xD6\xF1\xD3\x72\x95\xDB\x5E\xA0\xBD\x92\x48\x31\x4B\x6B\x8B\xAC\xCA\x15\x2D\x3D\xA4\xE6\x30\x80\x4B\xD6\xB0\xD3\x53\x33\xCB\x0A\xCB\x6A\x42\x79\x5A\xCF\xA7\xAE\x66\xBA\xD7\x4B\x15\x84\x98\xCA\xAA\x00\x67\xF9\x2C\x25\x97\x72\x67\x59\x26\xAB\x62\xE2\x4A\x89\x4A\xB8\x89\x34\x59\xCE\x44\x81\x59\x61\x26\x02\x97\x63\x8B\x0B\xB4\xA8\x39\x8E\x5F\xBE\xFC\x56\xFC\x1F\xFF\x7F\x65\x0A\x33\x2B\xCC\x33\xFA\x3D\x27\xB3\x9F\xAD\x4A\xE9\xF1\x59\x30\xE3\xD9\x42\x31\x3F\x76\x7F\x34\xCC\x62\x0A\x4C\xF9\x59\xFA\xCE\xBC\x35\x3C\x9D\x4F\x33\x62\x65\x9E\x1C\xC7\x51\x61\x2C\x92\x7B\x72\x66\xC5\x84\x9B\xAA\xE3\x49\xFA\xB2\xC9\x43\x83\x2B\x9B\xCB\x92\xE9\x2C\x98\x6A\xBF\x14\x7C\x0A\xD3\xE0\x9C\x59\xD1\xAF\x39\xCD\xEA\x79\xD2\xA2\x7B\x17\x56\xDC\x08\xEC\x4C\x3F\x3C\x83\x74\xC0\x40\xC9\xDC\x32\xBB\x06\xA6\x05\x30\xA8\xAD\x29\x2C\xAB\xC0\x1B\x58\x59\xE9\x11\x9C\x67\x5E\x72\x58\x61\x0E\x0E\xED\x12\x2D\xAC\xA4\x2F\x01\xAB\x1C\x85\x40\x91\x2C\xA4\x42\x07\x8D\x93\xBA\x80\x02\x35\x5A\xE6\x51\x40\xC5\xEC\x4D\x53\x03\xD9\x31\xF3\x93\x23\x39\x3C\x6F\xAC\x45\xED\xD5\x7A\x04\x81\x4D\x70\x4D\x4D\xE2\x26\xBD\x85\x77\xF8\x31\xC7\xA5\x14\x86\x7B\x63\x27\xC6\x16\x99\x30\xDC\x65\xAB\x92\xF9\xB1\x74\xE3\xAE\x33\x4B\xE7\xDD\x63\x70\x08\xA6\xC5\x6D\x4A\xDC\x54\x95\xD1\xC4\x6E\x20\x94\xCE\xE9\x51\x98\x95\x8E\x53\xDC\x5D\x92\x6C\xFF\xEC\xD8\xEF\x2B\x92\xA9\xFC\x26\x98\xF9\x75\x10\xE7\x9A\xB9\xEB\x16\xEC\x36\x04\x01\x73\x1D\xFC\xD3\xAC\xFC\xE6\x81\x4E\xF0\xA3\xB1\xC0\x34\x34\x35\x29\x4D\x30\x8F\x20\x0C\x6F\x2A\xD4\xA4\x39\xA3\xC1\x2C\xA0\x34\x2B\xEA\x6C\x1C\x82\x2F\xB1\x57\xB9\xC3\xF0\x7E\x1B\x95\xFA\xA6\x98\x44\x07\x39\x6C\x93\xF3\x77\x66\x67\x21\x0A\x22\x04\xDB\xA3\x51\xE2\x4A\x0E\x10\x3A\xBF\xFC\xF0\x08\x74\x86\x64\x80\x93\x6F\x1A\xBD\x44\xEB\x47\xBD\xB5\x8E\xC8\xF4\x33\x63\xA3\x7D\x03\x03\x21\x2D\x92\x69\xAD\xC1\x97\xCC\x43\x1B\x05\x5D\xD2\x99\x02\x59\xFC\x88\x48\xBE\xBF\xFA\xE9\x32\xBE\x1E\x34\x71\xE2\xF3\x9B\x20\xE4\xEB\x28\x64\xE3\x42\x4C\xF8\xD5\x05\xB0\xCA\xD7\xDB\xC3\x95\x74\x5E\xEA\x22\x57\x86\xDF\xEC\xC8\xD5\x6E\xBE\x41\x2A\x8B\xF3\xA7\xAD\x8F\xFC\x3E\xF6\x58\xD5\x8A\xD4\xFE\x4C\xF9\xEF\x89\x9B\x67\x85\xFF\xFE\x9F\x9D\x9C\xA1\x55\x48\x4B\x8D\xD3\x8C\x66\xDE\xA1\x92\x1D\xEC\x3E\x13\x5E\x64\x30\xCF\xBE\x7E\xFD\x3F\x81\x08\xFD\x82\x45\xDE\x58\x27\x97\xA8\xD6\x64\x55\x0B\xA9\x37\x0E\x1F\xA2\xC1\xF3\x09\x13\x86\x83\xB1\x30\xA9\xC4\x0B\x82\x37\xE9\xB8\xA1\x7C\x02\xA4\xF6\x66\x08\x5E\x72\x55\x22\x6C\xC4\xA0\x45\x82\x1C\x61\x35\xE9\xC0\xD4\xA4\x6B\xA6\x46\x20\x70\xC1\x1A\xE5\x69\xD9\x2E\x02\x45\x42\xDD\xEC\x83\x8A\x38\x21\xB2\xBF\x8F\x4D\xE3\x77\xDA\x3A\x63\x7A\x0C\xDC\xE7\x2D\x46\x4C\xA9\x6D\x3C\xA5\x76\x52\xE0\xC0\x2E\x0F\xEB\x83\xA2\xB8\x2F\x51\x27\x01\x9E\x01\xDC\x61\xE8\xA4\x30\x71\x8A\x33\x8D\xE5\x18\xC8\x47\xF4\x4D\xE3\x29\x59\xA1\x86\x83\x88\x87\xA9\xD1\x49\xE3\xAE\xD0\x93\x93\x1A\x78\x0C\xDE\x49\xCF\xE4\x7F\x47\x13\x4C\x08\x8B\xCE\x85\x76\x59\x9F\xD1\x76\x11\xBB\xA2\x33\x3F\xC6\xF8\x63\x34\x50\xEA\x38\x25\xC4\x1D\xF0\xFD\xD5\xD5\x2F\xED\x3E\x39\x82\x15\xF3\xBC\x4C\x78\xC9\x74\x81\x0E\x8C\xDE\x56\x2A\x69\x6A\xCB\x2B\xB6\x7C\x02\x58\xE3\x4D\x45\x20\x33\xA5\xD6\x93\xE4\xC3\x22\x38\x62\x27\xAE\x74\xA0\x8D\x07\x87\x7E\x04\xD2\xBB\x5B\x5A\x3A\xFB\xEE\xE5\x77\x2F\xA3\xCA\xEF\x08\xC6\x47\xC4\x64\xFC\xC2\xAA\x9A\x12\xB4\x8B\xF8\xF0\x88\x48\x7C\x55\x4A\x07\xED\x68\x10\xE8\xB8\x95\x39\xE5\x17\x58\x43\xBE\x8E\xBF\x52\x3B\x6F\x9B\x30\x22\x00\xD6\xEE\x5B\x79\x23\x95\x88\x80\x59\x64\x1E\x13\xDA\x69\x5C\xE8\x41\xA0\xA3\x0E\x0A\xDA\xA0\x41\x19\xCE\x14\x08\x5C\xA2\x32\x35\xED\x4B\xAD\x56\x6E\xA7\x3E\xC7\xA6\x24\x3F\x4A\xEB\xFC\x08\xB8\x32\x3A\x6E\x9D\x7B\xB2\x28\x4A\xE5\x9C\x0C\x26\x9F\x5C\xA2\x07\xC7\xD6\x2E\xEE\x1C\x2B\x84\x92\x2D\x31\x4E\x17\x7F\x3D\x3F\x78\x6A\x12\x87\x3C\x7D\xFF\xF1\xA7\x8B\xAC\x30\x99\xB3\xFC\xE0\x26\x7C\x4F\x69\xCE\x03\x7A\xA0\x71\x05\xEF\xCC\x70\x5D\x02\x2F\x80\x4C\x68\x31\x4A\xF6\x9C\xF4\x38\x49\x28\xB5\x68\x2D\x80\x0E\x8B\x03\xE3\xDF\xC3\xA4\xC5\xCA\x78\x9C\x78\x25\xB2\xC6\xA1\xCD\xAA\x35\x11\xE9\x58\x4C\x3E\x61\x38\xF4\x44\xFB\xDC\x3F\x36\x7A\x52\xC8\x2F\xD7\xA6\xB1\x50\x5B\xB9\x24\x86\x8D\x85\xBA\xC9\x95\xE4\x43\xAC\x4F\x11\x5E\xAA\x1B\x21\x2D\x8C\x6B\x38\x42\x90\xE4\x29\x70\x71\xCC\xC0\xFB\x06\x9A\x0F\xF1\x10\x2E\xFF\x8C\x16\xF6\xCE\x40\x65\x44\xA3\x70\x74\x0A\x09\x8B\x40\x2D\x1C\xF4\xE1\x34\xEC\xB6\x46\xD4\x99\x82\x44\x17\xCC\xC7\xD1\xB3\x2E\x80\xF4\xD6\x72\xD1\x1D\x29\x18\xBF\x61\x05\x42\x2E\x35\x3B\xB5\xDE\x78\x25\x36\xEA\xE9\x5B\xAF\xDB\x09\xEE\xFE\x3B\x70\xBD\x0E\x4A\x68\x8D\x9E\x72\xE7\x4D\xFD\x22\xC8\xB3\x95\x6D\xC0\xC2\x9A\xEA\x6E\x9F\x3E\x89\x16\xF9\x8E\x85\x1E\x8A\x07\x59\x17\xAF\x37\x25\x16\x98\x64\x3D\x1C\x59\xF2\x10\x52\xA5\xAF\xD4\xC4\x57\xB5\xDA\x22\xF5\x40\xB3\x61\x64\xDF\x6D\xF6\x41\x50\x51\x54\x21\xC0\xAD\x31\x7E\x18\x91\xBC\xD9\xA4\x54\xBB\xF9\x50\xD2\x6E\xC5\x7D\x66\xDA\x73\x15\x03\xC8\x26\x4A\x6D\x12\x4D\x52\x9D\x68\xAA\x3A\x10\x33\x61\x33\x95\xBD\xA1\xBA\x18\x78\xC3\xEC\x9D\xCC\x26\x66\x48\x97\xE8\x5D\x28\xDE\x08\x90\xBE\xD5\x77\xC7\xDE\x64\xA3\xF5\x10\xB9\x88\xAB\x85\x51\xCA\xAC\xC8\x1D\xDA\x15\x4E\x60\x05\x59\x56\x98\xB3\x1E\x93\xC2\x80\x6D\x34\x0C\x98\x48\x92\xCE\xD1\x2A\x26\x75\x92\xC8\x2A\x94\x0A\x9E\x27\xD0\xFE\x05\x81\x0E\xA9\x3B\x74\x26\x2F\x92\x64\xD1\x68\x1E\x28\x3C\x7F\x01\xFF\xEA\xE7\x86\x31\x93\x77\xED\x6A\xCF\x63\x66\x91\x6D\xE3\x3E\x6A\x13\xC4\xDE\x25\xB3\x6D\x2C\x47\x3B\xCA\xDA\x18\x56\x68\x7F\x91\xFC\xFB\x81\x36\x45\x88\x13\xC7\x64\x5A\xAD\x4D\x05\xFB\x68\xFD\xB2\xE7\xA7\x55\xD3\x09\x74\xF1\xB7\x23\x1D\x92\xCF\x3B\x61\xDE\x64\x78\x7F\x1B\xAE\x2C\x64\x0A\x03\xCF\x6B\x11\x8D\xB1\x85\x85\xBA\x4A\x5B\x48\x3A\xD2\x21\x27\xC9\x05\xE3\xE5\xE6\x7D\x44\x5B\xBB\x6B\xF2\x61\x8B\x2B\x4D\xA3\x44\x4C\x9C\x0E\xAD\x17\x33\x3F\x96\x2B\x04\x6F\x12\xC6\x39\x3A\x87\x82\xD2\xC8\xDC\x9A\x95\x43\x7B\x02\x1D\xCF\x00\xAE\xD0\xF9\xE4\x6C\xB3\x01\x9C\xC1\x20\x00\xF6\x11\x36\x49\xDE\xA3\x52\x66\x04\x2B\x63\x95\xF8\xEA\xBE\x60\x7F\x6A\x74\x17\x51\x36\x31\x6F\x20\x28\x1F\x9C\x0A\xE3\xD1\x20\x16\xE0\xBE\x78\xD4\x8E\x52\xE9\x38\x79\x0B\x9F\xE7\xC6\x76\xCD\x95\x68\x1B\x5F\x24\xDB\xD1\x30\x1E\x10\xA5\x8F\x41\xB2\x1D\x9D\xED\xF3\xDE\x96\xC0\x89\xD2\x91\x8D\x88\xF7\xC3\xE9\x67\xB3\x0A\x21\x6F\x93\x45\x0F\xB9\xDD\x0A\xC1\xB1\xFB\xED\xC5\x9B\x5F\xDF\x75\xED\xA8\x97\xD2\x1A\x1D\x0E\x05\x4B\x66\x65\xB0\x9D\x70\x1E\x39\x45\x6C\x7E\x0A\x61\xB1\xD9\xAB\x2E\x2E\x6F\xB3\x76\x3F\x41\x7F\xD0\x6B\xD0\x46\xC3\x9F\x68\x0D\x2C\x99\x6A\x10\x7A\x2D\x1F\x94\xC9\xC9\x42\x33\x15\xB7\xD2\x46\x6B\xDA\x7A\x86\xA5\xDC\xAD\x93\xA8\xD4\x6D\x35\x80\x4C\x66\xB3\xCD\xF5\x47\xDC\xBB\xF7\x54\xB2\x9D\x78\xB8\xA6\xC5\xFA\x92\x6F\x12\x0C\x2B\x0E\x53\x87\xB3\xBC\x1D\x59\x3F\xD6\x18\x35\xBA\xC2\xBC\xF3\x5D\x60\xDD\x16\x1B\xCE\x73\xA5\x71\x3E\x04\xBB\x81\x4F\xD0\x99\x6F\x7B\xF9\x60\xD7\x93\xE4\x37\xD3\x74\x01\xC4\x61\x2B\xC9\x96\x77\x46\x69\x98\xDB\x6C\x1A\x74\x96\x3C\x96\xDD\xAB\x92\x51\xB2\xE0\xBF\xBA\xDF\x81\xCB\x58\x68\xEA\x50\x99\x65\x7A\x7D\x0B\x7B\x63\x63\xD3\xA8\x12\xF7\x55\xC6\x28\x94\x6F\x3B\x45\xAF\xA4\x52\xDB\x75\x83\x01\x3E\x24\x6F\x5F\x5E\x18\x41\xDE\x78\x3A\x63\x81\xF3\x34\x49\x23\x8D\x30\x89\xC5\x85\x45\x57\xDE\xD2\x88\x37\xA0\x0C\x8B\x44\x68\x4F\xD8\x90\x0D\xF9\xD1\xD1\xC5\x4F\x81\xB5\x32\xEB\x2A\x78\xCE\xDB\xFE\xF9\x76\x19\x74\xFF\x29\x7C\x85\x1D\xA3\x50\xB1\x1B\x04\xD7\x58\x8C\x07\xED\x43\x15\xB3\x7D\xE0\x31\x8B\x5D\x35\x32\xD2\xEA\x23\xA2\xC5\x3E\xBD\x22\x41\xDB\x9B\x8F\x36\xA9\x38\xD5\x69\xEC\xA1\xE1\xEF\x33\x72\xA3\xC5\xA8\xAD\x85\x74\x0A\x6A\xAF\x64\x02\x08\xD1\x29\x5B\xAE\x07\x99\x6D\xEB\x4C\xBB\xE1\xFC\x54\x02\x45\x86\x1E\x1E\xEE\xAE\x4A\x69\xC5\x08\x3C\x3A\xBF\x2B\x56\xBE\xEE\x83\x19\xF5\x74\x76\x4E\xBA\x32\x35\xEA\x4D\x69\x67\x6F\xA0\x48\x8C\x1E\x1A\xF1\x49\xE4\x9D\x3C\x50\xC8\x1F\xA5\x26\x7F\x1C\x41\xF4\x80\x2D\x71\xBC\x89\xD5\x8E\x28\xF4\xB1\xA1\x68\xEA\xBC\x35\xBA\x98\xFF\xFC\xF1\xEA\xE2\x6C\x9A\xB5\x6F\xF0\x66\xDD\xD5\x02\x47\x9B\xFB\x3A\xA5\x80\x84\x44\x4D\x1B\x49\x57\x42\x8C\xC8\xBD\x9C\x84\xFF\x2E\xB6\x32\x0F\x21\x97\x8D\x9D\x03\x30\x43\x05\x92\x02\x46\xE7\x81\x8D\xA3\x68\x66\x7C\x89\x36\x4C\x19\x85\xDE\x70\xC1\x12\xB6\x19\xCA\x2C\xFA\x40\x3E\xC8\x2A\x28\xE6\x4E\x0E\x54\xAE\xEE\x5D\x9D\x54\xB2\x92\xF1\xF2\xC9\x5D\x33\x2D\xAE\x6F\xB4\x59\xE9\xEB\xBC\x29\x5C\x3A\xBF\xDC\xF4\x05\x8B\xF9\x07\xF5\xC1\x9B\xA6\x70\x8F\xA8\x61\xEE\x39\xEA\x07\x78\xB5\xF1\x50\x32\x2D\x14\x6E\x02\x30\x70\xB4\x5E\x2E\x24\xA7\x80\xF2\x1C\x27\xC5\xA4\x2D\x47\xD2\x69\xF2\x42\x73\xBB\xAE\xFD\x8B\x50\xBE\x4D\xE2\xDD\x1D\x13\x95\xD4\xD2\x79\xCB\xBC\xB1\xB1\xF6\x59\xE2\x16\x15\x66\x11\x0A\x46\x98\xA3\xA0\x7D\xA3\x8F\xC3\xC7\x5A\xCD\xAF\x81\x01\xB7\xAE\xC2\xF7\x19\x66\x7F\x22\xB0\x73\x7E\x8A\x3E\xD0\x89\x3B\x38\x5C\xB7\x35\xE8\x78\x49\x8B\x02\xD6\xF8\xE0\x0B\x52\x8B\xB1\x26\xE0\xD2\xF9\xA7\xEE\xF1\x51\xA5\x66\xDC\x2A\x32\x2C\x0C\xC5\x48\x4A\xF7\xCC\xC2\xAF\x08\x45\x32\xD4\x4D\xF1\x98\xED\xB9\x63\x3E\x7C\x19\xDF\xAE\x99\x33\x8B\xE9\xFC\xAF\x86\x87\x7B\xE4\x46\x0B\xB4\xC0\x1B\xE7\x4D\x05\x9B\xEF\x35\x2E\x3F\x9C\x5F\xFC\xFC\xF9\x22\x9D\xBF\xF9\xFC\x16\x94\xE4\xA8\xDD\xE3\x2E\x50\x2D\x0A\xE9\x5D\x4A\x79\x07\x3D\x3C\xBA\x58\xDF\xA1\x25\x5D\x57\x55\x47\x11\x53\xEC\x12\x55\xED\x62\xF9\x6B\xBB\xE0\xE1\x29\x9C\x43\xCD\xAC\x5F\x27\xF1\x9E\x59\xA2\x3B\x18\x7B\x1B\x8A\x4B\xB4\x62\xA3\xE6\xC9\x54\xC9\xE8\x64\x77\xA8\x22\x5F\x7B\x74\xC4\xAB\x90\x0E\x33\x25\xF3\xEE\x8A\x3F\x9D\x0F\x5E\x02\x88\x74\xA4\xBE\x45\xCA\xB2\xD5\x24\x92\x6B\x1C\xDA\xD6\xB6\xEF\xA6\x9C\x55\xCC\x79\xB4\x1B\x6D\x5D\xDE\xD6\x54\xE0\xFC\x08\xF6\xD7\x8D\xD4\x59\x61\x94\xA0\x6C\x25\x9D\x77\x4F\x0F\xE1\x77\x8B\xD4\xD1\x3C\x66\x01\xE8\x7D\x96\xB5\xF5\x89\x4E\x32\x78\xD9\xFD\x56\x68\x61\x8C\xEF\x3E\x09\xFA\xC5\xAC\x42\x2C\xCA\xD7\x30\x65\xE1\xDB\x9E\x23\x9D\x88\xC6\xC6\x8F\x81\xC2\x7B\x5C\x9D\x0D\x3E\x50\x7A\x32\xCD\xE2\xD7\x60\xD3\x2C\x7E\x30\xF7\x9F\x00\x00\x00\xFF\xFF\xE2\xD9\x2A\x05\x48\x27\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.html")
	node.SetSize(2823)
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
		V:               []byte("\x1F\x8B\x08\x00\x00\x00\x00\x00\x00\xFF\x8C\x54\x41\x6F\xDB\x3C\x0C\x3D\x3B\xBF\x42\x9F\xBE\xEB\x5C\xA1\x18\x86\x0D\x9B\x6C\x60\xEB\x7A\x1B\xB0\x02\xED\x65\xA7\x81\x96\x99\x88\x98\x2C\xA9\x12\x9D\x36\xFF\x7E\xB0\x65\xA7\x59\x16\x0C\x3D\x49\x24\xF5\xC8\xF7\x5E\x18\xEB\xFF\xBE\x7E\xBF\x79\xF8\x71\x77\x2B\x2C\x0F\xAE\xDD\xE8\x72\x54\xDA\x22\xF4\xED\xA6\xAA\xF4\x80\x0C\xC2\x32\xC7\x1A\x1F\x47\xDA\x37\xF2\x26\x78\x46\xCF\xF5\xC3\x21\xA2\x14\xA6\x44\x8D\x64\x7C\x66\x35\xA1\x3F\x09\x63\x21\x65\xE4\x66\xE4\x6D\xFD\x41\x0A\xF5\xD2\xC7\xC3\x80\x8D\xDC\x13\x3E\xC5\x90\xF8\x04\xFD\x44\x3D\xDB\xA6\xC7\x3D\x19\xAC\xE7\xE0\x8D\x20\x4F\x4C\xE0\xEA\x6C\xC0\x61\x73\x7D\xA1\x11\x5B\x1C\xB0\x36\xC1\x85\x74\xD2\xEB\xFF\xB7\xEF\xDF\xDD\x7E\xFE\x32\xBF\x9F\x00\x4C\xEC\xB0\xD5\xAA\x9C\x53\xC6\x91\xFF\x25\x12\xBA\x46\x66\x3E\x38\xCC\x16\x91\xA5\xB0\x09\xB7\x8D\x54\xE4\x7B\x7C\xBE\x32\x39\x97\x81\x5A\x2D\x56\xE8\x2E\xF4\x87\x19\xDE\xD3\x5E\x18\x07\x39\x37\x92\x43\xEC\x20\xC9\x29\xFD\x47\x7E\x22\x03\xE4\x71\x29\x9D\x63\xEA\xA9\x27\xF9\xDD\x52\xAD\x34\xAC\xD3\x65\x6B\x88\x76\x41\x2B\x58\x80\xAA\xA7\xFD\xDF\x3D\x06\xF4\xE3\x11\xBC\x0D\x69\x58\x0B\xC4\x38\x48\x01\x86\x29\xF8\x46\xAA\x9F\xE4\x19\x93\x07\xA7\x32\x42\x32\x76\xC5\x54\x9A\x7C\x1C\x59\xF0\x21\x62\xF9\xED\xE4\xE2\xE9\xA3\x14\xD1\x81\x41\x1B\x5C\x8F\xA9\x91\xF7\x05\x57\xCC\x9F\x09\x4D\xE3\x5E\x4F\xEE\xA8\x2C\x8F\x9D\x6C\xEF\xC7\xEE\x82\xB4\xE3\x6D\xB9\x9C\x99\x1C\x61\x87\xAF\xB0\xD8\x5E\xB7\x5A\xD9\xEB\x4B\x7C\x18\x56\x3E\x0B\xAB\x2E\x1D\x05\xFD\x4B\x49\x84\x04\xBB\x04\x71\xF2\x4D\xC7\xF6\xC1\x52\x16\x94\x05\x78\x81\xCF\x30\x44\x87\x22\x6C\xD7\xC5\x13\xE4\x45\x1E\x3B\xD1\x53\x42\xC3\x21\x1D\xC4\x98\xC9\xEF\x84\x19\x33\x87\x41\xBC\x6C\xDA\x95\x56\xB1\xDD\x5C\x50\x5F\x6D\x4E\x82\x73\x13\xB6\x21\xF0\xAA\xF5\x2E\x3C\x61\xC2\x5E\x74\x07\xA1\x61\x26\x5D\x3C\x9E\xFE\xA6\xF9\xA3\x52\x3B\x62\x3B\x76\x57\x26\x0C\x2A\xDB\xF1\x9B\x05\xAF\xE6\xAD\x92\xD3\xDB\xA2\x72\x8E\xCB\x74\x38\x71\xBE\xD2\xAA\xAC\xB9\x56\xE5\x4B\xF0\x3B\x00\x00\xFF\xFF\xF9\x57\xD4\x9E\x21\x04\x00\x00"),
	}
	node.SetMode(420)
	node.SetName("index.html")
	node.SetSize(506)
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
