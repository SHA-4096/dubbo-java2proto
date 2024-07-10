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

package main

import (
	"flag"
)

import (
	"github.com/dubbogo/dubbo-java2proto/generator"
	"github.com/dubbogo/dubbo-java2proto/parser"
)

var (
	// TODO: parse dir
	filepath = flag.String("file", "", "the path to the java file to be parsed")
	output   = flag.String("output", "", "the directory of the proto file to be generated")
)

func main() {
	flag.Parse()
	if *filepath == "" {
		panic("file path is empty")
	}

	jp, err := parser.NewParser(*filepath)
	if err != nil {
		panic(err)
	}
	jp.ParseFile()

	gen := generator.NewGenerator(*filepath, *output)
	err = gen.GenProto(jp)
	if err != nil {
		panic(err)
	}
}
