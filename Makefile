# for dev
clear:
	rm *.exe
	rm *.bin  
	rm *.zip
	rm -rf hydrogen_tmp

t:
	go build .
	cp hydrogen ~/go/bin 
t1: 
	rm hydrogen
	go build
	go run . -output hydrogen -encrypt -garbage -compress
	# ./hydrogen