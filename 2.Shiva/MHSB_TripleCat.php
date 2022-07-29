<?php
// Enter your code here, enjoy!
$in = trim(fgets(STDIN));
$datas = [];
$res = [];

for ($i = 0; $i < $in; $i++) {
	$datas[$i][0] = trim(fgets(STDIN));
	$datas[$i][1] = trim(fgets(STDIN));
}


foreach($datas as $data) {
	$length = $data[0];
	$arr = explode(' ', $data[1]);
	$flagO = ($arr[0] % 2 == 0);
	$flagE = ($arr[1] % 2 == 0);
	foreach ($arr as $k => $v) {
	  $check = ($v % 2 == 0);
		if ($k % 2 == 0) {
			if ($flagO == $check) {
			  if ($k == $length - 1) {
			    echo 'YES'.PHP_EOL;
		    }
				continue;
			} else {
				echo 'NO'.PHP_EOL;
				break;
			}
		} else {
			if ($flagE == $check) {
			  if ($k == $length - 1) {
			    echo 'YES'.PHP_EOL;
		    }
				continue;
			} else {
				echo 'NO'.PHP_EOL;
				break;
			}
		}
	}
}