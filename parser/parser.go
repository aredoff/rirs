package parser

import (
	"strings"
	"time"
)

const (
	TimeLayout = "2006-01-02T15:04:05Z"
)

type Storage interface {
	SaveASN(asn *ASN) error
	SaveInetNum(inetNum *InetNum) error
	SaveRoute(route *Route) error
	SaveRoute6(route6 *Route6) error
	SavePerson(person *Person) error
	SaveOrganization(org *Organization) error
	SaveDomain(domain *Domain) error
}

type Parser struct {
	storage Storage
}

func NewParser(storage Storage) *Parser {
	return &Parser{
		storage: storage,
	}
}

func (p *Parser) parseAndSaveObject(objType string, lines []string) error {
	base := p.parseBaseObject(lines)

	switch objType {
	case "aut-num":
		asn, err := p.parseASN(base, lines)
		if err != nil {
			return err
		}
		return p.storage.SaveASN(asn)
	case "inetnum":
		inetnum, err := p.parseInetNum(base, lines)
		if err != nil {
			return err
		}
		return p.storage.SaveInetNum(inetnum)
	case "route":
		route, err := p.parseRoute(base, lines)
		if err != nil {
			return err
		}
		return p.storage.SaveRoute(route)
	case "route6":
		route6, err := p.parseRoute6(base, lines)
		if err != nil {
			return err
		}
		return p.storage.SaveRoute6(route6)
	case "person":
		person, err := p.parsePerson(base, lines)
		if err != nil {
			return err
		}
		return p.storage.SavePerson(person)
	case "organisation":
		org, err := p.parseOrganization(base, lines)
		if err != nil {
			return err
		}
		return p.storage.SaveOrganization(org)
	case "domain":
		domain, err := p.parseDomain(base, lines)
		if err != nil {
			return err
		}
		return p.storage.SaveDomain(domain)
	}

	return nil
}

func (p *Parser) parseBaseObject(lines []string) BaseObject {
	base := BaseObject{
		MntBy: make([]string, 0),
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Key":
			base.Key = value
		case "created":
			if t, err := time.Parse(TimeLayout, value); err == nil {
				base.Created = t
			}
		case "last-modified":
			if t, err := time.Parse(TimeLayout, value); err == nil {
				base.LastModified = t
			}
		case "source":
			base.Source = value
		case "admin-c":
			base.AdminC = value
		case "tech-c":
			base.TechC = value
		case "mnt-by":
			base.MntBy = append(base.MntBy, value)
		}
	}

	return base
}

func (p *Parser) parseASN(base BaseObject, lines []string) (*ASN, error) {
	asn := &ASN{
		BaseObject:  base,
		Description: make([]string, 0),
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "aut-num":
			asn.ASNumber = value
		case "as-name":
			asn.ASName = value
		case "descr":
			asn.Description = append(asn.Description, value)
		case "org":
			asn.Org = value
		case "status":
			asn.Status = value
		case "notify":
			asn.Notify = value
		}
	}

	return asn, nil
}

func (p *Parser) parseInetNum(base BaseObject, lines []string) (*InetNum, error) {
	inetNum := &InetNum{
		BaseObject:  base,
		Description: make([]string, 0),
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "inetnum":
			inetNum.IPRange = value
		case "netname":
			inetNum.NetName = value
		case "descr":
			inetNum.Description = append(inetNum.Description, value)
		case "country":
			inetNum.Country = value
		case "status":
			inetNum.Status = value
		case "org":
			inetNum.Org = value
		}
	}

	return inetNum, nil
}

func (p *Parser) parseRoute(base BaseObject, lines []string) (*Route, error) {
	route := &Route{BaseObject: base}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "route":
			route.Prefix = value
		case "descr":
			route.Description = value
		case "origin":
			route.Origin = value
		case "org":
			route.Org = value
		}
	}

	return route, nil
}

func (p *Parser) parseRoute6(base BaseObject, lines []string) (*Route6, error) {
	route6 := &Route6{BaseObject: base}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "route6":
			route6.Prefix = value
		case "descr":
			route6.Description = value
		case "origin":
			route6.Origin = value
		case "org":
			route6.Org = value
		}
	}

	return route6, nil
}

func (p *Parser) parsePerson(base BaseObject, lines []string) (*Person, error) {
	person := &Person{
		BaseObject: base,
		Address:    make([]string, 0),
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "person":
			person.Name = value
		case "address":
			person.Address = append(person.Address, value)
		case "phone":
			person.Phone = value
		case "e-mail":
			person.Email = value
		case "nic-hdl":
			person.NicHdl = value
		}
	}

	return person, nil
}

func (p *Parser) parseOrganization(base BaseObject, lines []string) (*Organization, error) {
	org := &Organization{
		BaseObject: base,
		Address:    make([]string, 0),
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "organisation":
			org.OrgID = value
		case "org-name":
			org.Name = value
		case "org-type":
			org.Type = value
		case "address":
			org.Address = append(org.Address, value)
		case "e-mail":
			org.Email = value
		case "abuse-c":
			org.AbuseC = value
		}
	}

	return org, nil
}

func (p *Parser) parseDomain(base BaseObject, lines []string) (*Domain, error) {
	domain := &Domain{
		BaseObject:  base,
		Nameservers: make([]string, 0),
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "domain":
			domain.Domain = value
		case "descr":
			domain.Description = value
		case "nserver":
			domain.Nameservers = append(domain.Nameservers, value)
		case "zone-c":
			domain.ZoneC = value
		}
	}

	return domain, nil
}
