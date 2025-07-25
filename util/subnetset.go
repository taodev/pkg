package util

import "net/netip"

// IsLocallyServed checks if ip belongs to any network defined by [RFC 6303]:
//
//	10.0.0.0/8
//	127.0.0.0/8
//	169.254.0.0/16
//	172.16.0.0/12
//	192.0.2.0/24
//	192.168.0.0/16
//	198.51.100.0/24
//	203.0.113.0/24
//	255.255.255.255/32
//
//	::/128
//	::1/128
//	2001:db8::/32
//	fd00::/8
//	fe80::/10
//
// It may also be used as a [SubnetSetFunc].
//
// [RFC 6303]: https://datatracker.ietf.org/doc/html/rfc6303
func IsLocallyServed(ip netip.Addr) (ok bool) {
	if !ip.IsValid() {
		return false
	}

	if ip.Is4() {
		return isLocallyServedV4(ip.As4())
	}

	return isLocallyServedV6(ip.As16())
}

// isLocallyServedV6 returns true if ip belongs to at least one of networks
// listed in [RFC 6303].  The ip is expected to be a valid IPv6.
//
// See also [IsLocallyServed].
//
// [RFC 6303]: https://datatracker.ietf.org/doc/html/rfc6303
func isLocallyServedV6(ip [16]byte) (ok bool) {
	switch ip[0] {
	case 0x00:
		return string(ip[1:15]) == "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" &&
			ip[15]&0xFE == 0x00
	case 0x20:
		return string(ip[1:4]) == "\x01\x0D\xB8"
	case 0xFE:
		return ip[1]&0xC0 == 0x80
	default:
		return ip[0] == 0xFD
	}
}

// isLocallyServedV4 returns true if ip belongs to at least one of networks
// listed in [RFC 6303].  The ip is expected to be a valid IPv4.
//
// See also [IsLocallyServed].
//
// [RFC 6303]: https://datatracker.ietf.org/doc/html/rfc6303
func isLocallyServedV4(ip [4]byte) (ok bool) {
	switch ip[0] {
	case 10, 127:
		return true
	case 169:
		return ip[1] == 254
	case 172:
		return ip[1]&0xF0 == 16
	case 192:
		return ip[1] == 168 || string(ip[1:3]) == "\x00\x02"
	case 198:
		return string(ip[1:3]) == "\x33\x64"
	case 203:
		return string(ip[1:3]) == "\x00\x71"
	default:
		return string(ip[:]) == "\xFF\xFF\xFF\xFF"
	}
}

// IsSpecialPurpose checks if ip belongs to any network defined by IANA
// Special-Purpose Address Registry:
//
//	0.0.0.0/8          "This host on this network".
//	10.0.0.0/8         Private-Use.
//	100.64.0.0/10      Shared Address Space.
//	127.0.0.0/8        Loopback.
//	169.254.0.0/16     Link Local.
//	172.16.0.0/12      Private-Use.
//	192.0.0.0/24       IETF Protocol Assignments.
//	192.0.0.0/29       DS-Lite.
//	192.0.2.0/24       Documentation (TEST-NET-1)
//	192.88.99.0/24     6to4 Relay Anycast.
//	192.168.0.0/16     Private-Use.
//	198.18.0.0/15      Benchmarking.
//	198.51.100.0/24    Documentation (TEST-NET-2).
//	203.0.113.0/24     Documentation (TEST-NET-3).
//	240.0.0.0/4        Reserved.
//	255.255.255.255/32 Limited Broadcast.
//
//	::/128            Unspecified Address.
//	::1/128           Loopback Address.
//	64:ff9b::/96      IPv4-IPv6 Translation Address.
//	64:ff9b:1::/48    IPv4-IPv6 Translation Address.
//	100::/64          Discard-Only Address Block.
//	2001::/23         IETF Protocol Assignments.
//	2001::/32         TEREDO.
//	2001:1::1/128     Port Control Protocol Anycast.
//	2001:1::2/128     Traversal Using Relays around NAT Anycast.
//	2001:2::/48       Benchmarking.
//	2001:3::/32       AMT.
//	2001:4:112::/48   AS112-v6.
//	2001:10::/28      ORCHID.
//	2001:20::/28      ORCHIDv2.
//	2001:db8::/32     Documentation.
//	2002::/16         6to4.
//	2620:4f:8000::/48 Direct Delegation AS112 Service.
//	fc00::/7          Unique-Local.
//	fe80::/10         Linked-Scoped Unicast.
//
// See https://www.iana.org/assignments/iana-ipv4-special-registry and
// https://www.iana.org/assignments/iana-ipv6-special-registry.
//
// It may also be used as a [SubnetSetFunc].
func IsSpecialPurpose(ip netip.Addr) (ok bool) {
	if !ip.IsValid() {
		return false
	}

	if ip.Is4() {
		return isSpecialPurposeV4(ip.As4())
	}

	return isSpecialPurposeV6(ip.As16())
}

// isSpecialPurposeV6 returns true if ip belongs to at least one of networks
// from special-purpose address registries.  The ip is expected to be a valid
// IPv6.
//
// See also [IsSpecialPurpose].
func isSpecialPurposeV6(ip [16]byte) (ok bool) {
	switch ip[0] {
	case 0x00:
		ok = string(ip[1:5]) == "\x64\xFF\x9B\x00" &&
			(ip[5] == 0x01 || string(ip[5:12]) == "\x00\x00\x00\x00\x00\x00\x00")
	case 0x20:
		if ip[1] == 0x01 {
			ok = ip[2]&0xFE == 0x00
		} else {
			ok = ip[1] == 0x02
		}
	default:
		ok = string(ip[:6]) == "\x26\x20\x00\x4F\x80\x00" ||
			string(ip[:8]) == "\x01\x00\x00\x00\x00\x00\x00\x00" ||
			ip[0] == 0xFC || ip[0] == 0xFD
	}

	return ok || isLocallyServedV6(ip)
}

// isSpecialPurposeV4 returns true if ip belongs to at least one of networks
// from special-purpose address registries.  The ip is expected to be a valid
// IPv4.
//
// See also [IsSpecialPurpose].
func isSpecialPurposeV4(ip [4]byte) (ok bool) {
	switch ip[0] {
	case 0:
		return true
	case 100:
		ok = ip[1]&0xC0 == 64
	case 192:
		ok = string(ip[1:3]) == "\x00\x00" || string(ip[1:3]) == "\x58\x63"
	case 198:
		ok = ip[1]&0xFE == 18
	default:
		ok = ip[0] >= 0xF0
	}

	return ok || isLocallyServedV4(ip)
}
