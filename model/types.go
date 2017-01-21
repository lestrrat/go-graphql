package model

func NewSchema() *Schema {
	return &Schema{}
}

func (s Schema) Query() *ObjectDefinition {
	return s.query
}

func (s Schema) Types() chan *ObjectDefinition {
	return s.types.Iterator()
}

func (s *Schema) SetQuery(q *ObjectDefinition) {
	s.query = q
}

func (s *Schema) AddTypes(list ...*ObjectDefinition) {
	s.types.Add(list...)
}

func NewNamedType(name string) NamedType {
	return &namedType{
		nameComponent: nameComponent(name),
		nullable:      true,
	}
}

func NewListType(t Type) *ListType {
	return &ListType{
		nullable:      true,
		typeComponent: typeComponent{typ: t},
	}
}

func NewObjectFieldArgumentDefinition(name string, typ Type) *ObjectFieldArgumentDefinition {
	return &ObjectFieldArgumentDefinition{
		nameComponent: nameComponent(name),
		typeComponent: typeComponent{typ: typ},
	}
}

func NewObjectDefinition(name string) *ObjectDefinition {
	return &ObjectDefinition{
		nameComponent: nameComponent(name),
		nullable: nullable(true),
	}
}

func (t ObjectDefinition) Fields() chan *ObjectFieldDefinition {
	return t.fields.Iterator()
}

func (t *ObjectDefinition) AddFields(list ...*ObjectFieldDefinition) {
	t.fields.Add(list...)
}

func NewObjectFieldDefinition(name string, typ Type) *ObjectFieldDefinition {
	return &ObjectFieldDefinition{
		nameComponent: nameComponent(name),
		typeComponent: typeComponent{typ: typ},
	}
}

func (t *ObjectDefinition) SetImplements(typ NamedType) {
	t.hasImplements = true
	t.implements = typ
}

func (t ObjectDefinition) HasImplements() bool {
	return t.hasImplements
}

func (t ObjectDefinition) Implements() NamedType {
	return t.implements
}

func (t *ObjectFieldDefinition) AddArguments(list ...*ObjectFieldArgumentDefinition) {
	t.arguments.Add(list...)
}

func (t ObjectFieldDefinition) Arguments() chan *ObjectFieldArgumentDefinition {
	return t.arguments.Iterator()
}

func NewEnumDefinition(name string) *EnumDefinition {
	return &EnumDefinition{
		nameComponent: nameComponent(name),
		nullable:      nullable(true),
	}
}

func (t *EnumDefinition) AddElements(list ...*EnumElementDefinition) {
	t.elements.Add(list...)
}

func (t *EnumDefinition) Elements() chan *EnumElementDefinition {
	return t.elements.Iterator()
}

func NewEnumElementDefinition(name string, value Value) *EnumElementDefinition {
	return &EnumElementDefinition{
		nameComponent:  nameComponent(name),
		valueComponent: valueComponent{value: value},
	}
}

func NewInterfaceDefinition(name string) *InterfaceDefinition {
	return &InterfaceDefinition{
		nullable:      nullable(true),
		nameComponent: nameComponent(name),
	}
}

type Resolver interface {
	Resolve(interface{}) Type
}

func (iface *InterfaceDefinition) SetTypeResolver(v Resolver) {}
func (iface *InterfaceDefinition) TypeResolver() Resolver     { return nil }

func (iface InterfaceDefinition) Fields() chan *InterfaceFieldDefinition {
	return iface.fields.Iterator()
}

func (iface *InterfaceDefinition) AddFields(list ...*InterfaceFieldDefinition) {
	iface.fields.Add(list...)
}

func NewInterfaceFieldDefinition(name string, typ Type) *InterfaceFieldDefinition {
	return &InterfaceFieldDefinition{
		nameComponent: nameComponent(name),
		typeComponent: typeComponent{typ: typ},
	}
}

func (f *InterfaceFieldDefinition) Type() Type {
	return f.typ
}

func NewUnionDefinition(name string) *UnionDefinition {
	return &UnionDefinition{
		nameComponent: nameComponent(name),
	}
}

func (def UnionDefinition) Types() chan Type {
	return def.types.Iterator()
}
func (def *UnionDefinition) AddTypes(list ...Type) {
	def.types.Add(list...)
}

func NewInputDefinition(name string) *InputDefinition {
	return &InputDefinition{
		nameComponent: nameComponent(name),
	}
}

func NewInputFieldDefinition(name string) *InputFieldDefinition {
	return &InputFieldDefinition{
		nameComponent: nameComponent(name),
	}
}

func (def *InputDefinition) AddFields(list ...*InputFieldDefinition) {
	def.fields.Add(list...)
}

func (def InputDefinition) Fields() chan *InputFieldDefinition {
	return def.fields.Iterator()
}
