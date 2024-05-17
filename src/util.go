package main

// func isInputFromPipe() bool {
// 	fileInfo, _ := os.Stdin.Stat()
// 	return fileInfo.Mode()&os.ModeCharDevice == 0
// }

// func getStdin() (r string) {
// 	var arr []string
// 	if isInputFromPipe() {
// 		scanner := bufio.NewScanner(os.Stdin)
// 		for {
// 			scanner.Scan()
// 			text := scanner.Text()
// 			if len(text) != 0 {
// 				arr = append(arr, text)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	r = strings.Join(arr, " ")
// 	return
// }

// func pprint(i interface{}) {
// 	s, _ := json.MarshalIndent(i, "", "\t")
// 	fmt.Println(string(s))
// }
