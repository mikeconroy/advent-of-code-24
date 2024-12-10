use crate::Day;

pub struct DayThree;
impl Day for DayThree {
    fn part1(&self, input: &str) -> String {
        return get_result(input).to_string();
    }

    fn part2(&self, input: &str) -> String {
        let mut result = 0;
        //mul(1,2)don't()mul(1,2)mul(3,4)do()mul(2,2)
        // split by don't()
        // ["mul(1,2)", "mul(1,2)mul(3,4)do()mul(2,2)"
        // loop through each & split by do()
        // mul(1,2), ["mul(1,2)mul(3,4)", "mul(2,2)"
        let donts: Vec<&str> = input.split("don't()").collect();
        result += get_result(donts[0]);
        for index in 1..donts.len() {
            let mut dos: Vec<&str> = donts[index].split("do()").collect();
            if dos.len() == 1 {
                continue;
            }
            dos.remove(0);
            for do_mem in dos {
                result += get_result(do_mem);
            }
        }
        return result.to_string();
    }
}

fn get_result(input: &str) -> i32 {
    let mut result = 0;
    // abcmul(1,2)321mul(321mul( -> ["1,2)321", "321"]
    let muls: Vec<&str> = input.split("mul(").collect();
    for mul_index in 1..muls.len() {
        // println!("{}", mul_index);
        result += get_mul(muls[mul_index]);
        // println!("----------------");
    }
    return result;
}

fn get_mul(mul_str: &str) -> i32 {
    // println!("{}", mul_str);
    if !mul_str.contains(")") {
        return 0;
    }
    let brackets = mul_str.split(")").nth(0).unwrap();
    // println!("{}", brackets);
    let mul_vec: Vec<&str> = brackets.split(",").collect();
    if mul_vec.len() != 2 {
        // More than one or 0 commas exists inside the brackets.
        return 0;
    }

    // println!("{} {}", mul_vec[0], mul_vec[1]);
    if mul_vec.len() != 2 {
        // More than one or 0 commas exists inside the brackets.
        return 0;
    }

    if mul_vec[0].len() > 3 || mul_vec[1].len() > 3 {
        return 0;
    }

    let lhs: i32 = mul_vec[0].parse().unwrap_or(0);
    let rhs: i32 = mul_vec[1].parse().unwrap_or(0);
    // println!("Result increased {}", lhs * rhs);
    return lhs * rhs;
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils::file::read_file;

    #[test]
    fn test_part1() {
        let day = DayThree {};
        let input = read_file("data/day3_test");
        let result = day.part1(&input);
        assert_eq!(result, "161");
    }

    #[test]
    fn test_part2() {
        let day = DayThree {};
        let input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";
        let result = day.part2(&input);
        assert_eq!(result, "48");
    }
}
