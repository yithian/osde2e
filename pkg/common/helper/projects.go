package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/openshift/osde2e/pkg/common/runner"

	projectv1 "github.com/openshift/api/project/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

func (h *H) createProject(suffix string) (*projectv1.Project, error) {
	proj := &projectv1.Project{
		ObjectMeta: metav1.ObjectMeta{
			Name: "osde2e-" + suffix,
		},
	}

	project, err := h.Project().ProjectV1().Projects().Create(context.TODO(), proj, metav1.CreateOptions{})

	if err != nil {
		return project, err
	}

	wait.PollImmediate(5*time.Second, 60*time.Second, func() (done bool, err error) {
		project, err = h.Project().ProjectV1().Projects().Get(context.TODO(), project.Name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		if project != nil {
			return true, nil
		}

		return false, nil
	})

	return project, err
}

func (h *H) inspect(projectName string) error {
	inspectTimeoutInSeconds := 200
	h.SetServiceAccount("system:serviceaccount:%s:cluster-admin")
	r := h.Runner(fmt.Sprintf("oc adm inspect ns/%v --dest-dir=%v", projectName, runner.DefaultRunner.OutputDir))
	r.Name = "osde2e-project"
	r.Tarball = true
	stopCh := make(chan struct{})

	err := r.Run(inspectTimeoutInSeconds, stopCh)
	if err != nil {
		return fmt.Errorf("Error running project inspection: %s", err.Error())
	}

	gatherResults, err := r.RetrieveResults()
	if err != nil {
		return fmt.Errorf("Error retrieving project inspection results: %s", err.Error())
	}

	h.WriteResults(gatherResults)
	return nil
}
