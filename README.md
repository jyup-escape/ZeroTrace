# ZeroTrace
Tools to delete files so that they cannot be recovered by file recovery software
# Verified Recovery Software
```
RStudio9
```
# how to use
Put the full path of the file you want to delete in the “” of the filePath on line 23 of main.go.
```
23  filePath := "sample.txt"
```
Run go run main.go in a terminal.
Now all you have to do is wait for it to be deleted.
```
$ go run main.go
```
