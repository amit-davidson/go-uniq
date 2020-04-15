  # go-uniq
go-uniq is an implementation of Darwin's uniq command. 

  ## Installation

      go get github.com/amitdavidson234/go-uniq/cmd 


  ## Usage
	Usage of ./cmdCommandLinux:
	uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]
      -c	
		    Precede each output line with the count of the number of times the line occurred in the input, followed by a single space.
      -d	
		    Only output lines that are repeated in the input.
      -f int
        	Ignore the first num fields in each input line when doing comparisons.  A field is a string of non-blank characters separated from adjacent fields by blanks. Field numbers are one based, i.e., the first field is field one.
      -i	
		    Case insensitive comparison of lines.
      -s int
        	with the -f option, the first chars characters after the first num fields will be ignored.  Character numbers are one based, i.e., the first character is character one.
      -u	
		    Only output lines that are not repeated in the input.

  ## Performance
Admittedly, Darwin's implementation provides better results compared to go-uniq. I squeezed the most I could from the go compiler to get as close as I could to the real implementation. Comparison of all the flags is provided in a different markdown file.

| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |  
|:---|---:|---:|---:|---:|  
| `uniq ./input.txt` | 0.6 ± 0.0 | 0.5 | 0.9 | 1.00 |  
| `./goUniq ./input.txt` | 0.8 ± 0.0 | 0.7 | 1.5 | 1.39 ± 0.10 |
