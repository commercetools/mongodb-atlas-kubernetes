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
)

const (
	atlasClusterCRD = "atlasdeployments.atlas.mongodb.com"
)

func DeleteCRDs(data *model.TestDataProvider) {
	By(fmt.Sprintf("Deleting %s", atlasClusterCRD), func() {
		kubecli.DeleteClusterResource("crd", atlasClusterCRD)
		// TODO: check CRD deletion
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
