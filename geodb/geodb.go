package geodb

import (
	"fmt"
	"net/netip"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	DomainKeyFull    = "full"
	DomainKeyKeyword = "keyword"
	DomainKeySuffix  = "suffix"
	DomainKeyRegex   = "regex"
	DomainKeyGeoIP   = "geoip"
	DomainKeyGeoSite = "geosite"
)

var DataDir = "./"

type Param struct {
	Key   string
	Val   string
	Regex *regexp.Regexp
}

// 匹配器
type Matcher interface {
	Match(code string) bool
}

// 域名匹配器
type DomainMatcher struct {
	Code   string
	Params []*Param
}

// 域名匹配
func (m *DomainMatcher) Match(domain string) bool {
	domain = strings.ToLower(domain)
	for _, param := range m.Params {
		switch param.Key {
		case DomainKeyFull:
			if strings.EqualFold(param.Val, domain) {
				return true
			}
		case DomainKeyKeyword:
			if strings.Contains(domain, param.Val) {
				return true
			}
		case DomainKeySuffix:
			if domain == param.Val || strings.HasSuffix(domain, "."+strings.TrimPrefix(param.Val, ".")) {
				return true
			}
		case DomainKeyRegex:
			if param.Regex.MatchString(domain) {
				return true
			}
		}
	}
	return false
}

// IP 匹配器
type IPMatcher struct {
	Code   string
	Params []*Param
}

// IP 匹配
func (m *IPMatcher) Match(code string) bool {
	return false
}

func Site(filename string, code string) (Matcher, error) {
	if !strings.HasSuffix(filename, ".dat") {
		filename += ".dat"
	}
	filePath, err := filepath.Abs(filepath.Join(DataDir, filename))
	if err != nil {
		return nil, err
	}
	geoSite, err := UnmarshalGeoSite(filePath, code)
	if err != nil {
		return nil, err
	}
	var params []*Param
	code, attr, _ := strings.Cut(code, "@")
	for _, item := range geoSite.Domain {
		if attr != "" {
			// Filter by attr.
			attrHit := false
			for _, itemAttr := range item.Attribute {
				if strings.EqualFold(itemAttr.Key, attr) {
					attrHit = true
					break
				}
			}
			if !attrHit {
				continue
			}
		}

		switch item.Type {
		case Domain_Full:
			// Full.
			params = append(params, &Param{
				Key: string(DomainKeyFull),
				Val: strings.ToLower(item.Value),
			})
		case Domain_RootDomain:
			// Suffix.
			params = append(params, &Param{
				Key: string(DomainKeySuffix),
				Val: strings.ToLower(item.Value),
			})
		case Domain_Plain:
			// Keyword.
			params = append(params, &Param{
				Key: string(DomainKeyKeyword),
				Val: strings.ToLower(item.Value),
			})
		case Domain_Regex:
			// Regex.
			regex, err := regexp.Compile(strings.ToLower(item.Value))
			if err != nil {
				return nil, err
			}
			params = append(params, &Param{
				Key:   string(DomainKeyRegex),
				Val:   item.Value,
				Regex: regex,
			})
		}
	}
	return &DomainMatcher{Code: code, Params: params}, nil
}

func IP(filename string, code string) (Matcher, error) {
	if !strings.HasSuffix(filename, ".dat") {
		filename += ".dat"
	}
	filePath, err := filepath.Abs(filepath.Join(DataDir, filename))
	if err != nil {
		return nil, err
	}
	geoIp, err := UnmarshalGeoIp(filePath, code)
	if err != nil {
		return nil, err
	}
	if geoIp.InverseMatch {
		return nil, fmt.Errorf("not support inverse match yet")
	}
	var params []*Param
	for _, item := range geoIp.Cidr {
		ip, ok := netip.AddrFromSlice(item.Ip)
		if !ok {
			return nil, fmt.Errorf("bad geoip file: %v", filename)
		}
		params = append(params, &Param{
			Key: "",
			Val: netip.PrefixFrom(ip, int(item.Prefix)).String(),
		})
	}
	return &IPMatcher{Code: code, Params: params}, nil
}
