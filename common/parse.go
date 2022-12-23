package common

func ParseColonLine(line []byte, parseCols int) (entries [][]byte, remain []byte) {
	if parseCols <= 0 || len(line) == 0 {
		return entries, line
	}
	var parsed int
	var entryStart int
	for i := 0; i < len(line); i++ {
        if line[i] == ':' {
        	parsed += 1
			entries = append(entries, line[entryStart:i])
			entryStart = i + 1
        	if parsed > parseCols {
        		return entries, line[i+1:]
			}
		} else if line[i] == '\n' {
			entries = append(entries, line[entryStart:i])
			return entries, line[i+1:]
		}
	}
	entries = append(entries, line[entryStart:])
	return entries, line[len(line):]
}
