package main

func reformat(message string, formatter func(string) string) string {

	for i:= 0; i< 3; i++ {
		message = formatter(message)
	}

	return "TEXTIO: " + message 

}
