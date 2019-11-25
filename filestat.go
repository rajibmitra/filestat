package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Files represents a list of files with their metadata
type Files struct {
	Files []FileMetadata `json:"files"`
}

// ExtInfo represents informations about an extension
type ExtInfo struct {
	Extension      string `json:"extension"`
	NumOccurrences int64  `json:"num_occurrences"`
}

// FileMetadata represent a file with its metadata
type FileMetadata struct {
	Path     string `json:"path"`      // the file's absolute path
	Size     int64  `json:"size"`      //the file size
	IsBinary bool   `json:"is_binary"` // whether the file is a binary file or a simple text file
}

// FileStats represents statistics about files
type FileStats struct {
	Numfiles        int64        `json:"num_files"`
	LargestFile     FileMetadata `json:"largest_file"`
	AverageFileSize float64      `json:"avg_file_size"`
	MostFrequentExt ExtInfo      `json:"most_frequent_ext"`
	TextPercentage  float32      `json:"text_percentage"`
	MostRecentPaths []string     `json:"most_recent_paths"`
}

// AddFile adds a file in a Files list
func (f *Files) AddFile(filename ...string) error {
	for _, i := range filename {
		abspath, err := filepath.Abs(i)
		if err != nil {
			return err
		}
		stat, err := os.Stat(abspath)
		if err != nil {
			return err
		}
		f.Files = append(f.Files, FileMetadata{Path: abspath, Size: stat.Size()})
	}
	return nil
}

// getLargestFile retrieves largest file in list
func (f *Files) getLargestFile() FileMetadata {
	l := FileMetadata{Path: "", Size: 0}
	for _, i := range f.Files {
		if i.Size > l.Size {
			l = i
		}
	}
	return l
}

// getAverageFile retrieves average size of files in list
func (f *Files) getAverageFile() float64 {
	var sum int64
	for _, i := range f.Files {
		sum += i.Size
	}
	return float64(sum) / float64(len(f.Files))
}

// MostFrequentExt returns most frequent extension
func (f *Files) MostFrequentExt() ExtInfo {
	m := make(map[string]int64)
	for _, i := range f.Files {
		m[filepath.Ext(i.Path)]++
	}
	p := ExtInfo{}
	for i, j := range m {
		if j > p.NumOccurrences {
			p.Extension = i
			p.NumOccurrences = j
		}
	}
	return p
}

// GetStats returns statistics about files in list
func (f *Files) GetStats() FileStats {
	filestats := FileStats{
		Numfiles:        int64(len(f.Files)),
		LargestFile:     f.getLargestFile(),
		AverageFileSize: f.getAverageFile(),
		MostFrequentExt: f.MostFrequentExt(),
		// TextPercentage:  f.getTextPercentage(),
		// MostRecentPaths: f.getMostRecentPath(),
	}
	return filestats
}

func main() {
	//example
	// f := Files{Files: []FileMetadata{}}
	// f.AddFile("/tmp/tt.sh", "/tmp/aa.sh", "/tmp/bb.txt", "/tmp/cc.png", "/tmp/dd.png", "/tmp/ee.png", "/tmp/ff.jpeg")
	// fmt.Printf("%#v\n", f.GetStats())

	textPtr := flag.String("Addfile", "", "Files to add, (Required)")

	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1) 
	}else  *textPtr != nil  {
		f := Files{Files: []FileMetadata{}}
		f.AddFile(*textPtr)
	}
}

	fmt.Printf("textPtr: %s \n", *textPtr)
}
