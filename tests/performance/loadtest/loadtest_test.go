/*
Copyright 2019 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package loadtest

import (
	"net/http"
	"time"

	"github.com/kubeedge/kubeedge/tests/e2e/utils"
	. "github.com/kubeedge/kubeedge/tests/performance/common"

	"github.com/golang/glog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/api/core/v1"
)

var DeploymentTestTimerGroup *utils.TestTimerGroup = utils.NewTestTimerGroup()

//Run Test cases
var _ = Describe("Application deployment test in Perfronace test EdgeNodes", func() {
	var UID string
	var testTimer *utils.TestTimer
	var testDescription GinkgoTestDescription
	var podlist metav1.PodList
	Context("Test application deployment on Kubeedge EdgeNodes", func() {
		BeforeEach(func() {
			testDescription = CurrentGinkgoTestDescription()
			testTimer = DeploymentTestTimerGroup.NewTestTimer(testDescription.TestText)
		})
		AfterEach(func() {
			// End test timer
			testTimer.End()
			// Print result
			testTimer.PrintResult()
			IsAppDeployed := utils.HandleDeployment(false, false, http.MethodDelete, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler+"/"+UID, "", ctx.Cfg.AppImageUrl[1], nodeSelector, "", 10)
			Expect(IsAppDeployed).Should(BeTrue())

			utils.CheckPodDeleteState(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, podlist)
		})

		Measure("MEASURE_PERF_NODETEST_NODES_1: Create 1 KubeEdge Node Deployment, Measure time for application comes into Running state", func(b Benchmarker) {
			podlist = metav1.PodList{}
			runtime := b.Time("runtime", func() {
				var deploymentList v1.DeploymentList
				podlist = metav1.PodList{}
				replica := 1
				//Generate the random string and assign as a UID
				UID = "edgecore-app-" + utils.GetRandomString(5)
				IsAppDeployed := utils.HandleDeployment(false, false, http.MethodPost, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler, UID, ctx.Cfg.AppImageUrl[1], "", "", replica)
				Expect(IsAppDeployed).Should(BeTrue())
				err := utils.GetDeployments(&deploymentList, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler)
				Expect(err).To(BeNil())
				for _, deployment := range deploymentList.Items {
					if deployment.Name == UID {
						//label := nodeName
						time.Sleep(2 * time.Second)
						podlist, err = utils.GetPods(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, "")
						Expect(err).To(BeNil())
						break
					}
				}
				utils.CheckPodRunningState(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, podlist)
			})

			glog.Infof("Runtime stats: %+v", runtime)

		}, 5)

		Measure("MEASURE_PERF_NODETEST_NODES_10: Create 10 KubeEdge Node Deployment, Measure time for application comes into Running state", func(b Benchmarker) {
			podlist = metav1.PodList{}
			runtime := b.Time("runtime", func() {
				var deploymentList v1.DeploymentList
				podlist = metav1.PodList{}
				replica := 10
				//Generate the random string and assign as a UID
				UID = "edgecore-app-" + utils.GetRandomString(5)
				IsAppDeployed := utils.HandleDeployment(false, false, http.MethodPost, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler, UID, ctx.Cfg.AppImageUrl[1], "", "", replica)
				Expect(IsAppDeployed).Should(BeTrue())
				err := utils.GetDeployments(&deploymentList, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler)
				Expect(err).To(BeNil())
				for _, deployment := range deploymentList.Items {
					if deployment.Name == UID {
						//label := nodeName
						time.Sleep(2 * time.Second)
						podlist, err = utils.GetPods(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, "")
						Expect(err).To(BeNil())
						break
					}
				}
				utils.CheckPodRunningState(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, podlist)
			})

			glog.Infof("Runtime stats: %+v", runtime)

		}, 5)

		Measure("MEASURE_PERF_NODETEST_NODES_50: Create 50 KubeEdge Node Deployment, Measure time for application comes into Running state", func(b Benchmarker) {
			podlist = metav1.PodList{}
			runtime := b.Time("runtime", func() {
				var deploymentList v1.DeploymentList
				podlist = metav1.PodList{}
				replica := 50
				//Generate the random string and assign as a UID
				UID = "edgecore-app-" + utils.GetRandomString(5)
				IsAppDeployed := utils.HandleDeployment(false, false, http.MethodPost, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler, UID, ctx.Cfg.AppImageUrl[1], "", "", replica)
				Expect(IsAppDeployed).Should(BeTrue())
				err := utils.GetDeployments(&deploymentList, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler)
				Expect(err).To(BeNil())
				for _, deployment := range deploymentList.Items {
					if deployment.Name == UID {
						//label := nodeName
						time.Sleep(2 * time.Second)
						podlist, err = utils.GetPods(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, "")
						Expect(err).To(BeNil())
						break
					}
				}
				utils.CheckPodRunningState(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, podlist)
			})

			glog.Infof("Runtime stats: %+v", runtime)

		}, 5)
		Measure("MEASURE_PERF_NODETEST_NODES_75: Create 75 KubeEdge Node Deployment, Measure time for application comes into Running state", func(b Benchmarker) {
			podlist = metav1.PodList{}
			runtime := b.Time("runtime", func() {
				var deploymentList v1.DeploymentList
				podlist = metav1.PodList{}
				replica := 75
				//Generate the random string and assign as a UID
				UID = "edgecore-app-" + utils.GetRandomString(5)
				IsAppDeployed := utils.HandleDeployment(false, false, http.MethodPost, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler, UID, ctx.Cfg.AppImageUrl[1], "", "", replica)
				Expect(IsAppDeployed).Should(BeTrue())
				err := utils.GetDeployments(&deploymentList, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler)
				Expect(err).To(BeNil())
				for _, deployment := range deploymentList.Items {
					if deployment.Name == UID {
						//label := nodeName
						time.Sleep(2 * time.Second)
						podlist, err = utils.GetPods(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, "")
						Expect(err).To(BeNil())
						break
					}
				}
				utils.CheckPodRunningState(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, podlist)
			})

			glog.Infof("Runtime stats: %+v", runtime)

		}, 5)
		Measure("MEASURE_PERF_NODETEST_NODES_100: Create 100 KubeEdge Node Deployment, Measure time for application comes into Running state", func(b Benchmarker) {
			podlist = metav1.PodList{}
			runtime := b.Time("runtime", func() {
				var deploymentList v1.DeploymentList
				podlist = metav1.PodList{}
				replica := 100
				//Generate the random string and assign as a UID
				UID = "edgecore-app-" + utils.GetRandomString(5)
				IsAppDeployed := utils.HandleDeployment(false, false, http.MethodPost, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler, UID, ctx.Cfg.AppImageUrl[1], "", "", replica)
				Expect(IsAppDeployed).Should(BeTrue())
				err := utils.GetDeployments(&deploymentList, ctx.Cfg.K8SMasterForKubeEdge+DeploymentHandler)
				Expect(err).To(BeNil())
				for _, deployment := range deploymentList.Items {
					if deployment.Name == UID {
						//label := nodeName
						time.Sleep(2 * time.Second)
						podlist, err = utils.GetPods(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, "")
						Expect(err).To(BeNil())
						break
					}
				}
				utils.CheckPodRunningState(ctx.Cfg.K8SMasterForKubeEdge+AppHandler, podlist)
			})

			glog.Infof("Runtime stats: %+v", runtime)

		}, 5)
	})
})
