package daemon

import (
	"net"
	"time"

	"github.com/AliyunContainerService/terway/types"
)

// PodEipInfo store pod eip info
// NOTE: this is the type store in db
// DEPRECATED
type PodEipInfo struct {
	PodEip                   bool
	PodEipID                 string
	PodEipIP                 string
	PodEipBandWidth          int
	PodEipChargeType         InternetChargeType
	PodEipISP                string
	PodEipPoolID             string
	PodEipBandwidthPackageID string
}

// PodInfo store the pod info
// NOTE: this is the type store in db
type PodInfo struct {
	//K8sPod *v1.Pod
	Name            string
	Namespace       string
	TcIngress       uint64
	TcEgress        uint64
	PodNetworkType  string
	PodIP           string      // used for eip and mip
	PodIPs          types.IPSet // used for eip and mip
	SandboxExited   bool
	EipInfo         PodEipInfo    // deprecated
	IPStickTime     time.Duration // deprecated
	PodENI          bool
	PodUID          string
	NetworkPriority string
	ERdma           bool
}

// ExtraEipInfo store extra eip info
// To judge whether delete user eip instance
type ExtraEipInfo struct {
	Delete         bool   `json:"delete"` // delete related eip on pod deletion
	AssociateENI   string `json:"associate_eni"`
	AssociateENIIP net.IP `json:"associate_eniip"`
}

// ResourceItem to be store
type ResourceItem struct {
	Type         string        `json:"type"`
	ID           string        `json:"id"`
	ExtraEipInfo *ExtraEipInfo `json:"extra_eip_info"`

	ENIID  string `json:"eni_id"`
	ENIMAC string `json:"eni_mac"`
	IPv4   string `json:"ipv4"`
	IPv6   string `json:"ipv6"`
}

// PodResources pod resources related
type PodResources struct {
	Resources   []ResourceItem
	PodInfo     *PodInfo
	NetNs       *string
	ContainerID *string

	NetConf string
}

// GetResourceItemByType get pod resource by resource type
func (p PodResources) GetResourceItemByType(resType string) []ResourceItem {
	var ret []ResourceItem
	for _, r := range p.Resources {
		if resType == r.Type {
			ret = append(ret, ResourceItem{Type: resType, ID: r.ID, ENIID: r.ENIID, ENIMAC: r.ENIMAC, IPv4: r.IPv4, IPv6: r.IPv6})
		}
	}
	return ret
}
