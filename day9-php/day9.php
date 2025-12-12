<?php

function splitString($str)
{
    return(explode(",", $str));
}

$fich = fopen("day9-example.txt", "r") or die("Unable to open file");
$points = array();
while (!feof($fich)) {
    array_push($points, fgets($fich));
}
array_pop($points);
fclose($fich);

$points = array_map("splitString", $points);
print_r($points);
