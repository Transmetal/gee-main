## 🔖 Introduction
Gee is tool of stdin to each files and stdout. It is similar to the tee command, but there are more functions for convenience. In addition, it was written as go. which provides output to stdout and files. In this process, it has various processing functions for lines such as replace, prefix, and suffix, so it can be used as a pipeline configuration or as a utility for testing. For more information, see the usage and case of study below!

## 🚀 Installation
### From source
```
go install -v github.com/hahwul/gee@latest
```

### With homebrew (only macos)
```
brew tap hahwul/gee
brew install gee
```
### Download from [release](https://github.com/hahwul/gee/releases) page. (macos,linux,freebsd,windows)
```
wget https://github.com/hahwul/gee/releases/download/v1.0.4/gee_1.0.4_linux_amd64.tar.gz
tar -xvf gee_1.0.4_linux_amd64.tar.gz
cp ./gee /usr/bin
```

## ☄️ Usage
```
▶ ~/go/bin/gee -h (if you install from source)
▶ gee -h
```
```
Usage: ./gee [flags] [file1] [file2] ...
(If you do not specify a file, only stdout is output)

Flags:
  -append
        Append mode for files
  -chunked int
        Chuked files from line (e.g output / output_1 / output_2)
  -debug
        Show debug message!
  -distribute
        Distribution to files
  -find string
        Find string in line (colorize red)
  -format string
        Change output format (json, md-table, html-table) (default "line")
  -grep string
        Greping with Regular Expression (like grep)
  -grepv string
        Inverse greping with Regular Expression (like grep -v)
  -inject string
        Inject stdin into the format of the factor value (e.g: -inject='This is %%INJECT%% line!')
  -prefix string
        Prefix string
  -replace string
        Replace string in line with '-find' option
  -reverse
        Reverse string in line
  -rmnl
        Remove newline(\r\n)
  -split string
        Split string within line. (to line , to table, to md-table)
  -suffix string
        Suffix string
  -uncolor
        Uncolorize stdout
  -uniq
        Remove duplicated line
  -version
        Version of gee
  -with-lc
        With letters count (colorize magenta)
  -with-line
        With line number (colorize blue)
  -with-time
        With timestamp (colorize green)
```

## 📚 Case of Study
### gee with prefix and suffix
```
▶ cat urls | gee -prefix "curl -i -k " -suffix " -H 'Auth: abcd'" curls.sh
```
```
curl -i -k https://www.hahwul.com/?q=123 -H 'Auth: abcd'
curl -i -k http://testphp.vulnweb.com/listproducts.php?cat=asdf&ff=1 -H 'Auth: abcd'
curl -i -k https://xss-game.appspot.com/level1/frame  -H 'Auth: abcd'
```
### Find and replace
```
▶ cat raw.txt | gee -find keep-alive
▶ cat raw.txt | gee -find keep-alive -replace close
```

### Specify the maximum length of the file and save it in multiple files.
```
▶ wc -l http.txt
2278

▶ cat http.txt | gee -chunked 500 output
```

### Distribute each line sequentially to multiple files.
```
▶ wc -l http.txt
2278

▶ cat http.txt | gee -distribute alice.txt bob.txt charlie.txt
```
