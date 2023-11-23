package dnsbl

var (
	sourcesIpv4 = []string{
		"bl.nordspam.com",
		"proxies.dnsbl.sorbs.net",
		"aspews.ext.sorbs.net",
		"dnsbl.beetjevreemd.nl",
		"rbl.interserver.net",
		"all.s5h.net",
		"sbl-xbl.spamhaus.org",
		"white.dnsbl.brukalai.lt",
		"ip4.bl.zenrbl.pl",
		"combined.rbl.msrbl.net",
		"dnsbl-3.uceprotect.net",
		"grey.uribl.com",
		"web.rbl.msrbl.net",
		"free.v4bl.org",
		"nbl.0spam.org",
		"0spamtrust.fusionzero.com",
		"bl.nszones.com",
		"spam.rbl.msrbl.net",
		"b.barracudacentral.org",
		"rbl.iprange.net",
		"wadb.isipp.com",
		"dnsbl.rv-soft.info",
		"ips.backscatterer.org",
		"torexit.dan.me.uk",
		"escalations.dnsbl.sorbs.net",
		"dnsbl-2.uceprotect.net",
		"sbl.spamhaus.org",
		"dnsbl.justspam.org",
		"bl.spamcop.net",
		"black.junkemailfilter.com",
		"dnsbl.dronebl.org",
		"hostkarma.junkemailfilter.com",
		"bl.score.senderscore.com",
		"db.wpbl.info",
		"ips.whitelisted.org",
		"noptr.spamrats.com",
		"block.dnsbl.sorbs.net",
		"pbl.spamhaus.org",
		"smtp.dnsbl.sorbs.net",
		"ix.dnsbl.manitu.net",
		"ip.v4bl.org",
		"dnsbl-1.uceprotect.net",
		"dnsbl.madavi.de",
		"rbl.abuse.ro",
		"socks.dnsbl.sorbs.net",
		"web.dnsbl.sorbs.net",
		"abuse.spfbl.net",
		"sbl.nszones.com",
		"bl.spameatingmonkey.net",
		"dnsbl.rymsho.ru",
		"iadb2.isipp.com",
		"relays.dnsbl.sorbs.net",
		"wl.mailspike.net",
		"list.bbfh.org",
		"spam.rbl.blockedservers.com",
		"z.mailspike.net",
		"light.dnsbl.brukalai.lt",
		"rbl.megarbl.net",
		"auth.spamrats.com",
		"old.spam.dnsbl.sorbs.net",
		"dyna.spamrats.com",
		"spam.pedantic.org",
		"dnsrbl.swinog.ch",
		"truncate.gbudb.net",
		"spamguard.leadmon.net",
		"all.spamrats.com",
		"v4.fullbogons.cymru.com",
		"bl.fmb.la",
		"dnsbl-0.uceprotect.net",
		"dnsblchile.org",
		"dialups.visi.com",
		"backscatter.spameatingmonkey.net",
		"swl.spamhaus.org",
		"dul.dnsbl.sorbs.net",
		"new.spam.dnsbl.sorbs.net",
		"multi.surbl.org",
		"gl.suomispam.net",
		"dev.null.dk",
		"zombie.dnsbl.sorbs.net",
		"score.spfbl.net",
		"bogons.cymru.com",
		"0spam.fusionzero.com",
		"black.uribl.com",
		"access.redhawk.org",
		"psbl.surriel.com",
		"blackholes.five-ten-sg.com",
		"red.uribl.com",
		"reputation-ip.rbl.scrolloutf1.com",
		"dnsbl.sorbs.net",
		"misc.dnsbl.sorbs.net",
		"dnsbl.zapbl.net",
		"singular.ttk.pte.hu",
		"all.spam-rbl.fr",
		"bl.nosolicitado.org",
		"spamsources.fabel.dk",
		"safe.dnsbl.sorbs.net",
		"bl.suomispam.net",
		"bl.0spam.org",
		"rbl.metunet.com",
		"http.dnsbl.sorbs.net",
		"bl.mailspike.net",
		"origin.asn.spameatingmonkey.net",
		"rbl.blockedservers.com",
		"recent.spam.dnsbl.sorbs.net",
		"bl.blocklist.de",
		"zen.spamhaus.org",
		"bl.worst.nosolicitado.org",
		"korea.services.net",
		"mail-abuse.blacklist.jippg.org",
		"dnswl.spfbl.net",
		"rbl.schulte.org",
		"rep.mailspike.net",
		"virus.rbl.msrbl.net",
		"spam.dnsbl.sorbs.net",
		"mtawlrev.dnsbl.rediris.es",
		"netscan.rbl.blockedservers.com",
		"plus.bondedsender.org",
		"dnsbl.kempt.net",
		"images.rbl.msrbl.net",
		"niprbl.mailcleaner.net",
		"phishing.rbl.msrbl.net",
		"tor.dan.me.uk",
		"rbl.lugh.ch",
		"cbl.abuseat.org",
		"bl.scientificspam.net",
		"0spam-n.fusionzero.com",
		"problems.dnsbl.sorbs.net",
		"bb.barracudacentral.org",
		"bsb.spamlookup.net",
		"rbl.ircbl.org",
		"nobl.junkemailfilter.com",
		"xbl.spamhaus.org",
		"query.bondedsender.org",
		"spam.spamrats.com",
		"dnsbl.spfbl.net",
		"black.dnsbl.brukalai.lt",
	}

	sourcesIpv6 = []string{
		"dnsbl6.anticaptcha.net",
		"all.v6.ascc.dnsbl.bit.nl",
		"ipv6.all.dnsbl.bit.nl",
		"bitonly.dnsbl.bit.nl",
		"ipv6.rbl.choon.net",
		"ipv6.rwl.choon.net",
		"v6.fullbogons.cymru.com",
		"origin6.asn.cymru.com",
		"tor.dan.me.uk",
		"torexit.dan.me.uk",
		"dnsbl.beetjevreemd.nl",
		"list.dnswl.org",
		"dnsbl.dronebl.org",
		"bl.nordspam.com",
		"pofon.foobar.hu",
		"ispmx.pofon.foobar.hu",
		"bl6.rbl.polspam.pl",
		"ip6.white.polspam.pl",
		"lblip6.rbl.polspam.pl",
		"rblip6.rbl.polspam.pl",
		"eswlrev.dnsbl.rediris.es",
		"mtawlrev.dnsbl.rediris.es",
		"all.s5h.net",
		"bl.ipv6.spameatingmonkey.net",
		"pbl.spamhaus.org",
		"sbl.spamhaus.org",
		"sbl-xbl.spamhaus.org",
		"swl.spamhaus.org",
		"xbl.spamhaus.org",
		"zen.spamhaus.org",
		"abuse.spfbl.net",
		"dnsbl.spfbl.net",
		"score.spfbl.net",
		"dnswl.spfbl.net",
		"bl.suomispam.net",
		"gl.suomispam.net",
		"ipv6.blacklist.woody.ch",
	}
)
