# Take input, split into array and covert to ints
line = input()
raw_codes = line.split(sep=",")
const_codes = []
for code in raw_codes:
    const_codes.append(int(code))

codes = const_codes.copy()

pos = 0
for noun in range(100):
    for verb in range(100):
        codes = const_codes.copy()
        codes[1] = noun
        codes[2] = verb
        pos = 0

        while (pos < len(codes)):
            op_code = codes[pos]
            if op_code == 99:
                break

            num_one_pos = codes[pos+1]
            num_one = codes[num_one_pos]

            num_two_pos = codes[pos+2]
            num_two = codes[num_two_pos]

            dest = codes[pos+3]

            if op_code == 1:
                codes[dest] = num_one + num_two

            if op_code == 2:
                codes[dest] = num_one * num_two

            pos += 4

        if codes[0] == 19690720:
            print(100 * noun + verb)
