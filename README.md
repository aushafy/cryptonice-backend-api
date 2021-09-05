# Mock of Cryptonice RESTful API
## How To
1. Dependency Requirements
- Cryptonice
2. Install Dependency
- Pip install cryptonice
3. Clone this repo
- git clone https://github.com/aushafy/cryptonice-backend-api.git
4. Run the application
- cd cryptonice-backend-api && go run main.go

## Example Usage
### /api/v1/scan/:URL

* change URL with Public DNS
- curl -k -X GET localhost:8000/api/v1/scan/google.com

### Example Return JSON
```{
  "meta": {
    "message": "HTTPS Scan Result",
    "code": 200,
    "status": "success"
  },
  "data": {
    "scan_metadata": {
      "cryptonice_version": "1.4.2.3",
      "job_id": "default",
      "hostname": "google.com",
      "port": 443,
      "node_name": "sirclo-Lenovo-ideapad-330-14IKB",
      "site_pos": 0,
      "http_to_https": false,
      "status": "Successful",
      "start": "2021-09-05 15:29:30.803503",
      "end": "2021-09-05 15:29:52.117085"
    },
    "http": {
      "Connection": {
        "hostname": "www.google.com",
        "path": "/"
      },
      "Headers": {
        "Access-Control-Allow-Origin": "",
        "Access-Control-Allow-Credentials": "",
        "Access-Control-Expose-Headers": null,
        "Access-Control-Max-Age": null,
        "Access-Control-Allow-Methods": "",
        "Access-Control-Allow-Headers": "",
        "Allow": null,
        "Content-Encoding": null,
        "Content-Language": null,
        "Content-Length": null,
        "Content-Location": null,
        "Content-Type": "text/html; charset=ISO-8859-1",
        "ETag": null,
        "Location": null,
        "Origin": null,
        "Public-Key-Pins": null,
        "Server": "gws",
        "Strict-Transport-Security": null,
        "Transfer-Encoding": "chunked",
        "Tk": null,
        "Upgrade": null,
        "Via": null,
        "WWW-Authenticate": null,
        "X-Frame-Options": "SAMEORIGIN",
        "Content-Security-Policy": null,
        "X-Content-Security-Policy": null,
        "X-WebKit-CSP": null,
        "X-Powered-By": null,
        "X-XSS-Protection": "0"
      },
      "Cookies": {
        "cookie_1": {
          "pll_language": "",
          "expires": "05-Oct-2021 08:29:32 GMT",
          "Max-Age": "",
          "path": "/",
          "secure": "",
          "Secure": true
        },
        "cookie_2": {
          "SERVERID": "",
          "path": "/"
        }
      },
      "Components": {
        "Operating systems": null,
        "Blogs": null,
        "Font scripts": null,
        "Web server extensions": null,
        "SEO": null,
        "Web servers": [
          "Google Web Server"
        ],
        "Page builders": null,
        "UI frameworks": null,
        "Tag managers": null,
        "Programming languages": null
      },
      "Page": {
        "scripts": [],
        "metatags": {
          "generator": "",
          "msapplication-tileimage": "",
          "viewport": ""
        }
      }
    },
    "http2": {
      "http2": true
    },
    "tls": {
      "hostname": "www.google.com",
      "ip_address": "172.217.194.103",
      "cipher_suite_supported": "TLS_AES_256_GCM_SHA384",
      "client_authorization_requirement": "DISABLED",
      "highest_tls_version_supported": "TLS_1_3",
      "cert_recommendations": {},
      "certificate_info": {
        "leaf_certificate_has_must_staple_extension": false,
        "leaf_certificate_is_ev": false,
        "received_chain_has_valid_order": true,
        "received_chain_has_contains_root": false,
        "leaf_certificate_signed_certificate_timestamps_count": 2,
        "leaf_certificate_subject_matches_hostname": true,
        "ocsp_response": "",
        "ocsp_response_is_trusted": null,
        "certificate_0": {
          "common_name": "www.google.com",
          "issuer_name": "GTS CA 1C3",
          "serial_number": "288799108331868243045966257528948549806",
          "fingerprint": "bd46901febde968b713e111925436865bc98485246392c1da998f283bcb3da96",
          "public_key_algorithm": "EllipticCurve",
          "public_key_size": 256,
          "valid_from": "2021-08-16 03:56:32",
          "valid_until": "2021-11-08 03:56:31",
          "days_left": 63,
          "signature_algorithm": "sha256",
          "subject_alt_names": [
            "www.google.com"
          ],
          "certificate_errors": {
            "cert_trusted": true,
            "hostname_matches": true
          }
        }
      },
      "ssl_2_0": {
        "preferred_cipher_suite": null,
        "accepted_ssl_2_0_cipher_suites": []
      },
      "ssl_3_0": {
        "preferred_cipher_suite": null,
        "accepted_ssl_3_0_cipher_suites": []
      },
      "tls_1_0": {
        "preferred_cipher_suite": null,
        "accepted_tls_1_0_cipher_suites": [
          "TLS_RSA_WITH_AES_256_CBC_SHA",
          "TLS_RSA_WITH_AES_128_CBC_SHA",
          "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
          "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
          "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"
        ]
      },
      "tls_1_1": {
        "preferred_cipher_suite": null,
        "accepted_tls_1_1_cipher_suites": [
          "TLS_RSA_WITH_AES_256_CBC_SHA",
          "TLS_RSA_WITH_AES_128_CBC_SHA",
          "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
          "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
          "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"
        ]
      },
      "tls_1_2": {
        "preferred_cipher_suite": null,
        "accepted_tls_1_2_cipher_suites": [
          "TLS_RSA_WITH_AES_256_GCM_SHA384",
          "TLS_RSA_WITH_AES_256_CBC_SHA",
          "TLS_RSA_WITH_AES_128_GCM_SHA256",
          "TLS_RSA_WITH_AES_128_CBC_SHA",
          "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
          "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
          "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
          "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
          "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
          "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
          "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"
        ]
      },
      "tls_1_3": {
        "preferred_cipher_suite": null,
        "accepted_tls_1_3_cipher_suites": [
          "TLS_CHACHA20_POLY1305_SHA256",
          "TLS_AES_256_GCM_SHA384",
          "TLS_AES_128_GCM_SHA256"
        ]
      },
      "tests": {
        "http_headers": {}
      },
      "scan_information": {
        "tls_scan_start": "2021-09-05 15:29:33.496617",
        "tls_scan_end": "2021-09-05 15:29:51.435110",
        "scan_parameters": [
          "certificate_info",
          "ssl_2_0_cipher_suites",
          "ssl_3_0_cipher_suites",
          "tls_1_0_cipher_suites",
          "tls_1_1_cipher_suites",
          "tls_1_2_cipher_suites",
          "tls_1_3_cipher_suites",
          "http_headers"
        ],
        "commands_with_errors": {}
      },
      "tls_recommendations": {
        "HIGH - TLSv1.1": "Major browsers are disabling this TLS 1.1 immenently. Carefully monitor if clients still use this protocol. "
      },
      "jarm": {
        "fingerprint": "27d40d40d29d40d1dc42d43d00041d4689ee210389f4f6b4b5b1b93f92252d"
      }
    },
    "dns": {
      "Connection": "google.com",
      "dns_recommendations": {
        "Low - CAA": ""
      },
      "records": {
        "A": [
          "216.239.38.120"
        ],
        "CAA": [
          "0 issue \"pki.goog\""
        ],
        "TXT": [
          "\"facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95\"",
          "\"google-site-verification=wD8N7i1JTNTkezJ49swvWW48f8_9xveREV4oB-0Hf5o\"",
          "\"globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=\"",
          "\"docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e\"",
          "\"v=spf1 include:_spf.google.com ~all\"",
          "\"apple-domain-verification=30afIBcvSuDV2PLX\"",
          "\"MS=E4A68B9AB2BB9670BCE15412F62916164C0B20BB\"",
          "\"google-site-verification=TV9-DBe4R80X4v0M4U_bd_J9cpOJM0nikft0jAgjmsQ\"",
          "\"docusign=1b0a6754-49b1-4db5-8540-d2c12664b289\""
        ],
        "MX": [
          "40 alt3.aspmx.l.google.com.",
          "50 alt4.aspmx.l.google.com.",
          "10 aspmx.l.google.com.",
          "30 alt2.aspmx.l.google.com.",
          "20 alt1.aspmx.l.google.com."
        ]
      }
    }
  }
}```