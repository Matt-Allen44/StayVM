![StayVM](https://raw.githubusercontent.com/Matt-Allen44/StayVM/master/res/StayVM-Logo.png?token=AGDdTsoryIN2M6RflCp5xf1ghVOgrFPuks5VYamkwA%3D%3D)


#Running 
``` 
Stay <FILE_NAME>.stay
```

#Example Programs

**Infinite Incremental Loop**
```
0	|		PUSH, 1,	//PLACE 1 ON THE STACK
2	|		GET, 0,		//GET FROM 0 IN THE HEAP 
4	|		ADD,		//ADD 1 AND GET0 TOGETHER
5	|		MOV, 0,		//OVERWRITE HEAP0
7	|		CLRS,		//CLEAR STACK AND RESET
8	|		GOTO, 0,	//LOOP BACK TO START
10	|		HALT,		//EXIT (NEVER REACHED)
```

**Add/Subtract and Print**
```
0	|		PUSH, 1,	//PLACE 1 ON THE STACK
2	|		PUSH, 2,	//PLACE 2 ON THE STACK
4	|		ADD/SUB,	//ADD/SUB 1 AND 2
5	|		PRINT,		//PRINT SUM
6	|		HALT,	  	//EXIT
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

1 |   PUSH		- CALLS PUSH ON NEXT VALUE
2 |		12			- PUTS 12 INTO THE STACK
3 |		PUSH		- CALLS PUSH ON NEXT VALUE
4 |		2			- PUTS 12 I NTO THE STACK
5 |		ADD 		- CALLS ADD ON THE BOTTOM 2 VALUES IN STACK (12 AND 2)
6 | 	PRINT		- PRINTS VALUE FROM BOTTOM OF THE STACK (14)

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
PRINT |	NONE	  | PRINTS FROM CURRENT PLACE IN STACK
JIG   | LINE      | JUMP IF A > B
JIL   | LINE      | JUMP IF A < B
JIE   | LINE      | JUMP IF A = B
CLRS  | NONE     | CLEAR STACK
CLRH  | NONE      | CLEAR HEAP

GET   | MEM_LOCATION | GET FROM HEAP TO STACKcm
MOV  | MEM_LOCATION | MOVE FROM STACK TO HEAP

#Exit Codes
| EXIT CODE | DESCRIPTION |
| :-: 		|:------------|
| 601  		| STAY_VM STACK OVERFLOW (SC>LEN(STACK))|
| 602		| STAY_VM SYNTAX ERROR	|
| 603		| STAY_VM IO ERROR	|
