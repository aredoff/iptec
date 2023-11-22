package firehol

var (
	sources = map[string]string{}
)

func init() {
	sources["firehol_level1"] = "https://iplists.firehol.org/files/firehol_level1.netset"
	sources["firehol_level2"] = "https://iplists.firehol.org/files/firehol_level2.netset"
	sources["firehol_level3"] = "https://iplists.firehol.org/files/firehol_level3.netset"
	sources["firehol_level4"] = "https://iplists.firehol.org/files/firehol_level4.netset"
}
