package helpers

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

type AutoGenerated struct {
	ScanMetadata struct {
		CryptoniceVersion string `json:"cryptonice_version"`
		JobID             string `json:"job_id"`
		Hostname          string `json:"hostname"`
		Port              int    `json:"port"`
		NodeName          string `json:"node_name"`
		SitePos           int    `json:"site_pos"`
		HTTPToHTTPS       bool   `json:"http_to_https"`
		Status            string `json:"status"`
		Start             string `json:"start"`
		End               string `json:"end"`
	} `json:"scan_metadata"`
	HTTP struct {
		Connection struct {
			Hostname string `json:"hostname"`
			Path     string `json:"path"`
		} `json:"Connection"`
		Headers struct {
			AccessControlAllowOrigin      string      `json:"Access-Control-Allow-Origin"`
			AccessControlAllowCredentials string      `json:"Access-Control-Allow-Credentials"`
			AccessControlExposeHeaders    interface{} `json:"Access-Control-Expose-Headers"`
			AccessControlMaxAge           interface{} `json:"Access-Control-Max-Age"`
			AccessControlAllowMethods     string      `json:"Access-Control-Allow-Methods"`
			AccessControlAllowHeaders     string      `json:"Access-Control-Allow-Headers"`
			Allow                         interface{} `json:"Allow"`
			ContentEncoding               interface{} `json:"Content-Encoding"`
			ContentLanguage               interface{} `json:"Content-Language"`
			ContentLength                 interface{} `json:"Content-Length"`
			ContentLocation               interface{} `json:"Content-Location"`
			ContentType                   string      `json:"Content-Type"`
			ETag                          interface{} `json:"ETag"`
			Location                      interface{} `json:"Location"`
			Origin                        interface{} `json:"Origin"`
			PublicKeyPins                 interface{} `json:"Public-Key-Pins"`
			Server                        string      `json:"Server"`
			StrictTransportSecurity       interface{} `json:"Strict-Transport-Security"`
			TransferEncoding              string      `json:"Transfer-Encoding"`
			Tk                            interface{} `json:"Tk"`
			Upgrade                       interface{} `json:"Upgrade"`
			Via                           interface{} `json:"Via"`
			WWWAuthenticate               interface{} `json:"WWW-Authenticate"`
			XFrameOptions                 string      `json:"X-Frame-Options"`
			ContentSecurityPolicy         interface{} `json:"Content-Security-Policy"`
			XContentSecurityPolicy        interface{} `json:"X-Content-Security-Policy"`
			XWebKitCSP                    interface{} `json:"X-WebKit-CSP"`
			XPoweredBy                    interface{} `json:"X-Powered-By"`
			XXSSProtection                interface{} `json:"X-XSS-Protection"`
		} `json:"Headers"`
		Cookies struct {
			Cookie1 struct {
				PllLanguage string `json:"pll_language"`
				Expires     string `json:"expires"`
				MaxAge      string `json:"Max-Age"`
				Path        string `json:"path"`
				Secure      string `json:"secure"`
				SecureBool  bool   `json:"Secure"`
			} `json:"cookie_1"`
			Cookie2 struct {
				Serverid string `json:"SERVERID"`
				Path     string `json:"path"`
			} `json:"cookie_2"`
		} `json:"Cookies"`
		Components struct {
			OperatingSystems     []string `json:"Operating systems"`
			Blogs                []string `json:"Blogs"`
			FontScripts          []string `json:"Font scripts"`
			WebServerExtensions  []string `json:"Web server extensions"`
			Seo                  []string `json:"SEO"`
			WebServers           []string `json:"Web servers"`
			PageBuilders         []string `json:"Page builders"`
			UIFrameworks         []string `json:"UI frameworks"`
			TagManagers          []string `json:"Tag managers"`
			ProgrammingLanguages []string `json:"Programming languages"`
		} `json:"Components"`
		Page struct {
			Scripts  []string `json:"scripts"`
			Metatags struct {
				Generator              string `json:"generator"`
				MsapplicationTileimage string `json:"msapplication-tileimage"`
				Viewport               string `json:"viewport"`
			} `json:"metatags"`
		} `json:"Page"`
	} `json:"http"`
	HTTP2 struct {
		HTTP2 bool `json:"http2"`
	} `json:"http2"`
	TLS struct {
		Hostname                       string `json:"hostname"`
		IPAddress                      string `json:"ip_address"`
		CipherSuiteSupported           string `json:"cipher_suite_supported"`
		ClientAuthorizationRequirement string `json:"client_authorization_requirement"`
		HighestTLSVersionSupported     string `json:"highest_tls_version_supported"`
		CertRecommendations            struct {
		} `json:"cert_recommendations"`
		CertificateInfo struct {
			LeafCertificateHasMustStapleExtension           bool        `json:"leaf_certificate_has_must_staple_extension"`
			LeafCertificateIsEv                             bool        `json:"leaf_certificate_is_ev"`
			ReceivedChainHasValidOrder                      bool        `json:"received_chain_has_valid_order"`
			ReceivedChainHasContainsRoot                    bool        `json:"received_chain_has_contains_root"`
			LeafCertificateSignedCertificateTimestampsCount int         `json:"leaf_certificate_signed_certificate_timestamps_count"`
			LeafCertificateSubjectMatchesHostname           bool        `json:"leaf_certificate_subject_matches_hostname"`
			OcspResponse                                    string      `json:"ocsp_response"`
			OcspResponseIsTrusted                           interface{} `json:"ocsp_response_is_trusted"`
			Certificate0                                    struct {
				CommonName         string   `json:"common_name"`
				IssuerName         string   `json:"issuer_name"`
				SerialNumber       string   `json:"serial_number"`
				Fingerprint        string   `json:"fingerprint"`
				PublicKeyAlgorithm string   `json:"public_key_algorithm"`
				PublicKeySize      int      `json:"public_key_size"`
				ValidFrom          string   `json:"valid_from"`
				ValidUntil         string   `json:"valid_until"`
				DaysLeft           int      `json:"days_left"`
				SignatureAlgorithm string   `json:"signature_algorithm"`
				SubjectAltNames    []string `json:"subject_alt_names"`
				CertificateErrors  struct {
					CertTrusted     bool `json:"cert_trusted"`
					HostnameMatches bool `json:"hostname_matches"`
				} `json:"certificate_errors"`
			} `json:"certificate_0"`
		} `json:"certificate_info"`
		Ssl20 struct {
			PreferredCipherSuite      interface{}   `json:"preferred_cipher_suite"`
			AcceptedSsl20CipherSuites []interface{} `json:"accepted_ssl_2_0_cipher_suites"`
		} `json:"ssl_2_0"`
		Ssl30 struct {
			PreferredCipherSuite      interface{}   `json:"preferred_cipher_suite"`
			AcceptedSsl30CipherSuites []interface{} `json:"accepted_ssl_3_0_cipher_suites"`
		} `json:"ssl_3_0"`
		TLS10 struct {
			PreferredCipherSuite      interface{}   `json:"preferred_cipher_suite"`
			AcceptedTLS10CipherSuites []interface{} `json:"accepted_tls_1_0_cipher_suites"`
		} `json:"tls_1_0"`
		TLS11 struct {
			PreferredCipherSuite      interface{} `json:"preferred_cipher_suite"`
			AcceptedTLS11CipherSuites []string    `json:"accepted_tls_1_1_cipher_suites"`
		} `json:"tls_1_1"`
		TLS12 struct {
			PreferredCipherSuite      interface{} `json:"preferred_cipher_suite"`
			AcceptedTLS12CipherSuites []string    `json:"accepted_tls_1_2_cipher_suites"`
		} `json:"tls_1_2"`
		TLS13 struct {
			PreferredCipherSuite      interface{} `json:"preferred_cipher_suite"`
			AcceptedTLS13CipherSuites []string    `json:"accepted_tls_1_3_cipher_suites"`
		} `json:"tls_1_3"`
		Tests struct {
			HTTPHeaders struct {
			} `json:"http_headers"`
		} `json:"tests"`
		ScanInformation struct {
			TLSScanStart       string   `json:"tls_scan_start"`
			TLSScanEnd         string   `json:"tls_scan_end"`
			ScanParameters     []string `json:"scan_parameters"`
			CommandsWithErrors struct {
			} `json:"commands_with_errors"`
		} `json:"scan_information"`
		TLSRecommendations struct {
			HIGHTLSv11 string `json:"HIGH - TLSv1.1"`
		} `json:"tls_recommendations"`
		Jarm struct {
			Fingerprint string `json:"fingerprint"`
		} `json:"jarm"`
	} `json:"tls"`
	DNS struct {
		Connection         string `json:"Connection"`
		DNSRecommendations struct {
			LowCAA string `json:"Low - CAA"`
		} `json:"dns_recommendations"`
		Records struct {
			A   []string      `json:"A"`
			Caa []interface{} `json:"CAA"`
			Txt []string      `json:"TXT"`
			Mx  []string      `json:"MX"`
		} `json:"records"`
	} `json:"dns"`
}
