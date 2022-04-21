package railfence

func Encode(message string, rails int) string {
	var arrail = make([]string, rails)
	currentRail := 0
	directionDown := true
	for _, v := range message {
		arrail[currentRail] += string(v)
		if directionDown {
			currentRail++
		} else {
			currentRail--
		}
		if currentRail == rails-1 {
			directionDown = false
		}
		if currentRail == 0 {
			directionDown = true
		}
	}

	var res = ""
	for _, v := range arrail {
		res += v
	}
	return res
}

func Decode(message string, rails int) string {
	var arrail = make([]string, rails)

	/* The period is v-shape, repeating in a row, like:
	\
	 \    /
	  \  /
	   \/  */

	period := rails + rails - 2

	// Is a number of periods in the row
	chunks := len(message) / period

	//first row length depends if message lenght splits evenly to periods
	var firstRowLen int
	if chunks*period == len(message) {
		firstRowLen = chunks
	} else {
		firstRowLen = chunks + 1
	}
	//first row
	arrail[0] = message[:firstRowLen]
	//last row
	arrail[rails-1] = message[len(message)-(chunks):]
	//middle rows
	leftover := message[firstRowLen : len(message)-(chunks)]
	//number of middle rows
	divider := len(leftover) / (rails - 2)

	//split middle to rows
	for i := 1; i < rails-2+1; i++ {
		arrail[i] = leftover[(i-1)*divider : (i-1)*divider+divider]
	}

	currentRail := 0
	directionDown := true
	var res = ""
	//reconstruct result string
	for i := 0; i < len(message); i++ {

		res += arrail[currentRail][:1]
		arrail[currentRail] = arrail[currentRail][1:]
		if directionDown {
			currentRail++
		} else {
			currentRail--
		}
		if currentRail == rails-1 {
			directionDown = false
		}
		if currentRail == 0 {
			directionDown = true
		}
	}

	return res

}
