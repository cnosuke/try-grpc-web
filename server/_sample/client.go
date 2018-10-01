package main

import (
	"context"
	"github.com/cnosuke/try-grpc-web/server/pb"

	"github.com/k0kubun/pp"
	"google.golang.org/grpc"
)

//const address = "shire.alpha.api-theo.blue:443"
const address = "127.0.0.1:8888"

func main() {
//	ca_pem := `-----BEGIN CERTIFICATE-----
//	MIIEZTCCA02gAwIBAgIJAMbPO5z3Q0rdMA0GCSqGSIb3DQEBBQUAMH4xCzAJBgNV
//	BAYTAkpQMQ4wDAYDVQQIEwVUb2t5bzESMBAGA1UEBxMJTWluYXRvLWt1MRQwEgYD
//	VQQKEwtNb25leURlc2lnbjENMAsGA1UECxMEVGVjaDEmMCQGCSqGSIb3DQEJARYX
//	c3lzb3BzQG1vbmV5LWRlc2lnbi5jb20wHhcNMTcwNjA4MDk0NTMzWhcNMjcwNjA2
//	MDk0NTMzWjB+MQswCQYDVQQGEwJKUDEOMAwGA1UECBMFVG9reW8xEjAQBgNVBAcT
//	CU1pbmF0by1rdTEUMBIGA1UEChMLTW9uZXlEZXNpZ24xDTALBgNVBAsTBFRlY2gx
//	JjAkBgkqhkiG9w0BCQEWF3N5c29wc0Btb25leS1kZXNpZ24uY29tMIIBIjANBgkq
//	hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3trQcbkEH873YYLyWqkjKbVch1+5kxbq
//	TcRD6W855yVcCCAMUhEWKHnJFKQUwZcsSu+bSU3aw0IblaqL5QBoD2ztUVPIogtj
//	45otvp0t/YguL2+SJc35+Qv6+tHg2nHQMSLunTmzDdAfqYs9h9MPi+I/RMimUs7f
//	jDtO80Wtz6wImzPzeHnraqqwCsfrqyzACZRBItVUy0YY3L6FnEiOJLpsH/G3yOhY
//	/nVgaOXb+8H2FF8j0yEGXOb8ULH3L1YhOJNkIv+ZNS+IlJ9k0P0Jub6aARQJbsnb
//	4cjLt1i5+SEyi+kcONTmKBKD2rUB5suROYvG/ei/4F3zGMgnISCnWQIDAQABo4Hl
//	MIHiMB0GA1UdDgQWBBTcIIE8s8Q4PPe0VUHG7PWXkBqB6TCBsgYDVR0jBIGqMIGn
//	gBTcIIE8s8Q4PPe0VUHG7PWXkBqB6aGBg6SBgDB+MQswCQYDVQQGEwJKUDEOMAwG
//	A1UECBMFVG9reW8xEjAQBgNVBAcTCU1pbmF0by1rdTEUMBIGA1UEChMLTW9uZXlE
//	ZXNpZ24xDTALBgNVBAsTBFRlY2gxJjAkBgkqhkiG9w0BCQEWF3N5c29wc0Btb25l
//	eS1kZXNpZ24uY29tggkAxs87nPdDSt0wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B
//	AQUFAAOCAQEAsY0p6DsdPH4AA/0lC+bGnW94duNqWaxhSuH7UoHvvfAs8A34NcOT
//	hvEF9TS9ULCX4855/4rLpHHp6fnu59lmk+XZMTZ0xsEJvnYT/6yDV1HmstSfZ6gS
//	xvssZgBCy+B85ETMGjh0jVDY+700QVaDuO4kV5M6sR9XbRF1j3aRmgIalpAN2MRW
//	v39/fGJs4TLr2TOK3QQO04fb7EGV/EnubmlRrCHWuWSCRlS+wm/3VMNDgUgB9+fb
//	YFle/mneuG12DthCHvGjbA2FwjldFYIWoLq8xVbQS36wqx4u6llLo5ZQmsmI7tj9
//	cfOpM9Ml2oUzVtoHjnUQA4LxzTDHRwwM+Q==
//-----END CERTIFICATE-----`

	//roots := x509.NewCertPool()
	//roots.AppendCertsFromPEM([]byte(ca_pem))
	//
	//cred := credentials.NewTLS(&tls.Config{
	//	RootCAs: roots,
	//})
	//
	//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred))
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		pp.Fatalln(err)
	}

	defer conn.Close()

	c := todo.NewTodoServiceClient(conn)

	ctx := context.Background()

	req := todo.GetLatestRequest{}

	res, err := c.GetLatest(ctx, &req)

	if err != nil {
		pp.Fatalln(err)
	}


	pp.Println(res)
}
