package main

import (
	"fmt"
	"strconv"
	"strings"
)

// BulkOutput represents the output format from bulk jobs
type BulkOutput struct {
	MinSalePrice     float64  `json:"minSalePrice"`
	MaxSalePrice     float64  `json:"maxSalePrice"`
	OffersCount      int      `json:"offersCount"`
	StoresCount      int      `json:"storesCount"`
	LastRecordedAt   int64    `json:"lastRecordedAt"`
	CountryCode      string   `json:"countryCode"`
	Mpid             string   `json:"mpid"`
	Currency         string   `json:"currency"`
	Upcs             []string `json:"upcs"`
	Title            string   `json:"title"`
	BrandID          int      `json:"brandId"`
	CategoryIDPath   string   `json:"categoryIdPath"`
	BrandName        string   `json:"brandName"`
	CategoryID       int      `json:"categoryId"`
	CategoryName     string   `json:"categoryName"`
	Mpns             []string `json:"mpns"`
	ImageURL         string   `json:"imageUrl"`
	CategoryNamePath string   `json:"categoryNamePath"`
}

// BulkOutputHeader returns the all the headers of the record separated by fieldDelimiter
func BulkOutputHeader(fieldDelimiter string) string {
	return strings.Join(
		[]string{
			"MinSalePrice",
			"MaxSalePrice",
			"OffersCount",
			"StoresCount",
			"LastRecordedAt",
			"CountryCode",
			"Mpid",
			"Currency",
			"Upcs",
			"Title",
			"BrandID",
			"CategoryIDPath",
			"BrandName",
			"CategoryID",
			"CategoryName",
			"Mpns",
			"ImageURL",
			"CategoryNamePath",
		}, fieldDelimiter)
}

// ToRow converts the bulk output object to a row which can then written out to a file or STDOUT
func (b *BulkOutput) ToRow(multivalueDelimiter, fieldDelimiter string) string {
	return strings.Join(
		[]string{
			F64toa(b.MinSalePrice),
			F64toa(b.MaxSalePrice),
			Itoa(b.OffersCount),
			Itoa(b.StoresCount),
			I64toa(b.LastRecordedAt),
			b.CountryCode,
			b.Mpid,
			b.Currency,
			strings.Join(b.Upcs, multivalueDelimiter),
			fmt.Sprintf("\"%s\"", b.Title),
			Itoa(b.BrandID),
			b.CategoryIDPath,
			b.BrandName,
			Itoa(b.CategoryID),
			b.CategoryName,
			strings.Join(b.Mpns, multivalueDelimiter),
			b.ImageURL,
			b.CategoryNamePath},
		fieldDelimiter)
}

// F64toa converts float64 to string
func F64toa(number float64) string {
	return fmt.Sprintf("%f", number)
}

// Ftoa converts float64 to string
func Ftoa(number float32) string {
	return fmt.Sprintf("%f", number)
}

// Itoa converts int to string
func Itoa(number int) string {
	return strconv.Itoa(number)
}

// I64toa converts int64 to string
func I64toa(number int64) string {
	return fmt.Sprintf("%d", number)
}
