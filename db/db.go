// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package db

import "github.com/app-net-interface/awi-infra-guard/types"

const (
	vpcTable               = "vpcs"
	instanceTable          = "instances"
	subnetTable            = "subnets"
	clusterTable           = "clusters"
	podTable               = "pods"
	kubernetesServiceTable = "kubernetes_services"
	kubernetesNodeTable    = "kubernetes_nodes"
	namespaceTable         = "namespaces"
	accountTable           = "accounts"
	routeTableTable        = "route_tables"
	aclTable               = "acls"
	securityGroupTable     = "security_groups"
	syncTimeTable          = "sync_time"
)

var tableNames = []string{
	vpcTable,
	instanceTable,
	subnetTable,
	clusterTable,
	podTable,
	kubernetesServiceTable,
	kubernetesNodeTable,
	namespaceTable,
	accountTable,
	routeTableTable,
	aclTable,
	securityGroupTable,
	syncTimeTable,
}

type DbObject interface {
	DbId() string
	GetProvider() string
	SetSyncTime(string)
}

type Client interface {
	Open(filename string) error
	Close() error
	DropDB() error
	PutVPC(vpc *types.VPC) error
	GetVPC(id string) (*types.VPC, error)
	ListVPCs() ([]*types.VPC, error)
	DeleteVPC(id string) error
	PutInstance(instance *types.Instance) error
	GetInstance(id string) (*types.Instance, error)
	ListInstances() ([]*types.Instance, error)
	DeleteInstance(id string) error
	PutSubnet(subnet *types.Subnet) error
	GetSubnet(id string) (*types.Subnet, error)
	ListSubnets() ([]*types.Subnet, error)
	DeleteSubnet(id string) error
	PutACL(acl *types.ACL) error
	GetACL(id string) (*types.ACL, error)
	ListACLs() ([]*types.ACL, error)
	DeleteACL(id string) error
	PutRouteTable(routeTable *types.RouteTable) error
	GetRouteTable(id string) (*types.RouteTable, error)
	ListRouteTables() ([]*types.RouteTable, error)
	DeleteRouteTable(id string) error
	PutSecurityGroup(securityGroup *types.SecurityGroup) error
	GetSecurityGroup(id string) (*types.SecurityGroup, error)
	ListSecurityGroups() ([]*types.SecurityGroup, error)
	DeleteSecurityGroup(id string) error
	PutCluster(cluster *types.Cluster) error
	GetCluster(id string) (*types.Cluster, error)
	ListClusters() ([]*types.Cluster, error)
	DeleteCluster(id string) error
	PutPod(pod *types.Pod) error
	GetPod(id string) (*types.Pod, error)
	ListPods() ([]*types.Pod, error)
	DeletePod(id string) error
	PutKubernetesService(service *types.K8SService) error
	GetKubernetesService(id string) (*types.K8SService, error)
	ListKubernetesServices() ([]*types.K8SService, error)
	DeleteKubernetesService(id string) error
	PutKubernetesNode(node *types.K8sNode) error
	GetKubernetesNode(id string) (*types.K8sNode, error)
	ListKubernetesNodes() ([]*types.K8sNode, error)
	DeleteKubernetesNode(id string) error
	PutNamespace(namespace *types.Namespace) error
	GetNamespace(id string) (*types.Namespace, error)
	ListNamespaces() ([]*types.Namespace, error)
	DeleteNamespace(id string) error
	PutSyncTime(time *types.SyncTime) error
	GetSyncTime(id string) (*types.SyncTime, error)
	ListSyncTimes() ([]*types.SyncTime, error)
	DeleteSyncTime(id string) error
}