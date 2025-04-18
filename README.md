# Local copy WHOIS Database creator

A library for downloading and organizing WHOIS databases from Regional Internet Registries (RIRs) for further processing.

When I created this library, I realized that this approach does not suit me, but perhaps it will be useful to someone.

## Features

- Downloads WHOIS databases from all major RIRs
- Handles compressed (.gz) files automatically
- Provides structured access to WHOIS data
- Supports all RIR formats (AFRINIC, ARIN, APNIC, LACNIC, RIPE)
- Easy integration into data processing pipelines

## Supported Data Sources

The library fetches data from the following official RIR sources:

| RIR      | Database URLs |
|----------|---------------|
| **AFRINIC** | `https://ftp.afrinic.net/pub/dbase/afrinic.db.gz` |
| **ARIN** | `https://ftp.arin.net/pub/rr/arin.db.gz` |
| **LACNIC** | `https://ftp.lacnic.net/lacnic/dbase/lacnic.db.gz`<br>`https://ftp.lacnic.net/lacnic/irr/lacnic.db.gz` |
| **RIPE** | `https://ftp.ripe.net/ripe/dbase/ripe.db.gz` |
| **APNIC** | `https://ftp.apnic.net/apnic/whois/apnic.db.as-block.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.as-set.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.aut-num.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.domain.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.filter-set.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.inet-rtr.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.inet6num.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.inetnum.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.irt.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.key-cert.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.limerick.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.mntner.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.organisation.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.peering-set.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.role.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.route-set.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.route.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.route6.gz`<br>`https://ftp.apnic.net/apnic/whois/apnic.db.rtr-set.gz` |

## Installation

```bash
go get github.com/yourusername/whois-db-downloader
```

```go
func main() {
	folder, err := fs.New("/tmp/rirs")
	if err != nil {
		log.Fatal(err)
	}

	rir, err := rirs.New(folder)
	if err != nil {
		log.Fatal(err)
	}

	err = rir.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
```
