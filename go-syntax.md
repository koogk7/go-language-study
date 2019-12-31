# Go Syntax



### Variable

- 변수는 `var`  키워드로 선언하고 그 뒤에 변수명과 변수 타입을 적는다. 형식은 아래와 같다

  `var ${변수 이름} ${type}  `

- **선언된 변수가 사용되지 않는다면 에러를 발생시킨다.**

- 동일한 타입의 여러 개의 변수는 아래와 같이 한번에 지정할 수 있다.

  복수 개의 변수. 선언 - ```. var i,j,k int```

  복수 개의 변수 초기화 -  ```. var i,j,k int = 1, 2, 3``` 

- 선언 후 초기화를 하지 않으면  **Zero Value**를 기본적으로 할당한다. 즉 숫자형에는 0, bool 타입에는 fasle, string 타입에는 ""를 할당한다.

- 타입은 생략이 가능하며, 할당되는 값을 보고 타입 추론이 가능하다.

- 함수 내에서는 `:=` 를 사용해 var 키워드를 생략 할 수 있다.



### Const

- 상수는 const를 사용하여 선언한다.

  ```const ${상수 이름} ${type}```

- 여러 개의 상수는 아래와 같이 나열하여 사용 할 수 있다.

  ```go
  const (
      Visa = "Visa"
      Master = "MasterCard"
      Amex = "American Express"
  )
  
  const ( // ioto라는 identifier를 사용하여 0부터 1씩 증가시켜 값을 부여 할 수 있다. 
      Apple = iota // 0
      Grape        // 1
      Orange       // 2
  )
  ```
- 숫자형 상수는 var로 표현할 수 없는 범위를 저장하는 등 수를 정밀하게 표현 할 수 있다.


### String

- **Block Quote('')**로 둘러 싸인 문자열은 **Raw String Literal**이라고 불리며 이 안에 있는 문자열은 **별도로 해석되지 않는다**. 예를 들어 \n가 있을 경우 NewLine으로 해석되지 않는다. **복수 라인의 문자열을 표현 할 때 자주 사용된다.**
- **이중인용부호("")**로 둘러 싸인 문자열은 **Interpreted String Literal**이라 불리며 복수 라인에 걸쳐 쓸 수 없으며, 이 안의 **Escape 문자열들은 특별한 으미로 해석**된다. 여러 라인에 걸쳐 쓰기 위해서는 **+ 연산자를 이용해 결합**하여 사용한다.



### Type Conversion

- Go에서는 암묵적 변환이 일어나지 않으므로 반드시 명시적으로 변환타입을 지정해주어야 한다.

```go
var i int = 100
var u uint = uint(i)

str := "ABC"
bytes := []byte(str)
```



### Conditional Statement

- If 문에서 조건식 이전에 간단한 문장(Optional Statement)를 함께 실행 할 수 있다. 단 이때 정의된 변수는 if-else문 안에서만 사용 할 수 있다.

  ```go
  // Optional Statement는 switch, for문 등에서도 사용 할 수 있다.
  if val := i * 2; val < max {
      println(val)
  }
  ```
- **else 문은 if가 닫히는 } 와 같은 줄에 작성되어야 한다**
  ```go
    if val := i * 2; val < max {
        println(val)
    } else {
      // ...something
    }

- **swtich문에서 break를 사용하지 않아도 다음 case로 가지 않는다.**
- 다른 언어와 달리swtich문 뒤에 expression이 없을 수도 있다. 이 경우 true로 보고 첫번째 case문으로 이동하여 검사한다.



### Function

> Go 언어에서 함수는 일급객체이다.

```go
package main
func main() {
  	// 함수는 패키지 안에 정의되며, 호출되는 함수가 반드시 앞에 있을 필요는 없다.
    say("This", "is", "a", "book")
    say("Hi")
  
    count, total := sum(1, 7, 3, 5, 9)
    println(count, total)   
}
 
func say(msg ...string) { // ...는 가변 파라미터를 나타낸다.
    for _, s := range msg {
        println(s)
    }
}

// 리턴타입은 함수 이름 다음에 명시한다.
// 여러 개의 리턴이 가능하다.
func sum(nums ...int) (int, int) {
    s := 0      // 합계
    count := 0  // 요소 갯수
    for _, n := range nums {
        s += n
        count++
    }
    return count, s
}

// 아래와 같이 리턴값에 이름을 붙일 수 있다.
// 이 경우에는 리턴에 아무 값을 담아주지 않는다.
func sum(nums ...int) (count int, total int) {
    for _, n := range nums {
      total += n
    }
    count = len(nums)
    return
}
```

- type 문을 이용한 함수 원형 정의

```go
// 원형 정의
type calculator func(int, int) int
 
// calculator 원형 사용, 이를 Delegate라고 부른다.
func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
```



### Collection

**Array**

- **Go에서 배열크기는 Type을 구성하는 한 요소이다.** 즉 [3]int와 [5]int는 서로 다른 타입이다.

- 초기화

  ```go
  var a1 = [3]int{1, 2, 3}
  var a3 = [...]int{1, 2, 3} //배열크기 자동으로
  
  // 다차원 배열
  var multiArray [3][4][5]int  // 정의
  var a = [2][3]int{
          {1, 2, 3},
          {4, 5, 6}, 
      } // 초기화
  ```



**Slice**

> 슬라이스는 실제 배열을 가리키는 포인터 정보, 세그먼트의 길이, 용량 세가지 필드로 구성되어 있다.

- 내부적으로 배열로 구현되어 있지만, 고정된 크기를 지정하지 않고 크기를 동적으로 변경할 수 있으며 부분 배열을 추출 할 수 있다.

  ```go
  package main
  import "fmt"
   
  func main() {
      var a []int        //슬라이스 변수 선언
      a = []int{1, 2, 3} //슬라이스에 리터럴값 지정
      a[1] = 10
      fmt.Println(a)     // [1, 10, 3]출력
    
    // make(type, len, capacity)로 슬라이스 선언가능
    // length 만큼 Zero value가 할당된다.
  	  s := make([]int, 5, 10) 
      println(len(s), cap(s)) // len 5, cap 10
  }
  ```

  > 슬라이스에 별도의 길이와 용량을 지정하지 않으면, 기본적으로 길이와 용량이 0 인 슬라이스를 만드는데, 이를 *Nil Slice* 라 하고, nil 과 비교하면 참을 리턴한다.

- 활용

  ```go
  func main() {
    s := []int{0, 1, 2, 3, 4, 5}
    s[2:5]  // 2,3,4 마지막 인덱스는 +1 을 한다. 파이썬과 동일
    s[1:] // 1,2,3,4,5
    s[:3] // 0,1,2
    s = s[:] // 0,1,2,3,4,5
   		
    // 배열 추가
    s = append(s, 6, 7, 8) // 0,1,2,3,4,5,6,7,8
    
    sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}
  
    // 배열을 합치기 위해서는 ...(자바스크립트의 전개 연산자)를 사용한다.
    sliceA = append(sliceA, sliceB...)
  	fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
    
    // 배열 deep 복사
    source := []int{0, 1, 2}
    target := make([]int, len(source), cap(source)*2)
    copy(target, source)
    fmt.Println(target)  // [0 1 2 ] 출력
    println(len(target), cap(target)) // 3, 6 출력
  }
  ```

  > 추가시에 용량을 초과하는 경우 기존 용량의 2배에 해당하는 새로운 Underlying array를 생성하고 기존 배열 값들을 모두 복제하여 새로 할당한다.



### Struct

- struct는 필드 데이터만 가지며, 메서드를 갖지 않는다. 메서드는 별도로 분리하여 정의한다.

```go
package main
 
import "fmt"
 
// struct 정의
// 외부 패키지에서 사용할 수 있게 할려면 Person으로 정의하면 된다.
type person struct { 
    name string
    age  int
}
 
func main() {
    // person 객체 생성
    p := person{}
     
    // 필드값 설정
    p.name = "Lee"
    p.age = 10
     
    fmt.Println(p)
}
```



### Method

- 메서드는 특별한 형태의 함수이다.

```go
package main
 
//Rect - struct 정의
type Rect struct {
    width, height int
}
 
//Rect의 area() 메소드
//일반 함수와 시그니처가 다르다
func (r Rect) area() int {
    return r.width * r.height   
}
 
func main() {
    rect := Rect{10, 20}
    area := rect.area() //메서드 호출
    println(area)
}
```



### Interface

- 메서드들의 집합체

```go
type Shape interface {
    area() float64
    perimeter() float64
}
```

- Empty Interface는 어떠한 타입도 담을 수 있는 컨테이너로 Dynamic Type으로 볼 수 있다.

```go
package main
 
import "fmt"
 
func main() {
    var x interface{}
    x = 1 
    x = "Tom"
 
    printIt(x)
}
 
func printIt(v interface{}) {
    fmt.Println(v) //Tom
}

```

- 인터페이스 타입의 x에 대해서 x.(T)으로 표현할 경우, x가 nil이 아니며, T 타입일 경우 확인하는 것으로 `Type Assertion` 이라고 부른다. 확인이 시래하면 런타입 에러가 발생한다.

```go
func main() {
    var a interface{} = 1
 
    i := a       // a와 i 는 dynamic type, 값은 1
    j := a.(int) // j는 int 타입, 값은 1
 
    println(i)  // 포인터주소 출력
    println(j)  // 1 출력
}
```

