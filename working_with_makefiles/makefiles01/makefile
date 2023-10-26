all: say_hello generate write read clean

say_hello:
	@echo "Hello GO developer!"

generate:say_hello
	@echo "generating test file."
	@touch test-file00.txt
	@echo "test file generated."


write:generate
	@echo "writing to test-file00.txt."
	@echo And the light shineth in the darkness, and the darkness conprehended it not. > test-file00.txt
	@echo "successfully wrote to test-file00.txt."

read:write
	@echo "reading from test-file00.txt."
	@cat test-file00.txt

clean:read
	@echo "deleting test-file00.txt."
	@rm test-file00.txt
	@echo "test-file00.txt removed successfully!"
