### Jawaban soal

- Simple Database Querying
```mysql
SELECT child.ID as ID, child.UserName AS UserName, 
    (SELECT parent.UserName from USER as parent where parent.ID=child.ID) AS ParentUserName
FROM USER as child WHERE Parent IS NOT NULL
```

- Logic Test
```
func findFirstStringInBracket(str string) string {
	var result string
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				result = string(runes[1 : indexClosingBracketFound-1])
			}
		}
	} 
	return result
}
```