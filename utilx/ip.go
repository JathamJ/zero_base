package utilx

import (
	"net"
	"net/http"
	"strings"
)

// ClientIP 获取客户端ip
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func GetUserAgent(req *http.Request) string {
	return req.Header.Get("User-Agent")
}

// IpIsInternal 检测IP地址是否为内网IP
func IpIsInternal(ip string) bool {
	if ip == "" {
		return false
	}
	rip := net.ParseIP(ip)
	// 检查是否位于私有IP范围内
	privateRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"169.254.0.0/16", // 链路本地
		"127.0.0.0/8",    // 本地回环
		"::1/128",        // IPv6 本地回环
		"fe80::/10",      // IPv6 链路本地
		"fc00::/7",
	}

	for _, r := range privateRanges {
		_, network, _ := net.ParseCIDR(r)
		if network.Contains(rip) {
			return true
		}
	}

	return false
}

// DirectClientIP 直接获取客户端ip
func DirectClientIP(r *http.Request) string {

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
