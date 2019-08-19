<?php

function merge($array1,$array2)
{

    $len1 = count($array1);
    $len2 = count($array2);

    if ($len1 == 0 && $len2 == 0) {
        return [];
    }

    if ($len1 == 0) {
        return $array2;
    }

    if ($len2 == 0) {
        return $array1;
    }

    $k = 0;
    for ($i = 0; $i < $len1; $i++) {
        for ($j = $k; $j < $len2; $j++) {
            if ($array1[$i] > $array2[$j]) {
                $array3[] = $array2[$j];
                $k++;
            } else {
                $array3[] = $array1[$i];
                // reach the end
                if ($i == $len1 - 1) {
                    $fragment = array_slice($array2, $j);
                    // merge fragment
                    foreach($fragment as $value){
                        $array3[] = $value;
                    }
                }
                break;
            }
        }
        if ($k == $len2) {
            $array3[] = $array1[$i];
        }
    }

    return $array3;

}


    	$array = [706,100,188,9,11,25,133,452,364,5889,293,607,365,8633,555,18,66,78,15,43,67,96,155,387,972,1002,74,38,633,801];
	$arrayLen = count($array);
	for($i = 0 ; $i < $arrayLen ; $i++) {
        $array1[] = $array[$i];
		if(count($array1) == 2) {
            // swap element make it sequential
            if($array1[0] > $array1[1]) {
                $tmp = $array1[0];
				$array1[0] = $array1[1];
				$array1[1] = $tmp;
			}
            $array2 = merge($array1,$array2);
			$array1 = [];
		}

	}
	// maybe some elements left, merge it too
	$array2 = merge($array1,$array2);
	print_r($array2);
