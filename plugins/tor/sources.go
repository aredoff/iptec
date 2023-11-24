package tor

var (
	sources = map[string]string{}
)

func init() {
	sources["enkidu-6/tor-relay-lists/exits-v4.txt"] = "https://raw.githubusercontent.com/Enkidu-6/tor-relay-lists/main/exits-v4.txt"
	sources["enkidu-6/tor-relay-lists/main/exits-v6.txt"] = "https://raw.githubusercontent.com/Enkidu-6/tor-relay-lists/main/exits-v6.txt"
	sources["secops-Institute/tor-ip-addresses/master/tor-exit-nodes.lst"] = "https://raw.githubusercontent.com/SecOps-Institute/Tor-IP-Addresses/master/tor-exit-nodes.lst"
	sources["enkidu-6/tor-relay-lists/main/relays-v4.txt"] = "https://raw.githubusercontent.com/Enkidu-6/tor-relay-lists/main/relays-v4.txt"
	sources["enkidu-6/tor-relay-lists/main/relays-v6.txt"] = "https://raw.githubusercontent.com/Enkidu-6/tor-relay-lists/main/relays-v6.txt"
}
