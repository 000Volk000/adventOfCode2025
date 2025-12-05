use std::fs;

fn is_between(n: i64, range: (i64, i64)) -> bool {
    if n > range.0 && n < range.1 {
        return true;
    }
    false
}

fn is_between_equals(n: i64, range: (i64, i64)) -> bool {
    if n >= range.0 && n <= range.1 {
        return true;
    }
    false
}

fn check_ranges_overlap(ranges: &mut [(i64, i64)]) {
    for r in 0..(ranges.len() - 1) {
        for v in (r + 1)..ranges.len() {
            if ranges[v] != (-1, -1) {
                if is_between(ranges[r].0, ranges[v]) && is_between(ranges[r].1, ranges[v]) {
                    ranges[r] = ranges[v];
                    ranges[v] = (-1, -1)
                } else if is_between(ranges[r].0, ranges[v]) {
                    ranges[r].0 = ranges[v].0;
                    ranges[v] = (-1, -1)
                } else if is_between(ranges[r].1, ranges[v]) {
                    ranges[r].1 = ranges[v].1;
                    ranges[v] = (-1, -1)
                }
            }
        }
    }
}

pub fn ex1() {
    let full_fich = fs::read_to_string("src/day5.txt").expect("Could not open the file to read");
    let (ranges_fich, ids_fich) = full_fich
        .split_once("\n\n")
        .expect("Couldnt split full fich onto ranges and ids");
    let ranges_lines = ranges_fich.lines();
    let ids = ids_fich.lines();

    let mut ranges = Vec::<(i64, i64)>::with_capacity(ranges_lines.clone().count());

    for line in ranges_lines {
        let (start, end) = line
            .split_once("-")
            .expect("Couldnt split the range onto start and end");
        ranges.push((
            start.parse().expect("A range is not a Number"),
            end.parse().expect("A range is not a Number"),
        ));
    }

    check_ranges_overlap(&mut ranges);
    ranges.retain(|&x| x != (-1, -1));

    let mut cont = 0;
    for id in ids {
        for r in &ranges {
            if is_between_equals(id.parse().expect("And id is not a Number"), *r) {
                cont += 1;
                break;
            }
        }
    }

    println!("\nThere are {cont} fresh products");
}
