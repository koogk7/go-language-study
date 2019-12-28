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



