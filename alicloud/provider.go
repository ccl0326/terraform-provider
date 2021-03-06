package alicloud

import (
	"github.com/denverdino/aliyungo/common"
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for alicloud
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALICLOUD_ACCESS_KEY", nil),
				Description: descriptions["access_key"],
			},
			"secret_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALICLOUD_SECRET_KEY", nil),
				Description: descriptions["secret_key"],
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ALICLOUD_REGION", "cn-beijing"),
				Description: descriptions["region"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"alicloud_instance":        resourceAliyunInstance(),
			"alicloud_disk":            resourceAliyunDisk(),
			"alicloud_disk_attachment": resourceAliyunDiskAttachment(),
			"alicloud_security_group":  resourceAliyunSecurityGroup(),
			"alicloud_vpc":             resourceAliyunVpc(),
			"alicloud_nat_gateway":     resourceAliyunNatGateway(),
			//both subnet and vswith exists,cause compatible old version, and compatible aws habit.
			"alicloud_subnet":          resourceAliyunSubnet(),
			"alicloud_vswitch":         resourceAliyunSubnet(),
			"alicloud_eip":             resourceAliyunEip(),
			"alicloud_eip_association": resourceAliyunEipAssociation(),
			"alicloud_slb":             resourceAliyunSlb(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
		Region:    common.Region(d.Get("region").(string)),
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// This is a global MutexKV for use within this plugin.
var alicloudMutexKV = mutexkv.NewMutexKV()

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"access_key": "Access key of alicloud",
		"secret_key": "Secret key of alicloud",
		"region":     "Region of alicloud",
	}
}
