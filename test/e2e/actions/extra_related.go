package actions

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/controller/customresource"
	"github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/actions/deploy"
	kubecli "github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/cli/kubecli"
	mongocli "github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/cli/mongocli"
	"github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/model"
	"github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/utils"
)

const (
	atlasClusterCRD = "atlasdeployments.atlas.mongodb.com"
)

func DeleteClusterCRD(data *model.TestDataProvider) {
	By(fmt.Sprintf("Deleting %s", atlasClusterCRD), func() {
		kubecli.DeleteClusterResource("crd", atlasClusterCRD)

		By("Checking Cluster still existed", func() {
			state := mongocli.GetClusterStateName(data.Resources.ProjectID, data.Resources.Clusters[0].Spec.DeploymentSpec.Name)
			Expect(state).ShouldNot(Equal("DELETING"), "Cluster is being deleted despite the keep annotation")
		})
	})

	// Remove annotation so actions.AfterEachFinalCleanup can cleanup successfully
	By(fmt.Sprintf("Recreating %s", atlasClusterCRD), func() {
		deploy.NamespacedOperator(data)
		By("Removing keep annotation", func() {
			annotations := data.Resources.Clusters[0].ObjectMeta.GetAnnotations()
			// remove keep annotations from map
			delete(annotations, customresource.ResourcePolicyAnnotation)
			data.Resources.Clusters[0].ObjectMeta.SetAnnotations(annotations)
			UpdateCluster(data)
		})
	})
}

func Autoscaling(data *model.TestDataProvider) {
	By("Simulate autoscaling with changing cluster tier", func() {
		// we deployed cluster(default=M30) with Min=M30 and Max=M50 (data/atlascluster_with_autoscaling.yaml)
		// now we change the tier of the cluster to M40 to simulate the autoscaling
		newTier := "M40"
		mongocli.UpdateCluster(data.Resources.ProjectID, data.Resources.Clusters[0].Spec.DeploymentSpec.Name, "--tier", newTier)
		// Wait for updating to be done
		EventuallyWithOffset(1,
			func() string {
				return kubecli.GetK8sClusterStateName(data.Resources.Namespace, data.Resources.Clusters[0].GetClusterNameResource())
			},
			"45m", "1m",
		).Should(Equal("IDLE"), "Kubernetes resource: Cluster status should be IDLE")

		// Change sth in cluster to trigger reconciliation loop
		javaScriptEnabled := false
		data.Resources.Clusters[0].Spec.ProcessArgs.JavascriptEnabled = &javaScriptEnabled
		By("Update cluster\n", func() {
			utils.SaveToFile(
				data.Resources.Clusters[0].ClusterFileName(data.Resources),
				utils.JSONToYAMLConvert(data.Resources.Clusters[0]),
			)
			kubecli.Apply(data.Resources.Clusters[0].ClusterFileName(data.Resources), "-n", data.Resources.Namespace)
		})
		// Wait for updating to be done
		EventuallyWithOffset(1,
			func() string {
				return kubecli.GetK8sClusterStateName(data.Resources.Namespace, data.Resources.Clusters[0].GetClusterNameResource())
			},
			"45m", "1m",
		).Should(Equal("IDLE"), "Kubernetes resource: Cluster status should be IDLE")

		By("Checking operator won't scale down if 'instanceSize' conflicts with autoscaled size", func() {
			cluster := mongocli.GetClustersInfo(data.Resources.ProjectID, data.Resources.Clusters[0].Spec.DeploymentSpec.Name)
			clusterInstanceSize := cluster.ProviderSettings.InstanceSizeName
			Expect(clusterInstanceSize).Should(Equal(newTier), "Operator shouldn't change the instanceSize with autoscaling enabled")
		})
	})
}
