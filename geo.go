package main

import (
  "fmt"
  "context"
)

// Find clusters (For the last day)
// select id, st_clusterdbscan(location, 1, 1) over(order by id) as cluster from location;
// Return list of people in clusters
// Match on interests
// Create event

type ClusterRow struct {
  Id int
  Cluster int
  UserId int
}

type ClusterResult struct {
  ClusterId int
  UserIds []int
}

type ClusterContainer struct {
  ClusterData []ClusterResult
}

func GetDailyCluster() []ClusterRow {
  ctx := context.Background()
  c := getCursor()
  rows, err := c.QueryContext(ctx, "select id, st_clusterdbscan(location, .001, 2) over(order by id) as cluster, userid from location")
  if err != nil {
    panic(err)
  }
  cluster_rows := []ClusterRow{}
  for rows.Next() {
    var id int
    var cluster int
    var userid int
    _ = rows.Scan(&id, &cluster, &userid)
    cr := ClusterRow{Id: id, Cluster: cluster, UserId: userid}
    cluster_rows = append(cluster_rows, cr)
  }
  fmt.Println(cluster_rows)
  return cluster_rows
}

// Get uniq users from clusters
