package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `a.js`,
		FileModTime: time.Unix(1461760498, 0),
		Content:     string([]byte{0x28, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x20, 0x3d, 0x20, 0x27, 0x23, 0x23, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x23, 0x23, 0x27, 0x3b, 0xa, 0xa, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x71, 0x73, 0x20, 0x3d, 0x20, 0x5b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x27, 0x75, 0x72, 0x6c, 0x3d, 0x27, 0x20, 0x2b, 0x20, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x28, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x28, 0x29, 0x29, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x27, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x72, 0x3d, 0x27, 0x20, 0x2b, 0x20, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x28, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x29, 0x2c, 0xa, 0x20, 0x20, 0x5d, 0x2e, 0x6a, 0x6f, 0x69, 0x6e, 0x28, 0x27, 0x26, 0x27, 0x29, 0x3b, 0xa, 0xa, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x71, 0x20, 0x3d, 0x20, 0x5b, 0x5d, 0x3b, 0xa, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x3d, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x61, 0x2c, 0x20, 0x76, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x71, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x28, 0x5b, 0x61, 0x2c, 0x20, 0x76, 0x5d, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x7d, 0xa, 0xa, 0x20, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x28, 0x27, 0x76, 0x69, 0x65, 0x77, 0x27, 0x29, 0x3b, 0xa, 0xa, 0x20, 0x20, 0x74, 0x72, 0x79, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x73, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x57, 0x65, 0x62, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x28, 0x27, 0x77, 0x73, 0x3a, 0x2f, 0x2f, 0x31, 0x32, 0x37, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x3a, 0x35, 0x30, 0x30, 0x31, 0x2f, 0x61, 0x2f, 0x77, 0x73, 0x3f, 0x27, 0x20, 0x2b, 0x20, 0x71, 0x73, 0x29, 0x3b, 0xa, 0xa, 0x20, 0x20, 0x20, 0x20, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x28, 0x27, 0x6f, 0x70, 0x65, 0x6e, 0x27, 0x2c, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x3d, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x61, 0x2c, 0x20, 0x76, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x28, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x66, 0x79, 0x28, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x73, 0x3a, 0x20, 0x76, 0x20, 0x7c, 0x7c, 0x20, 0x7b, 0x7d, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x29, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x3b, 0xa, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x71, 0x2e, 0x66, 0x6f, 0x72, 0x45, 0x61, 0x63, 0x68, 0x28, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x61, 0x29, 0x20, 0x7b, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x28, 0x61, 0x5b, 0x30, 0x5d, 0x2c, 0x20, 0x61, 0x5b, 0x31, 0x5d, 0x29, 0x3b, 0x20, 0x7d, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x7d, 0x20, 0x63, 0x61, 0x74, 0x63, 0x68, 0x20, 0x28, 0x65, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x3d, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x61, 0x2c, 0x20, 0x76, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x20, 0x78, 0x68, 0x72, 0x20, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x58, 0x4d, 0x4c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x78, 0x68, 0x72, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x28, 0x27, 0x50, 0x4f, 0x53, 0x54, 0x27, 0x2c, 0x20, 0x27, 0x2f, 0x61, 0x2f, 0x65, 0x76, 0x3f, 0x27, 0x20, 0x2b, 0x20, 0x71, 0x73, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x78, 0x68, 0x72, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x28, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x69, 0x66, 0x79, 0x28, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x61, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x76, 0x61, 0x72, 0x73, 0x3a, 0x20, 0x76, 0x20, 0x7c, 0x7c, 0x20, 0x7b, 0x7d, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x29, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x3b, 0xa, 0xa, 0x20, 0x20, 0x20, 0x20, 0x71, 0x2e, 0x66, 0x6f, 0x72, 0x45, 0x61, 0x63, 0x68, 0x28, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x61, 0x29, 0x20, 0x7b, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x28, 0x61, 0x5b, 0x30, 0x5d, 0x2c, 0x20, 0x61, 0x5b, 0x31, 0x5d, 0x29, 0x3b, 0x20, 0x7d, 0x29, 0x3b, 0xa, 0xa, 0x20, 0x20, 0x20, 0x20, 0x28, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x70, 0x69, 0x6e, 0x67, 0x28, 0x29, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x28, 0x27, 0x70, 0x69, 0x6e, 0x67, 0x27, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x28, 0x70, 0x69, 0x6e, 0x67, 0x2c, 0x20, 0x31, 0x30, 0x30, 0x30, 0x20, 0x2a, 0x20, 0x33, 0x30, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x29, 0x28, 0x29, 0x3b, 0xa, 0x20, 0x20, 0x7d, 0xa, 0x7d, 0x29, 0x28, 0x29, 0x3b, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1461760083, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // a.js

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`public`, &embedded.EmbeddedBox{
		Name: `public`,
		Time: time.Unix(1461760928, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"a.js": file2,
		},
	})
}