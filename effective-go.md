## 명칭

> 이름의 첫 문자가 대문자인지 아닌지에 따라서 이름의 패키지 밖에서의 노출 여부가 결정된다.



#### 패키지명

- 소문자, 한 단어로만 부여

- 패키지 명은 소스 디렉토리 이름 기반

  ``src/encoding/base64``  로 사용,  `encoding_base64`로 쓰지 않는다.

- 최대한 간결하게

  

#### Getters

- getter의 이름에 Get을 넣는건 Go언어 답지 않다
- `owner` 라는 필드의 getter 메서드는 `GetOwner` 가 아니라 `Owner` (첫 문자가 대문자이며, 패키지 밖으로 노출됨 ) 라고 불러야 한다.
- 패키지 밖으로 노출하기 위해 대문자 이름을 사용하는 것은 메서드로부터 필드를 식별할 수 있는 훅을 제공한다.
- 필요하다면 setter함수는 `SetOwner` 라고 불리는 것이 좋다.



#### 대소문자 혼잡

- 언더바 대신 대소문자 혼합방식을 사용

  

## 세미콜론

> Go의 정식문법은 구문 구분을 위해 세미콜론을 사용한다. 그러나 구문분석기(lexer)가 자동으로 세미콜론을 삽입해줌으로 소스작성시 대부분 세미콜론을 사용하지 않는다.

- 세미콜론은 닫는 중괄호 바로 앞에서 생략 할 수 있다.

- **제어문(if, for 등)의 여는 중괄호를 다음 라인에 사용하지 않아야 한다.** 따라서 다음과 같이 사용하는 것이 옳다.

  ```go
  if i < f() {
    g()
  }
  ```

  반면에 다음과 같이 사용하면 세미콜론은 여는 중괄호({) 앞에 추가되어 예상하지 못한 결과를 낼 수 있다.

  ```go
  // void pattern
  if i < f() 
  {
    g()
  }
  ```



## 제어문

#### For

```go
// C언어와 같은 경우
for init; condition; post { }

// C언어의 while 처럼 사용
for condition { }

// key,value range 탐색
for key, value := range oldMap { 
    newMap[key] = value
}

// key range 탐색
for key := range m {
    //something
}

// value range 탐색
for _, value := range m {
    //something
}

// Reverse a, 여러 개의 변수를 사용할려면 병렬 할당을 사용
// 병렬 할당 --> i, j := 0, len(a)-1
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```



#### Switch

- Go 언어에서는 스위치는 C언어보다 더 일반적인 표현이 가능, 따라서 `if-else-if-else` 형태보다 `switch`  로 작성하는 것이 더 Go 언어답다.

> 이 부분은 이해가 잘 가지 않는다. 향후 이해가 가면 설명을 덧붙이도록 하자

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```



#### Type switch

- 스위치 구문은 **인터페이스 변수의 동적 타입**을 확인하는데 사용 할 수 있다.

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```



## 함수

>  함수에 대한 기본은 [go-syntax]() 참고

**Defer**

defer를 실행하는 함수가 반환되기 전에 즉각 함수 호출을 실행하도록 예약한다.

```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // Contens 함수가 반환되기 직전까지 실행이 미루어진다.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```

위와 같이 f.Close()와 같은 함수의 호출을 지연시키면 아래와 같은 장점을 얻게 된다.

- 파일을 닫는 것을 잊어버리는 실수를 하지 않도록 보장
- open 근처에 close가 위치하면 함수 맨 끝에 위치하는 것 보다 훨씬 명확한 코드가 되는것을 의미한다.

**defer 함수의 매개변수들은 함수의 호출이 실행될때가 아닌 defer가 실행될 때 평가된다. 지연된 함수들은 LIFO 순서에 의해 실행된다.**

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
// 4 3 2 1 0 출력
// 가장 늦게 들어온 마지막 defer가 먼저 출력되는데, defer문을 만났을 때 i는 4임으로 매개변수는 4로 확정되어 스택에 저장된다.
```

