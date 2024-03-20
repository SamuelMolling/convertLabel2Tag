package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func main() {
	getAllClusters()
}

func getAllClusters() {
	ctx := context.Background()

	apiKey := "YOUR_PUBLIC_KEY"
	apiSecret := "YOUR_PRIVATE_KEY"

	sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
	if err != nil {
		log.Fatalf("Error when instantiating new client: %v", err)
	}
	projects, _, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	if err != nil {
		log.Fatalf("Could not fetch projects: %v", err)
	}

	for _, project := range *projects.Results {
		clusters, _, err := sdk.ClustersApi.ListClusters(ctx, *project.Id).Execute()
		if err != nil {
			log.Fatalf("Could not fetch clusters: %v", err)
		}
		for _, cluster := range *clusters.Results {
			fmt.Println("CLUSTER:", *cluster.Name)

			if labels := cluster.GetLabels(); labels != nil && len(labels) > 0 {
				fmt.Println("LABELS:")
				for _, label := range labels {
					fmt.Printf("  %s: %s\n", label.GetKey(), label.GetValue())
				}

				if tags := cluster.GetTags(); tags == nil || len(tags) == 0 {
					fmt.Println("Creating tags from labels...")
					newResourceTags := make([]admin.ResourceTag, 0, len(labels))
					for _, label := range labels {
						tag := admin.NewResourceTag()
						key := label.GetKey()
						value := label.GetValue()
						tag.Key = &key
						tag.Value = &value
						newResourceTags = append(newResourceTags, *tag)
					}
					cluster.SetTags(newResourceTags)

					fmt.Println("Valida tags")
					for _, tag := range cluster.GetTags() {
						fmt.Printf("  %s: %s\n", *tag.Key, *tag.Value)
					}

					cluster.ConnectionStrings = nil

					_, _, err := sdk.ClustersApi.UpdateCluster(ctx, *project.Id, *cluster.Name, &cluster).Execute()
					if err != nil {
						fmt.Println("Could not update cluster: %v", err)
						continue
					}

					fmt.Println("New TAGS:")
					for _, tag := range newResourceTags {
						fmt.Printf("  %s: %s\n", *tag.Key, *tag.Value)
					}

				}
			} else {
				fmt.Println("LABELS: None")

			}
		}

	}

}
