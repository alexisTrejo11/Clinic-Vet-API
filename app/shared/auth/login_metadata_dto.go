package auth

type LoginMetadata struct {
	ip         string
	userAgent  string
	deviceInfo string
}

func NewLoginMetadata(ip, userAgent, deviceInfo string) LoginMetadata {
	return LoginMetadata{
		ip:         ip,
		userAgent:  userAgent,
		deviceInfo: deviceInfo,
	}
}

func (m LoginMetadata) IP() string         { return m.ip }
func (m LoginMetadata) UserAgent() string  { return m.userAgent }
func (m LoginMetadata) DeviceInfo() string { return m.deviceInfo }
