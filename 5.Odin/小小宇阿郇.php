<?php

    $dataSetLenght = (int) fgets(STDIN);

    for ($i=0; $i < $dataSetLenght; $i++) {
        fgets(STDIN);
        $numArray = explode(' ', trim(fgets(STDIN)));
        $filterNumArray = [];
        $tempMap = [0,0,0,0,0,0,0,0,0,0];
        for ($j=0; $j < count($numArray); $j++) {
            $single = $numArray[$j] % 10;
            $tempMap[$single]++;
            if ($tempMap[$single] <= 3) {
                $filterNumArray[] = $single;
            }
            if (array_sum($tempMap) > 30) {
                break;
            }
        }
        if (check3sum($filterNumArray)) {
            echo 'YES' . PHP_EOL;
        } else {
            echo 'NO' . PHP_EOL;
        }
    }

    function check3sum($numArray)
    {
        for ($i=0; $i < count($numArray); $i++) {
            $tempMap = [];
            for ($j=$i+1; $j < count($numArray); $j++) {
                $target = $numArray[$i] + $numArray[$j];
                if ($target <= 3) {
                    if (isset($tempMap[3 - $target])) {
                        return true;
                    }
                } else {
                    if (isset($tempMap[(13 - ($target%10))%10])) {
                        return true;
                    }
                }
                $tempMap[$numArray[$j]] = true;
            }
        }
        return false;
    }
