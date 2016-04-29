package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := []byte(`{"action":"get","node":{"key":"/rt/table","dir":true,"nodes":[{"key":"/rt/table/rt_mc","value":"{\"Tablename\":\"rt_mc\",\"ReplicaCount\":2,\"PartitionCount\":3,\"Status\":0,\"Ttl\":86400,\"LastChangeTime\":1460114949573048738,\"Bypass\":0,\"Trps\":[{\"ReplicaId\":0,\"PartitionId\":0,\"ChunkAddress\":{\"Hostname\":\"10.242.171.13\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":0},{\"ReplicaId\":0,\"PartitionId\":1,\"ChunkAddress\":{\"Hostname\":\"10.242.171.104\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":0},{\"ReplicaId\":1,\"PartitionId\":0,\"ChunkAddress\":{\"Hostname\":\"10.242.170.126\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":7,\"LastChangeTime\":0},{\"ReplicaId\":1,\"PartitionId\":2,\"ChunkAddress\":{\"Hostname\":\"10.242.171.13\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":0},{\"ReplicaId\":0,\"PartitionId\":2,\"ChunkAddress\":{\"Hostname\":\"10.242.170.126\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":7,\"LastChangeTime\":0},{\"ReplicaId\":1,\"PartitionId\":1,\"ChunkAddress\":{\"Hostname\":\"10.242.170.119\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":0}]}","modifiedIndex":9899155,"createdIndex":9899155},{"key":"/rt/table/tb_testnn","value":"{\"Tablename\":\"tb_testnn\",\"ReplicaCount\":3,\"PartitionCount\":3,\"Status\":0,\"Ttl\":86400,\"LastChangeTime\":1460641255512636206,\"Bypass\":0,\"Trps\":[{\"ReplicaId\":0,\"PartitionId\":0,\"ChunkAddress\":{\"Hostname\":\"10.242.170.119\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":1,\"PartitionId\":0,\"ChunkAddress\":{\"Hostname\":\"10.242.171.104\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":2,\"PartitionId\":0,\"ChunkAddress\":{\"Hostname\":\"10.242.171.13\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":0,\"PartitionId\":1,\"ChunkAddress\":{\"Hostname\":\"10.242.170.126\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":7,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":1,\"PartitionId\":1,\"ChunkAddress\":{\"Hostname\":\"10.242.170.119\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":2,\"PartitionId\":1,\"ChunkAddress\":{\"Hostname\":\"10.242.171.104\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":0,\"PartitionId\":2,\"ChunkAddress\":{\"Hostname\":\"10.242.171.13\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":1,\"PartitionId\":2,\"ChunkAddress\":{\"Hostname\":\"10.242.170.126\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":7,\"LastChangeTime\":1460641255512521570},{\"ReplicaId\":2,\"PartitionId\":2,\"ChunkAddress\":{\"Hostname\":\"10.242.170.119\",\"Port\":9091,\"PortNum\":6,\"ManagePort\":9090,\"Status\":0},\"Status\":0,\"LastChangeTime\":1460641255512521570}]}","modifiedIndex":9899156,"createdIndex":9899156}],"modifiedIndex":269,"createdIndex":269}}`)
	//fmt.Println((string)(data))

	//json反序列化
	var tableInfo interface{}
	err := json.Unmarshal(data, &tableInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(tableInfo)

	//json序列化，这种方式编码方便阅读(后两个参数表示每一行的前缀和缩进方式)
	d, err := json.MarshalIndent(tableInfo, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println((string)(d))

	for k, v := range tableInfo.(map[string]interface{}) {
		//fmt.Println(k, v)
		//fmt.Println("=========================================")
		switch v.(type) {
		case string:
			fmt.Println(k, "----------string type")
		case map[string]interface{}:
			fmt.Println(k, "----------map[string]interface{} type")
			//for sk, sv := range v.(map[string]interface{}) {
			//	fmt.Println(sk, sv)
			//}
			sv := v.(map[string]interface{})
			nodes := sv["nodes"].([]interface{})
			for nk, nv := range nodes {
				fmt.Println(nk, nv)
			}
		case int32:
			fmt.Println(k, "----------int32 type")
		case int64:
			fmt.Println(k, "----------int64 type")
		}
	}
}
