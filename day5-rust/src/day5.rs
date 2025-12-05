use std::fs;

fn check_ranges_overlap(ranges: &mut [(i32, i32)]) {
    for r in 0..ranges.len() {
        println!("{r}")
    }
}

pub fn ex1() {
    let full_fich =
        fs::read_to_string("day5-example.txt").expect("Could not open the file to read");
    let fich = full_fich.lines();

    let mut flag = true;
    let mut ranges = Vec::<(i32, i32)>::new();
    for line in fich {
        if flag && line.is_empty() {
            flag = false;
            check_ranges_overlap(&mut ranges);
            continue;
        }

        if flag {
            let min_max: Vec<&str> = line.split("-").collect();
            ranges.push((
                min_max[0].parse().expect("A range is not a Number"),
                min_max[1].parse().expect("A range is not a Number"),
            ));
        } else {
            continue;
        }
    }
}
