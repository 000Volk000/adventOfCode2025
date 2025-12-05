use std::fs;

fn check_ranges_overlap(ranges: &mut [(i32, i32)]) {
    for r in 0..ranges.len() {
        println!("{r}")
    }
}

pub fn ex1() {
    let full_fich =
        fs::read_to_string("day5-example.txt").expect("Could not open the file to read");
    let fich: Vec<&str> = full_fich.split("\n\n").collect();
    let ranges_lines = fich[0].lines();
    let ids = fich[1].lines();

    let mut ranges = Vec::<(i32, i32)>::new();

    for line in ranges_lines {
        let min_max: Vec<&str> = line.split("-").collect();
        ranges.push((
            min_max[0].parse().expect("A range is not a Number"),
            min_max[1].parse().expect("A range is not a Number"),
        ));
    }

    check_ranges_overlap(&mut ranges);

    for id in ids {
        println!("id:{id}");
    }
}
