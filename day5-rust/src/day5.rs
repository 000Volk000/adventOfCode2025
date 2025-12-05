use std::fs;

fn check_ranges_overlap(ranges: &mut [(i32, i32)]) {
    for r in 0..ranges.len() {
        println!("{r}")
    }
}

pub fn ex1() {
    let full_fich =
        fs::read_to_string("day5-example.txt").expect("Could not open the file to read");
    let (ranges_fich, ids_fich) = full_fich
        .split_once("\n\n")
        .expect("Couldnt split full fich onto ranges and ids");
    let ranges_lines = ranges_fich.lines();
    let ids = ids_fich.lines();

    let mut ranges = Vec::<(i32, i32)>::with_capacity(ranges_lines.clone().count());

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

    for id in ids {
        println!("id:{id}");
    }
}
