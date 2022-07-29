<?php
    $aaa = [
        [0, 0, 3],
        [0, 1, 2],
        [0, 4, 9],
        [0, 5, 8],
        [0, 6, 7],
        [1, 1, 1],
        [1, 3, 9],
        [1, 4, 8],
        [1, 5, 7],
        [1, 6, 6],
        [2, 2, 9],
        [2, 3, 8],
        [2, 4, 7],
        [2, 5, 6],
        [3, 6, 4],
        [3, 7, 3],
        [4, 4, 5],
        [5, 3, 5],
        [5, 9, 9],
        [6, 8, 9],
        [7, 7, 9],
        [7, 8, 8],
    ];

    $row = (int)trim(fgets(STDIN));
    $answer = [];

    for ($iii=0;$iii<$row;$iii++) {
        fgets(STDIN);
        $tmp = $aaa;
        $input = trim(fgets(STDIN));
        $input = explode(' ', $input);
        $flag = false;

        for ($i=0;$i<count($tmp);$i++) {
            foreach ($input as $v) {
                $v = (int)$v[strlen($v) - 1];
                $index = array_search($v, $tmp[$i]);
                if ($index !== false) {
                    unset($tmp[$i][$index]);
                }

                if (count($tmp[$i]) == 0) {
                    $flag = true;
                    $answer[] = 'YES';
                    break;
                }
            }

            if ($flag == true) {
                break;
            }
        }

        if ($flag == false) {
            $answer[] = 'NO';
        }
    }

    foreach ($answer as $v) {
        echo $v . PHP_EOL;
    }