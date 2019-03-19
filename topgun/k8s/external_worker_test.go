package k8s_test

import (
	"fmt"
	"time"

	"github.com/onsi/gomega/gexec"

	. "github.com/concourse/concourse/topgun"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("team external workers", func() {

	var (
		proxySession *gexec.Session
		releaseName  string
		namespace    string
		atcEndpoint  string
		workerKey    string
	)

	JustBeforeEach(func() {
		releaseName = fmt.Sprintf("topgun-xw-%d", randomGenerator.Int())
		namespace = releaseName

		Run(nil, "kubectl", "create", "namespace", namespace)

		configMapCreationArgs := []string{
			"create",
			"configmap",
			"main-worker-public-key",
			"--namespace=" + namespace,
			"--from-literal=main-worker-public-key=" + workerKey,
		}

		Run(nil, "kubectl", configMapCreationArgs...)

		helmArgs := []string{
			"--set=worker.replicas=1",
			"--set=concourse.worker.team=main",
			"--set=concourse.web.tsa.teamAuthorizedKeys=main:/authorized-team-keys/main-worker-public-key",
			"--set=web.additionalVolumes[0].name=main-worker-public-key",
			"--set=web.additionalVolumes[0].configMap.name=main-worker-public-key",
			"--set=web.additionalVolumeMounts[0].name=main-worker-public-key",
			"--set=web.additionalVolumeMounts[0].mountPath=/authorized-team-keys",
			"--set=web.env[0].name=CONCOURSE_TSA_AUTHORIZED_KEYS",
			"--set=web.env[0].value=",
		}
		deployConcourseChart(releaseName, helmArgs...)

		waitAllPodsInNamespaceToBeReady(namespace)

		By("Creating the web proxy")
		proxySession, atcEndpoint = startPortForwarding(namespace, "service/"+releaseName+"-web", "8080")

		By("Logging in")
		fly.Login("test", "test", atcEndpoint)

	})

	Context("web with correct public key to the worker", func() {
		BeforeEach(func() {
			workerKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC496FSYFcBAKgDtMsBAJiF/6/NxlXKP5UZecyEsedYuTt1GOgJTwaA1qZ1LmHsbfLDE68oDdiM4uvxfI4wtLhz57w3u0jOUxZ2JeF7SVwEf1nVqLn4Gh/f8GUNQGSyIp1zUD5Bx9fq0PAyQ47mt7Ufi84rcf8LKl7nzAIHTcdg2BvTkQN9bUGPaq/Pb1W2bKPAQy4OzXTSIyrAJ89TH2jFeaZfyxQFGbD9jVHH/yl0oiMrDeaRYgccE5II+KY7WoLjsBry/9Qf2ERELKTK4UeIGIqWci9lab1ti+GxFPPiC3krNFjo4jShV4eUs4cNIrjwNrxVaKPXmU6o7Y3Hpayx Concourse"
		})
		It("worker registers with team main", func() {
			By("waiting for a running worker")
			Eventually(func() []Worker {
				return getRunningWorkers(fly.GetWorkers())
			}, 2*time.Minute, 10*time.Second).
				ShouldNot(HaveLen(0))
			worker := getRunningWorkers(fly.GetWorkers())
			Expect(worker).To(HaveLen(1))
			Expect(worker[0].Team).To(Equal("main"))
		})
	})
	Context("web with invalid public key to the worker", func() {
		BeforeEach(func() {
			workerKey = "ssh-rsa 1234ABCD Concourse"
		})
		It("worker doesn't registers with team main", func() {
			Consistently(func() []Worker {
				return getRunningWorkers(fly.GetWorkers())
			}, 1*time.Minute, 10*time.Second).
				Should(HaveLen(0))
		})
	})

	AfterEach(func() {
		helmDestroy(releaseName)
		Wait(Start(nil, "kubectl", "delete", "namespace", namespace, "--wait=false"))
		Wait(proxySession.Interrupt())
	})

})
