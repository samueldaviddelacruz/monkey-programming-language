package compiler

type SymbolScope string

const (
	GlobalScope SymbolScope = "GLOBAL"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func (st *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Scope: GlobalScope, Index: st.numDefinitions}
	st.store[name] = symbol
	st.numDefinitions++
	return symbol
}

func (st *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := st.store[name]
	return obj, ok
}
