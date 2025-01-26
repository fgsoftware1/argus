package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"babywolf.io/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/exp/rand"
	"golang.org/x/term"
)

const (
	VERSION string = "1.0"
	AUTHOR  string = "Babywolf(original by Jason13)"
)

var tools = [1][52]map[string]utils.Any{
	{
		{
			"number":  1,
			"name":    "associated hosts",
			"script":  "associated_hosts.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  2,
			"name":    "DNS Over HTTPS",
			"script":  "dns_over_https.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  3,
			"name":    "DNS Records",
			"script":  "dns_records.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  4,
			"name":    "DNSSEC Check",
			"script":  "dnssec.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  5,
			"name":    "Domain Info",
			"script":  "domain_info.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  6,
			"name":    "Domain Reputation Check",
			"script":  "domain_reputation_check.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  7,
			"name":    "IP Info",
			"script":  "ip_info.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  8,
			"name":    "Open Ports Scan",
			"script":  "open_ports.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  9,
			"name":    "Server Info",
			"script":  "server_info.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  10,
			"name":    "Server Location",
			"script":  "server_location.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  11,
			"name":    "SSL Chain Analysis",
			"script":  "ssl_chain.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  12,
			"name":    "SSL Expiry Alert",
			"script":  "ssl_expiry.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  13,
			"name":    "TLS Cipher Suites",
			"script":  "tls_cipher_suites.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  14,
			"name":    "TLS Handshake Simulation",
			"script":  "tls_handshake.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  15,
			"name":    "Traceroute",
			"script":  "traceroute.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  16,
			"name":    "TXT Records",
			"script":  "txt_records.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  17,
			"name":    "WHOIS Lookup",
			"script":  "whois_lookup.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  18,
			"name":    "Zone Transfer",
			"script":  "zonetransfer.go",
			"section": "Network & Infrastructure",
		},
		{
			"number":  19,
			"name":    "Archive History",
			"script":  "archive_history.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  20,
			"name":    "Broken Links Detection",
			"script":  "broken_links.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  21,
			"name":    "Carbon Footprint",
			"script":  "carbon_footprint.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  22,
			"name":    "CMS Detection",
			"script":  "cms_detection.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  23,
			"name":    "Cookies Analyzer",
			"script":  "cookies.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  24,
			"name":    "Content Discovery",
			"script":  "content_discovery.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  25,
			"name":    "Crawler",
			"script":  "crawler.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  26,
			"name":    "Robots.txt Analyzer",
			"script":  "crawl_rules.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  27,
			"name":    "Directory Finder",
			"script":  "directory_finder.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  28,
			"name":    "Performance Monitoring",
			"script":  "performance_monitoring.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  29,
			"name":    "Quality Metrics",
			"script":  "quality_metrics.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  30,
			"name":    "Redirect Chain",
			"script":  "redirect_chain.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  31,
			"name":    "Sitemap Parsing",
			"script":  "sitemap.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  32,
			"name":    "Social Media Presence Scan",
			"script":  "social_media.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  33,
			"name":    "Technology Stack Detection",
			"script":  "technology_stack.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  34,
			"name":    "Third-Party Integrations",
			"script":  "third_party_integrations.go",
			"section": "Web Application Analysis",
		},
		{
			"number":  35,
			"name":    "Censys Reconnaissance",
			"script":  "censys.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  36,
			"name":    "Certificate Authority Recon",
			"script":  "certificate_authority_recon.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  37,
			"name":    "Data Leak Detection",
			"script":  "data_leak.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  38,
			"name":    "Firewall Detection",
			"script":  "firewall_detection.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  39,
			"name":    "Global Ranking",
			"script":  "global_ranking.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  40,
			"name":    "HTTP Headers",
			"script":  "http_headers.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  41,
			"name":    "HTTP Security Features",
			"script":  "http_security.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  42,
			"name":    "Malware & Phishing Check",
			"script":  "malware_phishing.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  43,
			"name":    "Pastebin Monitoring",
			"script":  "pastebin_monitoring.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  44,
			"name":    "Privacy & GDPR Compliance",
			"script":  "privacy_gdpr.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  45,
			"name":    "Security.txt Check",
			"script":  "security_txt.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  46,
			"name":    "Shodan Reconnaissance",
			"script":  "shodan.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  47,
			"name":    "SSL Labs Report",
			"script":  "ssl_labs_report.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  48,
			"name":    "SSL Pinning Check",
			"script":  "ssl_pinning_check.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  49,
			"name":    "Subdomain Enumeration",
			"script":  "subdomain_enum.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  50,
			"name":    "Subdomain Takeover",
			"script":  "subdomain_takeover.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  51,
			"name":    "VirusTotal Scan",
			"script":  "virustotal_scan.go",
			"section": "Security & Threat Intelligence",
		},
		{
			"number":  00,
			"name":    "BEAST MODE",
			"script":  "",
			"section": "Special Mode",
		},
	},
}

var count int = 0

var excludedSections = map[string]struct{}{
	"Run All Scripts": {},
	"Special Mode":    {},
}

var sectionOrder = []string{
	"Network & Infrastructure",
	"Web Application Analysis",
	"Security & Threat Intelligence",
}

func getHTTPHeaders(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	specificHeaders := []string{"X-XSS-Protection", "Content-Security-Policy", "Strict-Transport-Security"}
	for _, header := range specificHeaders {
		if value := resp.Header.Get(header); value != "" {
			fmt.Printf("%s: %s\n", header, value)
		} else {
			fmt.Printf("%s: Not present\n", header)
		}
	}
}

func promptForInput() {
	var input int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	if input == 40 {
		var url string
		fmt.Print("Enter URL or IP: ")
		fmt.Scan(&url)
		getHTTPHeaders(url)
	} else if input == 1 {
		var domain string
		fmt.Print("Enter domain: ")
		fmt.Scan(&domain)
		getAssociatedHosts(domain)
	}
}

func main() {
	Banner()

	sectionMap := make(map[string][]map[string]utils.Any)

	for _, tool := range tools[0] {
		section := tool["section"].(string)
		if _, excluded := excludedSections[section]; !excluded {
			sectionMap[section] = append(sectionMap[section], tool)
		}
	}

	var sections []string
	for _, section := range sectionOrder {
		if _, exists := sectionMap[section]; exists {
			sections = append(sections, section)
		}
	}

	maxRows := 0
	for _, toolsInSection := range sectionMap {
		if len(toolsInSection) > maxRows {
			maxRows = len(toolsInSection)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	headers := make([]interface{}, len(sections))
	for i, section := range sections {
		headers[i] = section
	}
	t.AppendHeader(headers)

	colorize := func(section string, tool map[string]utils.Any) string {
		var numColor, nameColor text.Colors
		switch section {
		case "Network & Infrastructure":
			numColor = text.Colors{text.FgHiBlue}
			nameColor = text.Colors{text.FgBlue}
		case "Web Application Analysis":
			numColor = text.Colors{text.FgHiGreen}
			nameColor = text.Colors{text.FgGreen}
		case "Security & Threat Intelligence":
			numColor = text.Colors{text.FgHiMagenta}
			nameColor = text.Colors{text.FgMagenta}
		default:
			numColor = text.Colors{text.FgWhite}
			nameColor = text.Colors{text.FgWhite}
		}

		numberStr := numColor.Sprint(tool["number"], ")")
		nameStr := nameColor.Sprintf("%s", tool["name"])

		return fmt.Sprintf("%s %s", numberStr, nameStr)
	}

	for row := 0; row < maxRows; row++ {
		toolRow := make([]interface{}, len(sections))
		for col, section := range sections {
			if row < len(sectionMap[section]) {
				tool := sectionMap[section][row]
				toolRow[col] = colorize(section, tool)
			} else {
				toolRow[col] = ""
			}
		}
		t.AppendRow(toolRow)
	}

	t.Render()

	promptForInput()
}

func CountModules() {
	for _, toolRow := range tools {
		for _, tool := range toolRow {
			if script, hasScript := tool["script"]; hasScript && script != nil {
				if section, hasSection := tool["section"]; hasSection {
					if _, excluded := excludedSections[section.(string)]; !excluded {
						count++
					}
				}
			}
		}
	}
}

func Banner() {
	rand.Seed(uint64(time.Now().UnixNano()))

	colors := []func(a ...interface{}) string{
		color.New(color.FgRed, color.Bold).SprintFunc(),
		color.New(color.FgGreen, color.Bold).SprintFunc(),
		color.New(color.FgYellow, color.Bold).SprintFunc(),
		color.New(color.FgBlue, color.Bold).SprintFunc(),
		color.New(color.FgMagenta, color.Bold).SprintFunc(),
		color.New(color.FgCyan, color.Bold).SprintFunc(),
		color.New(color.FgWhite, color.Bold).SprintFunc(),
	}

	cyan := colors[5]
	green := colors[1]
	//blue := colors[3]
	//magenta := colors[4]
	white := colors[6]

	asciiArt := `
		█████╗ ██████╗  ██████╗ ██╗   ██╗███████╗
		██╔══██╗██╔══██╗██╔════╝ ██║   ██║██╔════╝
		███████║██████╔╝██║  ███╗██║   ██║███████╗
		██╔══██║██╔══██╗██║   ██║██║   ██║╚════██║
		██║  ██║██║  ██║╚██████╔╝╚██████╔╝███████║
		╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚══════╝
	`

	lines := utils.SplitLines(asciiArt)

	CountModules()

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	for _, line := range lines {
		randomColor := colors[rand.Intn(len(colors))]
		coloredLine := randomColor(line)
		centeredLine := utils.CenterString(coloredLine, width)
		fmt.Println(centeredLine)
		time.Sleep(50 * time.Millisecond)
	}

	formatted := fmt.Sprintf(
		"%s%s",
		cyan("The Ultimate Information Gathering Tool\n\n"),
		white("\t\t\t\tVersion: "+green(VERSION))+white("\tModules: "+green(strconv.Itoa(count)))+white("\tCoded by: "+green(AUTHOR)),
	)

	centeredFormatted := utils.CenterString(formatted, width)

	fmt.Println(centeredFormatted)
}
