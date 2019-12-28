### Pass by value

- Pass by alue, 객체도 Pass by value로 넘어간다

  - for loop 안에서도 값이 복사가 됨

    ```go
    unc main() {
      employees := make([]Emp, 1)
      employees[0] = Emp{
        Name: "babokim",
        Salary: 1000,
      }
      for _, emp := range employees {
        emp.Salary = 2000
      } // for loop 이 방법 권장
      fmt.Println("After for each statement:", employees[0].Salary)
      
      for i := 0; i < len(employees); i++ {
        employees[i].Salary = 3000;
      } // 권장하지 않음 
      
      fmt.Println("After for with index:", employees[0].Salary)
    }
    ```

    위 코드 첫번째 출력에서 100이 출력된다. emp 변수는 employees의 실제 객체가 아닌 복사된 값이 전달되기 때문이다. (map 함수와 유사)

  - 객체를 넘겨주시 위해서는 C/C++의 포인터 개념을 이용