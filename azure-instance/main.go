package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/dkr290/go-devops/azure-instance/keys"
)

var sshk keys.SshKeys

func main() {

	var (
		location   = "northeurope"
		rg         = "test-rg"
		ctx        = context.Background()
		vnetName   = "test-vnet"
		subnetName = "subnet01"
	)
	if err := sshk.MyGenerateKeys(); err != nil {
		log.Fatalln("Error my generation of keys", err)
	}

	if err := sshk.GetToken(); err != nil {
		log.Fatalln("Error generation the token", err)
	}

	if err := launchInstance(ctx, rg, location, vnetName, subnetName); err != nil {
		log.Fatalln("Could not create resource group", rg)
	}

}
func getSubscriptionID() string {

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		log.Fatal("AZURE_SUBSCRIPTION_ID is not set.")
	}
	return subscriptionID
}

func launchInstance(ctx context.Context, resourceGroupName, location, vnetName, subnetName string) error {
	// create the client for azure resource group
	resourceGroupclient, err := armresources.NewResourceGroupsClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		log.Fatalln(err)
	}

	rgParams := armresources.ResourceGroup{
		Location:  to.Ptr(location),
		ManagedBy: to.Ptr("Managed by user01"),
	}
	//creaste the resource group
	resourceGroupREsponse, err := resourceGroupclient.CreateOrUpdate(ctx, resourceGroupName, rgParams, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("The resource group %s is creating... , please chgeck azure portal\n", resourceGroupName)
		fmt.Printf("Response ... - %v", resourceGroupREsponse)
	}
	//create the virtusl network
	virtualNetworkClient, err := armnetwork.NewVirtualNetworksClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		return err
	}

	vnetPollerResp, err := virtualNetworkClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		vnetName,
		armnetwork.VirtualNetwork{
			Location: to.Ptr(location),
			Properties: &armnetwork.VirtualNetworkPropertiesFormat{
				AddressSpace: &armnetwork.AddressSpace{
					AddressPrefixes: []*string{
						to.Ptr("10.1.0.0/16"),
					},
				},
			},
		},
		nil)

	if err != nil {
		return err
	}

	vnetResp, err := vnetPollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Println("Vnet is creating...\n", vnetResp)
	}

	//create subnet
	subnetsClient, err := armnetwork.NewSubnetsClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		return err
	}

	subnetPollerResp, err := subnetsClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		vnetName,
		subnetName,
		armnetwork.Subnet{
			Properties: &armnetwork.SubnetPropertiesFormat{
				AddressPrefix: to.Ptr("10.1.0.0/24"),
			},
		},
		nil,
	)

	if err != nil {
		return err
	}

	subnetResp, err := subnetPollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Println("Subnet is creating...\n", subnetResp)
	}

	return nil

}
