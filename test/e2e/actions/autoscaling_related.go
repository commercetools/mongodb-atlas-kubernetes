package actions

import (
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/util/compat"
	"github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/api/atlas"

	kubecli "github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/cli/kubecli"
	"github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/model"
	"github.com/mongodb/mongodb-atlas-kubernetes/test/e2e/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/atlas/mongodbatlas"
)

func Autoscaling(data *model.TestDataProvider) {
	// TODO: decompose this function like delete crd test
	By("Simulate autoscaling with changing deployment tier", func() {
		// we deployed atlas deployment(default=M30) with Min=M30 and Max=M50 (data/atlasdeployment_with_autoscaling.yaml)
		// now we change the tier of the deployment to M40 to simulate the autoscaling
		newTier := "M40"

		deployment := data.Resources.Deployments[0].Spec.DeploymentSpec
		deployment.ProviderSettings.InstanceSizeName = newTier

		newDeployment := &mongodbatlas.Cluster{}
		err := compat.JSONCopy(newDeployment, deployment)
		Expect(err).To(BeNil())

		atlasClient, err := atlas.AClient()
		Expect(err).To(BeNil())
		atlasClient.UpdateDeployment(data.Resources.ProjectID, data.Resources.Deployments[0].Spec.DeploymentSpec.Name, newDeployment)
		// Wait for updating to be done
		EventuallyWithOffset(1,
			func() string {
				return kubecli.GetK8sDeploymentStateName(data.Resources.Namespace, data.Resources.Deployments[0].GetDeploymentNameResource())
			},
			"45m", "1m",
		).Should(Equal("IDLE"), "Kubernetes resource: Deployment status should be IDLE")

		// Change sth in deployment to trigger reconciliation loop
		javaScriptEnabled := false
		data.Resources.Deployments[0].Spec.ProcessArgs.JavascriptEnabled = &javaScriptEnabled
		By("Update Deployment\n", func() {
			utils.SaveToFile(
				data.Resources.Deployments[0].DeploymentFileName(data.Resources),
				utils.JSONToYAMLConvert(data.Resources.Deployments[0]),
			)
			kubecli.Apply(data.Resources.Deployments[0].DeploymentFileName(data.Resources), "-n", data.Resources.Namespace)
		})
		// Wait for updating to be done
		EventuallyWithOffset(1,
			func() string {
				return kubecli.GetK8sDeploymentStateName(data.Resources.Namespace, data.Resources.Deployments[0].GetDeploymentNameResource())
			},
			"45m", "1m",
		).Should(Equal("IDLE"), "Kubernetes resource: Deployment status should be IDLE")

		By("Checking operator won't scale down if 'instanceSize' conflicts with autoscaled size", func() {
			deployment := atlasClient.GetDeployment(data.Resources.ProjectID, data.Resources.Deployments[0].Spec.DeploymentSpec.Name)
			deploymentInstanceSize := deployment.ProviderSettings.InstanceSizeName
			Expect(deploymentInstanceSize).Should(Equal(newTier), "Operator shouldn't change the instanceSize with autoscaling enabled")
		})
	})
}
