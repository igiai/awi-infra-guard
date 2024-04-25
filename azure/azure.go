// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
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

package azure

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
	"github.com/sirupsen/logrus"
)

const providerName = "Azure"

type ResourceClient struct {
	VNET        armnetwork.VirtualNetworksClient
	VNETPeering armnetwork.VirtualNetworkPeeringsClient
	NSG         armnetwork.SecurityGroupsClient
	Tag         armresources.TagsClient
}

func NewResourceClient(
	accountID string, credentials *azidentity.DefaultAzureCredential,
) (*ResourceClient, error) {
	vnetClient, err := armnetwork.NewVirtualNetworksClient(accountID, credentials, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create VNet Client: %v", err,
		)
	}
	if vnetClient == nil {
		return nil, errors.New(
			"failed to create VNet Client. Got empty client",
		)
	}
	vnetPeeringClient, err := armnetwork.NewVirtualNetworkPeeringsClient(accountID, credentials, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create VNet Peering Client: %v", err,
		)
	}
	if vnetPeeringClient == nil {
		return nil, errors.New(
			"failed to create VNet Peering Client. Got empty client",
		)
	}
	nsgClient, err := armnetwork.NewSecurityGroupsClient(accountID, credentials, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Security Group Client: %v", err,
		)
	}
	if nsgClient == nil {
		return nil, errors.New(
			"failed to create Security Group Client. Got empty client",
		)
	}
	tagClient, err := armresources.NewTagsClient(accountID, credentials, nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create Tag Client: %v", err,
		)
	}
	if tagClient == nil {
		return nil, errors.New(
			"failed to create Tag Client. Got empty client",
		)
	}
	return &ResourceClient{
		VNET:        *vnetClient,
		VNETPeering: *vnetPeeringClient,
		NSG:         *nsgClient,
		Tag:         *tagClient,
	}, nil
}

type Client struct {
	cred           *azidentity.DefaultAzureCredential
	logger         *logrus.Logger
	accountClients map[string]*ResourceClient
}

// NewClient initializes a new Azure client with all necessary clients for compute, network, and subscriptions.
func NewClient(ctx context.Context, logger *logrus.Logger) (*Client, error) {
	// Subscription ID from environment variable
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to obtain a credential: %w", err,
		)
	}
	client := &Client{
		cred:           cred,
		logger:         logger,
		accountClients: make(map[string]*ResourceClient),
	}

	if err = client.initializeClientsPerAccount(); err != nil {
		return nil, fmt.Errorf(
			"failed to initialize resource clients: %w", err,
		)
	}

	return client, nil
}

func (c *Client) initializeClientsPerAccount() error {
	accounts := c.ListAccounts()

	for _, account := range accounts {
		resClient, err := NewResourceClient(account.ID, c.cred)
		if err != nil {
			return fmt.Errorf(
				"failed to initialize Client for account ID '%s': %w",
				account.ID, err,
			)
		}
		if resClient == nil {
			return fmt.Errorf(
				"failed to initialize Client for account ID '%s'. Got nil object",
				account.ID,
			)
		}
		c.accountClients[account.ID] = resClient
	}

	return nil
}

func (c *Client) GetName() string {
	return providerName
}

func (c *Client) GetSyncTime(id string) (types.SyncTime, error) {
	return types.SyncTime{}, nil
}

func (c *Client) GetSubnet(ctx context.Context, input *infrapb.GetSubnetRequest) (types.Subnet, error) {
	// TBD
	return types.Subnet{}, nil
}

func (c *Client) GetVPCIDForCIDR(ctx context.Context, input *infrapb.GetVPCIDForCIDRRequest) (string, error) {
	// TBD
	return "", nil
}

func (c *Client) GetCIDRsForLabels(ctx context.Context, input *infrapb.GetCIDRsForLabelsRequest) ([]string, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetIPsForLabels(ctx context.Context, input *infrapb.GetIPsForLabelsRequest) ([]string, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetInstancesForLabels(ctx context.Context, input *infrapb.GetInstancesForLabelsRequest) ([]types.Instance, error) {
	// TBD
	return nil, nil
}

func (c *Client) GetVPCIDWithTag(ctx context.Context, input *infrapb.GetVPCIDWithTagRequest) (string, error) {
	// TBD
	return "", nil
}

func (c *Client) ListInternetGateways(ctx context.Context, params *infrapb.ListInternetGatewaysRequest) ([]types.IGW, error) {

	return nil, nil
}

/*
func getSubscriptionToken(ctx context.Context, subscriptionID string, credential *auth.DefaultAzureCredential) (string, error) {
    // Create a subscription-specific credential using azidentity
    subscriptionCredential, err := auth.NewSubscriptionCredential(credential, subscriptionID)
    if err != nil {
        return "", fmt.Errorf("failed to create subscription credential: %w", err)
    }

    // Acquire a token for the subscription
    token, err := subscriptionCredential.GetToken(ctx, "https://management.azure.com")
    if err != nil {
        return "", fmt.Errorf("failed to get token for subscription %s: %w", subscriptionID, err)
    }

    return token.AccessToken, nil
}

func getSubscriptionFactory (ctx context.Context) (subs []*armsubscriptions.Subscription, error) {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsubscriptions.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
	}
}

func getToken(ctx context.Context) (map[string]string, error) {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("Failed to obtain a credential:", err)
		return
	}

	subscriptionsClient, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		fmt.Println("Failed to create subscriptions client:", err)
		return
	}

	ctx := context.Background()

	pager := subscriptionsClient.NewListPager(nil)

	fmt.Println("Listing all VNets across all subscriptions:")

	for pager.More() {
		resp, err := pager.NextPage(ctx)
		if err != nil {
			fmt.Println("Failed to get the next page of subscriptions:", err)
			return
		}
		for _, sub := range resp.SubscriptionListResult.Value {
			fmt.Printf("Subscription: %s\n", *sub.SubscriptionID)
		}
	}
}*/
