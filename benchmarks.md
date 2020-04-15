#Benchmarks 

###Regular
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `uniq ./input.txt` | 0.6 ± 0.0 | 0.5 | 0.9 | 1.00 |
| `./goUniq ./input.txt` | 0.8 ± 0.0 | 0.7 | 1.5 | 1.39 ± 0.10 |

###Case sensitive -i
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `uniq -i ./input.txt` | 0.6 ± 0.0 | 0.5 | 1.3 | 1.00 |
| `./goUniq -i ./input.txt` | 0.9 ± 0.0 | 0.7 | 1.6 | 1.37 ± 0.10 |

###Count repeats -c
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `uniq -c ./input.txt` | 0.6 ± 0.0 | 0.5 | 1.2 | 1.00 |
| `./goUniq -c ./input.txt` | 0.9 ± 0.0 | 0.7 | 1.5 | 1.38 ± 0.09 |

###Output unrepeated lines -u
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `uniq -u ./input.txt` | 0.6 ± 0.0 | 0.5 | 1.2 | 1.00 |
| `./goUniq -u ./input.txt` | 0.9 ± 0.0 | 0.7 | 1.7 | 1.42 ± 0.10 |

###Output repeated lines -d
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `uniq -d ./input.txt` | 0.6 ± 0.0 | 0.5 | 1.2 | 1.00 |
| `./goUniq -d ./input.txt` | 0.9 ± 0.0 | 0.7 | 1.5 | 1.41 ± 0.10 |

###Skip fields -f
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `uniq -f 2 ./input.txt` | 0.6 ± 0.0 | 0.5 | 1.3 | 1.00 |
| `./goUniq -f 2 ./input.txt` | 0.9 ± 0.0 | 0.7 | 1.6 | 1.39 ± 0.09 |

###Skip chars -s
| Command | Mean [ms] | Min [ms] | Max [ms] | Relative |
|:---|---:|---:|---:|---:|
| `uniq -s 5 ./input.txt` | 0.6 ± 0.0 | 0.5 | 1.3 | 1.00 |
| `./goUniq -s 5 ./input.txt` | 0.8 ± 0.0 | 0.7 | 1.6 | 1.39 ± 0.09 |


###input.txt
```I love eating.
 I love eating.
 I love eating.
 
 I love eating Pizza.
 I love eating Pizza.
 
 Thanks.
 Than
 
 I love eating.
 
 I love eating.
 I love eating.
 
 I love eating.
 I love eating.
```
