package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
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
		log.Fatalln("Launch instance error", err)
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

		return err
	}

	rgParams := armresources.ResourceGroup{
		Location:  to.Ptr(location),
		ManagedBy: to.Ptr("Managed by user01"),
	}
	//creaste the resource group
	_, err = resourceGroupclient.CreateOrUpdate(ctx, resourceGroupName, rgParams, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("The resource group %s is creating... , please chgeck azure portal\n", resourceGroupName)

	}
	//create the virtusl network
	virtualNetworkClient, err := armnetwork.NewVirtualNetworksClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		return err
	}

	found, err := findVnet(ctx, resourceGroupName, vnetName, virtualNetworkClient)

	if err != nil {
		return err
	}
	if !found {
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

		vnetPolerResponse, err := vnetPollerResp.PollUntilDone(ctx, nil)
		if err != nil {
			return err
		} else {
			fmt.Printf("Vnet %v is creating...", *vnetPolerResponse.ID)
		}

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

	subnetResponse, err := subnetPollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("Subnet %v is creating...\n", *subnetResponse.Name)
	}
	// create the public ip address

	ipClient, err := armnetwork.NewPublicIPAddressesClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		return err
	}
	polerIPAddressResponse, err := ipClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		"pubIPaddress-net01",
		armnetwork.PublicIPAddress{
			Location: to.Ptr(location),
			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
			},
		},
		nil,
	)
	if err != nil {
		return err
	}

	ipAddressPolResponse, err := polerIPAddressResponse.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("Public IP  %v is creating...\n", *ipAddressPolResponse.Name)
	}

	//Network Security Group

	nsgSecurityClient, err := armnetwork.NewSecurityGroupsClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		return err
	}
	nsgSecurityResponse, err := nsgSecurityClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		"nsg-demo",
		armnetwork.SecurityGroup{
			Location: to.Ptr(location),
			Properties: &armnetwork.SecurityGroupPropertiesFormat{
				SecurityRules: []*armnetwork.SecurityRule{
					{
						Name: to.Ptr("allow-ssh"),
						Properties: &armnetwork.SecurityRulePropertiesFormat{
							SourceAddressPrefix:      to.Ptr("0.0.0.0/0"),
							SourcePortRange:          to.Ptr("*"),
							DestinationAddressPrefix: to.Ptr("0.0.0.0/0"),
							DestinationPortRange:     to.Ptr("22"),
							Protocol:                 to.Ptr(armnetwork.SecurityRuleProtocolTCP),
							Access:                   to.Ptr(armnetwork.SecurityRuleAccessAllow),
							Direction:                to.Ptr(armnetwork.SecurityRuleDirectionInbound),
							Description:              to.Ptr("allow ssh on port 22"),
							Priority:                 to.Ptr(int32(1001)),
						},
					},
				},
			},
		},
		nil,
	)
	if err != nil {
		return err
	}

	nsgSecurityGroupPResponse, err := nsgSecurityResponse.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("NSG %v is creating...\n", *nsgSecurityGroupPResponse.Name)
	}

	interfaceClient, err := armnetwork.NewInterfacesClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		return err
	}

	netInterfacePolerResponse, err := interfaceClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		"vm-network-interface-01",
		armnetwork.Interface{
			Location: to.Ptr(location),
			Properties: &armnetwork.InterfacePropertiesFormat{
				NetworkSecurityGroup: &armnetwork.SecurityGroup{
					ID: nsgSecurityGroupPResponse.ID,
				},
				IPConfigurations: []*armnetwork.InterfaceIPConfiguration{
					{
						Name: to.Ptr("pubIPaddress-net01"),
						Properties: &armnetwork.InterfaceIPConfigurationPropertiesFormat{
							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
							Subnet: &armnetwork.Subnet{
								ID: subnetResponse.ID,
							},
							PublicIPAddress: &armnetwork.PublicIPAddress{
								ID: ipAddressPolResponse.ID,
							},
						},
					},
				},
			},
		},
		nil,
	)
	if err != nil {
		return err
	}

	netInterfaceResponse, err := netInterfacePolerResponse.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("Network Interface %v is creating...\n", *netInterfaceResponse.Name)
	}

	// Create the vm

	fmt.Println("Creating the vm")
	vmClient, err := armcompute.NewVirtualMachinesClient(getSubscriptionID(), sshk.Token, nil)
	if err != nil {
		return err
	}

	parameters := armcompute.VirtualMachine{
		Location: to.Ptr(location),
		Identity: &armcompute.VirtualMachineIdentity{
			Type: to.Ptr(armcompute.ResourceIdentityTypeNone),
		},
		Properties: &armcompute.VirtualMachineProperties{
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{

					Offer:     to.Ptr("0001-com-ubuntu-server-focal"),
					Publisher: to.Ptr("canonical"),
					SKU:       to.Ptr("20_04-lts-gen2"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr("disk-01"),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS), // OSDisk type Standard/Premium HDD/SSD
					},
					DiskSizeGB: to.Ptr[int32](50), // default 127G
				},
			},
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_B2s")), // VM size include vCPUs,RAM,Data Disks,Temp storage.
			},
			OSProfile: &armcompute.OSProfile{ //
				ComputerName:  to.Ptr("server01"),
				AdminUsername: to.Ptr("azureadmin"),
				LinuxConfiguration: &armcompute.LinuxConfiguration{
					DisablePasswordAuthentication: to.Ptr(true),
					SSH: &armcompute.SSHConfiguration{
						PublicKeys: []*armcompute.SSHPublicKey{
							{
								Path:    to.Ptr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", "azureadmin")),
								KeyData: to.Ptr(string(sshk.PublicKey)),
							},
						},
					},
				},
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: netInterfaceResponse.ID,
					},
				},
			},
		},
	}

	pollerResponse, err := vmClient.BeginCreateOrUpdate(ctx, resourceGroupName, "Server01", parameters, nil)
	if err != nil {
		return err
	}

	vmResponse, err := pollerResponse.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	} else {
		fmt.Printf("Virtual Machine %v is creating...\n", *vmResponse.Name)
	}

	return nil

}

func findVnet(ctx context.Context, rg, vnetName string, vnetClient *armnetwork.VirtualNetworksClient) (bool, error) {

	_, err := vnetClient.Get(ctx, rg, vnetName, nil)
	if err != nil {
		var errResponse *azcore.ResponseError
		if errors.As(err, &errResponse) && errResponse.ErrorCode == "ResourceNotFound" {
			return false, nil
		}
		return false, err

	}
	return true, nil

}
