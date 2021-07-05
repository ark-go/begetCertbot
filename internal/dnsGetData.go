package internal

type dnsDataVal struct {
	Value    string
	Priority int64
}

type dnsData struct {
	Result struct {
		Is_under_control int64    // 1, // домен на обслуживании BeGet (0 - нет / 1 - да)
		Is_beget_dns     int64    //1,     // домен на DNS-серверах BeGet (0 - нет / 1 - да)
		Is_subdomain     int64    // 0,     // является ли домен поддоменом (0 - нет / 1 - да)
		Fqdn             string   // переданное доменное имя
		Records          struct { // текущие используемые DNS-записи
			DNS    []dnsDataVal
			DNS_IP []dnsDataVal
			A      []struct {
				Ttl     int64
				Address string
			}
			MX []struct {
				Ttl        int64
				Exchange   string
				Preference int64
			}
			TXT []struct {
				Ttl     int64
				Txtdata string
			}
		}
		//  тип текущих используемых настроек:
		// 1 - используются A, MX, TXT-записи;
		// 2 - используются NS-записи (для поддоменов);
		// 3 - используются CNAME-записи (для поддоменов).
		Set_type int64 //1
	}
}
