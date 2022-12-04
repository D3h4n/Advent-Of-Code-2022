import sys
from typing import List

def main(input_file: str):
    with open(input_file, 'r') as f:
        cals_per_elf = parse_cals_per_elf(f.read())
        print(f"Part 1: {max_cals(cals_per_elf)}")
        print(f"Part 2: {total_top_n_cals(cals_per_elf, 3)}")

def max_cals(cals_per_elf: List[int]) -> int:
    return max(*cals_per_elf)

def total_top_n_cals(cals_per_elf: List[int], n: int) -> int:
    cals_per_elf.sort(reverse=True)
    return sum(cals_per_elf[:n])

def parse_cals_per_elf(data: str) -> List[int]:
    cals_per_elf = [0]
    index = 0

    for line in data.strip().splitlines():
        line = line.strip()

        if line:
            cals_per_elf[index] += int(line)
        else:
            cals_per_elf.append(0)
            index += 1

    return cals_per_elf

if __name__ == '__main__':
    main(sys.argv[1])