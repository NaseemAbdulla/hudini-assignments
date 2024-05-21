// 1.===========
package main

func monkeyCount(n int) []int {
   // Your code here, happy coding!
  slice_a := []int{};
  for i:=1;i<=n;i++{
   slice_a = append(slice_a,i)
}
  return slice_a
   
}

// 2.=========================================
package kata

func MakeNegative(x int) int {
  if x>0{
    return -x
  }else {
    return x
  }
    
}
//3. ====================================
package kata

func FindMultiples(integer, limit int) []int {
  // Your code here!
  arr:=[]int{}
  for i:=1;i<=limit;i++{
    if(i%integer==0){
      arr=append(arr,i)
    }
}
  return arr
}
//4. ==========================
package kata


func CountBy(x, n int) []int {
  count:=[]int{}
  for i:=1;i<=n;i++{
    count=append(count,i*x)
}
  return count
}
//5. ===============================================
package kata
import "math"
func PowersOfTwo(n int) []uint64 {
  // your code goes here
  power:=[]uint64{}
  
  for i:=0;i<=n;i++{
    res:=math.Pow(2,float64(i))
    power=append(power,uint64(res))
  }
  return power
}
//6. =========================================
package kata
import "strings"

func Points(games []string) int {
  // your code here!
  var score int
  for i:=0;i<len(games);i++{
    res:=strings.Split(games[i],":")
    if res[0]>res[1]{
      score=score+3
    }else if res[0]<res[1]{
      score+=0
}else{
      score+=1
    }
}
  return score
}
// 7.===========================
package kata

func GetSum(a, b int) int {
  sum:=0;
  max:=0;
  min:=0;
  if(a>b){
    max=a;
    min=b
}else{
    max=b;
    min=a
  }
  if(a==b){
    return a
}
  
  for i:=min;i<=max;i++{
    sum+=i;
}
  return sum
}

//8. ===========================
package kata
import "strings"

func FindShort(s string) int {
  // your code
  shortest:=[]string{}
  shortest=strings.Split(s," ")
  min:=shortest[0];
  for i:=1;i<len(shortest);i++{
    if(len(shortest[i])<len(min)){
      min=(shortest[i])
    }
  }
  return len(min)
}
// 9.==========================
package kata

func SquareSum(numbers []int) int {
    // your code here
  sum:=0;
  for i:=0;i<len(numbers);i++{
    sum+=numbers[i]*numbers[i]
  }
  return sum
}
// 10.============================
package kata

func OddCount(n int) int{
  //your code here
  if n % 2 == 0 {
    // handle even cases
    return n/2
  } else {
    // handle odd cases
    return (n-1)/2
  }
}
// 11.====================================
package kata
import ("strings"
        "fmt"
       )

func SortMyString(s string) string {
  var even string
  var odd string
  str:=strings.Split(s,"")
  for i:=0;i<len(str);i++{
    if i%2==0{
      even+=str[i]
    }else{
      odd+=str[i]
    }
  }
  result:=fmt.Sprintf("%s %s",even,odd)
  return result
}
// 12.===========================================
package kata


func LeastLarger(a []int, i int) int {
  least:=10000;
  index:=-1;
  for j:=0;j<len(a);j++{
    if(a[j]>a[i]){
      if(a[j]<least){
        least=a[j];
        index=j
}
    }
}
  return index
}
//13 ===============================
