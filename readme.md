## Gencoder
Gencoder is a command line tool for encoding and decoding text using funny encodings. 
It is a homework project for Yalantis Golang School


## Test
```
go test ./encodings/
```

## Usage
#### Encode
to encode string run next command
```
go run main.go encode -e <encoding> <text>
```
Supported encodings:
 - numbers
 - piglatin
 
#### Decode
to decode string run next command
```
go run main.go decode -e <encoding> <text>
```
Supported encodings:
 - numbers
 
#### Examples
```
go run main.go encode -e numbers hello world
go run main.go decode -e numbers h2ll4 w4rld
go run main.go encode -e piglatin Hello. How are you. I am fine!
```