use crate::Day;

pub struct DaySeven;

impl Day for DaySeven {
    fn part1(&self, input: &str) -> String {
        let eqs = parse_input(input);
        let mut result = 0;
        for eq in eqs {
            // println!("EQ: {:?}", eq);
            if can_be_solved(&eq, false) {
                // println!("Solved");
                result += eq.ans;
            }
            // println!("{}", result);
            // println!("-----");
        }
        return result.to_string();
    }

    fn part2(&self, input: &str) -> String {
        let eqs = parse_input(input);
        let mut result = 0;
        for eq in eqs {
            // println!("EQ: {:?}", eq);
            if can_be_solved(&eq, true) {
                // println!("Solved");
                result += eq.ans;
            }
            // println!("{}", result);
            // println!("-----");
        }
        return result.to_string();
    }
}

// Recursive function to check all permutations of operations.
fn can_be_solved(eq: &Equation, enable_concatanation: bool) -> bool {
    // println!("{:?}", eq);
    if eq.vals.len() == 1 {
        if eq.vals[0] == eq.ans {
            return true;
        } else {
            return false;
        }
    }
    if eq.vals[0] > eq.ans {
        return false;
    }

    let mut add_vals: Vec<i64> = eq.vals.clone();
    let new_val = add_vals[0] + add_vals[1];
    add_vals[0] = new_val;
    add_vals.remove(1);
    let add_eq: Equation = Equation {
        ans: eq.ans,
        vals: add_vals,
    };
    if can_be_solved(&add_eq, enable_concatanation) {
        return true;
    } else {
        let mut mul_vals: Vec<i64> = eq.vals.clone();
        let mul_val = mul_vals[0] * mul_vals[1];
        mul_vals[0] = mul_val;
        mul_vals.remove(1);
        let mul_eq: Equation = Equation {
            ans: eq.ans,
            vals: mul_vals,
        };
        if can_be_solved(&mul_eq, enable_concatanation) {
            return true;
        } else if enable_concatanation {
            // println!("CONCATANATING");
            let mut concat_vals: Vec<i64> = eq.vals.clone();
            let concat_val: i64 = (concat_vals[0]
                * i64::pow(10, concat_vals[1].to_string().len().try_into().unwrap()))
                + concat_vals[1];
            // println!("POW: {}", concat_vals[1].to_string().len());
            concat_vals[0] = concat_val;
            concat_vals.remove(1);
            let concat_eq: Equation = Equation {
                ans: eq.ans,
                vals: concat_vals,
            };
            return can_be_solved(&concat_eq, enable_concatanation);
        } else {
            return false;
        }
    }
}

#[derive(Debug)]
struct Equation {
    ans: i64,
    vals: Vec<i64>,
}

// enum Ops {
// ADD,
// MUL,
// }

// impl Ops {
// fn exec(&self, a: i64, b: i64) -> i64 {
// match &self {
// Ops::ADD => a + b,
// Ops::MUL => a * b,
// }
// }
// }

fn parse_input(input: &str) -> Vec<Equation> {
    let mut eqs: Vec<Equation> = Vec::new();

    for line in input.lines() {
        let split: Vec<&str> = line.split(":").collect();
        let lhs: i64 = split[0].parse().unwrap();
        let rhs: Vec<&str> = split[1].split(" ").collect::<Vec<&str>>();
        let mut rhs_nums: Vec<i64> = Vec::new();
        for val in rhs.clone().into_iter() {
            if val != "" {
                rhs_nums.push(val.parse().unwrap());
            }
        }
        eqs.push(Equation {
            ans: lhs,
            vals: rhs_nums,
        })
    }
    return eqs;
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils;

    #[test]
    fn test_part1() {
        let test_input = utils::read_file("data/day7_test");
        let day = DaySeven {};
        let result = day.part1(&test_input);
        assert_eq!(result, "3749");
    }

    #[test]
    fn test_part2() {
        let test_input = utils::read_file("data/day7_test");
        let day = DaySeven {};
        let result = day.part2(&test_input);
        assert_eq!(result, "11387");
    }
}
