package rirs

var (
	sources []source
)

type source struct {
	Name          string
	httpDatabases []string
	ftpDatabases  []string
}

func init() {

	sources = append(sources, source{
		Name:          "afrinic",
		httpDatabases: []string{"https://ftp.afrinic.net/pub/dbase/afrinic.db.gz"},
	})

	sources = append(sources, source{
		Name:          "arin",
		httpDatabases: []string{"https://ftp.arin.net/pub/rr/arin.db.gz"},
	})

	sources = append(sources, source{
		Name: "lacnic",
		httpDatabases: []string{
			"https://ftp.lacnic.net/lacnic/dbase/lacnic.db.gz",
			"https://ftp.lacnic.net/lacnic/irr/lacnic.db.gz",
		},
	})

	sources = append(sources, source{
		Name:          "ripe",
		httpDatabases: []string{"https://ftp.ripe.net/ripe/dbase/ripe.db.gz"},
	})

	sources = append(sources, source{
		Name: "apnic",
		httpDatabases: []string{
			"https://ftp.apnic.net/apnic/whois/apnic.db.as-block.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.as-set.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.aut-num.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.domain.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.filter-set.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.inet-rtr.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.inet6num.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.inetnum.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.irt.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.key-cert.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.limerick.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.mntner.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.organisation.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.peering-set.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.role.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.route-set.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.route.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.route6.gz",
			"https://ftp.apnic.net/apnic/whois/apnic.db.rtr-set.gz",
		},
	})
}
