<?php

$datasetCount = (int) trim(fgets(STDIN));

for ($j = 0; $j < $datasetCount; $j++) {

    $arrayCount = (int) fgets(STDIN);

    $even = -1; $odd = -1;
    foreach (explode(" ", fgets(STDIN)) as $i => $v) {

        if ($i % 2) {
            if ($odd == -1) {
                $odd = $v % 2;
            } else {
                if ($odd !== $v % 2) {
                    echo "NO\n"; continue 2;
                }
            }
        } else {
            if ($even == -1) {
                $even = $v % 2;
            } else {
                if ($even !== $v % 2) {
                    echo "NO\n"; continue 2;
                }
            }
        }
    }
    echo "YES\n";
}