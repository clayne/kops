/*
Copyright 2017 The Kubernetes Authors.

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

package cmd

import (
	"fmt"
	"net/http"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/kops/util/pkg/vfs"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type ChannelsFactory struct {
	configFlags      genericclioptions.ConfigFlags
	cachedRESTConfig *rest.Config
	cachedHTTPClient *http.Client
	vfsContext       *vfs.VFSContext
	restMapper       *restmapper.DeferredDiscoveryRESTMapper
	dynamicClient    dynamic.Interface
}

func NewChannelsFactory() *ChannelsFactory {
	return &ChannelsFactory{}
}

func (f *ChannelsFactory) RESTConfig() (*rest.Config, error) {
	if f.cachedRESTConfig == nil {
		clientGetter := genericclioptions.NewConfigFlags(true)

		restConfig, err := clientGetter.ToRESTConfig()
		if err != nil {
			return nil, fmt.Errorf("cannot load kubecfg settings: %w", err)
		}

		restConfig.UserAgent = "kops"
		restConfig.Burst = 50
		restConfig.QPS = 20
		f.cachedRESTConfig = restConfig
	}
	return f.cachedRESTConfig, nil
}

func (f *ChannelsFactory) HTTPClient() (*http.Client, error) {
	if f.cachedHTTPClient == nil {
		restConfig, err := f.RESTConfig()
		if err != nil {
			return nil, err
		}
		httpClient, err := rest.HTTPClientFor(restConfig)
		if err != nil {
			return nil, fmt.Errorf("getting http client: %w", err)
		}
		f.cachedHTTPClient = httpClient
	}
	return f.cachedHTTPClient, nil
}

func (f *ChannelsFactory) RESTMapper() (*restmapper.DeferredDiscoveryRESTMapper, error) {
	if f.restMapper == nil {
		discoveryClient, err := f.configFlags.ToDiscoveryClient()
		if err != nil {
			return nil, err
		}

		restMapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)

		f.restMapper = restMapper
	}

	return f.restMapper, nil
}

func (f *ChannelsFactory) DynamicClient() (dynamic.Interface, error) {
	if f.dynamicClient == nil {
		restConfig, err := f.RESTConfig()
		if err != nil {
			return nil, err
		}
		httpClient, err := f.HTTPClient()
		if err != nil {
			return nil, err
		}
		dynamicClient, err := dynamic.NewForConfigAndClient(restConfig, httpClient)
		if err != nil {
			return nil, err
		}
		f.dynamicClient = dynamicClient
	}
	return f.dynamicClient, nil
}

func (f *ChannelsFactory) VFSContext() *vfs.VFSContext {
	if f.vfsContext == nil {
		// TODO vfs.NewVFSContext()
		f.vfsContext = vfs.Context
	}
	return f.vfsContext
}
