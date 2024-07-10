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

package util

// JavaTypeToProtoType converts Java type to Proto type
// return prototype name and whether it is a wrapper type
func JavaTypeToProtoType(javaType string) (string, bool) {
	// TODO: provide extend for byte, short .etc?
	switch javaType {
	//case "byte":
	//	return "int32", false
	//case "short":
	//	return "int32", false
	case "int":
		return "int32", false
	case "long":
		return "int64", false
	case "char":
		return "string", false
	case "float":
		return "float", false
	case "double":
		return "double", false
	case "boolean":
		return "bool", false
	case "Integer":
		return "int32", true
	case "Float":
		return "float", true
	case "Boolean":
		return "bool", true
	//case "Character":
	//case "Short":
	case "Long":
		return "int64", true
	case "Double":
		return "double", true
	case "String":
		return "string", false
	}
	return javaType, false
}
