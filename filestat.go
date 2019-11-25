package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Files : type which is part of Filemetadata
type Files struct {
	Files []FileMetadata `json:"files"`
}

// ExtInfo : extension info
type ExtInfo struct {
	Extension      string `json:"extension"`
	NumOccurrences int64  `json:"num_occurrences"`
}

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

// FileMetadata : holds the files metadata
type FileMetadata struct {
	Path     string `json:"path"`      // the file's absolute path
	Size     int64  `json:"size"`      //the file size
	IsBinary bool   `json:"is_binary"` // whether the file is a binary file or a simple text file
}

// FileStats is File's file stats/metadatas
type FileStats struct {
	Numfiles        int64        `json:"num_files"`
	LargestFile     FileMetadata `json:"largest_file"`
	AverageFileSize float64      `json:"avg_file_size"`
	MostFrequentExt ExtInfo      `json:"most_frequent_ext"`
	TextPercentage  float32      `json:"text_percentage"`
	MostRecentPaths []string     `json:"most_recent_paths"`
}

// AddFile : adds a fle based on metadata
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

// this function returns the largest file received including name and size
func (f *Files) getLargestFile() FileMetadata {
	l := FileMetadata{Path: "", Size: 0}
	for _, i := range f.Files {
		if i.Size > l.Size {
			l = i
		}
	}
	return l
}

// this function returns the average file size which is (total_size / num_of_files)
func (f *Files) getAverageFile() float64 {
	var sum int64
	for _, i := range f.Files {
		sum += i.Size
	}
	return float64(sum) / float64(len(f.Files))
}

func (f *Files) MostFrequentExt() ExtInfo {
	m := make(map[string]int64)
	for _, i := range f.Files {
		m[filepath.Ext(i.Path)] += 1
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

// GetStats return File's stats
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
	f := Files{Files: []FileMetadata{}}
	f.AddFile("/tmp/tt.sh", "/tmp/aa.sh", "/tmp/bb.txt", "/tmp/cc.png", "/tmp/dd.png", "/tmp/ee.png", "/tmp/ff.jpeg")
	fmt.Printf("%#v\n", f.GetStats())
}
