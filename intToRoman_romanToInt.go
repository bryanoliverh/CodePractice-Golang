func intToRoman(num int) string {
    var roman string = ""
    var numbers = [] int {
        1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000
    }
    var romans = [] string {
        "I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"
    }
    var index = len(romans) - 1
    for num > 0 {
        for numbers[index] <= num {
            roman += romans[index]
            num -= numbers[index]
        }
        index -= 1
    }
    return roman
}

func romanToInt(s string) int {
    romanList: = map[byte] int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    input_len: = len(s)
    if input_len == 0 {
        return 0
    }
    if input_len == 1 {
        return romanList[s[0]]
    }
    numeric_digit: = romanList[s[input_len - 1]]
    for i: = input_len - 2;i >= 0;i--{
        if romanList[s[i]] < romanList[s[i + 1]] {
            numeric_digit -= romanList[s[i]]
        } else {
            numeric_digit += romanList[s[i]]
        }
    }
    return numeric_digit
}
