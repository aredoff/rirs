package parser

import "time"

// BaseObject contains common fields for all RIPE objects
type BaseObject struct {
	Key          string
	Created      time.Time
	LastModified time.Time
	Source       string
	AdminC       string
	TechC        string
	MntBy        []string
}

// ASN represents an Autonomous System Number object
type ASN struct {
	BaseObject
	ASNumber    string
	ASName      string
	Description []string
	Org         string
	Status      string
	Notify      string
}

// InetNum represents an IP address range object
type InetNum struct {
	BaseObject
	IPRange     string
	NetName     string
	Description []string
	Country     string
	Status      string
	Org         string
}

// Route represents a route object
type Route struct {
	BaseObject
	Prefix      string
	Description string
	Origin      string
	Org         string
}

// Route6 represents an IPv6 route object
type Route6 struct {
	BaseObject
	Prefix      string
	Description string
	Origin      string
	Org         string
}

// Person represents a person object
type Person struct {
	BaseObject
	Name    string
	Address []string
	Phone   string
	Email   string
	NicHdl  string
}

// Organization represents an organization object
type Organization struct {
	BaseObject
	Name    string
	Type    string
	Address []string
	Email   string
	AbuseC  string
	OrgID   string
}

// Domain represents a domain object
type Domain struct {
	BaseObject
	Domain      string
	Description string
	Nameservers []string
	ZoneC       string
}

// RipeDatabase represents the complete database
type RipeDatabase struct {
	ASNs          map[string]*ASN
	InetNums      map[string]*InetNum
	Routes        map[string]*Route
	Routes6       map[string]*Route6
	Persons       map[string]*Person
	Organizations map[string]*Organization
	Domains       map[string]*Domain
}
