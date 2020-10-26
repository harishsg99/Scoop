package repl

import {
    "bufio"
    "fmt"
    "io"
    "scoopml/lexer"
    "scoopml/token"
}

const PROMPT = ">>"

func start(in io.Reader  , out io.Writter){
  scanner := bufio.NewScanner(in)
   for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()
    if !scanned {
    return
} 
line := scanner.Text()
l := lexer.New(line)
for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
fmt.Printf("%+v\n", tok)

} 
}
}
