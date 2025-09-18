; Assembly Language Calculator
; This program demonstrates basic assembly programming concepts
; including system calls, arithmetic operations, and control flow

section .data
    ; Messages
    prompt db 'Enter two numbers: ', 0
    result db 'Result: ', 0
    newline db 10, 0
    error_msg db 'Error: Invalid input', 10, 0
    menu db 'Choose operation (1=add, 2=sub, 3=mul, 4=div): ', 0
    
    ; Constants
    prompt_len equ $ - prompt - 1
    result_len equ $ - result - 1
    error_len equ $ - error_msg - 1
    menu_len equ $ - menu - 1

section .bss
    ; Input buffers
    num1_str resb 16
    num2_str resb 16
    op_str resb 16
    result_str resb 16

section .text
    global _start

; Convert ASCII string to integer
; Input: RSI = pointer to string, RCX = length
; Output: RAX = integer value
string_to_int:
    push rbx
    push rcx
    push rdx
    
    xor rax, rax        ; Clear result
    xor rbx, rbx        ; Clear digit accumulator
    
    .loop:
        cmp rcx, 0
        je .done
        
        mov bl, [rsi]   ; Get current character
        sub bl, '0'     ; Convert ASCII to digit
        
        ; Check if valid digit (0-9)
        cmp bl, 0
        jl .error
        cmp bl, 9
        jg .error
        
        ; Multiply current result by 10 and add new digit
        mov rdx, 10
        mul rdx
        add rax, rbx
        
        inc rsi         ; Move to next character
        dec rcx         ; Decrement counter
        jmp .loop
    
    .done:
        pop rdx
        pop rcx
        pop rbx
        ret
    
    .error:
        mov rax, -1     ; Return -1 for error
        pop rdx
        pop rcx
        pop rbx
        ret

; Convert integer to ASCII string
; Input: RAX = integer value, RDI = output buffer
; Output: RDI = pointer to null-terminated string
int_to_string:
    push rbx
    push rcx
    push rdx
    
    mov rbx, 10         ; Base 10
    mov rcx, 0          ; Digit counter
    
    ; Handle zero case
    cmp rax, 0
    jne .convert
    mov byte [rdi], '0'
    mov byte [rdi + 1], 0
    jmp .done
    
    .convert:
        cmp rax, 0
        je .reverse
        
        xor rdx, rdx    ; Clear remainder
        div rbx         ; Divide by 10
        add dl, '0'     ; Convert to ASCII
        push rdx        ; Store digit
        inc rcx         ; Increment counter
        jmp .convert
    
    .reverse:
        cmp rcx, 0
        je .done
        
        pop rdx         ; Get digit
        mov [rdi], dl   ; Store in buffer
        inc rdi         ; Move to next position
        dec rcx         ; Decrement counter
        jmp .reverse
    
    .done:
        mov byte [rdi], 0   ; Null terminate
        pop rdx
        pop rcx
        pop rbx
        ret

; Print string
; Input: RSI = string pointer, RDX = length
print_string:
    mov rax, 1          ; sys_write
    mov rdi, 1          ; stdout
    syscall
    ret

; Read string
; Input: RSI = buffer pointer, RDX = buffer size
; Output: RAX = number of bytes read
read_string:
    mov rax, 0          ; sys_read
    mov rdi, 0          ; stdin
    syscall
    ret

; Get string length
; Input: RSI = string pointer
; Output: RCX = string length
strlen:
    push rsi
    mov rcx, 0
    
    .loop:
        cmp byte [rsi], 0
        je .done
        cmp byte [rsi], 10  ; Also stop at newline
        je .done
        inc rsi
        inc rcx
        jmp .loop
    
    .done:
        pop rsi
        ret

; Main program
_start:
    ; Print welcome message
    mov rsi, prompt
    mov rdx, prompt_len
    call print_string
    
    ; Read first number
    mov rsi, num1_str
    mov rdx, 16
    call read_string
    mov r8, rax         ; Save length
    
    ; Read second number
    mov rsi, num2_str
    mov rdx, 16
    call read_string
    mov r9, rax         ; Save length
    
    ; Print operation menu
    mov rsi, menu
    mov rdx, menu_len
    call print_string
    
    ; Read operation choice
    mov rsi, op_str
    mov rdx, 16
    call read_string
    
    ; Convert first number
    mov rsi, num1_str
    mov rcx, r8
    dec rcx             ; Exclude newline
    call string_to_int
    cmp rax, -1
    je .error
    mov r10, rax        ; Store first number
    
    ; Convert second number
    mov rsi, num2_str
    mov rcx, r9
    dec rcx             ; Exclude newline
    call string_to_int
    cmp rax, -1
    je .error
    mov r11, rax        ; Store second number
    
    ; Get operation choice
    mov rsi, op_str
    call strlen
    dec rcx             ; Exclude newline
    call string_to_int
    cmp rax, -1
    je .error
    mov r12, rax        ; Store operation
    
    ; Perform operation
    mov rax, r10        ; Load first number
    mov rbx, r11        ; Load second number
    
    cmp r12, 1
    je .add
    cmp r12, 2
    je .sub
    cmp r12, 3
    je .mul
    cmp r12, 4
    je .div
    jmp .error
    
    .add:
        add rax, rbx
        jmp .print_result
    
    .sub:
        sub rax, rbx
        jmp .print_result
    
    .mul:
        mul rbx
        jmp .print_result
    
    .div:
        cmp rbx, 0
        je .error       ; Division by zero
        xor rdx, rdx    ; Clear remainder
        div rbx
        jmp .print_result
    
    .print_result:
        ; Convert result to string
        mov rdi, result_str
        call int_to_string
        
        ; Print result message
        mov rsi, result
        mov rdx, result_len
        call print_string
        
        ; Print result value
        mov rsi, result_str
        call strlen
        call print_string
        
        ; Print newline
        mov rsi, newline
        mov rdx, 1
        call print_string
        
        jmp .exit
    
    .error:
        mov rsi, error_msg
        mov rdx, error_len
        call print_string
        jmp .exit
    
    .exit:
        mov rax, 60     ; sys_exit
        mov rdi, 0      ; exit status
        syscall
