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

