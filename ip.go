package govalidator

import "net"

const (
	// IP4 represents rule name which will be used to find the default error message.
	IP4 = "ip4"
	// IP4Msg is the default error message format for fields with IP4 validation rule.
	IP4Msg = "%s should be a valid ipv4"
)

// IP4 checks if given string is a valid IPV4.
//
// Example:
//
//	v := validator.New()
//	v.IP4("127.0.0.1", "server_ip", "server_ip should be a valid ipv4.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v *Validator) IP4(s, field, msg string) *Validator {
	ip := net.ParseIP(s)

	v.Check(ip != nil && ip.To4() != nil, field, v.msg(IP4, msg, field))

	return v
}
