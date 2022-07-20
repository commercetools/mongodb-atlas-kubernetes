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

func Deploy(data *model.TestDataProvider) {
	// Remove annotation so actions.AfterEachFinalCleanup can cleanup successfully
	By(fmt.Sprintf("Creating %s", atlasClusterCRD), func() {
		deploy.NamespacedOperator(data)
	})
}

func DeleteClusterCRD(data *model.TestDataProvider) {
	By(fmt.Sprintf("Deleting %s", atlasClusterCRD), func() {
		kubecli.DeleteClusterResource("crd", atlasClusterCRD)
	})
}

func DeleteCluster(data *model.TestDataProvider) {
	By("Deleting cluster", func() {
		DeleteCluster(data)
	})
}

func AnnotateKeep(data *model.TestDataProvider) {
	By("Adding keep annotation", func() {
		annotations := data.Resources.Clusters[0].ObjectMeta.GetAnnotations()
		// remove keep annotations from map
		annotations[customresource.ReconciliationPolicyAnnotation] = "keep"
		data.Resources.Clusters[0].ObjectMeta.SetAnnotations(annotations)
		UpdateCluster(data)
	})
}

func ClusterExists(data *model.TestDataProvider) {
	By("Checking Cluster still existed", func() {
		state := mongocli.GetClusterStateName(data.Resources.ProjectID, data.Resources.Clusters[0].Spec.DeploymentSpec.Name)
		Expect(state).ShouldNot(Equal("DELETING"), "Cluster is being deleted despite the keep annotation")
	})
}

func RemoveAnnotateKeep(data *model.TestDataProvider) {
	By("Removing keep annotation", func() {
		annotations := data.Resources.Clusters[0].ObjectMeta.GetAnnotations()
		// remove keep annotations from map
		delete(annotations, customresource.ResourcePolicyAnnotation)
		data.Resources.Clusters[0].ObjectMeta.SetAnnotations(annotations)
		UpdateCluster(data)
	})
}
