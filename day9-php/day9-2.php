<?php

require 'day9-functions.php';

$fich = fopen("day9-example.txt", "r") or die("Unable to open file");
$points = array();
while (!feof($fich)) {
    array_push($points, fgets($fich));
}
array_pop($points);
fclose($fich);

$points = array_map("splitString", $points);
$connected_matrix = createMatrix($points);
printMatrix($connected_matrix);

$max_area = 0;
foreach ($points as $current_point) {
    foreach ($points as $compare_point) {
        if (hasOnlyGreenTiles($points, $current_point, $compare_point)) {
            if (($area = area($current_point, $compare_point)) > $max_area) {
                $max_area = $area;
            }
        }
    }
}

print_r("The maximum area is: " . $max_area . "\n");
