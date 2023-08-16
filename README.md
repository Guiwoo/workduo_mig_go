워크듀오 의 기존 Spring Boot base 코드를 go echo 기반으로 변경하고자 한다.

# 왜 코드 마이그레이션을 하는가?
    - Go 언어가 메인이 되었고, 이제는 자바보다 Go 가 더편하다. 예전에 했던 작업을 다시 Go 로 해보고 싶다.
    - 지난번 MSA 를 적용하기위해 작성하다 취직이 되어 미루어졌던 작업을 Go 언어로 다시 시작하고자 한다.

## 1. ERD
[link](https://www.erdcloud.com/d/2FgvGBc45wFqw6qMX)
<a href="https://www.erdcloud.com/d/2FgvGBc45wFqw6qMX" target="_blank">ERD 보러가기 </a>
## 2. 기존 Spring Boot base 코드
[link](https://github.com/Guiwoo/WorkDuo_dev)
<a href="https://github.com/Guiwoo/WorkDuo_dev" target="_blank">기존 Spring Boot base 코드 보러가기 </a>
## 3. API 명세서
[link](https://alive-tern-b83.notion.site/WORKDUO-55b0477f47c74e0683678ba35d311968)
<a href="https://alive-tern-b83.notion.site/WORKDUO-55b0477f47c74e0683678ba35d311968" target="_blank">API 명세서 보러가기 </a>

## 4. 개발환경 
- Go 1.21.0
- MySql 8.0.25
- Echo, GORM 은 Module 은 work 로 관리

## 5. 프로젝트 구조
모듈이 추가될때마다 업데이트 예정
```
├── README.md
├── go.work
```