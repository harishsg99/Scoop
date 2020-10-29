package object

type ObjectType string

type object interface {
    Type() ObjectType
    Inspect() string
}
import{
    "fmt"
}
type Integer struct {
    Value int64
}
const(
    BOOLEAN_OBJ = "BOOLEAN"
    NULL_OBJ = "NULL"
    RETURN_VALUE_OBJ = "RETURN_VALUE"
    ERROR_OBJ = "ERROR"
    FUNCTION_OBJ = "FUNCTION"
)

type Boolean struct { 
    Value bool
}

type Null struct{}
type ReturnValue struct { 
    Value Object
}
type Error struct { 
    Message string
}
type Function struct { 
Parameters []*ast.Identifier
Body *ast.BlockStatement 
Env *Environment
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
func (f *Function) type() ObjectType{ return FUNCTION_OBJ }
func (b *Function) Inspect() string{
    var out bytes.Buffer
    Params := []string{}

    for _, p:= range f.Parameters {
        params = append(params, p.String()) 
    }
    out.WriteString("fn")
    out.WriteString("(") 
    out.WriteString(strings.Join(params, ", ")) 
    out.WriteString(") {\n") 
    out.WriteString(f.Body.String()) 
    out.WriteString("\n}")
    
    return out.String() 
}