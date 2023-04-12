package main

import (
	"log"
)

/*
입력 형식
 - dgrep {keyword} {relative path}

출력 형식
 - 파일의 각 line에 keyword가 있는 경우, 해당 파일과 줄 번호를 출력한다. <- 파일 명을 말하는 거겠지?
*/

/*
	- 멀티 쓰레드를 사용한다
		- golang에서는 goroutine이라는 경량쓰레드를 사용한다.
		- 일단 경량 쓰레드라고 생각하자

	- command line parse
		- os.Args가 3이 아니면 error
		- os.Args의 1: keyword
		- os.Args의 2: relative path

	- path가 dir인지 file인지 확인
		- file일 경우 단일 실행
		- dir인 경우 복수 실행

	- 단일 파일이라면
	- 복수의 파일이라면

	- target files path를 구함
	- 해당 파일들에 대하여 grep을 실행
	- 단일 파일이 끝나면, 출력쓰레드로 결과 전달
*/

func main() {
	if err := parseCommandLine(); err != nil {
		log.Panic(err)
	}

	ok, err := isDir()
	if err != nil {
		log.Panic(err)
	}

	if !ok {
		single()
		return
	}

	multi()
}

func single() {
	if err := dgrep(relativePath); err != nil {
		log.Panic(err)
	}
}

func multi() {
	if err := findTarget(); err != nil {
		log.Panic(err)
	}
}
