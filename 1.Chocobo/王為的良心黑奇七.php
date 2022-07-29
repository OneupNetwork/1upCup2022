<?php

$datasetCount = (int) trim(fgets(STDIN));

for ($j = 0; $j < $datasetCount; $j++) {

    fgets(STDIN);
    $countAry = [];
    $target = 0; $prev = 0;
    for ($i = 0; $i < 8; $i++) {
                $prev = $target;
        $line = fgets(STDIN);
        $target = count(explode("#", $line)) - 1;

        if ($target === 1 && $prev === 2) {
            printf("%d %d\n", $i + 1, strpos($line, "#") + 1);
            break;
        }

    }

    for ($i = $i+1; $i <8 ;$i++) {
        $line = fgets(STDIN);
    }

}