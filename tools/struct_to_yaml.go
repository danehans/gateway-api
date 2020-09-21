/*

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
package main

import (
	"fmt"
	//"io"
	"os"
	"path"
	"path/filepath"

	"github.com/ghodss/yaml"
)

func findTypes(dirPath string) []string {
	var filenames []string

	for _, pattern := range []string{path.Join(dirPath, "*.go")} {
		fmt.Printf("pattern: %s\n", pattern)
		matches, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Errorf("malformed pattern: %v", err)
			os.Exit(1)
		}
		for _, m := range matches {
			fmt.Printf("matched: %s\n", m)
		}

		filenames = append(filenames, matches...)
	}

	return filenames
}

func encodeType(filename string) []byte {
	//parts := [][]byte{}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open file %s: %w", filename, err)
		os.Exit(1)
	}

	defer f.Close()

	y, err := yaml.Marshal(f)
	if err != nil {
		fmt.Errorf("failed to marshal file %s: %w", filename, err)
		os.Exit(1)
	}
	return y
	/*for {
		buf := make([]byte, 4096)
		nread, err := splitter.Read(buf)
		switch err {
		case nil:
			parts = append(parts, buf[:nread])
		case io.EOF:
			return parts
		default:
			t.Fatalf("failed to read YAML from %q: %s", filename, err)
		}
	}*/
}

func main() {
	// Marshal a Person struct to YAML.
	for _, filename := range findTypes("../apis/v1alpha1/") {
		for buf := range encodeType(filename) {
			fmt.Println(string(buf))
		}
	}
}
