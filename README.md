# WASM ssh-keygen example

これはWASMを利用して、ブラウザ上でSSH鍵ペアを生成するサンプルです。

## デモページ

[Demo](https://chibiegg.github.io/wasm-ssh-keygen/index.html)


## 実行方法

```bash
docker build -t wasm-ssh-keygen .
docker run -p 8080:80 wasm-ssh-keygen
```

access to http://127.0.0.1:8080/ .
