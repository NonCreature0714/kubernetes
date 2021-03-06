/*
Copyright 2016 The Kubernetes Authors.

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

package filters

import (
	"net/http"

	"k8s.io/kubernetes/pkg/apiserver/request"
	"k8s.io/kubernetes/pkg/util/sets"
)

// LongRunningRequestCheck is a predicate which is true for long-running http requests.
type LongRunningRequestCheck func(r *http.Request, requestInfo *request.RequestInfo) bool

// BasicLongRunningRequestCheck returns true if the given request has one of the specified verbs or one of the specified subresources
func BasicLongRunningRequestCheck(longRunningVerbs, longRunningSubresources sets.String) LongRunningRequestCheck {
	return func(r *http.Request, requestInfo *request.RequestInfo) bool {
		if longRunningVerbs.Has(requestInfo.Verb) {
			return true
		}
		if requestInfo.IsResourceRequest && longRunningSubresources.Has(requestInfo.Subresource) {
			return true
		}
		return false
	}
}
