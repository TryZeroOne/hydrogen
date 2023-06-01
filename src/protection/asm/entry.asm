.globl _start

_start:
    call main
    mov $0, %ebx
    mov $1, %eax
    int $0x80