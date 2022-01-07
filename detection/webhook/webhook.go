package webhook

import (
	"encoding/json"
	"engine/engine"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	v1 "k8s.io/apiserver/pkg/apis/audit/v1"
)

func HookStart() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(r.Body)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}

		var events v1.EventList
		err = json.Unmarshal(body, &events)
		if err != nil {
			http.Error(w, "failed to unmarshal audit events", http.StatusBadRequest)
			return
		}
		// Iterate and filter audit events
		for _, event := range events.Items {
			// capture curl event
			engine.Fire(&event)
		}
	})

	log.Println("Kubernetes: WebServer Hook started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// isPodCreation returns true if the given event is of a pod creation
func isPodCreation(event v1.Event) bool {
	return event.Verb == "create" &&
		event.Stage == v1.StageResponseComplete &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "pods"
}

// isPodCreation returns true if the given event is of a pod creation
func isPodDelete(event v1.Event) bool {
	return event.Verb == "delete" &&
		event.Stage == v1.StageResponseComplete &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "pods"
}

func isPodEvent(event v1.Event) bool {
	return event.Stage == v1.StageResponseComplete &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "pods"
}

// isPodCreation returns true if the given event is of a pod creation
func isConfigMap(event v1.Event) bool {
	return event.Verb != "watch" &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "configmaps"
}

// isPodCreation returns true if the given event is of a pod creation
func isSecretList(event v1.Event) bool {
	return event.Verb == "list" &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "secrets"
}

// isPodCreation returns true if the given event is of a pod creation
func isCronjobChange(event v1.Event) bool {
	return event.Verb == "patch" &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "cronjobs"
}

func test(event v1.Event) bool {
	return strings.Split(event.UserAgent, "/")[0] == "kubectl" &&
		event.ObjectRef != nil &&
		event.ObjectRef.Resource == "podsecuritypolicies"
}

func isUnauthenticated(event v1.Event) bool {
	return event.User.Groups != nil &&
		event.User.Groups[0] == "system:unauthenticated"
}
