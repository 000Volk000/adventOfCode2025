<?php

function splitString($str)
{
    return(explode(",", trim($str)));
}

function area($p1, $p2)
{
    $base = abs($p1[0] - $p2[0]) + 1;
    $height = abs($p1[1] - $p2[1]) + 1;
    return ($base * $height);
}

function hasOnlyGreenTiles($points, $p1, $p2)
{
    if ($p1[0] < $p2[0]) {
        $x1 = $p1[0];
        $x2 = $p2[0];
    } else {
        $x1 = $p2[0];
        $x2 = $p1[0];
    }
    if ($p1[1] < $p2[1]) {
        $y1 = $p1[1];
        $y2 = $p2[1];
    } else {
        $y1 = $p2[1];
        $y2 = $p1[1];
    }
    $found1 = false;
    $found2 = false;
    $found3 = false;
    $found4 = false;

    foreach ($points as $p) {
        if ($p[0] <= $x1 && $p[1] <= $y1) {
            $found1 = true;
        }
        if ($p[0] <= $x1 && $p[1] >= $y2) {
            $found2 = true;
        }
        if ($p[0] >= $x2 && $p[1] >= $y2) {
            $found3 = true;
        }
        if ($p[0] >= $x2 && $p[1] <= $y1) {
            $found4 = true;
        }
    }

    if ($found1 && $found2 && $found3 && $found4) {
        return true;
    }
    return false;
}

function maxX($points)
{
    $max = $points[0][0];
    foreach ($points as $p) {
        if ($p[0] > $max) {
            $max = $p[0];
        }
    }
    return $max;
}

function maxY($points)
{
    $max = $points[0][1];
    foreach ($points as $p) {
        if ($p[1] > $max) {
            $max = $p[1];
        }
    }
    return $max;
}

function createMatrix($points)
{
    $cols = maxX($points) + 3;
    $rows = maxY($points) + 2;
    $matrix = [];
    for ($i = 0; $i < $rows; $i++) {
        $matrix[$i] = array_fill(0, $cols, ".");
    }

    foreach ($points as $p) {
        $matrix[$p[1]][$p[0]] = "#";
    }
    return $matrix;
}

function printMatrix($matrix)
{
    for ($row = 0; $row < sizeof($matrix); $row++) {
        for ($col = 0; $col < sizeof($matrix[$row]); $col++) {
            print_r($matrix[$row][$col]);
        }
        print_r("\n");
    }
}
