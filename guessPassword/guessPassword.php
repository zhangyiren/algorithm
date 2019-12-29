<?php


// 筛选符合条件的数组
function seekOut($ca){
	
	LOOP:
	foreach($ca as $ka=>$va){

		// 初始化检测值
		$ck_1=1;
		$ck_2=0;
		$ck_3=0;
		$ck_4=0;
		$ck_5=0;
		$len_va=count($va);

		// 检测一 条件四
		for($i=0;$i<$len_va;$i++){
			if($va[$i]==1 || $va[$i]==4 || $va[$i]==7){
				$ck_1=0;
				goto LOOP;
			}
		}

		// 检测二 条件一
		if($va[0]==2 || $va[2]==6){
			$ck_2=1;
		}

		// 检测三 条件二
		for($i=0;$i<$len_va;$i++){
			if(($va[$i]==2 && $i!=0) || ($va[$i]==5 && $i!=1) || ($va[$i]==8 && $i!=2)){
				$ck_3=1;
			}
		}
		
		// 检测四 条件三
		$correctNumber=0;
		foreach($va as $k=>$v){
			if($v==6 || $v==9 || $v==2){
				$correctNumber+=1;
			}
		}
		// 2个号码正确
		if($correctNumber==2){
			if(($va[0]!=6 && $va[1]!=9 && $va[2]!=2) || ($va[1]!=9 && $va[2]!=2 && $va[0]!=6) || ($va[0]!=6 && $va[2]!=2 && $va[1]!=9)){
				$ck_4=1;
			}
		}

		// 检测五 条件五
		if($va[0]==9 || $va[1]==9){
			$ck_5=1;
		}

		// 筛选输出符合5个条件的数组
		$result=$ck_1*$ck_2*$ck_3*$ck_4*$ck_5;
		if($result==1){
			return $va;
		}

	}

}


// 排列组合
function sequence(&$va,$ma){

	// 前两位号码
	$clt[]=$ma[0].$ma[1];
	$clt[]=$ma[1].$ma[0];
	for($i=2;$i<strlen($ma);$i++){
		for($j=0;$j<count($clt);$j++){
			// 组成新序列
			$cb=$ma[$i].$clt[$j];
			for($k=0;$k<strlen($cb);$k++){
				$cm=substr($cb,1,strlen($cb)-1).$cb[0];
				$cb=$cm;
				$tmp[]=$cb;
			}	
		}
		unset($clt);
		$clt=$tmp;
		unset($tmp);
	}
	foreach($clt as $v){
		$va[]=str_split($v,1);
	}
}


//5个已知条件
$cd[]=[2,4,6]; //条件一: 1个号码正确位置正确
$cd[]=[2,5,8]; //条件二: 1个号码正确位置不正确
$cd[]=[6,9,2]; //条件三: 2个号码正确位置都不正确
$cd[]=[1,7,4]; //条件四: 没有一个号码正确
$cd[]=[4,1,9]; //条件五: 1个号码正确位置不正确

//$ca存放可能号码
$ca=[];

// 合并
foreach($cd as $v){
	$ca=array_merge($v,$ca);
}
// 去重
$ca = array_unique($ca);

// 由条件四可知，排除1,4,7
foreach($ca as $k=>$v){
	if($v==1 || $v==4 || $v==7){
		unset($ca[$k]);
	}
}

// 排序
sort($ca);
// 转成字符串
$caStr=implode('', $ca);

// 组合-n取3
$lenCa=strlen($caStr);
for($i=0;$i<$lenCa;$i++){
	for($j=$i+1;$j<$lenCa;$j++){
		for($k=$j+1;$k<$lenCa;$k++){
			$tmp=$ca[$i].$ca[$j].$ca[$k];
			$mtva[]=$tmp;
		}
	}
}

foreach($mtva as &$ma){

	$va=array();
	// 排列组合
	sequence($va,$ma);
	$answer=seekOut($va);
	if($answer!=false){
		print_r($answer);
		break;
	}

}
