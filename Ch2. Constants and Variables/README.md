## println, print
Go 언어에 내장된 콘솔 출력 함수.
print -> 개행 X
println -> 개행 O

## fmt
일반적으로 import 해서 사용하는 콘솔 입출력 패키지. 

```fmt.Print("Hello goorm!", num1, num2, "\n")```
:arrow_forward: 개행 X

``` fmt.Println("Hello goorm!", num1, num2) ```
:arrow_forward: 개행 O, 콤마 사이에 띄어쓰기 O

``` fmt.Printf("num1의 값은:%d num2의 값은:%d\n", num1, num2) ```
:arrow_forward: C언어와 동일한 함수. %d 데이터를 채움.

## 변수 선언 방식
``` var 변수이름 변수형 ``` 
+ 변수를 선언한 곳에서 바로 초기값을 설정할 수 있다
+ 초기값을 설정하지 않으면 'Zero value'로 설정됨 (bool 타입은 false, 숫자 타입은 0, string 타입은 "")
+ ``` := ```를 사용하면 별도의 선언 없이 타입 추론 가능 **(단, 함수 내에서만 사용해야 함)**
+ 선언만 하고 쓰지 않으면 에러 발생 (패키지, 함수, 변수 등에 모두 적용)

## 상수의 선언과 초기화
``` const 상수이름 상수형 ```
+ 반드시 선언과 동시에 초기화 해야 함.
+ ``` := ``` 용법 사용할 수 없음.

##### 이렇게 초기화 할 수도 있음
``` 
const (
	상수이름1 = 값1
	상수이름2 = 값2
	...
)
```
+ 괄호로 같이 묶여있는 상수들은 다른 형으로 초기화될 수 있다.
+ 반드시 각 상수들은 개행하여 초기화해야 함.
+ 각 상수들 사이에 콤마(,)를 입력하면 안 됨.
+ 첫번째 값은 꼭 선언되어야 하고, 선언되지 않은 값은 바로 전 상수의 값을 가짐.
+ ``` iota ```라는 식별자를 값으로 초기화하면 그 후에 초기화하지 않고 이어지는 상수들은 순서(index)가 값으로 저장됨.


