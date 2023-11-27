package blacklist

var (
	sources = []*source{}
)

func newSource(name, url string, extractor func(string) []string, points int) *source {
	return &source{
		Name:      name,
		Url:       url,
		Points:    points,
		extractor: extractor,
	}
}

type source struct {
	Name      string
	Url       string
	Points    int
	extractor func(string) []string
}

func init() {
	sources = append(sources, newSource("firehol_level1", "https://iplists.firehol.org/files/firehol_level1.netset", simpleExtracotor, 5))
	sources = append(sources, newSource("firehol_level2", "https://iplists.firehol.org/files/firehol_level2.netset", simpleExtracotor, 5))
	sources = append(sources, newSource("firehol_level3", "https://iplists.firehol.org/files/firehol_level3.netset", simpleExtracotor, 5))
	sources = append(sources, newSource("firehol_level4", "https://iplists.firehol.org/files/firehol_level4.netset", simpleExtracotor, 5))
	sources = append(sources, newSource("blocklist.de", "https://lists.blocklist.de/lists/all.txt", simpleExtracotor, 5))
	sources = append(sources, newSource("myip.ms", "http://myip.ms/files/blacklist/csf/latest_blacklist.txt", simpleExtracotor, 5))
	sources = append(sources, newSource("cinsscore.com", "http://cinsscore.com/list/ci-badguys.txt", simpleExtracotor, 5))
	sources = append(sources, newSource("stopforumspam.com", "https://www.stopforumspam.com/downloads/toxic_ip_cidr.txt", simpleExtracotor, 5))
	sources = append(sources, newSource("danger.rulez.sk", "https://danger.rulez.sk/projects/bruteforceblocker/blist.php", firstAddressSpaсeExtracotor, 5))
	sources = append(sources, newSource("spamhaus.org", "https://www.spamhaus.org/drop/drop.lasso", firstAddressSpaсeExtracotor, 5))
	sources = append(sources, newSource("emergingthreats.net", "https://rules.emergingthreats.net/fwrules/emerging-Block-IPs.txt", simpleExtracotor, 5))
	sources = append(sources, newSource("team-cymru.org/fullbogons-ipv4.txt", "https://www.team-cymru.org/Services/Bogons/fullbogons-ipv4.txt", simpleExtracotor, 5))
	sources = append(sources, newSource("team-cymru.org/fullbogons-ipv6.txt", "https://www.team-cymru.org/Services/Bogons/fullbogons-ipv6.txt", simpleExtracotor, 5))
}
