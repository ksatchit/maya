/*
Copyright 2019 The OpenEBS Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openebs/maya/pkg/client/generated/clientset/versioned/typed/openebs.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOpenebsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeOpenebsV1alpha1) BackupCStors(namespace string) v1alpha1.BackupCStorInterface {
	return &FakeBackupCStors{c, namespace}
}

func (c *FakeOpenebsV1alpha1) BackupCStorLasts(namespace string) v1alpha1.BackupCStorLastInterface {
	return &FakeBackupCStorLasts{c, namespace}
}

func (c *FakeOpenebsV1alpha1) CASTemplates() v1alpha1.CASTemplateInterface {
	return &FakeCASTemplates{c}
}

func (c *FakeOpenebsV1alpha1) CStorPools() v1alpha1.CStorPoolInterface {
	return &FakeCStorPools{c}
}

func (c *FakeOpenebsV1alpha1) CStorVolumes(namespace string) v1alpha1.CStorVolumeInterface {
	return &FakeCStorVolumes{c, namespace}
}

func (c *FakeOpenebsV1alpha1) CStorVolumeReplicas(namespace string) v1alpha1.CStorVolumeReplicaInterface {
	return &FakeCStorVolumeReplicas{c, namespace}
}

func (c *FakeOpenebsV1alpha1) Disks() v1alpha1.DiskInterface {
	return &FakeDisks{c}
}

func (c *FakeOpenebsV1alpha1) RunTasks(namespace string) v1alpha1.RunTaskInterface {
	return &FakeRunTasks{c, namespace}
}

func (c *FakeOpenebsV1alpha1) StoragePools() v1alpha1.StoragePoolInterface {
	return &FakeStoragePools{c}
}

func (c *FakeOpenebsV1alpha1) StoragePoolClaims() v1alpha1.StoragePoolClaimInterface {
	return &FakeStoragePoolClaims{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOpenebsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}