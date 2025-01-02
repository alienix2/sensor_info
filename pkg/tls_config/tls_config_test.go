package tlsconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCertificates(t *testing.T) {
	certFile, keyFile, caFile := "test-cert.pem", "test-key.pem", "test-ca.pem"

	defer func() {
		os.Remove(certFile)
		os.Remove(keyFile)
		os.Remove(caFile)
	}()

	writeTestFile(t, certFile, testCert)
	writeTestFile(t, keyFile, testKey)
	writeTestFile(t, caFile, testCA)

	tests := []struct {
		name          string
		certFile      string
		keyFile       string
		caFile        string
		expectedError bool
		expectedCerts int
	}{
		{"Valid certificates", certFile, keyFile, caFile, false, 1},
		{"Invalid certificate file", "nonexistent-cert.pem", keyFile, caFile, true, 0},
		{"Invalid key file", certFile, "nonexistent-key.pem", caFile, true, 0},
		{"Invalid CA file", certFile, keyFile, "nonexistent-ca.pem", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tlsConfig, err := LoadCertificates(tt.certFile, tt.keyFile, tt.caFile)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, tlsConfig)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, tlsConfig)
				assert.Len(t, tlsConfig.Certificates, tt.expectedCerts)
			}
		})
	}
}

func writeTestFile(t *testing.T, filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	assert.NoError(t, err)

	// Verify the file content
	writtenContent, err := os.ReadFile(filename)
	assert.NoError(t, err)

	assert.Equal(t, content, string(writtenContent))
}

const testCert = `-----BEGIN CERTIFICATE-----
MIID3DCCAsSgAwIBAgIULXDxL+UPevmYg45jmctf6Fya/zAwDQYJKoZIhvcNAQEL
BQAwgYUxCzAJBgNVBAYTAklUMQ4wDAYDVQQIDAVJdGFseTEQMA4GA1UEBwwHRmly
ZW56ZTEOMAwGA1UECgwFVW5pZmkxDDAKBgNVBAsMA1NTVDESMBAGA1UEAwwJbWF0
dGVtb25pMSIwIAYJKoZIhvcNAQkBFhNtYXR0ZW1vbmlAZ21haWwuY29tMB4XDTI0
MTIyMDE0Mzc0N1oXDTM0MTIxODE0Mzc0N1owgYUxCzAJBgNVBAYTAklUMQ4wDAYD
VQQIDAVJdGFseTEQMA4GA1UEBwwHRmlyZW56ZTEOMAwGA1UECgwFVW5pZmkxDDAK
BgNVBAsMA1NTVDESMBAGA1UEAwwJbWF0dGVtb25pMSIwIAYJKoZIhvcNAQkBFhNt
YXR0ZW1vbmlAZ21haWwuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEApH7JZ/20MWzPSSfEcgjhkt1K8t4/TqcQgLpZ+5MqwZvY8YRb6V9dPQUbqz1L
4EFz1KBL4A3D6lqqXFsNwYzxW9X/r73UEZI+qgDI06q7E3gdJMFHR4qded3ukPPq
f28HSfNZGMiF7yTWFS/Cbev3qB9ytZD4vtV4EZ2ZTtjdE7qtc0XArGtTnzF/Fcuy
1tbnckOvteWeC1IA2HvCCWHcm0RskMTK7+53cF/wJS1hWg/NSmiD5nG9WBGAJUgg
/Gu4j4y9puvBbbd90BR42nqzx4DktXroGVm0iV9aRrTlA5QFANO9huzCu72255Br
Ky7l9Tbvb8SBanCvEnYrZN0WBwIDAQABo0IwQDAdBgNVHQ4EFgQUePEikMgIX+wv
k+LBzEzZ5RkzrGUwHwYDVR0jBBgwFoAUz2YVS5qK6MycyTn7vI+N0FwOHRAwDQYJ
KoZIhvcNAQELBQADggEBAI2Z4wuzK4I+K7pTYQ0AiWGiKnKBztBUWPfXm4rHfw0b
1/nLmR8n35gMiOzLe8Ni5XhSk+zqmX5ZfaPvHR1/64iaCuGhN39qqlIMlMaEfbYh
fMJBEAxASIijk5Tb0ClXEXeZjAIsw5p17/6O/1XWywS9SNCuOjZ0bxEOq33B+4tn
iijZpvOVVb00gjU63FVN6pE+XdzgIUPnAyxW0En9OeW8gt7TnENUaODuBJfcBwLo
Yrs7lmoUihISfkr+raMiv51WdLIlRROQC79nFun/bmRBkaHZHqijdKj4KmFGgOWr
+ad+ztDsfldvAbz68EIRdBc+K0PpgsHATUqo9NeMT30=
-----END CERTIFICATE-----`

const testKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCkfsln/bQxbM9J
J8RyCOGS3Ury3j9OpxCAuln7kyrBm9jxhFvpX109BRurPUvgQXPUoEvgDcPqWqpc
Ww3BjPFb1f+vvdQRkj6qAMjTqrsTeB0kwUdHip153e6Q8+p/bwdJ81kYyIXvJNYV
L8Jt6/eoH3K1kPi+1XgRnZlO2N0Tuq1zRcCsa1OfMX8Vy7LW1udyQ6+15Z4LUgDY
e8IJYdybRGyQxMrv7ndwX/AlLWFaD81KaIPmcb1YEYAlSCD8a7iPjL2m68Ftt33Q
FHjaerPHgOS1eugZWbSJX1pGtOUDlAUA072G7MK7vbbnkGsrLuX1Nu9vxIFqcK8S
ditk3RYHAgMBAAECggEAFCmaRE3bpWdB7SSbtqKSC69mPWiyd9CQfEbbOrPqPLr+
l3Py1CMlSZJztsSXpVBAg2J2imsFhZdBZHqgsAW3m9saMQ9ijBjR54KbGx7FOXiy
kcvDEejSrppeYJQVPTN9wDm8xQXnLM3mE8i72z9eJtRb+hdf9fwycG5R2VFJn95C
cnW8MsYPGvNHo1qM+Jl7KXBCXFN1bpuodd0mi+gJJt4hw1vRe7jJCyJFzQTyIqJJ
XdrKqYRNUz8rDbxNdZKFyn8R10gT7ITcWqZfiyc4YvZPkyd1mtVRb9wl5oA4wTdx
LJ5XhUiAzw+57njT17gIukjcX3E/yh7fSKZfuDz0eQKBgQDbkqmTSubklPuEP7FX
M858nbSoarbOoMbcWokkwVsab8k2ZQudOCTAvcE9HV9/SqjYt+RPtzhmWfBdOD8X
o1qDeOK74H5UrnLCHUUs1toI1FLRjmYzH26XT8mcj9GE2/vZdjjM0JJ/f1zfnwX1
BQ+OCa1W9nHSN7ZgjWXjAvJV8wKBgQC/yPUybosZHztYy+YlnWjb7y/4lWkS2Q+N
3Ve4XigPKtW64zxaw0ivd1WYULo/QdYtMr8Ul8kq5qiRcvepdZ874Dyv5MTVapGa
qpirB4XpsXoIRUqKY/q4nDs1zmc9TvIfXZ9NRUAYe23X8j+fyJ/CaiHoAhU8A3IU
aCNL5oAgnQKBgHYmkCsa9e1wIXtDTqkOzoCN2AV8DsxXBUrTSNLHXL94AXzMmJhL
+rLgKJg7MwTq5rpfEXK3s5iXsthmiMSueOkf5lmUbkYg7M15NJzxK7fukHYEuwet
VTQEkgc1+FcWjImyrNPBM+N5ZD9McccrpvgWSvjtecfVhMSlsXqbPk1zAoGBAJkO
uVkhTdOj0EpCmA9m+7uivWcnXq8TGk2+23Yhdtj794zqM03AUm7uzxn7O1imo1Z5
DHRT2tFpOhiyZyMP4x/3CpfZ/JjSLxf/lE3SeDYUVO7q1d1ygzL4RGzhqBUOvz73
Cd6yKMAhKX3RMKPFinKvHxY5K1c07MOKhLjbYWAlAoGBAMsDqNfx98A6TR85Cx7g
PeB9rtDgTSk+g6Wyke6avaVisj7H1JrHHu7isRZy7lEWdxdbE7nWqGi0hohm4DuO
Tyi3YNCQkVMT7hVYG4hIrJF8+K5ppbzBxaJLKVjFeMS/NZxXdsNavLHRNQm5rGvE
gXwlpLwo6osAiBTlBs/LgHkh
-----END PRIVATE KEY-----`

const testCA = `-----BEGIN CERTIFICATE-----
MIID7TCCAtWgAwIBAgIUbU1txy4YamE1yC405O33WNjttQowDQYJKoZIhvcNAQEL
BQAwgYUxCzAJBgNVBAYTAklUMQ4wDAYDVQQIDAVJdGFseTEQMA4GA1UEBwwHRmly
ZW56ZTEOMAwGA1UECgwFVW5pZmkxDDAKBgNVBAsMA1NTVDESMBAGA1UEAwwJbWF0
dGVtb25pMSIwIAYJKoZIhvcNAQkBFhNtYXR0ZW1vbmlAZ21haWwuY29tMB4XDTI0
MTIyMDE0MzUzM1oXDTM0MTIxODE0MzUzM1owgYUxCzAJBgNVBAYTAklUMQ4wDAYD
VQQIDAVJdGFseTEQMA4GA1UEBwwHRmlyZW56ZTEOMAwGA1UECgwFVW5pZmkxDDAK
BgNVBAsMA1NTVDESMBAGA1UEAwwJbWF0dGVtb25pMSIwIAYJKoZIhvcNAQkBFhNt
YXR0ZW1vbmlAZ21haWwuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEA6l7IWS8p0VSv5J980NLUbUtw1AobxBD6TFZ6YUyjTNbwCSh8U2QP6LYSuAOw
TyL3fWEX4dg8bbYxBRTT8urzfXpBI4jwDOJ6V0Nxjmm+0+xQ50NcvwM2xkFbjEJh
mKf92O1Op3I9K3jCgjnNVxlzv7tWoMUYtEIlWraKM+fJsZm4RTaDIK0N6Xow6Nxj
8PPH2zzuGhhgytLtsInVWWKKchgiqBEeHdzvljItQTeB/FkxE9ozRJtASclKldQz
aINSe0YW6QBaLLH+KBeYWazkqAKlhkAAfRYYOy30euXBh2mlqLZOlZpc13ZkgJO0
nmiQuVJn8ChVfcJxyzDjb8mLpwIDAQABo1MwUTAdBgNVHQ4EFgQUz2YVS5qK6Myc
yTn7vI+N0FwOHRAwHwYDVR0jBBgwFoAUz2YVS5qK6MycyTn7vI+N0FwOHRAwDwYD
VR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEA0NHCnzVo8bhoY54I0+3o
nez+dLqbtbrgNgk08XjZSsoamaiCgHZlVn6lSHctSsET5F0uQU633v/khW6G0Fns
Ig/Sg7O5UzW0xtkp6TouJ8hb7f5g50QdGH2vpSQXeCn3QbmNpIptywfQQ7nFB7gx
5EeHd+tYPlOaGXSiiFjaEf0/EWdD/K96MMwrahQZh8405UropfsU2XiYxRVQXrBc
1zXobRIQcMzQjAZOHni7q7lQUOqp2dG7UmxdQNpMZrYoelSJ9IhIWd/3FzNZ7Jq6
UCQ54F1EpHqrNKhFK2LudYx7uz9uCMTKfqADxSRujRw0mVVvCrxyJ0QVsICp6izZ
rA==
-----END CERTIFICATE-----`
