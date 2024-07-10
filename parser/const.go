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

const (
	InterfaceDecl   = "interface_declaration"
	PackageDecl     = "package_declaration"
	MethodDecl      = "method_declaration"
	ClassDecl       = "class_declaration"
	FieldDecl       = "field_declaration"
	ConstructorDecl = "constructor_declaration"
	VariableDecl    = "variable_declarator"

	Identifier       = "identifier"
	TypeIdentifier   = "type_identifier"
	ScopedIdentifier = "scoped_identifier"

	ClassBody = "class_body"

	Modifiers     = "modifiers"
	InterfaceBody = "interface_body"

	FormalParameters = "formal_parameters"
	FormalParameter  = "formal_parameter"

	IntegralType      = "integral_type"       // byte, short, int, long, char
	FloatingPointType = "floating_point_type" // float, double
	BooleanType       = "boolean_type"
	GenericType       = "generic_type" // List, Map ...
)
