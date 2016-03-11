package main

type room struct {
	// forwardは他のクライントに転送するためのメッセージを保存するチャネルです
	forward chan []byte
}
