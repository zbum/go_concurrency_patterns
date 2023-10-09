# go_concurrency_patterns

## 왜 Concurrency (동시성) 을 사용해야 하는가?
* 사람을은 말하면서도 트윗을 한다. 
* 실제 세상에는 여러 사람들, 달리는 차들 등 .. 독립적인 아이템들로 가득차 있다. 
* 이런 세상과 상호작용하거나 시뮬레이션하고자 한다면 순차적인 실행은 적절한 접근법이 아닐 것이다.
* Concurrency(동시성) 은 실세계를 다루는 프로그램을 작성하는 방법일 것이다.

## Concurrency (동시성) 이란 무엇인가?
* 독립적인 계산실행의 집합이다.
* 소프트웨어어를 구조화하는 방법이며, 특히 실세계와 적절한 상호작용을 위한 클린코드를 작성하는 방법이다.
* Concurrency (동시성) 은 Parallelism(병렬성) 이 아니다. 

## Concurrency is not parallelism
* go 프로그램을 처음 접하는 분은 대부분 go 는 병렬 언어다 라고 착각하게 된다. 
* go는 병렬언어가 아니고 동시성 언어이다.
* 만약 단일 프로세서만 가진 시스템을 가지고 있다면 당신의 프로그램은 병렬성을 가질 수 없고 동시성을 가진다고 할 수 있다.
* 반면에 잘 작성된 동시성 프로그램은 멀티 프로세서에서 병렬성을 가진다.

## Concurrency 는 소프트웨어 구축의 모델
* 이해하기 쉽고
* 사용하기 쉽고
* 추론하기 쉬워야 함
* 전문가가 될 필요가 없음

## History
* Go가 처음 나왔을때, 많은 사람들은 새로운 것이라 생각했습니다. 
* 하지만 오랜 역사를 기반으로 한 것임
* Tony Hoare's CSP (1978)의 논문을 기반으로 한 것이며 모든 실제 아이디어가 그 논문에 있습니다.
### 비슷한 기능을 가진 언어들
* Occam
* Erlang
* Newqueak
* Concurrent ML
* Alef
* Limbo

## 구분
* Go 는 Newsqueak-Alef-Limbo 의 가장 최신 브랜치입니다. 
* 다른 언어와 구분되는 점은 channel을 최고 가치로 생각한다는 것입니다. 
* Erlang은 원래 CSP와 더 가깝습니다. 프로세스와 통신하기 위해 채널대신 이름을 사용하고 있습니다. 
* Go 는 프로세스간 통신이 없습니다. 채널과 대화하면 다른 프로세스에서 보내는 값을 읽을 수 있습니다.

## 예제
* 0_basic
* 1_slightly_less_boring
* 2_ignoring_it
* 3_ignoreing_it_a_little_less
## Goroutines
* go 문을 이용해서 독립적으로 함수를 실행하는 것
* 필요에 따라 늘었다 줄었다 하는 자체 호출 스택을 가지고 있음
* 아주 싸다. 보통 몇천개에서 몇십만개의 고루틴을 가진 경우가 일반적이다. 
* 쓰레드가 아니다.

## Communicating
* boring 예제에서 다른 고루틴은 아웃풋을 볼 수 없습니다. 
* 프로세스간 커뮤니케이션이 필요합니다.

## Channel
* 두개의 고루틴이 통신하기 위한 Channel 이 제공됩니다. 
```go
// 선언과 초기화
var c chan int
c = make(chan int)
// 또는
c := make(chan int)
```

```go
// 신호 보내기
c <- 1
```
```go
// 채널로 부터 수신
// 화살표는 데이터 흐름을 표시한다.
value = <-c 
```

## 예제
* 4_using_channels

## Synchronization (동기화)
* main 함수에서 <-c 를 실행할때, 값이 전달될때까지 기다리게 됩니다. 
* 비슷한 방식으로, boring 함수에서 c <- value 를 실행하면, 수신자가 받을 때 까지 기다리게 됩니다. 
* 수신자와 발신자는 각자의 역할을 수행할때 까지 기다리게 됩니다. 
* 따라서 채널은 커뮤니케이션과 동기화를 수행합니다.

## Buffered channels
* go 채널은 버퍼와 함께 생성할 수 있습니다. 
* 버퍼링은 동기화를 제거합니다. 
* 버퍼링은 Erlang 의 메일박스와 비슷합니다.

## Go 접근법
* 공유메모리로 커뮤니케이션 하지 말것

## "패턴"
* gof 의 패턴과 혼동하지 않기 위해 " " 으로 감쌌음

## Generator 
* 채널을 반환하는 함수
## 예제 
* 5_pattern_generator (ex1, ex2)
* Joe, Ann 예제는 Ann 이 수신 되었더라도 Joe 가 수신되지 않았으면 기다려야 한다. 

## Multiplexing
* (ex3)
* fan-in 함수를 사용하면 준비된 정보를 바로 사용할 수 있다.

```
Joe ---> 
          FanIn --->
Ann --->
```

## 시퀀스 복원(Restoring sequence)
* 채널은 first-class 값이기 때문에 채널에 채널을 보내는 것이 가능하다. 
* 채널에 채널을 보내서 고루틴이 그 순서를 기다리게 할 수 있다.
* 모든 메시지를 수신하고, 내부 채널에 메시지를 보내서 다시 가능하게 만든다.

```go
type Message struct {
	str string
	wait chan bool // 신호원... "go-ahead" 라고 말할때 까지 채널을 차단
}
```

## 시퀀스 복원(Restoring sequence)
* 각 발신자는 go-ahead 를 기다립니다. 
```go
for i := 0 ; i < 5 ; i++ {
	msg1 := <-c; fmt.Println(msg1.str)
    msg2 := <-c; fmt.Println(msg2.str)
	msg1.wait <- true
    msg2.wait <- true
}
```
* boring function
```go
waitForIt := make(chan bool)
```

```go
c <- Message{
    fmt.Sprintf("%s, %d", msg, i),
    waitForIt,
}
time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
<-waitForIt
```

## Select
* Select 문은 다중 채널을 처리하는 방법을 제공합니다.
```go
select {
case v1 := <-c1:
	fmt.Printf("received %v from c1\n", v1)
case v2 := <-c2:
    fmt.Printf("received %v from c2\n", v2)
case c3 <- 23
    fmt.Printf("sent %v to c3\n", 23)
default:
    fmt.Printf("no one was ready to communicate\n")
}
```
* 모든 채널을 점검합니다. 
* Select 문의 하나의 커뮤니케이션이 처리될때 까지 블로킹 상태입니다. 
* 여러 커뮤니케이션이 진행되면, 의사랜덤으로 선택합니다.
* default 가 존재하면 채널이 준비되지 않았을때 바로 실행합니다. 

## Fan-in using Select
* 기존 fanIn 펑션을 수정합니다. 
* go 루틴 하나로 동작한다는 것 외에 모두 동일한 기능을 제공합니다. 
```go
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
```

(ex5)

## Timeout using select

