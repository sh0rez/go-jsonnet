/*
Copyright 2019 Google Inc. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package formatter

import "github.com/sh0rez/go-jsonnet/internal/formatter"

// Options is a set of parameters that control the reformatter's behaviour.
type Options = formatter.Options

// DefaultOptions returns the recommended formatter behaviour.
func DefaultOptions() Options {
	return formatter.DefaultOptions()
}

// Format returns code that is equivalent to its input but better formatted
// according to the given options.
func Format(filename string, input string, options Options) (string, error) {
	return formatter.Format(filename, input, options)
}
