package k8s_chart_support_upgrade_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"

	"github.com/rancher/hosted-providers-e2e/hosted/aks/helper"
	"github.com/rancher/hosted-providers-e2e/hosted/helpers"
)

var _ = Describe("K8sChartSupportUpgradeImport", func() {
	var cluster *management.Cluster

	BeforeEach(func() {
		err := helper.CreateAKSClusterOnAzure(location, clusterName, k8sVersion, "1", helpers.GetCommonMetadataLabels())
		Expect(err).To(BeNil())
		cluster, err = helper.ImportAKSHostedCluster(ctx.RancherAdminClient, clusterName, ctx.CloudCredID, location, helpers.GetCommonMetadataLabels())
		Expect(err).To(BeNil())
		cluster, err = helpers.WaitUntilClusterIsReady(cluster, ctx.RancherAdminClient)
		Expect(err).To(BeNil())
	})
	AfterEach(func() {
		if ctx.ClusterCleanup && cluster != nil {
			err := helper.DeleteAKSHostCluster(cluster, ctx.RancherAdminClient)
			Expect(err).To(BeNil())
			err = helper.DeleteAKSClusteronAzure(clusterName)
			Expect(err).To(BeNil())
		} else {
			fmt.Println("Skipping downstream cluster deletion: ", clusterName)
		}
	})

	It("should successfully test k8s chart support import in an upgrade scenario", func() {
		GinkgoLogr.Info(fmt.Sprintf("Testing K8s %s chart support for import on Rancher upgraded from %s to %s", helpers.K8sUpgradedMinorVersion, helpers.RancherFullVersion, helpers.RancherUpgradeFullVersion))

		testCaseID = 253 // Report to Qase
		commonchecks(&ctx, cluster, clusterName, helpers.RancherUpgradeFullVersion, helpers.K8sUpgradedMinorVersion)
	})

})
