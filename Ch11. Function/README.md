## 함수
+ 특정 기능을 위해 만든 여러 문장을 묶어서 실행하는 코드 블록 단위
+ ```func 함수이름 (매개변수이름 매개변수형) 반환형```으로 선언

## 함수의 특징

(5번 빼고는 다른 언어와 비슷)
1. 함수를 선언할 때 쓰는 키워드는 'func'이다.
2. '반환형'이 괄호(()) 뒤에 명시된다. 물론 '매개변수형'도 '매개변수이름' 뒤에 명시된다.
3. 함수는 패키지 안에서 정의되고 호출되는 함수가 꼭 호출하는 함수 앞에 있을 필요는 없다. 
4. '반환값'이 여러 개일 수 있다. 이럴 때는 '반환형'을 괄호로 묶어 개수만큼 입력해야한다. ((반환형1, 반환형2)형식, 두 형이 같더라도 두 번 써야 한다)
5. 블록 시작 브레이스({)가 함수 선언과 동시에 첫 줄에 있어야 한다.

## 전역변수와 지역변수
차이점
1. 메모리에 존재하는 시간
   - 지역변수는 함수가 호출되고 끝날때까지
   - 전역변수는 프로그램이 끝날때까지(가급적 피해야 함)
2. 변수에 접근할 수 있는 범위
   - 지역변수는 실행되고 있는 지역에서만 유효
   - 전역변수는 프로그램 전체에서 유효

## 매개변수
### Pass by value
인자의 값을 복사해서 전달하는 방식. 다른 언어와 동일하게 함수를 호출할 때 ```함수이름(변수이름)```만 입력하면 됨.

### Pass by reference
함수를 호출할 때 주솟값 전달을 위해 ```함수이름(&변수이름)```를 입력하고 
함수에서 매개변수 이름을 입력할 때 값을 직접 참조하기 위해 *를 매개변수형 앞에 붙임. 
함수 안에서 매개변수 앞에 모두 *를 붙여야 함.

### 가변 인자 함수
전달하는 매개변수의 개수를 고정한 함수가 아니라 함수를 호출할 때마다 매개변수의 개수를 다르게 전달.

+ ```func 함수이름(매개변수이름 ...매개변수형) 반환형```형식으로 선언
+ n개의 동일한 형의 매개변수 전달
+ 슬라이스 형태
+ ```함수이름(슬라이스이름...)```로 슬라이스 전달할 수 있음.

## 반환값
복수개의 반환값을 반환할 수 있음. 
```func add(num ...int) (int, int)```처럼 동일한 반환형이더라도 괄호 안에 모두 명시해야 함.

### Named Return Parameter
+ 이름이 있는 반환값.
+ ```(반환값이름1 반환형1, 반환값이름2 반환형2, 반환값이름3 반환형3, ...)``` 형식
+ "반환값이름 반환형" 자체가 변수 선언이므로 함수 안에서 따로 선언하면 안됨.
+ 함수 안에서는 꼭 return을 명시해야 함!!
+ 필요에 따라 함수를 변수처럼 사용할 수 있지만 코드 가독성이 나빠질 수 있음...

## 익명 함수
코드를 기능별로 함수화 하는 것의 단점은 '프로그램의 속도 저하'이다.
이유는 다음과 같다.
+ 함수 선언 자체가 프로그래밍 전역으로 초기화되면서 메모리를 잡아먹음.
+ 기능을 수행할 때마다 함수를 찾아서 호출해야 함.

이러한 단점을 보완하기 위해 '익명 함수'가 필요.

### 익명 함수의 특징
1. 함수의 이름만 없고 그 외의 형태는 동일
2. 함수의 블록 마지막 중괄호 뒤에 괄호를 사용해 바로 함수를 호출. 괄호 안에 매개변수를 넣을 수도 있음.
3. 선언 함수보다 익명 함수가 나중에 읽힘

## 일급 함수
함수를 기본 타입과 동일하게 사용할 수 있어 함수 자체를 다른 함수의 매개변수로 전달하거나 다른 함수의 반환 값으로 사용될 수 있다는 것.
함수를 매개변수형으로 사용할 때는 ```매개변수함수이름 func(전달받는함수의매개변수형) 전달받는함수의반환형``` 형태로 선언.

## type문을 사용한 함수 원형 정의
C언어의 구조체 개념과 유사. 함수의 원형을 정의하고 사용자가 정의한 이름을 형으로써 사용.
