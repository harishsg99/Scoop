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

)

type Boolean struct { 
    Value bool
}

type Null struct{}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string { return "null" }