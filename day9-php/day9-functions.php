<?php

function splitString($str)
{
    return(explode(",", $str));
}

function area($p1, $p2)
{
    $base = abs($p1[0] - $p2[0]) + 1;
    $height = abs($p1[1] - $p2[1]) + 1;
    return ($base * $height);
}
