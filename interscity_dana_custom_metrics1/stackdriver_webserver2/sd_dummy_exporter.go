/*
Copyright 2018 Google Inc

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

package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"net/http"

	"strings"

	"strconv"

	gce "cloud.google.com/go/compute/metadata"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	monitoring "google.golang.org/api/monitoring/v3"
)

// SD Dummy Exporter is a testing utility that exports a metric of constant value to Stackdriver
// in a loop. Metric name and value can be specified with flags 'metric-name' and 'metric-value'.
// SD Dummy Exporter assumes that it runs as a pod in GCE or GKE cluster, and the pod id is passed
// to it with 'pod-id' flag (which can be passed to a pod via Downward API).
func main() {
	// Gather pod information
	// podId := flag.String("pod-id", "", "pod id")
	// metricName := flag.String("metric-name", "foo", "custom metric name")
	// metricValue := flag.Int64("metric-value", 124, "custom metric value")
	// flag.Parse()
	//
	// if *podId == "" {
	// 	log.Fatalf("No pod id specified.")
	// }
	//
	// stackdriverService, err := getStackDriverService()
	// if err != nil {
	// 	log.Fatalf("Error getting Stackdriver service: %v", err)
	// }
	//
	// labels := getResourceLabels(*podId)
	//
	//
	// //-----------
	// err2 := exportMetric(stackdriverService, *metricName, *metricValue, labels)
	// 	if err2 != nil {
	// 		log.Printf("Failed to write time series data: %v\n", err2)
	// 	} else {
	// 		log.Printf("Finished writing time series with value: %v\n", metricValue)
	// 	}
	// 	time.Sleep(5000 * time.Millisecond)
	//-----------



	http.HandleFunc("/", handler)
    	log.Fatal(http.ListenAndServe(":8083", nil))

	/*for {
		err := exportMetric(stackdriverService, *metricName, *metricValue, labels)
		if err != nil {
			log.Printf("Failed to write time series data: %v\n", err)
		} else {
			log.Printf("Finished writing time series with value: %v\n", metricValue)
		}
		time.Sleep(5000 * time.Millisecond)
	}*/
}

func handler(w http.ResponseWriter, r *http.Request) {
	// podId := flag.String("pod-id", "", "pod id")
	// 	metricName := flag.String("metric-name", "foo", "custom metric name")
	// 	metricValue := flag.Int64("metric-value", 112, "custom metric value")
	// 	flag.Parse()
	//
	// 	if *podId == "" {
	// 		log.Fatalf("No pod id specified.")
	// 	}
	//
	// 	stackdriverService, err := getStackDriverService()
	// 	if err != nil {
	// 		log.Fatalf("Error getting Stackdriver service: %v", err)
	// 	}
	//
	// 	labels := getResourceLabels(*podId)
	// 	for {
	// 		err := exportMetric(stackdriverService, *metricName, *metricValue, labels)
	// 		if err != nil {
	// 			log.Printf("Failed to write time series data: %v\n", err)
	// 		} else {
	// 			log.Printf("Finished writing time series with value: %v\n", metricValue)
	// 		}
	// 		time.Sleep(5000 * time.Millisecond)
	// 	}


fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	// converting the r.URL.Path[1:] variable into an int using Atoi method
	custom_metric_value, err := strconv.Atoi(r.URL.Path[1:])
 	custom_metric_value_int64 := int64(custom_metric_value)
	// Gather pod information
	podId := flag.String("pod-id", "", "pod id")
	metricName := flag.String("metric-name", "foo", "custom metric name")
	//metricValue := flag.Int64("metric-value", 197, "custom metric value")
	metricValue := flag.Int64("metric-value", custom_metric_value_int64, "custom metric value")
	flag.Parse()

	if *podId == "" {
		log.Fatalf("No pod id specified.")
	}

	stackdriverService, err := getStackDriverService()
	if err != nil {
		log.Fatalf("Error getting Stackdriver service: %v", err)
	}

	labels2 := getResourceLabels(*podId)


	//-----------
	err2 := exportMetric(stackdriverService, *metricName, *metricValue, labels2)
		if err2 != nil {
			log.Printf("Failed to write time series data: %v\n", err2)
		} else {
			log.Printf("Finished writing time series with value: %v\n", metricValue)
		}
		time.Sleep(5000 * time.Millisecond)
	//-----------

	//fmt.Fprintf(w, "Hi thereeeeee..., I love %s!", r.URL.Path[1:])

    	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	/*for {
		err := exportMetric(stackdriverService, *metricName, *metricValue, labels)
		if err != nil {
			log.Printf("Failed to write time series data: %v\n", err)
		} else {
			log.Printf("Finished writing time series with value: %v\n", metricValue)
		}
		time.Sleep(5000 * time.Millisecond)
	}*/
}

func getStackDriverService() (*monitoring.Service, error) {
	oauthClient := oauth2.NewClient(context.Background(), google.ComputeTokenSource(""))
	return monitoring.New(oauthClient)
}

// getResourceLabels returns resource labels needed to correctly label metric data
// exported to StackDriver. Labels contain details on the cluster (name, zone, project id)
// and pod for which the metric is exported (id)
func getResourceLabels(podId string) map[string]string {
	projectId, _ := gce.ProjectID()
	zone, _ := gce.Zone()
	clusterName, _ := gce.InstanceAttributeValue("cluster-name")
	clusterName = strings.TrimSpace(clusterName)
	return map[string]string{
		"project_id":   projectId,
		"zone":         zone,
		"cluster_name": clusterName,
		// container name doesn't matter here, because the metric is exported for
		// the pod, not the container
		"container_name": "",
		"pod_id":         podId,
		// namespace_id and instance_id don't matter
		"namespace_id": "default",
		"instance_id":  "",
	}
}

func exportMetric(stackdriverService *monitoring.Service, metricName string,
	metricValue int64, resourceLabels map[string]string) error {


		//var metricValue2 int64 = 1411
		dataPoint := &monitoring.Point{
			Interval: &monitoring.TimeInterval{
				EndTime: time.Now().Format(time.RFC3339),
			},
			Value: &monitoring.TypedValue{
				Int64Value: &metricValue,
			},
			// Value: &monitoring.TypedValue{
			// 	Int64Value: &metricValue2,
			// },
		}
		// Write time series data.
		request := &monitoring.CreateTimeSeriesRequest{
			TimeSeries: []*monitoring.TimeSeries{
				{
					Metric: &monitoring.Metric{
						Type: "custom.googleapis.com/" + metricName,
					},
					Resource: &monitoring.MonitoredResource{
						Type:   "gke_container",
						Labels: resourceLabels,
					},
					Points: []*monitoring.Point{
						dataPoint,
					},
				},
			},
		}
		projectName := fmt.Sprintf("projects/%s", resourceLabels["project_id"])
		_, err := stackdriverService.Projects.TimeSeries.Create(projectName, request).Do()
		return err









	// var metricValue2 int64 = 1409
	// dataPoint := &monitoring.Point{
	// 	Interval: &monitoring.TimeInterval{
	// 		EndTime: time.Now().Format(time.RFC3339),
	// 	},
	// 	//Value: &monitoring.TypedValue{
	// 	//	Int64Value: &metricValue,
	// 	//},
	// 	Value: &monitoring.TypedValue{
	// 		Int64Value: &metricValue2,
	// 	},
	// }
	// // Write time series data.
	// request := &monitoring.CreateTimeSeriesRequest{
	// 	TimeSeries: []*monitoring.TimeSeries{
	// 		{
	// 			Metric: &monitoring.Metric{
	// 				Type: "custom.googleapis.com/" + metricName,
	// 			},
	// 			Resource: &monitoring.MonitoredResource{
	// 				Type:   "gke_container",
	// 				Labels: resourceLabels,
	// 			},
	// 			Points: []*monitoring.Point{
	// 				dataPoint,
	// 			},
	// 		},
	// 	},
	// }
	// projectName := fmt.Sprintf("projects/%s", resourceLabels["project_id"])
	// _, err := stackdriverService.Projects.TimeSeries.Create(projectName, request).Do()
	// return err
}
