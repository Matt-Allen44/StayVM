![StayVM](https://raw.githubusercontent.com/Matt-Allen44/StayVM/master/res/StayVM-Logo-Small.png?token=AGDdTkUX8QW8exkdW7nRJOvf3pFEyt2Kks5UxM8JwA%3D%3D)


#Running 
``` 
Stay <FILE_NAME>.stay
```

#Runtime

**Launch**

Check that the file is valid, hashing, valid commands, etc.
Syntax checking, error checking, etc.

Parse <FILE_NAME>.stay to an int[]
Execute through commans and their arguments, 

````
LOAD <FILE_NAME>.stay
CHECK

PARSE
EXECUTE

--> LOAD COMMAND TO INT ARRAY
--> EXECUTE COMMAND AND MOVE TO VALID LOCATION IN CODE
--> EG. Execute a PUSH then jump an extra line so the argument is not called as a command

````

**Execution**

*Code*

````

1 |   	PUSH		- CALLS PUSH ON NEXT VALUE
2 |		12			- PUTS 12 INTO THE STACK
3 |		PUSH		- CALLS PUSH ON NEXT VALUE
4 |		2			- PUTS 12 I NTO THE STACK
5 |		ADD 		- CALLS ADD ON THE BOTTOM 2 VALUES IN STACK (12 AND 2)
6 |		PUT 0x2		-
7 | 	PRINT		- PRINTS VALUE FROM BOTTOM OF THE STACK (14)

````

*Stack*

````

1 |   	[0,0,0,0]
2 |		[12,0,0,0]
3 |		[12,0,0,0]
4 |		[12,2,0,0]
5 |		[14,0,0,0]
6 | 	[14,0,0,0]
7 | 	[14,0,0,0]

```

*Heap*

````

1 |   	[0,0,0,0 ]
2 |		[0,0,0,0 ]
3 |		[0,0,0,0 ]
4 |		[0,0,0,0 ]
5 |		[0,0,0,0 ]
6 | 	[0,0,14,0]
7 | 	[0,0,0,0 ]

```

#Instruction Set
| OP  | ARGS      | DESCRIPTION |
| :-: |:----------| :--------:|
PUSH  |	NONE	  | PUSH TO STACK
ADD   |	NONE	  | ADD TWO VALUES FROM STACK
SUB   |	NONE 	  | TAKE TWO VALUES B-A 
GOTO  | LINE 	  | JUMP TO LINE
JIG   | LINE      | JUMP IF A > B
PRINT |	NONE	  | PRINTS FROM CURRENT PLACE IN STACK

GET   | MEM_LOCATION | LOC VAL //PUSH TO HEAP FROM STACK
PUT   | MEM_LOCATION | GET FROM HEAP TO STACK
