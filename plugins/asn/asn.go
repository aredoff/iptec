package asn

import (
	"fmt"
	"net/netip"
	"sort"
)

type AS struct {
	Start       netip.Addr
	End         netip.Addr
	Number      int
	CountryCode string
	Description string
}

func (a AS) String() string {
	ip := "[invalid]"
	if a.Start.IsValid() && a.End.IsValid() {
		ip = fmt.Sprintf("[%s-%s]", a.Start, a.End)
	}
	return fmt.Sprintf("AS%d(%s)@%s%s", a.Number, a.Description, a.CountryCode, ip)
}

func (a AS) Contains(ip netip.Addr) bool {
	return ip.Compare(a.Start) >= 0 && ip.Compare(a.End) <= 0
}

type asnDb []AS

func (a asnDb) Len() int {
	return len(a)
}

func (a asnDb) Less(i, j int) bool {
	return a[i].Start.Less(a[j].Start)
}

func (a asnDb) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a asnDb) Sort() {
	sort.Sort(a)
}

func (a asnDb) Find(ip netip.Addr) *AS {
	for _, v := range a {
		if v.Contains(ip) {
			return &v
		}
	}
	return nil
}
