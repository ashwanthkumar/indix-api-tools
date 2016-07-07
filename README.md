# Indix API Tools
Set of (CLI) tools to make working with Indix API easy.

### JsonToCSV
`json2csv` helps you convert the bulk output's json format to a CSV / TSV file. Example

```json
{"minSalePrice":12.99,"maxSalePrice":50.64,"offersCount":36,"storesCount":9,"lastRecordedAt":1467784679225,"countryCode":"US","mpid":"0c7b2d9462f910b3987e84295d0a6a68","currency":"USD","upcs":["UPC1","UPC2","UPC3"],"title":"Women's Sandals gold","brandId":4722,"categoryIdPath":"10179 > 16238 > 20822","brandName":"Ellie","categoryId":20822,"categoryName":"Sandals","mpns":["null","014-MIRIAM - M"],"imageUrl":"http://images10.newegg.com/ProductImage/productId.jpg","categoryNamePath":"Shoes > Women > Sandals"}
```

The output file generated by the program would be something like
```csv
MinSalePrice,MaxSalePrice,OffersCount,StoresCount,LastRecordedAt,CountryCode,Mpid,Currency,Upcs,Title,BrandID,CategoryIDPath,BrandName,CategoryID,CategoryName,Mpns,ImageURL,CategoryNamePath
12.990000,50.640000,36,9,1467784679225,US,0c7b2d9462f910b3987e84295d0a6a68,USD,UPC1|UPC2|UPC3,"Women's Sandals gold",4722,10179 > 16238 > 20822,Ellie,20822,Sandals,null|014-MIRIAM - M,http://images10.newegg.com/ProductImage/productId.jpg,Shoes > Women > Sandals
```

#### Usage
Download the latest version of `json2csv` from the releases based on your operating system.

```
Used to convert Bulk Job's JSON output to CSV

Usage:
  json2csv <input.json> <output.csv> [flags]

Flags:
      --field-delimiter string        Field delimiters for the output file (default ",")
      --multivalue-delimiter string   Value delimiters for multivalued fields (default "|")
```

#### Example
Once you download the binary (rename it to json2csv), you run the command as
```
$ ./json2csv /path/to/bulk-output.json /path/to/output.csv
```

## License
http://www.apache.org/licenses/LICENSE-2.0