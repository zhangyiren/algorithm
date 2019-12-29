import java.util.ArrayList;
import java.util.Arrays;
import java.util.Set;
import java.util.TreeSet;

public class GuessPassword {
	
		// 筛选符合条件的数组
		public static String seekOut(int al,ArrayList<String> ca) {
			
			// 当前字符串
			String bkp="";
			
			LOOP:
			for(int i=0;i<ca.size();i++) {
				bkp=ca.get(i);
				int[] va=new int[al];
				// 转整型数组
				for(int ii=0;ii<al;ii++) {
					va[ii]=Integer.parseInt(ca.get(i).substring(ii, ii+1));
				}

				// 初始化检测值
				int ck_1=1;
				int ck_2=0;
				int ck_3=0;
				int ck_4=0;
				int ck_5=0;
				int lenVa=va.length;
				
				// 检测一 条件四
				for(int j=0;j<lenVa;j++){
					if(va[j]==1 || va[j]==4 || va[j]==7){
						ck_1=0;
						break LOOP;
					}
				}
				

				// 检测二 条件一
				if(va[0]==2 || va[2]==6){
					ck_2=1;
				}
				
				// 检测三 条件二
				for(int j=0;j<lenVa;j++){
					if((va[j]==2 && j!=0) || (va[j]==5 && j!=1) || (va[j]==8 && j!=2)){
						ck_3=1;
					}
				}
							
				// 检测四 条件三
				int correctNumber=0;
				for(int k=0;k<va.length;k++){
					if(va[k]==6 || va[k]==9 || va[k]==2){
						correctNumber+=1;
					}
				}
				// 2个号码正确
				if(correctNumber==2){
					if((va[0]!=6 && va[1]!=9 && va[2]!=2) || (va[1]!=9 && va[2]!=2 && va[0]!=6) || (va[0]!=6 && va[2]!=2 && va[1]!=9)){
						ck_4=1;
					}
				}
				
				// 检测五 条件五
				if(va[0]==9 || va[1]==9){
					ck_5=1;
				}
				
				// 筛选输出同时符合5个条件的数组
				int result=ck_1*ck_2*ck_3*ck_4*ck_5;
				if(result==1){
					return bkp;
				}
				
			}		
		
			return "";
		
		}
		
		// 排列
		public static ArrayList<String> sequence(ArrayList<Integer> ma) {
							 
			 ArrayList<String> clt=new ArrayList<String>();
			 ArrayList<String> yet=new ArrayList<String>();
			 String ce1 = String.valueOf(ma.get(0))+String.valueOf(ma.get(1));
			 String ce2 = String.valueOf(ma.get(1))+String.valueOf(ma.get(0));			 
			 clt.add(ce1);
			 clt.add(ce2);
			 for(int i=2;i<ma.size();i++){
				for(int j=0;j<clt.size();j++){
					// 新序列
					String cb=String.valueOf(ma.get(i))+String.valueOf(clt.get(j));
					for(int k=0;k<cb.length();k++){
						String cm=cb.substring(1)+cb.substring(0, 1);
						cb=cm;
						yet.add(cb);
					}	
				}
				clt=yet;
			}		
			return clt;
			
		}
		
		
		public static void main(String[] args) {
			
			// 密码有3个
			int al=3;
			// 5个已知条件
			int num=5;
			
			// 5个已知条件数组
			ArrayList[] cds = new ArrayList[num];
			// 1个号码正确位置正确
			cds[0] = new ArrayList<Integer>(Arrays.asList(2,4,6));
			// 1个号码正确位置不正确
			cds[1] = new ArrayList<Integer>(Arrays.asList(2,5,8));
			// 有2个号码正确位置都不正确
			cds[2] = new ArrayList<Integer>(Arrays.asList(6,9,2));
			// 没有一个号码正确
			cds[3] = new ArrayList<Integer>(Arrays.asList(1,7,4));
			// 有1个号码正确位置不正确
			cds[4] = new ArrayList<Integer>(Arrays.asList(4,1,9));
	
			Set<Integer> set = new TreeSet<Integer>();
			for(int i=0;i<5;i++) {
				set.addAll(cds[i]);
			}
			
	        // 去重取合并结果
	        ArrayList<Integer> ca = new ArrayList<Integer>();
	        for(Integer inte:set){
	        	ca.add(inte);       
	        }
	      
	        // 没有一个号码正确，删除1,4,7
	        for (int i = 0; i < ca.size(); i++) {
	            Integer v = ca.get(i);
	            if(v==1 || v==4 || v==7) {
	            	ca.remove(v);
	            }
	        }
        
	        // 组合
	        int lenCa=ca.size();
	        ArrayList<ArrayList<Integer>> mtva=new ArrayList<ArrayList<Integer>>();
			for(int i=0;i<lenCa;i++){
				for(int j=i+1;j<lenCa;j++){
					for(int p=j+1;p<lenCa;p++){
						ArrayList<Integer> tmp=new ArrayList<Integer>();
						tmp.add(ca.get(i));
						tmp.add(ca.get(j));
						tmp.add(ca.get(p));
						mtva.add(tmp);
					}
				}
			}
			
					
			// 执行计算
			for(int i=0;i<mtva.size();i++) {				
				ArrayList<String> va=new ArrayList<String>();
				va=GuessPassword.sequence(mtva.get(i));		
				String answer=GuessPwd.seekOut(al,va);
				if(answer!=""){
					System.out.print("password is : "+answer);
					System.exit(0);
				}					
			}
			
			
		}
}
