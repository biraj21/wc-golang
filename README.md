# Unix's `wc` in Golang

My implementaion of Unix's `wc` tool in Golang. Challenge - [Write Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc/).

```
usage: wc [OPTION]... [FILE]...
```

## Without options

```
$ go run . test.txt
```

## With options

```
$ go run . -c -l test.txt
```
