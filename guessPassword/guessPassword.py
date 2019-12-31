#!/usr/bin/python
# -*- coding: utf-8 -*-   

import os

# 筛选符合条件的数组
def seekOut(ca):

	for i in range(len(ca)):
		ck_1=1
		ck_2=0
		ck_3=0
		ck_4=0
		ck_5=0
		va=ca[i]
		lenVa=len(va)
		# 检测一 条件四
		for i in range(lenVa):
			if (va[i]==1) or (va[i]==4) or (va[i]==7):
				ck_1=0
				continue

		# 检测二 条件一
		if (va[0]==2) or (va[2]==6):
			ck_2=1
	
		# 检测三 条件二
		for i in range(lenVa):
			if (va[i]==2 and i!=0) or (va[i]==5 and i!=1) or (va[i]==8 and i!=2):
				ck_3=1
		
		# 检测四 条件三
		correctNumber=0
		for i in range(lenVa):
			if (va[i]==6) or (va[i]==9) or (va[i]==2):
				correctNumber+=1

		# 2个号码正确
		if correctNumber==2:
			if (va[0]!=6 and va[1]!=9 and va[2]!=2) or (va[1]!=9 and va[2]!=2 and va[0]!=6) or (va[0]!=6 and va[2]!=2 and va[1]!=9):
				ck_4=1
		
		# 检测五 条件五
		if va[0]==9 or va[1]==9:
			ck_5=1
		
		# 筛选输出符合5个条件的数组
		result=ck_1*ck_2*ck_3*ck_4*ck_5
		if result==1:
			return va
		
	return []


# 排列组合
def sequence(ma):
	clt=[]
	clt.append([ma[0],ma[1]])
	clt.append([ma[1],ma[0]])
	for i in range(2,len(ma)):
		tmp=[]
		for j in range(0,len(clt)):
			cb=clt[j]
			cb.append(ma[i]);
			for k in range(len(cb)):
				cm=cb[1:]
				cm.append(cb[0])
				cb=cm
				tmp.append(cb)
		clt=tmp

	return clt



#ca存放可能号码
ca=[]
#5个已知条件
cd=[
	[2,4,6],  #条件一: 1个号码正确位置正确
	[2,5,8],  #条件二: 1个号码正确位置不正确
	[6,9,2],  #条件三: 2个号码正确位置都不正确
	[1,7,4],  #条件四: 没有一个号码正确
	[4,1,9],  #条件五: 1个号码正确位置不正确
]

for i in range(len(cd)):
	ca.extend(cd[i])

ca=list(set(ca))

# 由条件四可知，排除1,4,7
ca.remove(1)
ca.remove(4)
ca.remove(7)
# 排序
ca.sort()
mtva=[]
la=len(ca)
# 组合-n取3
for i in range(la):
	for j in range(i+1,la):
			for k in range(j+1,la):
				tmp=[]
				tmp.append(ca[i])
				tmp.append(ca[j])
				tmp.append(ca[k])
				mtva.append(tmp)


for i in range(len(mtva)):
	va=sequence(mtva[i])
	answer=seekOut(va);
	if len(answer)>0:
		print(answer)
		exit(1)
