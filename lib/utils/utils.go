package utils

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

var DryRun bool = false

func Chunk(items []string) (retval [][]string) {
	numCPUs := runtime.NumCPU()
	chunkSize := (len(items) + numCPUs - 1) / numCPUs

	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize

		if end > len(items) {
			end = len(items)
		}

		retval = append(retval, items[i:end])
	}

	return retval
}

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func ByteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func SecondsToHuman(input int) (result string) {
	years := math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
	seconds := input % (60 * 60 * 24 * 7 * 30 * 12)
	months := math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
	seconds = input % (60 * 60 * 24 * 7 * 30)
	weeks := math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
	seconds = input % (60 * 60 * 24 * 7)
	days := math.Floor(float64(seconds) / 60 / 60 / 24)
	seconds = input % (60 * 60 * 24)
	hours := math.Floor(float64(seconds) / 60 / 60)
	seconds = input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)
	seconds = input % 60

	result = ""
	if years > 0 {
		result += strconv.Itoa(int(years)) + "y"
	}
	if months > 0 {
		result += strconv.Itoa(int(months)) + "m"
	}
	if weeks > 0 {
		result += strconv.Itoa(int(weeks)) + "w"
	}
	if days > 0 {
		result += strconv.Itoa(int(days)) + "d"
	}
	if hours > 0 {
		result += strconv.Itoa(int(hours)) + "h"
	}
	if minutes > 0 {
		result += strconv.Itoa(int(minutes)) + "m"
	}
	if seconds > 0 {
		result += strconv.Itoa(int(seconds)) + "s"
	}

	return
}
