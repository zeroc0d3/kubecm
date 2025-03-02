/*
Copyright © 2020 Guo Xudong

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
	"os"

	"github.com/sunny0826/kubecm/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure" // required for Azure
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"   // required for GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	baseCommand := cmd.NewBaseCommand()
	if err := baseCommand.CobraCmd().Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
