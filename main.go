package main

import (
	"os"
	"path/filepath"
)

// Files : type which is part of Filemetadata
type Files struct {
	files []FileMetadata `json:"files"`
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

// util func that will return the max number from a map
func max(numbers map[int]bool) int {
	var maxNumber int
	for maxNumber = range numbers {
		break
	}
	for n := range numbers {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

// FileMetadata : holds the files metadata
type FileMetadata struct {
	Path     string `json:"path"`      // the file's absolute path
	Size     int64  `json:"size"`      //the file size
	IsBinary bool   `json:"is_binary"` // whether the file is a binary file or a simple text file
}

// FileStats is File's file stats/metadatas
type FileStats struct {
	Numfiles        int64           `json:"num_files"`
	LargestFile     LargestFileInfo `json:"largest_file"`
	AverageFileSize float64         `json:"avg_file_size"`
	MostFrequentExt ExtInfo         `json:"most_frequent_ext"`
	TextPercentage  float32         `json:"text_percentage"`
	MostRecentPaths []string        `json:"most_recent_paths"`
}

// LargestFileInfo : struct of largest file info
type LargestFileInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

// AddFile : adds a fle based on metadata
func (files *Files) AddFile(filename string) error {
	/*
		The function receives a structure containing the metadata of one file.
		This file should be taken into account when calculating statistics.
		The function can return an error if the input is invalid or processing of the file fails.
		Need to throw error here if input is invalid or processing of the file fails
	*/

	abspath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	stat, err := os.Stat(abspath)
	if err != nil {
		return err
	}
	files = append(files, FileMetadata{Path: abspath, Size: state.Size()})
}

// this function returns the largest file received including name and size
func (files *Files) getLargestFile() string {
	l := FileMetadata{Path: "", Size: 0}
	for _, i := range files {
		if i.Size > l.Size {
			latest = i
		}
	}
	return l
}

// this function returns the average file size which is (total_size / num_of_files)
func (file *Files) getAverageFile() int64 {
	var sum int64
	for _, i := range file {
		sum += i.Size
	}
	return sum / len(file)
}

func (file *Files) MostFrequentExt() GetExtInfo {
	m := make(map[string]int64)
	for _, i := range file {
		m[filepath.Ext(i.Path)] += 1
	}
	// p :=
	// for i, j := range m {
	// 	if i > j
	// }
	// need to make it simple :rajibmitra
	return max(m)
}

// GetStats return File's stats
func (files *Files) GetStats() FileStats {

	/*
		This function returns statistics for all files added until that point. The following statistics should be returned:
		Number of files received
		Largest file received (including name and size)
		Average file size
		Most frequent file extension (including number of occurences)
		Percentage of text files of all files received
		List of latest 10 file paths received
	*/

	filestats := FileStats{
		Num:                len(files),
		GetLargestFileInfo: getLargestFile(files),
		AverageFileSize:    getAverageFile(files),
		MostFreqExt:        MostFrequentExt(files),
		TextPercentage:     getTextPercentage(files),
		MostRecentPaths:    getMostRecentPath(files),

		//rajibmitra: need to implement https://stackoverflow.com/questions/17133590/how-to-get-file-length-in-go

	}
	return &filestats, nil

}
