use std::fs;

fn check_ranges_overlap(ranges: &mut [(i64, i64)]) {
    ranges.sort_by_key(|start| start.0);

    for r in 0..(ranges.len() - 1) {
        for v in (r + 1)..ranges.len() {
            if ranges[r].1 >= ranges[v].0 {
                if ranges[r].1 < ranges[v].1 {
                    ranges[r].1 = ranges[v].1;
                }
                ranges[v] = (-1, -1);
            }
        }
    }
}

pub fn ex2() {
    let full_fich = fs::read_to_string("src/day5.txt").expect("Could not open the file to read");
    let (ranges_fich, _) = full_fich
        .split_once("\n\n")
        .expect("Couldnt split full fich onto ranges and ids");
    let ranges_lines = ranges_fich.lines();

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
    for r in ranges {
        cont += r.1 - r.0 + 1;
    }

    println!("\nThere are {cont} fresh products");
}
