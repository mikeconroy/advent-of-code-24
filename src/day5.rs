use crate::Day;

use std::collections::HashMap;
use std::collections::HashSet;

pub struct DayFive;
impl Day for DayFive {
    // Input is 2 Sections
    // Page Ordering Rules
    //  X|Y -> if X and Y are to be produced in an update, X must be printed before Y.
    // Updates
    //  a,b,c,X,Y -> In order - Take middle value
    //  a,b,Y,c,X -> Out of order drop
    fn part1(&self, input: &str) -> String {
        let (pages, orders) = parse_input(input);
        let mut safe_orders: Vec<Vec<i32>> = Vec::new();
        for order in orders {
            if is_valid_ordering(&pages, &order) {
                safe_orders.push(order);
            }
            // println!("{}", safe_orders.len());
            // println!("---------------");
        }

        return get_sum_of_middle_nums(safe_orders).to_string();
    }

    fn part2(&self, input: &str) -> String {
        let (pages, orders) = parse_input(input);
        let mut unsafe_orders: Vec<Vec<i32>> = Vec::new();
        for order in orders {
            // println!("{:?}", order);
            if !is_valid_ordering(&pages, &order) {
                let new_order = reorder(&pages, &order);
                // println!("{:?}", order);
                unsafe_orders.push(new_order);
            }
            // println!("{}", unsafe_orders.len());
            // println!("---------------");
        }
        return get_sum_of_middle_nums(unsafe_orders).to_string();
    }
}

fn get_sum_of_middle_nums(orders: Vec<Vec<i32>>) -> i32 {
    let mut result = 0;
    for order in orders {
        let len = order.len();
        result += order[len / 2];
    }
    return result;
}

fn is_valid_ordering(pages: &Pages, order: &Vec<i32>) -> bool {
    let mut prev_nums: HashSet<i32> = HashSet::new();
    for num in order {
        prev_nums.insert(*num);
        let page: &Page = pages.get(num).unwrap();
        if !prev_nums.is_disjoint(&page.comes_before) {
            return false;
        }
    }
    return true;
}

fn reorder(pages: &Pages, order: &Vec<i32>) -> Vec<i32> {
    let mut vals: HashSet<i32> = HashSet::new();
    for val in order {
        vals.insert(*val);
    }
    let mut new_order: Vec<i32> = Vec::new();
    while new_order.len() != order.len() {
        // println!(
        // "New Order Len: {} Order Len: {}",
        // new_order.len(),
        // order.len()
        // );
        // println!("New Order: {:?}", new_order);
        for (_index, val) in vals.clone().into_iter().enumerate() {
            let page: &Page = pages.get(&val).unwrap();
            // println!(
            // "Index: {} Val: {} Comes After: {:?}",
            // index, val, page.comes_after
            // );
            if page.comes_after.is_disjoint(&vals) {
                new_order.push(val);
                vals.remove(&val);
                break;
            }
        }
    }
    return new_order;
}
struct Page {
    comes_before: HashSet<i32>,
    comes_after: HashSet<i32>,
}

type Pages = HashMap<i32, Page>;

fn parse_input(input: &str) -> (Pages, Vec<Vec<i32>>) {
    let mut pages: Pages = HashMap::new();
    let mut orders: Vec<Vec<i32>> = Vec::new();

    for line in input.lines() {
        if line.contains("|") {
            let split: Vec<&str> = line.split("|").collect();
            let lhs_num: i32 = split[0].parse().unwrap();
            let rhs_num: i32 = split[1].parse().unwrap();
            let lhs_page = pages.entry(lhs_num).or_insert(Page {
                comes_before: HashSet::new(),
                comes_after: HashSet::new(),
            });
            lhs_page.comes_before.insert(rhs_num);
            let rhs_page = pages.entry(rhs_num).or_insert(Page {
                comes_before: HashSet::new(),
                comes_after: HashSet::new(),
            });
            rhs_page.comes_after.insert(lhs_num);
        } else if line.contains(",") {
            let ordering: Vec<i32> = line
                .split(",")
                .map(|s| s.trim().parse::<i32>().expect("Invalid Number"))
                .collect();
            orders.push(ordering);
        }
    }

    return (pages, orders);
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::utils;

    #[test]
    fn test_part1() {
        let day = DayFive {};
        let input = utils::read_file("data/day5_test");
        let result = day.part1(&input);
        assert_eq!(result, "143");
    }

    #[test]
    fn test_part2() {
        let day = DayFive {};
        let input = utils::read_file("data/day5_test");
        let result = day.part2(&input);
        assert_eq!(result, "123");
    }
}
