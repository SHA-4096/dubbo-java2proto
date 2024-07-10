/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package parser

import (
	"context"
	"os"
	"strings"
)

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/java"
)

type JavaParser struct {
	content []byte
	ast     *sitter.Node

	PackageName string
	Interfaces  []*JavaInterface
	Classes     []*JavaClass
}

type JavaInterface struct {
	Name    string
	Methods []*JavaMethod
}

type JavaClass struct {
	Name            string
	Fields          []*JavaField
	Methods         []*JavaMethod
	GetterMethodMap map[string]bool
}

type JavaField struct {
	Name string
	Type string

	Modifier string
}

type JavaMethod struct {
	Name string

	Modifier   string
	ReturnType string
	Args       []*Arg
}

type Arg struct {
	Name string
	Type string
}

func (jp *JavaParser) AddPackagePrefix(s string) string {
	return jp.PackageName + "." + s
}

func (jp *JavaParser) ParseFile() {
	for i := 0; i < int(jp.ast.ChildCount()); i++ {
		child := jp.ast.Child(i)
		typ := child.Type()

		switch typ {
		case PackageDecl:
			jp.parsePackage(child)
		case InterfaceDecl:
			jp.parseInterface(child)
		case ClassDecl:
			jp.parseClass(child)
		}
	}
}

func (jp *JavaParser) parsePackage(node *sitter.Node) {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		if typ == ScopedIdentifier {
			jp.PackageName = child.Content(jp.content)
		}
	}
}

func (jp *JavaParser) parseInterface(node *sitter.Node) {
	ji := new(JavaInterface)
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		// check if the interface is public
		// don't generate code for non-public interface
		switch typ {
		case Modifiers:
			if child.Content(jp.content) != "public" {
				break
			}
		case Identifier:
			identifier := child.Content(jp.content)
			ji.Name = identifier
		case InterfaceBody:
			jp.parseInterfaceBody(child, ji)
			jp.Interfaces = append(jp.Interfaces, ji)
		}
	}
}

func (jp *JavaParser) parseInterfaceBody(node *sitter.Node, ji *JavaInterface) {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		if typ == MethodDecl {
			m := jp.parseMethod(child)
			ji.Methods = append(ji.Methods, m)
		}
	}
}

func (jp *JavaParser) parseMethod(node *sitter.Node) *JavaMethod {
	jm := new(JavaMethod)

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		switch typ {
		case Modifiers:
			jm.Modifier = child.Content(jp.content)
		case Identifier:
			jm.Name = child.Content(jp.content)
		case TypeIdentifier:
			jm.ReturnType = child.Content(jp.content)
		case FormalParameters:
			jp.parseParams(child, jm)
		}
	}
	return jm
}

func (jp *JavaParser) parseParams(node *sitter.Node, jm *JavaMethod) {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		switch typ {
		case FormalParameter:
			jp.parseParam(child, jm)
		}
	}
}

func (jp *JavaParser) parseParam(node *sitter.Node, jm *JavaMethod) {
	var args []*Arg
	var a = new(Arg)
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		nodeType := child.Type()

		switch nodeType {
		case TypeIdentifier: // type
			// Integer, Float, Boolean, Character, Short, Long, Double, String, CustomType
			a.Type = child.Content(jp.content)
		case IntegralType, FloatingPointType, BooleanType:
			// basic type
			a.Type = child.Content(jp.content)
		case GenericType:
			// List, Map ...
			// TODO: handle generic type
		case Identifier:
			a.Name = child.Content(jp.content)
			args = append(args, a)
			a = new(Arg)
		}
	}
	jm.Args = args
}

func (jp *JavaParser) parseClass(node *sitter.Node) {
	var classes []*JavaClass
	var class = new(JavaClass)
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		switch typ {
		case Identifier:
			class.Name = child.Content(jp.content)
		case ClassBody:
			jp.parseClassBody(child, class)
			classes = append(classes, class)
			class = new(JavaClass)
		}
	}
	jp.Classes = append(jp.Classes, classes...)
}

func (jp *JavaParser) parseClassBody(node *sitter.Node, class *JavaClass) {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		switch typ {
		case FieldDecl:
			jf := jp.parseField(child)
			class.Fields = append(class.Fields, jf)
		case MethodDecl:
			jm := jp.parseMethod(child)
			class.Methods = append(class.Methods, jm)
		}
	}
	class.GetterMethodMap = parseGetterMethod(class.Methods)
}

func (jp *JavaParser) parseField(node *sitter.Node) *JavaField {
	var jf = new(JavaField)
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		typ := child.Type()

		switch typ {
		case Modifiers:
			jf.Modifier = child.Content(jp.content)
		case TypeIdentifier:
			jf.Type = child.Content(jp.content)
		case IntegralType, FloatingPointType, BooleanType:
			jf.Type = child.Content(jp.content)
		case GenericType:
			// TODO: handle generic type
		case VariableDecl:
			jf.Name = child.Content(jp.content)
		}
	}
	return jf
}

func parseGetterMethod(methods []*JavaMethod) map[string]bool {
	getterMethods := make(map[string]bool)
	for _, method := range methods {
		if strings.HasPrefix(method.Name, "get") {
			getterMethods[strings.TrimPrefix(method.Name, "get")] = true
		}
	}
	return getterMethods

}

func NewParser(filepath string) (*JavaParser, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	parser := sitter.NewParser()
	parser.SetLanguage(java.GetLanguage())

	tree, err := parser.ParseCtx(context.Background(), nil, bytes)
	if err != nil {
		return nil, err
	}

	jp := &JavaParser{
		content: bytes,
		ast:     tree.RootNode(),
	}

	return jp, nil
}
