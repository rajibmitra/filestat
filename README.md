# filestat

a library that has basically two func

# func AddFile(metadata FileMetadata) error

The function receives a structure containing the metadata of one file. This file should be taken into account
when calculating statistics. The function can return an error if the input is invalid or processing of the file fails.

# func GetStats() FileStats

This function returns statistics for all files added until that point. The following statistics should be returned:
Number of files received
Largest file received (including name and size)
Average file size
Most frequent file extension (including number of occurences)
Percentage of text files of all files received
List of latest 10 file paths received
