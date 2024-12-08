use std::collections::HashMap;
#[derive(Hash, Eq, PartialEq, Debug)]
struct Point {
    x: i64,
    y: i64,
}

pub struct Grid {
    cells: HashMap<Point, char>,
    row_count: i64,
    col_count: i64,
}
impl Grid {
    pub fn new() -> Self {
        Grid {
            cells: HashMap::new(),
            row_count: 0,
            col_count: 0,
        }
    }

    pub fn len_to_string(&self) -> String {
        self.cells.len().to_string()
    }

    fn insert(&mut self, point: Point, value: char) {
        self.cells.insert(point, value);
    }

    fn get_char(&self, point: &Point) -> &char {
        return self.cells.get(point).unwrap();
    }

    pub fn load(&mut self, input: &str) {
        let mut x = 0;
        let mut y = 0;
        for line in input.lines() {
            for val in line.chars() {
                self.insert(Point { x, y }, val);
                x += 1;
            }
            x = 0;
            y += 1;
        }
        self.col_count = x;
        self.row_count = y;
    }
}
