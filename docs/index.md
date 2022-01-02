# VEN 
Welcome to VEN! A script language to help you automate capturing Venari. For those more advanced and wanting to look at our interpreter, please check the /interpreter file of our project and read through the structuring. **For those that are new or learning, read below for a simple tutorial**

------------

### Keywords
*VEN runs on keywords(s)!* 
VEN runs on a list of commands that lets you control your Venari farmer. Heres a list of them, you can click each one for more detail!
- EQUALS
- ISEMPTY
- GREATER
- GREATEREQUALS
- LESSER
- LESSEREQUALS
- END
- SET
- ACTION
- GOTO
- NONE
- SLEEP
- PRINT
- TIME
- ADD

------------
### Execution
VEN is parsed at runtime into a list of functions, and then executed in order. Keywords like GOTO can modify or change the order of execution and END can stop execution altogether.  The script will login to the account beforehand, so no login/cookie handling is needed.

**Variables must be SET before they are called/used.**

------------
### Formatting
Formatting in VEN is similar to a lot of languages, but has some guidelines.
- **SPACES MATTER**
- 1 Keyword Per Line, unless using a comparative like EQUALS, LESSER, etc...
- Comparatives cannot be called from inside a comparative, the resulting case of an EQUALS, LESSER, GREATER, etc... **cannot** be a LESSER, GREATER, EQUALS, etc...
- Code/Wrapping is formatted as follows
```
SET[varName,1] //The outer wrapping is done with []
EQUALS[x,y,GOTO(1),GOTO(2)] //The wrapping of inner Keywords is done with ()
EQUALS[x,y,ACTION(x.y.z),GOTO(1)] //Pair [ , ] and ( . )
```

**If you are confused, do not worry, you can come back to this part later for reference**

------------
### Reference
##### Set
Set will create a variable with a value, this variable can then be used in functions
```
SET[x,1]
```
```
SET[x,1] //sets x = 1
SET[x,test] sets x = test
SET[x,] sets x to empty
```


##### Print
Print will print the given variable name or input to the console
```
PRINT[x]
```
```
PRINT[test] //test
SET[test,1]
PRINT[test] //1
```


##### Sleep
Sleep will sleep the execution of the script for the given amount of **Milliseconds**
**Sleep will only take integer values, not variables**
```
SLEEP[1000]
```
```
SET[x,1]
SLEEP[5000] //waits 5 seconds
PRINT[x]
```


##### Time
Time will set the given variable name to the Current **Unix Seconds**. Calling time on an existing variable name will overwrite it.
```
TIME[x]
```
```
TIME[x] //sets to the current unix timestamp
PRINT[x]
```


##### Goto
Goto will set the next line that executes. Setting it to a number lower than 1 or higher than the amount of lines in the script will cause an error
```
GOTO[1]
```
```
PRINT[here]
SLEEP[1000]
GOTO[1] //goes back to line 1.
```


##### End
End will stop the execution of the script, you cannot restart after calling End.
```
END[]
```
```
PRINT[1] //1
PRINT[2] //2
END[] //script stops here
PRINT[3]
```

##### Add
Add will add a number to the given variable name, if the given variable name is not a number, an error is thrown.
```
ADD[x,1]
```
```
SET[x,1]
ADD[x,5]
PRINT[x] //6
```


##### Action
Action is the bread, butter, and the meat of the sandwich, Actions are premade functions that allow you to easily do things on Venari. A List of Actions are provided below.
```
ACTION[actionName,...variables]
```
```
SET[energy,] //create energy variable for storage
ACTION[getUserEnergy,energy] //call getUserEnergy Action
PRINT[energy] //ex: 10
```
Actions allow you to call lots of complicated code very easily and simply.


##### Equals
Equals will compare to variables to see if they are equal. It will convert both variables to string then compare. The ifCase is executed if they are equal, otherwise the elseCase is executed.

```
EQUALS[x,y,ifCase,elseCase]
```
```
SET[x,1]
SET[y,2]
EQUALS[x,y,PRINT(x and y equal),PRINT(x and y not equal)] //x and y not equal
EQUALS[x,y,GOTO(5),GOTO(6)]
PRINT[x]
PRINT[y] //2
```


##### None
None does nothing, literally, its as placeholders for GOTO functions, or when you want to do nothing in a ifCase or elseCase when using comparatives. It is called to help exit comparatives without having to do tedious PRINT/GOTO(s).

```
NONE[]
```
```
SET[x,1]
SET[y,1]
EQUALS[x,y,NONE(),GOTO(1)] //none() is called, which will do nothing and move on
PRINT[here] //here
```


##### IsEmpty
Checks if the given variable is empty
```
ISEMPTY[x,ifCase,elseCase]
```
```
SET[x,]
ISEMPTY[x,PRINT(empty),PRINT(not empty)] //empty
```


##### GREATER, LESSER, GREATEREQUALS, LESSEREQUALS
Greater: checks if x is greater than y
```
GREATER[x,y,ifCase,elseCase]
```
Lesser: checks if x is less than y
```
LESSER[x,y,ifCase,elseCase]
```
GreaterEquals: checks if x is greater than or equal to y
```
GREATEREQUALS[x,y,ifCase,elseCase]
```
LesserEquals: checks if x is less than or equal to y
```
LESSEREQUALS[x,y,ifCase,elseCase]
```


------------
### Actions
Here is a list of all possible actions, please remember to use the formatting below when calling them.
```
ACTION[actionName,variables of the action]
```

*Follow the example when using actions*
```
SET[energy,]
ACTION[getUserEnergy,energy]
PRINT[energy]
```
##### User
- `getUserEnergy[resVar]` Returns the amount of energy the user has

##### Expedition
- `getCurrentExpedition[area,resVar]` Returns the current expedition ID
- `createExpedition[area,baitId,resVar]`  Creates a new expedition ID

##### Search
- `searchForVenariByName[area,name,resVar]` Returns a Venari that matches the given name (case sensitive)
- `searchForVenariByTier[area,tier,resVar]` Returns a Venari that matches the given tier
- `searchForVenariByAny[area,resVar]` Returns any present Venari

##### Inventory
- `getInventoryByName[name,resVar]` Returns the ID of an item from the inventory if its name matches

##### Battle
- `startBattle[expeditionId,venariId,baitId,rigId,resVar]` Starts a new battle and returns the Venari's name
- `battleAction[expeditionId,baitId,rigId,action,resVar]` Does a battle action **(Play, Feed, Fight)**
- `battleCatch[expeditionId,resVar]` Attempts a catch, returns the amount of coins or 0 depending if caught

##### Shop
- `getItemByName[city,itemName,resVar]` Gets the ID if any item's match itemName
- `buyItem[itemId,amount,resVar]` Buys the listed amount of the item ID.
