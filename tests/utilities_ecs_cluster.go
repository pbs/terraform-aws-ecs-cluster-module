package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func testECSCluster(t *testing.T, variant string) {
	t.Parallel()

	terraformDir := fmt.Sprintf("../examples/%s", variant)

	clusterName := RandomString(5)
	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
		Vars: map[string]interface{}{
			"name": fmt.Sprintf("example-%s", clusterName),
		},
		LockTimeout: "5m",
	}

	var outputClusterName string

	defer func() {
		if variant == "ec2" {
			DeleteClusterInstances(outputClusterName)
		}
		terraform.Destroy(t, terraformOptions)
	}()

	terraform.InitAndApply(t, terraformOptions)

	clusterARN := terraform.Output(t, terraformOptions, "arn")
	outputClusterName = terraform.Output(t, terraformOptions, "name")

	region := getAWSRegion(t)
	accountID := getAWSAccountID(t)

	expectedARN := fmt.Sprintf("arn:aws:ecs:%s:%s:cluster/example-%s", region, accountID, clusterName)

	assert.Contains(t, expectedARN, clusterARN)
	assert.Equal(t, outputClusterName, fmt.Sprintf("example-%s", clusterName))
}
