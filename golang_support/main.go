package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// regexes
var url_regex, _ = regexp.Compile(`https?://[^\s"\'<>]+`)
var ip_regex, _ = regexp.Compile(`b(?:\d{1,3}\.){3}\d{1,3}\b`)
var email_regex, _ = regexp.Compile(`b[\w.-]+@[\w.-]+\.\w+\b`)
var md5_regex, _ = regexp.Compile(`\b[a-fA-F0-9]{32}\b`)
var sha1_regex, _ = regexp.Compile(`\b[a-fA-F0-9]{40}\b`)
var sha256_regex, _ = regexp.Compile(`https?://[^\s"\'<>]+`)

func defang(text string) string {
	log.Println("[+] defanging data")

	var t = strings.Replace(text, "http", "hxxp", -1)
	t = strings.Replace(t, ".", "[.]", -1)
	t = strings.Replace(t, "@", "[@]", -1)

	return t
}

func refang(text string) string {
	var t = strings.Replace(text, "hxxp", "http", -1)
	t = strings.Replace(t, "[.]", ".", -1)
	t = strings.Replace(t, "[@]", "@", -1)

	return t
}

func extract_iocs(text_blob string) []string {
	var iocs []string = make([]string, 0)
	var urls = url_regex.FindAllString(text_blob, -1)
	var ips = ip_regex.FindAllString(text_blob, -1)
	var emails = email_regex.FindAllString(text_blob, -1)
	var md5s = md5_regex.FindAllString(text_blob, -1)
	var sha1s = sha1_regex.FindAllString(text_blob, -1)
	var sha256s = sha256_regex.FindAllString(text_blob, -1)

	iocs = append(iocs, urls...)
	iocs = append(iocs, ips...)
	iocs = append(iocs, emails...)
	iocs = append(iocs, md5s...)
	iocs = append(iocs, sha1s...)
	iocs = append(iocs, sha256s...)

	fmt.Println(iocs)
	return iocs
}

func get_input() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func get_lines(file_name string) (result []string, e error) {
	log.Printf("[+] reading lines from file %s\n", file_name)
	var file, err = os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Unable to read file!")
		return nil, err
	}

	var file_data = string(file)

	var file_lines = strings.Split(file_data, "\n")
	var temp_lines []string
	for _, v := range file_lines {
		log.Println(v)
		temp_lines = append(temp_lines, strings.Trim(v, " "))
	}
	file_lines = temp_lines
	return file_lines, nil
}

func _get_manual_lines() []string {
	fmt.Println("Paste text or enter [done] to quit input loop")
	var lines []string
	for {
		var line = get_input()
		if line == "done" {
			break
		}
		var fline = strings.Replace(line, "\n", "", -1)
		lines = append(lines, fline)
	}

	return lines
}

func save_output(data []string, output string) {

	var f, e = os.Create(output)

	if e != nil {
		log.Fatalln(e)
	}

	defer f.Close()

	for _, v := range data {
		var _, we = fmt.Fprintln(f, v)
		if we != nil {
			log.Fatalln(we)
		}
	}

}

func main() {
	fmt.Println("[+] Golang FangShepherd Version 1.0 by mwcsur")
	fmt.Println("")

	fmt.Println("Choose action:\n1. Extract IOCs + Defang\n2. Extract IOCs + Refang\n3. Just Defang\n4. Just Refang\n> ")

	var input string
	input = get_input()

	log.Printf("[+] selected %s", input)

	var input_source string
	fmt.Println("Enter 1 for text and 2 for text file")
	input_source = get_input()

	var input_lines []string
	var final_input_lines []string

	if input[0] == '1' || input[0] == '2' {

		log.Printf("[+] Input source choice: %s\n", input_source)

		if input_source == "1" {
			log.Println("[+] Using manual line input source")

			input_lines = _get_manual_lines()
		} else {
			log.Println("[+] Using file input source")
			fmt.Println("Enter text file")
			var input_file = get_input()
			input_lines, _ = get_lines(input_file)
		}

		if input == "1" {
			//defang
			log.Println("[+] defanging data")

			for _, line := range input_lines {
				var ioc_lines = extract_iocs(line)
				for _, v := range ioc_lines {
					var defang_line = defang(v)
					final_input_lines = append(final_input_lines, defang_line)
				}
			}
		} else {
			//refang
			log.Println("[+] refanging data")

			for _, line := range input_lines {
				var ioc_lines = extract_iocs(line)
				for _, v := range ioc_lines {
					var refang_line = refang(v)
					fmt.Println(refang_line)
					final_input_lines = append(final_input_lines, refang_line)
				}
			}
		}
	} else if input[0] == '3' || input[0] == '4' {

		if input_source == "1" {
			log.Println("[+] Using manual line input source")

			input_lines = _get_manual_lines()
		} else {
			log.Println("[+] Using file input source")
			fmt.Println("Enter text file")
			var input_file = get_input()
			input_lines, _ = get_lines(input_file)
		}

		if input[0] == '3' {
			//just defang lines
			log.Println("[+] defanging data")

			for _, line := range input_lines {
				var defang_line = defang(line)
				final_input_lines = append(final_input_lines, defang_line)
			}
		} else {
			//just refang lines
			log.Println("[+] refanging data")

			for _, line := range input_lines {
				var defang_line = refang(line)
				final_input_lines = append(final_input_lines, defang_line)
			}
		}
	} else {
		log.Fatalln("[-] Invalid  choice")
	}

	log.Println("[+] Results")
	for _, v := range final_input_lines {
		fmt.Println(v)
	}

	fmt.Println("Save file? enter y/n")
	var save_file = get_input()

	if save_file[0] == 'y' {
		//save output
		fmt.Println("Enter output file name")
		var output_file = get_input()

		log.Println("[+] saving file output!")
		save_output(final_input_lines, output_file)
	} else if save_file[0] == 'n' {

	} else {
		log.Printf("[-] Unknown option %s", save_file)
	}

}
