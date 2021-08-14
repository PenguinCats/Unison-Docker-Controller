package main

import (
	"fmt"
	"github.com/PenguinCats/Unison-Docker-Controller/api/types/docker_controller"
	"github.com/PenguinCats/Unison-Docker-Controller/pkg/controller"
)

func main() {
	fmt.Println("UDC: Starting UDC...")

	ctr, errController := controller.NewDockerController(&docker_controller.DockerControllerCreatBody{
		MemoryReserveRatio:          5,
		StorageReserveRatioForImage: 5,
		CoreAvailableList:           []string{"2", "3", "4", "5"},
		StoragePoolName:             "docker-thinpool",
		ContainerStopTimeout:        5,
	})
	fmt.Println(errController)

	fmt.Println(ctr)
	//
	//volumes := []string{"zhangbinjie", "tdd"}
	//_ = ctr.VolumeCreate("zhangbinjie")
	//_ = ctr.VolumeCreate("tdd")
	//
	//cID, err1 := ctr.ContainerCreat(container.ContainerCreateBody{
	//	ImageName:     "ubuntu:latest",
	//	CoreCnt:       2,
	//	MemorySize:    524288000,
	//	ContainerName: "pcat",
	//	Mounts:        volumes,
	//})
	//fmt.Println(cID, err1)
	//
	//err2 := ctr.ContainerStart(cID)
	//fmt.Println(err2)
	//
	////cID := "0dd144293cc4"
	//
	//usage, _ := ctr.ContainerResourceUsage(cID)
	//memUsage := usage.Memory / 1024 / 1024
	//fmt.Println(memUsage)
	//
	//err3 := ctr.ContainerStop(cID)
	//fmt.Println(err3)
	//
	//err4 := ctr.ContainerRemove(cID)
	//fmt.Println(err4)
	//
	//_ = ctr.VolumeRemove("zhangbinjie", false)
	//_ = ctr.VolumeRemove("tdd", false)
}
